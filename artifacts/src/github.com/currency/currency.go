
package main


import (
	"fmt"
  "encoding/json"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type TaskChaincode struct {
}


type payTX struct {
	ObjectType     string    `json:"objectType"`
	TxTime				 time.Time `json:"txTime"`
	TxID           string    `json:"txID"`
	TaskId         string    `json:"taskId"`
  Payer          string    `json:"payer"`
  Payee          string    `json:"payee"`
	Value          float64   `json:"value"`
}

type agreement struct {
	Requester      string      `json:"requester"`
	TaskId         string      `json:"taskId"`
	Provider       string      `json:"provider"`
	ObjectType     string      `json:"objectType"`    //objectType is used to distinguish the various types of objects in state database
	Url            string      `json:"url"`
	BeginTime      time.Time   `json:"beginTime"`
	ExpireTime     time.Time   `json:"expireTime"`
	ResponseTime   float64     `json:"responseTime"`
	Throughput     float64     `json:"throughput"`
	FinalPrice     float64     `json:"finalPrice"`
}

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
//       Init - initializes chaincode
// =========================================
func (t *TaskChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
  //初始化snowflake的IdWorker
	//iw, _ = NewIdWorker(1)
	//if err!= nil {
	//	fmt.Println(err) UTC
	//}
	return shim.Success(nil)
}


// ======================================================
//       Invoke - Our entry point for Invocations
// ======================================================
func (t *TaskChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	if function == "regist" {
		return t.regist(stub, args)
	} else if function == "pay" {
		return t.pay(stub, args)
	} else if function == "pendingPay" {
		return t.pendingPay(stub, args)
  } else if function == "confirmPay" {
		return t.confirmPay(stub, args)
  } else if function == "getBalance" {
		return t.getBalance(stub, args)
	} else if function == "queryPayTxByTaskId" {
		return t.queryPayTxByTaskId(stub, args)
	} else if function == "queryPayTxByPayer" {
		return t.queryPayTxByPayer(stub, args)
	} else if function == "queryPayTxByPayee" {
		return t.queryPayTxByPayee(stub, args)
	} else if function == "queryMembers" {
		return t.queryMembers(stub)
	} else {
		return shim.Error("Function " + function + " doesn't exits, make sure function is right!")
	}
}

// ===========================================================================
//      Transaction makes payment of X units from requester to provider
// ===========================================================================
func (t *TaskChaincode) pay(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//        0                 1             2          3
	// "$payer account", "$payee acount", "$taskId",  "$value"
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}


	fmt.Println("Payment begins!")
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
		return shim.Error("3rd argument must be a non-empty string")
	}

	payer := args[0]
	payee := args[1]
	taskId := args[2]
	valueStr := args[3]

	value, err := strconv.ParseFloat(valueStr, 64);
  if err != nil {
		return shim.Error("Failed to decode the value " + valueStr)
	}

	// ==== Check if payer exists ====
	payerAsBytes, err := stub.GetState(payer)
	if err != nil {
		return shim.Error("Failed to get the payer " + payer)
	}
	if payerAsBytes == nil {
		return shim.Error("Payer " + payer + " not found")
	}

	//balanceResponse := t.balance(stub, []string{requester})
  //balance, err := strconv.ParseFloat(string(balanceResponse.Payload), 64)
	//if err != nil {
	//	return shim.Error("Failed to get the balance")
	//}
  //fmt.Println(balance)

	// ==== Check if payee exists ====
  payeeAsBytes, err := stub.GetState(payee)
	if err != nil {
		return shim.Error("Failed to get the payee " + payee)
	}
	if payeeAsBytes == nil {
		return shim.Error("payee " + payee + " not found")
	}

	var idStr string
  id, err := iw.NextId()
	if err != nil {
		fmt.Println(err)
	} else {
		idStr = strconv.FormatInt(id, 10)
	}

  txTime := time.Now()
	//txTimeStr := time.Now().Format("02 Jan 2006 15:04:05 -0700")
	//txTime, err := time.Parse("02 Jan 2006 15:04:05 -0700", txTimeStr)
	//if err != nil {
	//	return shim.Error("Failed to convert the txTime into type time.Time")
	//}

  txJSON := &payTX{"PayTX", txTime, idStr, taskId, payer, payee, value}
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
	fmt.Println("Payment ends!")
	return shim.Success(txJSONasBytes)
}

// ===========================================================================
//      pending pay
// ===========================================================================
func (t *TaskChaincode) pendingPay(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//       0                 1
	// "$payer account", "$value"
	/*if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	fmt.Println("PendingPay starts!")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

  payer := args[0]
	value, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return shim.Error("2nd argument must be a numeric string")
	}*/

	//   0      1      2
	// payer, value, taskId
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	//fmt.Println("PendingPay starts!")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	//payer := agree.Requester
	//value := agree.FinalPrice
	//taskId := agree.TaskId

	payer := args[0]
	valueStr := args[1]
  value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return shim.Error(err.Error());
	}

	taskId := args[2]

	// ==== Check if payer exists ====
	payerAsBytes, err := stub.GetState(payer)
	if err != nil {
		return shim.Error("Failed to get the payer " + payer)
	}
	if payerAsBytes == nil {
		//fmt.Println(shim.Error("Payer not found"))
		return shim.Error("Payer " + payer + " not found")
	}

	var idStr string
  id, err := iw.NextId()
	if err != nil {
		fmt.Println(err)
	} else {
		idStr = strconv.FormatInt(id, 10)
	}

  txTime := time.Now()
	//txTimeStr := time.Now().Format("02 Jan 2006 15:04:05 -0700")
	//txTime, err := time.Parse("02 Jan 2006 15:04:05 -0700", txTimeStr)
	//if err != nil {
	//	return shim.Error("Failed to convert the txTime into type time.Time")
	//}

  payee := "contract";
  txJSON := &payTX{"PayTX", txTime, idStr, taskId, payer, payee, value}
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
	fmt.Println("PendingPay ends!")
	return shim.Success(txJSONasBytes)
}

// ===========================================================================
//      pending pay
// ===========================================================================
func (t *TaskChaincode) confirmPay(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//       0
	// "$taskId",
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	/*fmt.Println("ConfirmPay starts!")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	payee := args[0]
	value, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return shim.Error("2nd argument must be a numeric string")
	}*/

	taskId := args[0];

	objectType := "agreement"

	/*queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"%s\",\"taskId\":\"%s\"}}", objectType, taskId)
	queryResults, err := getResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}*/

	scasArgs := [][]byte{[]byte("queryByObjectType"),[]byte(taskId),[]byte(objectType)}
	agreementResponse := stub.InvokeChaincode("task", scasArgs, "softwarechannel");

	if agreementResponse.Status != 200 {
		return agreementResponse;
	}
	queryResults := agreementResponse.Payload;


	var agreements []agreement
	err := json.Unmarshal(queryResults, &agreements)
	if err != nil {
		return shim.Error(err.Error())
	}

	if len(agreements) <= 0 {
		return shim.Error("no agreement for the task " + taskId)
	}

	agreement := agreements[0]
	//fmt.Println(agreement.Provider)
	payee := agreement.Provider
	value := agreement.FinalPrice
	expireTime := agreement.ExpireTime

	// ==== Check if payee exists ====
  payeeAsBytes, err := stub.GetState(payee)
	if err != nil {
		return shim.Error("Failed to get the payee " + payee)
	}
	if payeeAsBytes == nil {
		return shim.Error("payee " + payee + " not found")
	}

	var idStr string
  id, err := iw.NextId()
	if err != nil {
		fmt.Println(err)
	} else {
		idStr = strconv.FormatInt(id, 10)
	}

  txTime := time.Now()

  //允许提前几秒确定支付
  threshold := 5.0
	//离expireTime还有多久
	//timeDiff := expireTime.Sub(txTime).Seconds

  if (expireTime.Sub(txTime).Seconds()) >= threshold {
		txTimeString := txTime.Format("02 Jan 2006 15:04:05 -0700")
		expireTimeString := expireTime.Format("02 Jan 2006 15:04:05 -0700")
		return shim.Error("cannot confirm pay! Now is " + txTimeString + ", but the expireTime is " + expireTimeString)
	}

	//txTimeStr := time.Now().Format("02 Jan 2006 15:04:05 -0700")
	//txTime, err := time.Parse("02 Jan 2006 15:04:05 -0700", txTimeStr)
	//if err != nil {
	//	return shim.Error("Failed to convert the txTime into type time.Time")
	//}

  payer := "contract";
  txJSON := &payTX{"PayTX", txTime, idStr, taskId, payer, payee, value}
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
	fmt.Println("confirmPay ends!")
	return shim.Success(txJSONasBytes)
}

// ===========================================================================
//      calculate balance
// ===========================================================================
func (t *TaskChaincode) getBalance(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//       0
	// "$account"
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("cacluate begins!");
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	account := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"PayTX\",\"payer\":\"%s\"}}", account)
	queryResults, err := getResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	var payerTXs []payTX
	err = json.Unmarshal(queryResults, &payerTXs)
	if err != nil {
		shim.Error(err.Error())
	}

	//fmt.Println(len(payTXs))
	var i int
	outcomeVal := 0.0
  for i=0;i<len(payerTXs);i=i+1 {
		payerTX := payerTXs[i]
		outcomeVal = outcomeVal + payerTX.Value
	}
  //fmt.Println(outcomeVal)

	queryString = fmt.Sprintf("{\"selector\":{\"objectType\":\"PayTX\",\"payee\":\"%s\"}}", account)
	queryResults, err = getResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	var payeeTXs []payTX
	err = json.Unmarshal(queryResults, &payeeTXs)
	if err != nil {
		shim.Error(err.Error())
	}

	incomeVal := 0.0
  for i=0;i<len(payeeTXs);i=i+1 {
		payeeTX := payeeTXs[i]
		incomeVal = incomeVal + payeeTX.Value
	}
  //fmt.Println(incomeVal)

	balance := incomeVal - outcomeVal
	//fmt.Println(balance)
  balanceStr := strconv.FormatFloat(balance, 'f', 6, 64)

  return shim.Success([]byte(balanceStr))
}

// ===========================================================================
//      get payTX by TaskId
// ===========================================================================
func (t *TaskChaincode) queryPayTxByTaskId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//       0
	// "$taskId"
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}


	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	taskId := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"PayTX\",\"taskId\":\"%s\"}}", taskId)
	queryResults, err := getResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

// ===========================================================================
//      get payTX by payer
// ===========================================================================
func (t *TaskChaincode) queryPayTxByPayer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//       0
	// "$payer"
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}


	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	payer := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"PayTX\",\"payer\":\"%s\"}}", payer)
	queryResults, err := getResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

// ===========================================================================
//      get payTX by payee
// ===========================================================================
func (t *TaskChaincode) queryPayTxByPayee(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//       0
	// "$payee"
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}


	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	payee := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"PayTX\",\"payee\":\"%s\"}}", payee)
	queryResults, err := getResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}
