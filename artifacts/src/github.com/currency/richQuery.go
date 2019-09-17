
package main


import (
	"fmt"
	"bytes"
	//"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)


// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// ==== buffer is a JSON array containing QueryRecords ====
	var buffer bytes.Buffer
	buffer.WriteString("[")

	providerAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// ==== Add a comma before array members, suppress it for the first array member =====
		if providerAlreadyWritten == true {
			buffer.WriteString(",")
		}
		//buffer.WriteString("{\"Key\":")
		//buffer.WriteString("\"")
		//buffer.WriteString(queryResponse.Key)
		//buffer.WriteString("\"")

		//buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		////buffer.WriteString("}")
		providerAlreadyWritten = true
	}
	buffer.WriteString("]\n")

	fmt.Printf("- getResultForQueryString queryResult:\n%s\n", buffer.String())
	return buffer.Bytes(), nil
}
