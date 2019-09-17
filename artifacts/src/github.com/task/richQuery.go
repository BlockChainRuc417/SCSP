
package main


import (
	"fmt"
	"bytes"
	"encoding/json"

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


// =========================================================================================
// getArrayForRequest executes the passed in query string.
// Result set is built and returned as a agreement array containing the JSON results.
// =========================================================================================
func getJSONForAgreement(stub shim.ChaincodeStubInterface, queryString string) (agreement, error) {
	var agr agreement
	var agrArray agreement

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return agrArray, err
	}
	defer resultsIterator.Close()

	fmt.Println("agreementArray:" )
	for resultsIterator.HasNext() {
		queryRequest, err := resultsIterator.Next()
		if err != nil {
			return agrArray, err
		}

		errResult := json.Unmarshal([]byte(queryRequest.Value), &agr)
		if errResult != nil {
			return agrArray, errResult
		}

		fmt.Println(string(queryRequest.Value))
	}
	return agr, nil
}


// =========================================================================================
// getArrayForRequest executes the passed in query string.
// Result set is built and returned as a request array containing the JSON results.
// =========================================================================================
func getArrayForRequest(stub shim.ChaincodeStubInterface, queryString string) ([]request, error) {
	var req request
	reqArray := make([]request, 0)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return reqArray, err
	}
	defer resultsIterator.Close()

	fmt.Println("requestArray:" )
	for resultsIterator.HasNext() {
		queryRequest, err := resultsIterator.Next()
		if err != nil {
			return reqArray, err
		}

		errResult := json.Unmarshal([]byte(queryRequest.Value), &req)
		if errResult != nil {
			return reqArray, errResult
		}

		reqArray = append(reqArray, req)
		fmt.Println(string(queryRequest.Value))
	}
	return reqArray, nil
}


// =========================================================================================
// getArrayForResponse executes the passed in query string.
// Result set is built and returned as a response array containing the JSON results.
// =========================================================================================
func getArrayForResponse(stub shim.ChaincodeStubInterface, queryString string) ([]response, error) {
	var res response
	resArray := make([]response, 0)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return resArray, err
	}
	defer resultsIterator.Close()

	fmt.Println("responseArray:" )
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return resArray, err
		}

		errResult := json.Unmarshal([]byte(queryResponse.Value), &res)
		if errResult != nil {
			return resArray, errResult
		}

		resArray = append(resArray, res)
		fmt.Println(string(queryResponse.Value))
	}
	return resArray, nil
}


// =========================================================================================
// getArrayForServiceTX executes the passed in query string.
// Result set is built and returned as a response array containing the JSON results.
// =========================================================================================
func getArrayForServiceTX(stub shim.ChaincodeStubInterface, queryString string) ([]serviceTX, error) {
	var servTX serviceTX
	serviceTXArray := make([]serviceTX, 0)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return serviceTXArray, err
	}
	defer resultsIterator.Close()

	fmt.Println("serviceTXArray:" )
	for resultsIterator.HasNext() {
		queryServiceTX, err := resultsIterator.Next()
		if err != nil {
			return serviceTXArray, err
		}

		errResult := json.Unmarshal([]byte(queryServiceTX.Value), &servTX)
		if errResult != nil {
			return serviceTXArray, errResult
		}

		serviceTXArray = append(serviceTXArray, servTX)
		fmt.Println(string(queryServiceTX.Value))
	}

	return serviceTXArray, nil
}
