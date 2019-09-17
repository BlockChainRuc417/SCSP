
port=4000
cc_name=currency

jim_gfe_token=$(cat ../token/jim_gfe.token)
tim_deke_token=$(cat ../token/tim_deke.token)

qunar_gfe_token=$(cat ../token/qunar_gfe.token)
ali_deke_token=$(cat ../token/ali_deke.token)
ctrip_deke_token=$(cat ../token/ctrip_deke.token)
dbiir_deke_token=$(cat ../token/dbiir_deke.token)

starttime=$(date +%s)

echo "export port=$port;"
echo "export cc_name=$cc_name"
#echo "export org_token=$org_token;"

echo
echo "================ join chaincode ================="
echo
curl -s -X POST \
  "http://localhost:$port/joinchaincode" \
  -H "authorization: Bearer $jim_gfe_token" \
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
  -H "authorization: Bearer $tim_deke_token" \
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
  -H "authorization: Bearer $qunar_gfe_token" \
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
  -H "authorization: Bearer $ctrip_deke_token" \
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
  -H "authorization: Bearer $ali_deke_token" \
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
  -H "authorization: Bearer $dbiir_deke_token" \
  -H "content-type: application/json" \
  -d "{
		   \"chaincode\":\"$cc_name\",
		   \"description\":\"dbiir service\"
     }"
echo
echo

echo
echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
