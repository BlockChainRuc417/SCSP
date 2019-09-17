
package main


import (
	"fmt"
	"time"
	"strconv"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type TaskChaincode struct {
}


type task struct {
	ObjectType     string    `json:"objectType"`
  ID             string    `json:"id"`
  SignString     string    `json:"signString"`
	SigR           string    `json:"sigR"`
	SigS           string    `json:"sigS"`
	Certificate    string    `json:"certificate"`
}


type taskSign struct {
	TaskName       string    `json:"taskName"`
	Requester      string    `json:"requester"`
  Description    string    `json:"description"`
}


type request struct {
	ObjectType     string    `json:"objectType"`
	TaskId         string    `json:"taskId"`
	ReqId          string    `json:"reqId"`
	SignString     string    `json:"signString"`
	SigR           string    `json:"sigR"`
	SigS           string    `json:"sigS"`
	Certificate    string    `json:"certificate"`
}


type requestSign struct {
	Requester      string    `json:"requester"`
	ResponseTime   float64   `json:"responseTime"`
	Throughput     float64   `json:"throughput"`
	Budget         float64   `json:"budget"`
}


type response struct {
	ObjectType     string      `json:"objectType"`
	TaskId         string      `json:"taskId"`
	SignString     string      `json:"signString"`
	SigR           string      `json:"sigR"`
	SigS           string      `json:"sigS"`
	Certificate    string      `json:"certificate"`
}


type responseSign struct {
	ReqId          string      `json:"reqId"`
	Requester      string      `json:"requester"`
	Provider       string      `json:"provider"`
	Url            string      `json:"url"`
	ExpireTime     string      `json:"expireTime"`
	ResponseTime   float64     `json:"responseTime"`
	Throughput     float64     `json:"throughput"`
	Price          float64     `json:"price"`
}


const MIN = 0.001
var iw, _ = NewIdWorker(1)


// ==============
//      Main
// ==============
func main() {
	err := shim.Start(new(TaskChaincode))
	if err != nil {
		fmt.Printf("Error starting Task chaincode: %s", err)
	}
}


// =========================================
//       Init - Initializes chaincode
// =========================================
func (t *TaskChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}


// ======================================================
//       Invoke - Our entry point for Invocations
// ======================================================
func (t *TaskChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	if function == "addTask" {
		return t.addTask(stub, args)
	} else if function == "queryTask" {
    return t.queryTask(stub)
  } else if function == "queryTaskByName" {
		return t.queryTaskByName(stub, args)
	} else if function == "queryTaskByNameAndRequester" {
		return t.queryTaskByNameAndRequester(stub, args)
	} else if function == "queryTaskByRequester" {
		return t.queryTaskByRequester(stub, args)
	} else if function == "queryTaskByDescription" {
		return t.queryTaskByDescription(stub, args)
	} else if function == "queryTaskById" {
		return t.queryTaskById(stub, args)
	} else if function == "queryStateByTaskId" {
		return t.queryStateByTaskId(stub, args)
	} else if function == "writeRequest" {
		return t.writeRequest(stub, args)
  } else if function == "readRequest" {
	 	return t.readRequest(stub, args)
 	} else if function == "writeResponse" {
	 	return t.writeResponse(stub, args)
 	} else if function == "readResponse" {
	 	return t.readResponse(stub, args)
 	} else if function == "queryByObjectType" {
    return t.queryByObjectType(stub, args)
  } else if function == "drop" {
		return t.drop(stub, args)
	} else if function == "delete" {
		return t.delete(stub, args)
	} else if function == "deleteTask" {
		return t.deleteTask(stub, args)
	} else if function == "deleteServiceTX" {
		return t.deleteServiceTX(stub)
	} else if function == "check" {
		return t.check(stub, args)
	} else if function == "changeStateToInstantiation" {
		return t.changeStateToInstantiation(stub, args)
	} else if function == "readState" {
		return t.readState(stub, args)
	} else if function == "queryServiceTx" {
		return t.queryServiceTx(stub, args)
	} else if function == "queryServiceTxByTaskId" {
		return t.queryServiceTxByTaskId(stub, args)
	} else if function == "queryServiceTxByDate" {
		return t.queryServiceTxByDate(stub, args)
	} else if function == "invokeRestAPI" {
		return t.invokeRestAPI(stub, args)
	} else if function == "saveServiceTX" {
		return t.saveServiceTX(stub, args)
	} else {
		return shim.Error("Function " + function + "  doesn't exits, make sure function is right!")
	}
}


// ====================================================================
//      generateUniqueId - Generate unique id for task and request
// ====================================================================
func (t *TaskChaincode) generateUniqueId(stub shim.ChaincodeStubInterface) pb.Response {
	var idStr string

  id, err := iw.NextId()
	if err != nil {
		fmt.Println(err)
	} else {
		idStr = strconv.FormatInt(id, 10)
	}

	return shim.Success([]byte(idStr))
}


// ==================================================
//      addTask - Add tasks to chaincode state
// ==================================================
func (t *TaskChaincode) addTask(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//        0         1       2       3
	// "signString", "sigR", "sigS", "cert"
  if len(args) != 4 {
    return shim.Error("Incorrect number of arguments. Expecting 4")
  }

  fmt.Println("- start add task")
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

	// ==== Generate unique id for task ====
	var idStr string
  id, err := iw.NextId()
	if err != nil {
		fmt.Println(err)
	} else {
		idStr = strconv.FormatInt(id, 10)
	}

	objectType := "Task"
	taskId := idStr
	signString := args[0]
	sigR := args[1]
  sigS := args[2]
	cert := args[3]

  state := "instantiation"

  // ==== Save task state ====
  err = stub.PutState(taskId, []byte(state))
  if err != nil {
    return shim.Error(err.Error())
  }

  fmt.Println("Current State is:", state)

	// ==== Verify signature ====
	publicKey := getPublicKey(cert)
	verifyResult := verifySignature(publicKey, sigR, sigS, signString)
	if verifyResult == false {
		return shim.Error("Verify Failed!")
	}

	// ==== Create task compositekey ====
	indexName := "task"
	indexKey, err := stub.CreateCompositeKey(indexName, []string{taskId})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}

	// ==== Check if task already exists ====
	taskAsBytes, err := stub.GetState(indexKey)
	if err != nil {
		return shim.Error("Failed to get task: " + err.Error())
	} else if taskAsBytes != nil {
		fmt.Println("The task already exists")
		return shim.Error("The task already exists")
	}
	stub.PutState(indexKey, value)    // Save index entry to state.

	// ==== Create task object and marshal to JSON ====
  addTaskJSON := &task{objectType, taskId, signString, sigR, sigS, cert}
  taskJSONasBytes, err := json.Marshal(addTaskJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Save task to state ====
  err = stub.PutState(indexKey, taskJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

  fmt.Println("add task successfully", taskId)
  fmt.Println("- end add task\n")
	return shim.Success(taskJSONasBytes)
}


// ==============================================================
//       writeRequest - Write request from chaincode state
// ==============================================================
func (t *TaskChaincode) writeRequest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var requestSignJSON requestSign

	//      0           1          2       3       4
	// "$taskId", "signString", "sigR", "sigS", "cert"
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	fmt.Println("- start write request")
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

	// ==== Generate unique id for request ====
	var idStr string
  id, err := iw.NextId()
	if err != nil {
		fmt.Println(err)
	} else {
		idStr = strconv.FormatInt(id, 10)
	}

	objectType := "request"
	taskId := args[0]
	reqId := idStr
	signString := args[1]
	unmarshalError := json.Unmarshal([]byte(signString), &requestSignJSON)
	if unmarshalError != nil {
		return shim.Error("Failed to unmarshal the requestSign string")
	}
	requester := requestSignJSON.Requester
	sigR := args[2]
	sigS := args[3]
	cert := args[4]

	// ==== Get current state of task ====
	stateAsByte, err := stub.GetState(taskId)
	if err != nil {
		fmt.Println("Failed to read state!")
	}
	state := string(stateAsByte)

	if state != "instantiation" && state != "rejection" {
		fmt.Println("Current state is " + state + ", stop writeRequest! You can write request only state is instantiation or rejection!\n")
		return shim.Error("Can't write request. Current state is " + state + ", you can write request only state is instantiation or rejection!")
	}

	fmt.Println("Current state is " + state + ", continue......")

	// ==== Verify signature ====
	publicKey := getPublicKey(cert)
	verifyResult := verifySignature(publicKey, sigR, sigS, signString)
	if verifyResult == false {
		return shim.Error("Verify Failed!")
	}

	// ==== Create task compositekey ====
	indexTaskName := "task"
	indexTaskKey, err := stub.CreateCompositeKey(indexTaskName, []string{taskId})

	if err != nil {
		return shim.Error(err.Error())
	}

	taskAsBytes, err := stub.GetState(indexTaskKey)
	if err != nil {
		return shim.Error("Failed to get task: " + err.Error())
	} else if taskAsBytes == nil {
		return shim.Error("The task doesn't exist")
	} else {
    var taskJSON task
		err = json.Unmarshal([]byte(taskAsBytes), &taskJSON)
		if err != nil {
			return shim.Error("Failed to unmarshal the taskAsBytes")
		} else {
			var taskSignJSON taskSign
			taskSignString := taskJSON.SignString
			err = json.Unmarshal([]byte(taskSignString), &taskSignJSON)
			if err != nil {
				return shim.Error("Failed to unmarshal the signature string")
			} else if taskSignJSON.Requester != requester {
 				return shim.Error("The task is not published by the requester " +requester);
 			}
		}
	}

	// ==== Create request compositekey ====
	indexRequestName := "request"
	indexRequestKey, err := stub.CreateCompositeKey(indexRequestName, []string{reqId, requester})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}

	// ==== Check if request already exists ====
	requestAsBytes, err := stub.GetState(indexRequestKey)
	if err != nil {
		return shim.Error("Failed to get request: " + err.Error())
	} else if requestAsBytes != nil {
		fmt.Println("The request already exists")
		return shim.Error("The request already exists")
	}
	stub.PutState(indexRequestKey, value)    // Save index entry to state.

	// ==== Create request object and marshal to JSON ====
	requestJSON := &request{objectType, taskId, reqId, signString, sigR, sigS, cert}
	requestJSONasBytes, err := json.Marshal(requestJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Save request to state ====
	err = stub.PutState(indexRequestKey, requestJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("write request successfully", requester)
	fmt.Println("- end write request\n")
	return shim.Success(requestJSONasBytes)
}


// =============================================================
//       readRequest - Read request from chaincode state
// =============================================================
func (t *TaskChaincode) readRequest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string

	//       0         1       2
	// "$requestId", "u1", "$taskId"
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	fmt.Println("- start read request")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	reqId := args[0]
	requester := args[1]
	taskId := args[2]

	// ==== Get current state of task ====
	stateAsByte, err := stub.GetState(taskId)
	if err != nil {
		fmt.Println("Failed to read state!")
	}
	state := string(stateAsByte)

	if state != "instantiation" {
		fmt.Println("Current state is " + state + ", stop read! You can only read request only state is instantiation!\n")
		return shim.Error("Can't read request. Current state is " + state + ", you can read request only state is instantiation!")
	}

	fmt.Println("Current state is " + state + ", continue......")

	// ==== Create request compositekey ====
	indexName := "request"
	indexKey, err := stub.CreateCompositeKey(indexName, []string{reqId, requester})
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Check if request exists ====
	requestAsbytes, err := stub.GetState(indexKey)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for request\"}"
		return shim.Error(jsonResp)
	} else if requestAsbytes == nil {
		jsonResp = "{\"Error\":\"Agenttask does not exist\"}"
		return shim.Error(jsonResp)
	}

	fmt.Println("readRequestResults:", string(requestAsbytes))
	fmt.Println("- end read request\n")
	return shim.Success(requestAsbytes)
}


// ================================================================
//       writeResponse - Write response from chaincode state
// ================================================================
func (t *TaskChaincode) writeResponse(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var responseSignJSON responseSign

	//     0            1          2       3       4
	// "$taskId", "signString", "sigR", "sigS", "cert"
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	fmt.Println("- start write response")
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

	objectType := "response"
	taskId := args[0]
	signString := args[1]
	unmarshalError := json.Unmarshal([]byte(signString), &responseSignJSON)
	if unmarshalError != nil {
		return shim.Error("Failed to unmarshal the responseSign string")
	}
	reqId := responseSignJSON.ReqId
	requester := responseSignJSON.Requester
	provider := responseSignJSON.Provider
	expireTime, err := time.Parse("02 Jan 2006 15:04:05 -0700", responseSignJSON.ExpireTime)
	if err != nil {
		return shim.Error("Failed to convert the expireTime argument into type time.Time")
	}
	nowTime := time.Now()
	if expireTime.Sub(nowTime).Seconds() <= 0.0 {
		return shim.Error("Expire time must be after current time!")
	}
	sigR := args[2]
	sigS := args[3]
	cert := args[4]

	// ==== Get current state of task ====
	stateAsByte, err := stub.GetState(taskId)
	if err != nil {
		fmt.Println("Failed to read state for task " + taskId)
		return shim.Error("Failed to read state for task " + taskId)
	}
	state := string(stateAsByte)

	if state != "instantiation" {
		fmt.Println("Current state is " + state + ", stop writeResponse! You can only write response only state is instantiation!\n")
		return shim.Error("Can't to write response. Current state is " + state + ", you can write response only state is instantiation!")
	}

	fmt.Println("Current state is " + state + ", continue......")

	// ==== Verify signature ====
	publicKey := getPublicKey(cert)
	verifyResult := verifySignature(publicKey, sigR, sigS, signString)
	if verifyResult == false {
		return shim.Error("Verify Failed!")
	}

	// ==== Create request compositekey ====
	indexRequestName := "request"
	indexRequestKey, err := stub.CreateCompositeKey(indexRequestName, []string{reqId, requester})
	if err != nil {
		return shim.Error(err.Error())
	}

	requestAsBytes, err := stub.GetState(indexRequestKey)
	if err != nil {
		return shim.Error("Failed to get request: " + err.Error())
	} else if requestAsBytes == nil {
		return shim.Error("The request " + reqId + " & requester " + requester + " doesn't exist")
	}

	// ==== Create response compositeKey ====
	indexResponseName := "response"
	indexResponseKey, err := stub.CreateCompositeKey(indexResponseName, []string{reqId, provider})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}

	// ==== Check if response already exists ====
	provideAsBytes, err := stub.GetState(indexResponseKey)
	if err != nil {
		return shim.Error("Failed to get response: " + err.Error())
	} else if provideAsBytes != nil {
		fmt.Println("You have already wrote response " + string(provideAsBytes) + " to this request, please choose another request!")
		return shim.Error("You have already wrote response"  + string(provideAsBytes) + " to this request, please choose another request!")
	}
	stub.PutState(indexResponseKey, value)    // Save index entry to state.

	// ==== Create response object and marshal to JSON ====
	provideJSON := &response{objectType, taskId, signString, sigR, sigS, cert}
	provideJSONasBytes, err := json.Marshal(provideJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Save response to state ====
	err = stub.PutState(indexResponseKey, provideJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("write response successfully", provider)
	fmt.Println("- end write response\n")
	return shim.Success(provideJSONasBytes)
}


// =============================================================
//      readResponse - Read response from chaincode state
// =============================================================
func (t *TaskChaincode) readResponse(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string

	//       0         1       2
	// "$requestId", "s1", "$taskId"
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	fmt.Println("- start read response")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	reqId := args[0]
	provider := args[1]
	taskId := args[2]

	// ==== Get current state of task ====
	stateAsByte, err := stub.GetState(taskId)
	if err != nil {
		shim.Error("Failed to read state of task " + taskId);
	}
	state := string(stateAsByte)

	if state != "instantiation" {
		fmt.Println("Current state is " + state + ", stop readResponse! You can only read response only state is instantiation!")
		return shim.Error("Can't to read response. Current state is " + state + ", you can read response only state is instantiation!")
	}

	fmt.Println("Current state is " + state + ", continue......")

	// ==== Create response compositeKey ====
	indexName := "response"
	indexKey, err := stub.CreateCompositeKey(indexName, []string{reqId, provider})
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Check if response exists ====
	responseAsbytes, err := stub.GetState(indexKey)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for response\"}"
		return shim.Error(jsonResp)
	} else if responseAsbytes == nil {
		jsonResp = "{\"Error\":\"(Response does not exist\"}"
		return shim.Error(jsonResp)
	}

	fmt.Println("readResponseResults:",string(responseAsbytes))
	fmt.Println("- end read response\n")
	return shim.Success(responseAsbytes)
}


// =========================================================================
//        drop - Remove all requests and all responses ignore state
// =========================================================================
func (t *TaskChaincode) drop(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	requestArray := make([]request, 0)
	responseArray := make([]response, 0)
	var i, j int
	var requestErr, responseErr error

	//     0
	// "$taskId"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	taskId := args[0]

	// ==== Query all requests ====
	queryRequestString := fmt.Sprintf("{\"selector\":{\"objectType\":\"request\",\"taskId\":\"%s\"}}", taskId)
	requestArray, requestErr = getArrayForRequest(stub, queryRequestString)
	if requestErr != nil {
		return shim.Error(requestErr.Error())
	}

	for i = 0; i < len(requestArray); i ++ {
		var reqSignJSON requestSign
		req := requestArray[i]
		reqId := req.ReqId
		reqSignString := req.SignString
		unmarshalReqError := json.Unmarshal([]byte(reqSignString), &reqSignJSON)
		if unmarshalReqError != nil {
			return shim.Error("Failed to unmarshal the reqSign string")
		}
		requester := reqSignJSON.Requester

		indexName := "request"
		requestIndexKey, err := stub.CreateCompositeKey(indexName, []string{reqId, requester})
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.DelState(requestIndexKey)
		if err != nil {
			return shim.Error("Failed to delete state:" + err.Error())
		}
	}

	// ==== Query all responses ====
	queryResponseString := fmt.Sprintf("{\"selector\":{\"objectType\":\"response\",\"taskId\":\"%s\"}}", taskId)
	responseArray, responseErr = getArrayForResponse(stub, queryResponseString)
	if responseErr != nil {
		return shim.Error(responseErr.Error())
	}

	for j = 0; j < len(responseArray); j ++ {
		var resSignJSON responseSign
		res := responseArray[j]
		resSignString := res.SignString
		unmarshalResError := json.Unmarshal([]byte(resSignString), &resSignJSON)
		if unmarshalResError != nil {
			return shim.Error("Failed to unmarshal the resSign string")
		}
		reqId := resSignJSON.ReqId
		provider := resSignJSON.Provider

		indexName := "response"
		responseIndexKey, err := stub.CreateCompositeKey(indexName, []string{reqId, provider})
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.DelState(responseIndexKey)
		if err != nil {
			return shim.Error("Failed to delete state:" + err.Error())
		}
	}

	// ==== Query agreement ====
	queryAgreementString := fmt.Sprintf("{\"selector\":{\"objectType\":\"agreement\",\"taskId\":\"%s\"}}", taskId)
	agreementByte, agreementErr := getJSONForAgreement(stub, queryAgreementString)
	if agreementErr != nil {
		return shim.Error(agreementErr.Error())
	}

	requester := agreementByte.Requester

	indexName := "agreement"
	agreementIndexKey, err := stub.CreateCompositeKey(indexName, []string{taskId, requester})
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.DelState(agreementIndexKey)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	return shim.Success([]byte(taskId))
}


// =====================================================================================
//        delete - Remove all requests and all responses when state is rejection
// =====================================================================================
func (t *TaskChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	requestArray := make([]request, 0)
	responseArray := make([]response, 0)
	var i, j int
	var requestErr, responseErr error

	//     0
	// "$taskId"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	taskId := args[0]

	// ==== get current state of task ====
	stateAsByte, err := stub.GetState(args[0])
	if err != nil {
		fmt.Println("Failed to read state!")
	}
	state := string(stateAsByte)

	if state != "rejection" {
		fmt.Println("Failed to delete!")
		return shim.Error("Failed to delete!")
	}

	// ==== Query all requests ====
	queryRequestString := fmt.Sprintf("{\"selector\":{\"objectType\":\"request\",\"taskId\":\"%s\"}}", taskId)
	requestArray, requestErr = getArrayForRequest(stub, queryRequestString)
	if requestErr != nil {
		return shim.Error(requestErr.Error())
	}

	for i = 0; i < len(requestArray); i ++ {
		var reqSignJSON requestSign
		req := requestArray[i]
		reqId := req.ReqId
		reqSignString := req.SignString
		unmarshalReqError := json.Unmarshal([]byte(reqSignString), &reqSignJSON)
		if unmarshalReqError != nil {
			return shim.Error("Failed to unmarshal the reqSign string")
		}
		requester := reqSignJSON.Requester

		indexName := "request"
		requestIndexKey, err := stub.CreateCompositeKey(indexName, []string{reqId, requester})
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.DelState(requestIndexKey)
		if err != nil {
			return shim.Error("Failed to delete state:" + err.Error())
		}
	}

	// ==== Query all responses ====
	queryResponseString := fmt.Sprintf("{\"selector\":{\"objectType\":\"response\",\"taskId\":\"%s\"}}", taskId)
	responseArray, responseErr = getArrayForResponse(stub, queryResponseString)
	if responseErr != nil {
		return shim.Error(responseErr.Error())
	}

	for j = 0; j < len(responseArray); j ++ {
		var resSignJSON responseSign
		res := responseArray[j]
		resSignString := res.SignString
		unmarshalResError := json.Unmarshal([]byte(resSignString), &resSignJSON)
		if unmarshalResError != nil {
			return shim.Error("Failed to unmarshal the resSign string")
		}
		reqId := resSignJSON.ReqId
		provider := resSignJSON.Provider

		indexName := "response"
		responseIndexKey, err := stub.CreateCompositeKey(indexName, []string{reqId, provider})
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.DelState(responseIndexKey)
		if err != nil {
			return shim.Error("Failed to delete state:" + err.Error())
		}
	}

	return shim.Success([]byte(taskId))
}


// =============================================
//        delete task - Remove task by id
// =============================================
func (t *TaskChaincode) deleteTask(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//     0
	// "$taskId"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	taskId := args[0]

	// ==== Create task compositekey ====
	indexName := "task"
	indexKey, err := stub.CreateCompositeKey(indexName, []string{taskId})

	if err != nil {
		return shim.Error(err.Error())
	}

	taskAsBytes, err := stub.GetState(indexKey)
	if err != nil {
		return shim.Error("Failed to get task: " + err.Error())
	} else if taskAsBytes == nil {
		return shim.Error("no task " + taskId);
	}

  // ==== Delete request, response and agreement ====
	t.drop(stub, args);

  // ==== Delete task ====
	err = stub.DelState(indexKey)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	// ==== Delete state of task ====
	err = stub.DelState(taskId)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	return shim.Success([]byte(taskId))
}


// ============================================================
//      queryStateByTaskId - Query task state from chaincode
// ============================================================
func (t *TaskChaincode) queryStateByTaskId(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//     0
	// "$taskId"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	taskId := args[0]

	// ==== Get current state of task ====
	stateAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to read state of task " + taskId);
	}

	return shim.Success(stateAsBytes)
}


// ============================================================
//      queryTask - Query all tasks from chaincode state
// ============================================================
func (t *TaskChaincode) queryTask(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("- start query all tasks")

	queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"Task\"}}")
	queryResults, err := getResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end query all tasks\n")
	return shim.Success(queryResults)
}


// ==============================================================================
//      queryTaskByName - Query all tasks from chaincode state by taskName
// ==============================================================================
func (t *TaskChaincode) queryTaskByName(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//        0
	// "ticket-airline"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	fmt.Println("- start query task by name")

	taskName := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"taskName\":{\"$regex\":\"(?i)%s\"}}}", taskName)
	queryResults, err := getResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end query task by name\n")
	return shim.Success(queryResults)
}


// ====================================================================================
//      queryTaskByRequester - Query all tasks from chaincode state by requester
// ====================================================================================
func (t *TaskChaincode) queryTaskByRequester(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "Jim"
	fmt.Println("- start query all tasks by requester")
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	requester := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"Task\",\"requester\":\"%s\"}}", requester)
	queryResults, err := getResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end query all tasks by requester\n")
	return shim.Success(queryResults)
}


// ==========================================================================================================
//      queryTaskByNameAndRequester - Query all tasks from chaincode state by requester and task name
// ==========================================================================================================
func (t *TaskChaincode) queryTaskByNameAndRequester(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//    0          1
	// "Jim", "ticket-airline"
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	requester := args[0]
	taskName := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"Task\",\"signString\":{\"$regex\":\"(?i)%s.*(?i)%s\"}}}", taskName, requester)
	queryResults, err := getResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end query all tasks by requester\n")
	return shim.Success(queryResults)
}


// ========================================================================================
//      queryTaskByDescription - Query all tasks from chaincode state by description
// ========================================================================================
func (t *TaskChaincode) queryTaskByDescription(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "000"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	fmt.Println("- start query task by description")

	description := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"description\":{\"$regex\":\"(?i)%s\"}}}", description)
	queryResults, err := getResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end query task by description\n")
	return shim.Success(queryResults)
}


// =========================================================================
//       queryTaskById - Query the task from chaincode state by taskId
// =========================================================================
func (t *TaskChaincode) queryTaskById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//     0
	// "$taskId"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	fmt.Println("- start query task by ID")

  taskId := args[0];

	// ==== Create task compositekey ====
	indexName := "task"
	indexKey, err := stub.CreateCompositeKey(indexName, []string{taskId})

	if err != nil {
		return shim.Error(err.Error())
	}

	taskAsBytes, err := stub.GetState(indexKey)
	if err != nil {
		return shim.Error("Failed to get task: " + err.Error())
	} else{
		fmt.Println("- end query task by ID\n")
		return shim.Success(taskAsBytes)
	}
}


// ==========================================================================================================
//      queryByObjectType - Query all requests or all responses or all agreements from chaincode state
// ==========================================================================================================
func (t *TaskChaincode) queryByObjectType(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//     0          1
	// "$taskId", "request"
	if len(args) !=2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	taskId := args[0]
	objectType := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"%s\",\"taskId\":\"%s\"}}", objectType, taskId)
	queryResults, err := getResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}
