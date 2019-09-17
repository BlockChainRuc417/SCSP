#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000
cc_name=task

admin_org1_token=$(cat ../token/admin_org1.token)

starttime=$(date +%s)

echo "export port=$port;"
#echo "export org_token=$org_token;"

#echo
#echo
#echo
#curl -s -X POST \
#  "http://localhost:$port/test" \
#  -H "authorization: Bearer $org_token" \
#  -H "content-type: application/json" \
#  -d "{
#       \"channel\":\"mychannel\",
#       \"chaincode\":\"mycc\",
#       \"peer\":\"peer0.org1.example.com\",
#       \"name\":\"analyze\",
#       \"requester\":\"u1\",
#       \"description\":\"000\"
#     }"

echo
echo "================ schedule ================="
echo
curl -s -X POST \
  "http://localhost:$port/schedule" \
  -H "authorization: Bearer $admin_org1_token" \
  -H "content-type: application/json" \
  -d "{
       \"channel\":\"mychannel\",
       \"chaincode\":\"$cc_name\",
       \"peer\":\"peer0.org1.example.com\",
       \"taskId\":\"1787375553826259968\"
     }"
echo
echo
#echo "============ query balance ============"
#echo
#balance=$(curl -s -X POST \
#  "http://localhost:$port/query" \
#  -H "authorization: Bearer $org_token" \
#  -H "content-type: application/json" \
#  -d "{
#       \"channel\":\"mychannel\",
#       \"chaincode\":\"mycc\",
#       \"peer\":\"peer0.org1.example.com\",
#       \"fcn\":\"getBalance\",
#       \"args\":[\"s1\"]
#   }")
#echo $balance

#     echo
#     echo
#     echo "============ query agreement ============"
#     echo
#     agreementJson=$(curl -s -X POST \
#       "http://localhost:$port/query" \
#       -H "authorization: Bearer $org_token" \
#       -H "content-type: application/json" \
#       -d "{
#             \"channel\":\"mychannel\",
#             \"chaincode\":\"mycc\",
#             \"peer\":\"peer0.org1.example.com\",
#             \"fcn\":\"queryByObjectType\",
#             \"args\":[\"1769919787821433856\",\"request\"]
#     }")
#     echo $agreementJson


echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
