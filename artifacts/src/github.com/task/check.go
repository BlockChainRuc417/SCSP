
package main


import (
	"fmt"
	"encoding/json"
	"strconv"
	"math"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type agreement struct {
	Requester             string      `json:"requester"`
	TaskId                string      `json:"taskId"`
	Provider              string      `json:"provider"`
	ObjectType            string      `json:"objectType"`
	Url                   string      `json:"url"`
	BeginTime             time.Time   `json:"beginTime"`
	ExpireTime            time.Time   `json:"expireTime"`
	ResponseTime          float64     `json:"responseTime"`
	Throughput            float64     `json:"throughput"`
	FinalPrice            float64     `json:"finalPrice"`
	RequestSigR           string      `json:"requestSigR"`
	RequestSigS           string      `json:"requestSigS"`
	RequestCertificate    string      `json:"certificate"`
	ResponseSigR          string      `json:"responseSigR"`
	ResponseSigS          string      `json:"responseSigS"`
	ResponseCertificate   string      `json:"responseCertificate"`
}


// ========================================================================================
//        Check - check responses that meet all requirements and write into agreement
// ========================================================================================
func (t *TaskChaincode) check(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	requestArray := make([]request, 0)
	responseArray := make([]response, 0)
	var requ request
	var requestErr, responseErr error
	var i, j, k int
	var maxQoS float64 = 0
	var secondLowPrice, lowestPrice, reservePrice, aFinalPrice float64
	var maxQoSReqIndex int
	var minPriceResIndex string
	responseTimeSlice := make([]float64, 0)
	throughputSlice := make([]float64, 0)

	var satResponseMap map[int][]response
	satResponseMap = make(map[int][]response)      //init satResponseMap

	//     0
	// "$taskId"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	taskId := args[0]

	stateAsByte, err := stub.GetState(taskId)
	if err != nil {
		fmt.Println("Failed to read state!")
	}
	state := string(stateAsByte)

	if state != "instantiation" {
		fmt.Println("task " + taskId + ". Current state is " + state + ", you can execute check only state is instantiation!\n")
		return shim.Error("task " + taskId + ". Can't execute check. Current state is " + state + ", you can execute check only state is instantiation!")
	}

	fmt.Println("Current state is " + state + ", continue......")
	fmt.Println("- start check")

	// ==== Query all requests ====
	queryRequestString := fmt.Sprintf("{\"selector\":{\"objectType\":\"request\",\"taskId\":\"%s\"}}", taskId)
	requestArray, requestErr = getArrayForRequest(stub, queryRequestString)
	if requestErr != nil {
		return shim.Error(requestErr.Error())
	}

	// ==== Check if requestArray is null ====
	if len(requestArray) == 0 {
		fmt.Println("requestArray is null!!!")
		return shim.Error("No requests!")
	}

	// ==== Query all responses ====
	queryResponseString := fmt.Sprintf("{\"selector\":{\"objectType\":\"response\",\"taskId\":\"%s\"}}", taskId)
	responseArray, responseErr = getArrayForResponse(stub, queryResponseString)
	if responseErr != nil {
		return shim.Error(responseErr.Error())
	}

	// ==== Check if responseArray is null ====
	if len(responseArray) == 0 {
		fmt.Println("task " + taskId + ". ResponseArray is null!!!")
		return shim.Error("task " + taskId + ". No responses!")
	}

	// ==== Select all satisfactory responses and write into satResponseArray ====
	for i = 0; i < len(requestArray); i++ {
		satResponseArray := make([]response, 0)
		var requestSignJSON requestSign
		req := requestArray[i]
		reqId := req.ReqId
		reqSignString := req.SignString
		unmarshalReqError := json.Unmarshal([]byte(reqSignString), &requestSignJSON)
		if unmarshalReqError != nil {
			return shim.Error("Failed to unmarshal the reqSign string")
		}

		for j = 0; j < len(responseArray); j++ {
			var responseSignJSON responseSign
			res := responseArray[j]
			resSignString := res.SignString
			unmarshalReqError := json.Unmarshal([]byte(resSignString), &responseSignJSON)
			if unmarshalReqError != nil {
				return shim.Error("Failed to unmarshal the reqSign string")
			}
			resId := responseSignJSON.ReqId

			condition1 := (reqId == resId)
			condition2 := math.Max(requestSignJSON.ResponseTime, responseSignJSON.ResponseTime) == requestSignJSON.ResponseTime && math.Abs(requestSignJSON.ResponseTime - responseSignJSON.ResponseTime) > MIN
			condition3 := math.Max(requestSignJSON.Throughput, responseSignJSON.Throughput) == responseSignJSON.Throughput && math.Abs(responseSignJSON.Throughput - requestSignJSON.Throughput) > MIN
			condition4 := math.Max(requestSignJSON.Budget, responseSignJSON.Price) == requestSignJSON.Budget && math.Abs(requestSignJSON.Budget - responseSignJSON.Price) > MIN

			if condition1 && condition2 && condition3 && condition4 {
				satResponseArray = append(satResponseArray, res)
			}
		}
		if len(satResponseArray) != 0 {
			satResponseMap[i + 1] = satResponseArray
		}
	}

	// ==== Check if satResponseMap is null ====
	if len(satResponseMap) <= 0 {
		state = changeStateToRejection(stub, taskId)
		fmt.Println("Current state is " + state + ", stop execute check!\n")
		fmt.Println("task " + taskId + ". No satisfactory response in satResponseMap!!!")
		return shim.Error("task " + taskId + ". No satisfactory responses!")
	}

	var finalRequest request
	var satResArray []response
	var requestSigR, requestSigS, requestCertificate string
	if len(satResponseMap) == 1 {
		var finalRequestSignJSON requestSign

		for key := range satResponseMap {
			finalRequest = requestArray[key - 1]
			requestSigR = finalRequest.SigR
			requestSigS = finalRequest.SigS
			requestCertificate = finalRequest.Certificate

			finalRequestSign := finalRequest.SignString
			unmarshalError := json.Unmarshal([]byte(finalRequestSign), &finalRequestSignJSON)
			if unmarshalError != nil {
				return shim.Error("Failed to unmarshal the requSign string")
			}
			reservePrice = finalRequestSignJSON.Budget

			satResArray = satResponseMap[key]
		}
	} else {
		// ==== Select max, min responseTime and max, min throughput ====
		for key := range satResponseMap {
			var requSignJSON requestSign
			requ = requestArray[key - 1]
			requSignString := requ.SignString
			unmarshalRequError := json.Unmarshal([]byte(requSignString), &requSignJSON)
			if unmarshalRequError != nil {
				return shim.Error("Failed to unmarshal the requSign string")
			}
			responseTimeSlice = append(responseTimeSlice, requSignJSON.ResponseTime)
			throughputSlice = append(throughputSlice, requSignJSON.Throughput)
		}
		var m, n int
		var minResponseTime, maxResponseTime, minThroughput, maxThroughput float64
		minResponseTime = responseTimeSlice[0]
		maxResponseTime = responseTimeSlice[0]
		for m = 1; m < len(responseTimeSlice); m ++ {
			if responseTimeSlice[m] > maxResponseTime {
				maxResponseTime = responseTimeSlice[m]
			} else if responseTimeSlice[m] <= minResponseTime {
				minResponseTime = responseTimeSlice[m]
			}
		}
		minThroughput = throughputSlice[0]
		maxThroughput = throughputSlice[0]
		for n = 1; n < len(throughputSlice); n ++ {
			if throughputSlice[n] > maxThroughput {
				maxThroughput = throughputSlice[n]
			} else if throughputSlice[n] <= minThroughput {
				minThroughput = throughputSlice[n]
			}
		}

		// ==== Range satResponseMap and select final request ====
		for key := range satResponseMap {
			var requSignJSON requestSign
			requ = requestArray[key - 1]
			requSignString := requ.SignString
			unmarshalRequError := json.Unmarshal([]byte(requSignString), &requSignJSON)
			if unmarshalRequError != nil {
				return shim.Error("Failed to unmarshal the requSign string")
			}
			reservePrice = requSignJSON.Budget
			//QoS := requSignJSON.ResponseTime + requSignJSON.Throughput
			QoS := float64((maxResponseTime - requSignJSON.ResponseTime) / (maxResponseTime - minResponseTime)) + float64((requSignJSON.Throughput - minThroughput) / (maxThroughput - minThroughput))
			if maxQoS < QoS {
				maxQoS = QoS
				maxQoSReqIndex = key
			}
		}
		satResArray = satResponseMap[maxQoSReqIndex]

		finalRequest = requestArray[maxQoSReqIndex - 1]
		requestSigR = finalRequest.SigR
		requestSigS = finalRequest.SigS
		requestCertificate = finalRequest.Certificate
	}

	// ==== Select final agreement price ====
	if len(satResArray) == 1 {
		minPriceResIndex = strconv.Itoa(0)
		aFinalPrice = reservePrice
	} else {
    var l int
    priceSlice := make([]float64, 0)

    for l = 0; l < len(satResArray); l ++ {
			var satResSignJSON responseSign
      satResSignString := satResArray[l].SignString
      unmarshalSatResError := json.Unmarshal([]byte(satResSignString), &satResSignJSON)
      if unmarshalSatResError != nil {
  			return shim.Error("Failed to unmarshal the satResSign string")
  		}
      priceSlice = append(priceSlice, satResSignJSON.Price)
    }

    if math.Max(priceSlice[0], priceSlice[1]) == priceSlice[1] && math.Abs(priceSlice[1] - priceSlice[0]) > MIN {
      lowestPrice = priceSlice[0]
      secondLowPrice = priceSlice[1]
			minPriceResIndex = strconv.Itoa(0)
    } else if math.Max(priceSlice[0], priceSlice[1]) == priceSlice[0] || math.Abs(priceSlice[0] - priceSlice[1]) < MIN {
			lowestPrice = priceSlice[1]
			secondLowPrice = priceSlice[0]
			minPriceResIndex = strconv.Itoa(1)
		}

		for k = 2; k < len(priceSlice); k ++ {
	    if math.Max(priceSlice[k], lowestPrice) == lowestPrice || math.Abs(lowestPrice - priceSlice[k]) < MIN {
	      secondLowPrice = lowestPrice
	      lowestPrice = priceSlice[k]
				minPriceResIndex = strconv.Itoa(k)
	    } else if (math.Max(priceSlice[k], lowestPrice) == priceSlice[k] || math.Abs(priceSlice[k] - lowestPrice) < MIN) && (math.Max(priceSlice[k], secondLowPrice) == secondLowPrice || math.Abs(secondLowPrice - priceSlice[k]) < MIN) {
	      secondLowPrice = priceSlice[k]
	    }
	  }
		aFinalPrice = secondLowPrice
	}

	minPriceResIndexInt, _ := strconv.Atoi(minPriceResIndex)
	resp := satResArray[minPriceResIndexInt]
	var respSignJSON responseSign
	respSignString := resp.SignString
	unmarshalRespError := json.Unmarshal([]byte(respSignString), &respSignJSON)
	if unmarshalRespError != nil {
		return shim.Error("Failed to unmarshal the respSign string")
	}

	aRequester := respSignJSON.Requester
	aProvider := respSignJSON.Provider
	aObjectType := "agreement"
	aResponseTime := respSignJSON.ResponseTime
	aThroughput := respSignJSON.Throughput
	aUrl := respSignJSON.Url
	aExpireTime, err := time.Parse("02 Jan 2006 15:04:05 -0700", respSignJSON.ExpireTime)
	if err != nil {
		return shim.Error("Failed to convert the aExpireTime argument into type time.Time")
	}
	responseSigR := resp.SigR
	responseSigS := resp.SigS
	responseCertificate := resp.Certificate

	// ==== Create CreateCompositeKey ====
	indexName := "agreement"
	indexKey, err := stub.CreateCompositeKey(indexName, []string{taskId, aRequester})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}
	stub.PutState(indexKey, value)  // Save index entry to state.

  beginTime := time.Now()

	agreementJSON := &agreement{aRequester, taskId, aProvider, aObjectType, aUrl, beginTime, aExpireTime, aResponseTime, aThroughput, aFinalPrice, requestSigR, requestSigS, requestCertificate, responseSigR, responseSigS, responseCertificate}
  fmt.Println(agreementJSON)
	agreementJSONAsBytes, err := json.Marshal(agreementJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	//==== 删除agreementMap原来的记录====/
	delete(agreementMap, taskId)

	// ==== Save agreement to state ====
	err = stub.PutState(indexKey, agreementJSONAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

  aFinalPriceStr := strconv.FormatFloat(aFinalPrice, 'f', 6, 64)

  currencyArgs:=[][]byte{[]byte("pendingPay"),[]byte(aRequester),[]byte(aFinalPriceStr),[]byte(taskId)}
	pendingResponse := stub.InvokeChaincode("currency", currencyArgs, "softwarechannel");

	if pendingResponse.Status != 200 {
		return pendingResponse;
	}

	state = changeStateToAcception(stub, taskId)

	fmt.Println("task" + taskId + ". Current state is:", state)
	fmt.Println("- end check\n")
	return shim.Success(agreementJSONAsBytes)
}
