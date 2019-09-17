
package main


import (
	"fmt"
  "encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type user struct {
	ObjectType     string    `json:"objectType"`
  Account        string    `json:"account"`
  OrgName        string    `json:"orgName"`
	Description    string    `json:"description"`
}


// ===================================================
//       regist - new user need to regist first
// ===================================================
func (t *TaskChaincode) regist(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0            1            2
	//"$username", "$orgname"  "$description"
  if len(args) != 3 {
    return shim.Error("Incorrect number of arguments. Expecting 3")
  }

  fmt.Println("- start regist")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3nd argument must be a non-empty string")
	}

	objectType := "User"
	account := args[0]
	orgname := args[1]
	description := args[2]

	// ==== Check if account already exists ====
	registAsBytes, err := stub.GetState(account)
	if err != nil {
		return shim.Error("Failed to regist: " + err.Error())
	} else if registAsBytes != nil {
		fmt.Println("The account already exists! Please change your account.")
		return shim.Error("The account " + account + " already exists! Please modify your account.")
	}

	registJSON := &user{objectType, account, orgname, description}
  registJSONAsBytes, err := json.Marshal(registJSON)
	if err != nil {
		return shim.Error(err.Error())
	}
	// ==== Save regist to state ====
  err = stub.PutState(account, registJSONAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

  fmt.Println("regist successfully", account)
  fmt.Println("- end regist\n")
	return shim.Success(registJSONAsBytes)
}


// ===========================================================
//       signIn - users need to sign in before add task
// ===========================================================
/*func (t *TaskChaincode) signIn(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var requesterJSON user
	var jsonResp string

	//  0         1
	//"Tom", "123456789"
  if len(args) != 2 {
    return shim.Error("Incorrect number of arguments. Expecting 2")
  }

  fmt.Println("- start signIn")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	account := args[0]
	password := args[1]

	// ==== Check if account exists ====
	registAsBytes, err := stub.GetState(account)
	if err != nil {
		return shim.Error("Failed to get account: " + err.Error())
	} else if registAsBytes == nil {
		fmt.Println("The account doesn't exists! Please check your account.")
		return shim.Error("The account doesn't exists! Please check your account.")
	}

	err = json.Unmarshal([]byte(registAsBytes), &requesterJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + account + "\"}"
		return shim.Error(jsonResp)
	}

	// ==== Check if password is right ====
	if password != requesterJSON.Password {
		fmt.Println("Failed to sign in! Please check your password!")
		return shim.Error("Failed to sign in! Please check your password!")
	}

	fmt.Println("Successfully!")
	fmt.Println("- end signIn")
	return shim.Success([]byte(requesterJSON.Account))
}*/



// ============================================================
//      queryTask - query all tasks from chaincode state
// ============================================================
func (t *TaskChaincode) queryMembers(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("- start query all users")

	queryString := fmt.Sprintf("{\"selector\":{\"objectType\":\"User\"}}")
	queryResults, err := getResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end query all users\n")
	return shim.Success(queryResults)

}
