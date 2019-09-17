#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000
cc_name=task

admin_gfe_token=$(cat ../token/admin_gfe.token)
admin_deke_token=$(cat ../token/admin_deke.token)
jim_gfe_token=$(cat ../token/jim_gfe.token)
tim_deke_token=$(cat ../token/tim_deke.token)
qunar_gfe_token=$(cat ../token/qunar_gfe.token)
ali_deke_token=$(cat ../token/ali_deke.token)
ctrip_deke_token=$(cat ../token/ctrip_deke.token)
dbiir_deke_token=$(cat ../token/dbiir_deke.token)

starttime=$(date +%s)

echo "export port=$port;"
echo "export cc_name=$cc_name"


echo
echo "============ write task ============"
echo
taskJson=$(curl -s -X POST \
  "http://localhost:$port/writetask" \
  -H "authorization: Bearer $jim_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"chaincode\":\"$cc_name\",
        \"name\":\"analyze12\",
        \"description\":\"000\"
}")
echo $taskJson
taskId=$(echo $taskJson | jq ".id" | sed "s/\"//g")
echo "taskId is $taskId"

echo
echo
echo "============ write request ============"
echo
requestJson=$(curl -s -X POST \
  "http://localhost:$port/writerequest" \
  -H "authorization: Bearer $jim_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"chaincode\":\"$cc_name\",
        \"taskId\":\"$taskId\",
        \"request\":[\"1.0\",\"100\",\"15\"]
}")
echo $requestJson
requestId=$(echo $requestJson | jq ".reqId" | sed "s/\"//g")
echo "requestId is $requestId"

echo
echo
echo "============ write request ============"
echo
requestJson=$(curl -s -X POST \
  "http://localhost:$port/writerequest" \
  -H "authorization: Bearer $jim_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"chaincode\":\"$cc_name\",
        \"taskId\":\"$taskId\",
        \"request\":[\"1.5\",\"150\",\"30\"]
}")
echo $requestJson
requestId=$(echo $requestJson | jq ".reqId" | sed "s/\"//g")
echo "requestId is $requestId"

echo
echo
echo "============ write response ============"
echo
curl -s -X POST \
  "http://localhost:$port/writeresponse" \
  -H "authorization: Bearer $ali_deke_token" \
  -H "content-type: application/json" \
  -d "{
        \"chaincode\":\"$cc_name\",
        \"requestId\":\"$requestId\",
        \"requester\":\"Jim\",
        \"taskId\":\"$taskId\",
        \"url\":\"http://localhost:8080/hotel3-0.0.1-SNAPSHOT/hotels/\",
        \"expireTime\":\"30 May 2018 18:43:00 +0800\",
        \"response\":[\"0.7\",\"160\",\"28\"]
}"

echo
echo
echo "============ check ============"
echo
agreementJson=$(curl -s -X POST \
  "http://localhost:$port/check" \
  -H "authorization: Bearer $admin_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"taskId\":\"$taskId\"
}")
echo $agreementJson
expireTime=$(echo $agreementJson | jq ".expireTime" | sed "s/\"//g")
echo "expireTime is $expireTime"

echo
echo "============ confirm pay ============"
echo
agreementJson=$(curl -s -X POST \
  "http://localhost:$port/invoke" \
  -H "authorization: Bearer $admin_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"currency\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"fcn\":\"confirmPay\",
        \"args\":[\"$taskId\"]
}")
echo $agreementJson

echo
echo
echo "============ new round ============"
echo
curl -s -X POST \
  "http://localhost:$port/new_round" \
  -H "authorization: Bearer $admin_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"taskId\":\"$taskId\"
}"

#echo
#echo "============ query balance ============"
#echo
#balance=$(curl -s -X POST \
#  "http://localhost:$port/query" \
#  -H "authorization: Bearer $org_token" \
#  -H "content-type: application/json" \
#  -d "{
#        \"channel\":\"mychannel\",
#        \"chaincode\":\"mycc\",
#        \"peer\":\"peer0.gfe.example.com\",
#        \"fcn\":\"getBalance\",
#        \"args\":[\"u1\"]
#}")
#echo $balance

#echo
#echo "============ query balance ============"
#echo
#balance=$(curl -s -X POST \
#  "http://localhost:$port/query" \
#  -H "authorization: Bearer $org_token" \
#  -H "content-type: application/json" \
#  -d "{
#        \"channel\":\"mychannel\",
#        \"chaincode\":\"mycc\",
#        \"peer\":\"peer0.gfe.example.com\",
#        \"fcn\":\"getBalance\",
#        \"args\":[\"s1\"]
#}")
#echo $balance

#echo
#echo "================ scheduled job ================="
#echo
#curl -s -X POST \
#  "http://localhost:$port/schedule" \
#  -H "authorization: Bearer $org_token" \
#  -H "content-type: application/json" \
#  -d "{
#       \"channel\":\"mychannel\",
#       \"chaincode\":\"mycc\",
#       \"peer\":\"peer0.gfe.example.com\",
#       \"taskId\":\"$taskId\"
#     }"

#echo
#echo "============ confirm pay ============"
#echo
#agreementJson=$(curl -s -X POST \
#  "http://localhost:$port/invoke" \
#  -H "authorization: Bearer $org_token" \
#  -H "content-type: application/json" \
#  -d "{
#        \"channel\":\"mychannel\",
#        \"chaincode\":\"mycc\",
#        \"peer\":\"peer0.gfe.example.com\",
#        \"fcn\":\"confirmPay\",
#        \"args\":[\"$taskId\"]
#}")
#echo $agreementJson

echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
