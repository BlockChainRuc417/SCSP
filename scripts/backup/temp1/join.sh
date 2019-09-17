
port=4000
cc_name=currency

jim_org1_token=$(cat ../token/jim_org1.token)
tim_org2_token=$(cat ../token/tim_org2.token)

qunar_org1_token=$(cat ../token/qunar_org1.token)
ali_org2_token=$(cat ../token/ali_org2.token)
ctrip_org2_token=$(cat ../token/ctrip_org2.token)
dbiir_org2_token=$(cat ../token/dbiir_org2.token)

starttime=$(date +%s)

echo "export port=$port;"
echo "export cc_name=$cc_name"
#echo "export org_token=$org_token;"

echo
echo "================ join chaincode ================="
echo
curl -s -X POST \
  "http://localhost:$port/joinchaincode" \
  -H "authorization: Bearer $jim_org1_token" \
  -H "content-type: application/json" \
  -d "{
		   \"chaincode\":\"$cc_name\",
		   \"description\":\"travel agent\"
     }"
echo
echo
echo "================ join chaincode ================="
echo
curl -s -X POST \
  "http://localhost:$port/joinchaincode" \
  -H "authorization: Bearer $tim_org2_token" \
  -H "content-type: application/json" \
  -d "{
		   \"chaincode\":\"$cc_name\",
		   \"description\":\"travel agent\"
     }"
echo
echo
echo "================ join chaincode ================="
echo
curl -s -X POST \
  "http://localhost:$port/joinchaincode" \
  -H "authorization: Bearer $qunar_org1_token" \
  -H "content-type: application/json" \
  -d "{
		   \"chaincode\":\"$cc_name\",
		   \"description\":\"qunar service\"
     }"
echo
echo
echo "================ join chaincode ================="
echo
curl -s -X POST \
  "http://localhost:$port/joinchaincode" \
  -H "authorization: Bearer $ctrip_org2_token" \
  -H "content-type: application/json" \
  -d "{
		   \"chaincode\":\"$cc_name\",
		   \"description\":\"ctrip service\"
     }"
echo
echo
echo
echo "================ join chaincode ================="
echo
curl -s -X POST \
  "http://localhost:$port/joinchaincode" \
  -H "authorization: Bearer $ali_org2_token" \
  -H "content-type: application/json" \
  -d "{
		   \"chaincode\":\"$cc_name\",
		   \"description\":\"ali service\"
     }"
echo
echo
echo "================ join chaincode ================="
echo
curl -s -X POST \
  "http://localhost:$port/joinchaincode" \
  -H "authorization: Bearer $dbiir_org2_token" \
  -H "content-type: application/json" \
  -d "{
		   \"chaincode\":\"$cc_name\",
		   \"description\":\"dbiir service\"
     }"
echo
echo

echo
echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
