
package main


import (
	"fmt"
	"time"
	"bytes"
	"strings"
	"strconv"
	"io/ioutil"
	"net/http"
  "encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type serviceTX struct {
	ObjectType     string    `json:"objectType"`
	TxID           string    `json:"txID"`
	TaskId         string    `json:"taskId"`
	Requester      string    `json:"requester"`
	Provider       string    `json:"provider"`
	Url            string    `json:"url"`
	Success        string    `json:"success"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"endTime"`
}


type serviceRes  struct {
	Message        string   	 `json:"message"`
	StartTime      time.Time   `json:"startTime"`
	EndTime        time.Time   `json:"endTime"`
	ReqResTime     float64     `json:"reqResTime"`
	ReqThroughput  float64     `json:"reqThroughput"`
}


//var agreementMap = make(map[string]agreement)
var agreementMap = make(map[string]agreement)

// ==========================================
//      saveServiceTX - Save serviceTX
// ==========================================
func (t *TaskChaincode) saveServiceTX(stub shim.ChaincodeStubInterface, args []string) pb.Response{

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return shim.Error("6th argument must be a non-empty string")
	}

  taskId := args[0];
	requester := args[1];
	provider := args[2];
	url := args[3];
	success := args[4];
	startTimeString := args[5];
	endTimeString  := args[6];

	startTime, startErr := time.Parse("2006-01-02T15:04:05Z", startTimeString)
	if startErr != nil {
		return shim.Error("Failed to convert the startTimeString " + startTimeString + " into type time.Time")
	}

	endTime, endErr := time.Parse("2006-01-02T15:04:05Z", endTimeString)
	if endErr != nil {
		return shim.Error("Failed to convert the endTimeString " + endTimeString + " into type time.Time")
	}

	var idStr string
	id, err := iw.NextId()
	if err != nil {
		fmt.Println(err)
	} else {
		idStr = strconv.FormatInt(id, 10)
	}

	txJSON := &serviceTX{"ServiceTX", idStr, taskId, requester, provider, url, success, startTime, endTime}
	txJSONasBytes, err := json.Marshal(txJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Save Transaction ====
	err = stub.PutState(idStr, txJSONasBytes)
	if err != nil {
		return shim.Error("Failed to save transaction " + string(txJSONasBytes))
	}

	return shim.Success(txJSONasBytes)
}


// =================================================
//      invokeRestAPI - Client of Web service
// =================================================
func (t *TaskChaincode) invokeRestAPI(stub shim.ChaincodeStubInterface, args []string) pb.Response{

	//       0           1           2             3          4        5         6 (method为post才有)
	// "$requester", "$taskId", "$provider", "$tolerance", "$url", "$method", "$args"
	if len(args) < 6 {
		return shim.Error("Incorrect number of arguments. Expecting 7+")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th argument must be a non-empty string")
	}

	requester := args[0]
	taskId := args[1]
	provider := args[2]

	// Tolerance，if real responseTime is 1.5 times of responseTime in agreement，we should find a new service
  tolerance, err := strconv.ParseFloat(args[3], 32)
	if err != nil {
		return shim.Error("strconv.ParseFloat: parsing the 4th argument " + args[3] + ": invalid synatx")
	}
	if tolerance < 1.0 {
		tolerance = 1.0
	}

	url := args[4]
  method := args[5]

	agree, exist := agreementMap[taskId]
	if !exist {
		// ==== Get current state of task ====
		stateAsByte, errState := stub.GetState(taskId)
		if errState != nil {
			fmt.Println("Failed to read state!")
			return shim.Error("Failed to read state!")
		}

		if stateAsByte == nil {
			fmt.Println("no state! task: " + taskId)
			return shim.Error("no state! task: " + taskId)
		} else {
			state := string(stateAsByte)

		  if state != "acception" {
			  fmt.Println("Current state is " + state + ", stop execute! You can execute only state is acception!\n")
			  return shim.Error("Can't execute. Current state is " + state + ", you can execute only state is acception!")
      } else {
		  	objectType := "agreement"
			  queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"%s\",\"taskId\":\"%s\"}}", objectType, taskId)
			  queryResults, err := getResultForQueryString(stub, queryString)
			  if err != nil {
			  	return shim.Error(err.Error())
			  }

		  	var agreements []agreement
			  err = json.Unmarshal(queryResults, &agreements)
			  if err != nil {
			  	return shim.Error(err.Error())
		  	}

			  if len(agreements) <= 0 {
			  	return shim.Error("no agreement for the task " + taskId)
			  }

			  agree = agreements[0]

        // ==== If the length of agreementMap > 2000, clear up agreementMap ====
			  if len(agreementMap) > 2000 {
					agreementMap = make(map[string]agreement)
				}

			  agreementMap[taskId] = agree
		  }
	  }
  }

	aRequester := agree.Requester
	aProvider := agree.Provider
	aUrl := agree.Url
	aResponseTime := agree.ResponseTime
	aThroughput := agree.Throughput

	if aRequester != requester {
		return shim.Error("no valid requester " + requester)
	} else if aProvider != provider {
		return shim.Error("no valid provider " + provider)
	} else if !strings.Contains(url, aUrl) {
		return shim.Error("no valid url " + url + " Expected " + aUrl)
	} else {
		aBeginTime := agree.BeginTime
		aExpireTime := agree.ExpireTime
		nowTime := time.Now()

		if nowTime.Sub(aBeginTime).Seconds() < 0.0 || nowTime.Sub(aExpireTime).Seconds() > 0.0 {
			aBeginTimeString := aBeginTime.Format("02 Jan 2006 15:04:05 -0700")
			aExpireTimeString := aExpireTime.Format("02 Jan 2006 15:04:05 -0700")
			return shim.Error("no valid time. Expected between " + aBeginTimeString + " and " + aExpireTimeString)
		}
	}

	var idStr string
	id, err := iw.NextId()
	if err != nil {
		fmt.Println(err)
	} else {
		idStr = strconv.FormatInt(id, 10)
	}

	var startTime, endTime time.Time

	if method == "get" {
		// ==== Record start time ====
		startTime = time.Now()

		client := &http.Client{
			Timeout: time.Duration(2.0*aResponseTime*tolerance)*time.Second,
		}

		response, resterr := client.Get(url)

		// ==== Record end time ====
		endTime = time.Now()

		var success string
		if resterr != nil {
			success = "false"
		} else {
			success = "true"
		}

		txJSON := &serviceTX{"ServiceTX", idStr, taskId, requester, provider, url, success, startTime, endTime}
		txJSONasBytes, err := json.Marshal(txJSON)
		if err != nil {
			return shim.Error(err.Error())
		}

		// ==== Save Transaction ====
		err = stub.PutState(idStr, txJSONasBytes)
		if err != nil {
			return shim.Error("Failed to save transaction " + string(txJSONasBytes))
		}
		fmt.Println(txJSON)

		if resterr != nil {
			return shim.Success([]byte("Failed to invoke service. Error code: " + resterr.Error()));
		}

		if response == nil {
			return shim.Success([]byte("Failed to get response from service at url: " + url));
		} else {
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			message := string(body)

			serviceResJSON := &serviceRes{message, startTime, endTime, aResponseTime, aThroughput}
			serviceResAsBytes, err := json.Marshal(serviceResJSON)
			if err != nil {
				return shim.Error(err.Error())
			}

			return shim.Success(serviceResAsBytes)
		}

	} else if method == "post" {
		if len(args) < 7 {
			return shim.Error("Incorrect number of arguments. Expecting 7")
		}
		if len(args[6]) <= 0 {
			return shim.Error("7th argument must be a non-empty string")
		}

		postJson := args[6]
		postStr := []byte(postJson)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(postStr))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{
			Timeout: time.Duration(2.0*aResponseTime*tolerance)*time.Second,
		}

		// ==== Record start time ====
		startTime = time.Now()

		response, resterr := client.Do(req)

		// ==== Record end time ====
		endTime = time.Now()

		var success string
		if resterr != nil {
			success = "false"
		} else {
			success = "true"
		}

		txJSON := &serviceTX{"ServiceTX", idStr, taskId, requester, provider, url, success, startTime, endTime}
		txJSONasBytes, err := json.Marshal(txJSON)
		if err != nil {
			return shim.Error(err.Error())
		}

		// ==== Save Transaction ====
		err = stub.PutState(idStr, txJSONasBytes)
		if err != nil {
			return shim.Error("Failed to save transaction " + string(txJSONasBytes))
		}
		fmt.Println(txJSON)

		if resterr != nil {
			return shim.Success([]byte("Failed to invoke service. Error code: " + resterr.Error()));
		}

		if response == nil {
			return shim.Success([]byte("Failed to get response from service at url: " + url));
		} else {
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			message := string(body)

			serviceResJSON := &serviceRes{message, startTime, endTime, aResponseTime, aThroughput}
			serviceResAsBytes, err := json.Marshal(serviceResJSON)
			if err != nil {
				return shim.Error(err.Error())
			}

			return shim.Success(serviceResAsBytes)
		}

	} else {
		return shim.Error("the method should be get or post")
	}

}


// ==============================================================
//      queryServiceTxByTaskId - Get serviceTX by TaskId
// ==============================================================
func (t *TaskChaincode) queryServiceTxByTaskId(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//     0
	// "$taskId"
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	taskId := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"ServiceTX\",\"taskId\":\"%s\"}}", taskId)
	queryResults, err := getResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

// ===========================================================================
//      queryServiceTxByDate - Get serviceTX by date
// ===========================================================================
func (t *TaskChaincode) queryServiceTxByDate(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//         0                1
	// "2018-08-08T00", "2018-08-09T24"
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	gteTime := args[0]
	lteTime := args[1]

  selectorString :=
    "{" +
			"\"selector\":{" +
				"\"objectType\":\"ServiceTX\"," +
        "\"$and\": [" +
           "{" +
					    "\"startTime\":{\"$gte\":\"%s\"}" +
					 "}," +
           "{" +
						  "\"startTime\":{\"$lte\":\"%s\"}" +
					 "}" +
				"]"+
			"}"+
		"}";
  queryString := fmt.Sprintf(selectorString, gteTime, lteTime)
	queryResults, err := getResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}


// ===================================================================================
//      queryServiceTx - Get serviceTX by TaskId, Requester, Provider and date
// ===================================================================================
func (t *TaskChaincode) queryServiceTx(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//     0        1       2           3                 4
	// "$taskId", "Jim", "Deke", "2018-08-08T00", "2018-08-09T24"
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}

	taskId := args[0]
	requester := args[1]
	provider := args[2]
	gteTime := args[3]
	lteTime := args[4]

  selectorString :=
    "{" +
			"\"selector\":{" +
				"\"objectType\":\"ServiceTX\"," +
				"\"taskId\":\"%s\"," +
				"\"requester\":\"%s\"," +
				"\"provider\":\"%s\"," +
        "\"$and\": [" +
           "{" +
					    "\"startTime\":{\"$gte\":\"%s\"}" +
					 "}," +
           "{" +
						  "\"startTime\":{\"$lte\":\"%s\"}" +
					 "}" +
				"]"+
			"}"+
		"}";
  queryString := fmt.Sprintf(selectorString, taskId, requester, provider, gteTime, lteTime)
	queryResults, err := getResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

// =====================================================================================
//        delete - remove serviceTX
// =====================================================================================
func (t *TaskChaincode) deleteServiceTX(stub shim.ChaincodeStubInterface) pb.Response {

	serviceTXArray := make([]serviceTX, 0)
	var i int
	var txErr error

	// ==== Query all serviceTX ====
	queryServiceTXString := fmt.Sprintf("{\"selector\":{\"objectType\":\"ServiceTX\"}}")
  serviceTXArray, txErr = getArrayForServiceTX(stub, queryServiceTXString)

	if txErr != nil {
		return shim.Error("test1" + txErr.Error())
	}

	for i = 0; i < len(serviceTXArray); i ++ {
		serviceTX := serviceTXArray[i]
		txID := serviceTX.TxID

		//fmt.Sprintf(txID)

		//删除serviceTX
		err := stub.DelState(txID)
		if err != nil {
			return shim.Error("Failed to delete state:" + err.Error())
		}
	}

	return shim.Success([]byte("deleted!"))
}
