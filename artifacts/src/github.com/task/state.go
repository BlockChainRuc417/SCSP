
package main


import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


// ===============================================================================================
//      changeStateToInstantiation - Change task's state from instantiation to instantiation
// ===============================================================================================
func (t *TaskChaincode) changeStateToInstantiation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var stateAsByte []byte

	//     0
	// "$taskId"
	fmt.Println("- start changeStateToInstantiation")
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	taskId := args[0]

	stateAsByte, err = stub.GetState(taskId)
	if err != nil {
		return shim.Error("Failed to get state")
	}

	state := string(stateAsByte)
	state = "instantiation"

	err = stub.PutState(string(taskId), []byte(state))
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end changeStateToInstantiation\n")
	return shim.Success(nil)
}


// =======================================================================================
//      changeStateToRejection - Change task's state from instantiation to rejection
// =======================================================================================
func changeStateToRejection(stub shim.ChaincodeStubInterface, taskId string) (string) {
	var err error
	var stateAsByte []byte

	stateAsByte, err = stub.GetState(taskId)
	if err != nil {
		return "0"
	}

	state := string(stateAsByte)
	state = "rejection"

	// ==== save task state ====
	err = stub.PutState(string(taskId), []byte(state))
	if err != nil {
		return "0"
	}

	fmt.Println("State is:", state)
	return state
}


// =====================================================================================
//      changeStateToAcception - Change task state from instantiation to acception
// =====================================================================================
func changeStateToAcception(stub shim.ChaincodeStubInterface, taskId string) (string) {
	var err error
	var stateAsByte []byte

	stateAsByte, err = stub.GetState(taskId)
	if err != nil {
		return "0"
	}

	state := string(stateAsByte)
	state = "acception"

	// ==== Save task state ====
	err = stub.PutState(taskId, []byte(state))
	if err != nil {
		return "0"
	}

	return state
}

// =======================================
//      readState - Read task state
// =======================================
func (t *TaskChaincode) readState(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("- start readState")
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	taskId := args[0]
	state, err := stub.GetState(taskId)
	if err != nil {
		fmt.Println("Failed to read state!")
	}

	fmt.Println("Current state is:", string(state))
	fmt.Println("- end readState\n")
	return shim.Success(state)
}
