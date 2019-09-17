#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000
cc_name=task

jim_org1_token=$(cat ../token/jim_org1.token)
tim_org2_token=$(cat ../token/tim_org2.token)
perry_org2_token=$(cat ../token/perry_org2.token)

qunar_org1_token=$(cat ../token/qunar_org1.token)
ali_org2_token=$(cat ../token/ali_org2.token)
ctrip_org2_token=$(cat ../token/ctrip_org2.token)
dbiir_org2_token=$(cat ../token/dbiir_org2.token)

starttime=$(date +%s)

echo "export port=$port;"
echo "export cc_name=$cc_name;"
#echo "export org_token=$org_token;"



echo
echo
echo "============ test request strategy ============"
echo
curl -s -X POST \
  "http://localhost:$port/request_strategy" \
  -H "authorization: Bearer $perry_org2_token" \
  -H "content-type: application/json" \
  -d "{
       \"chaincode\":\"$cc_name\",
       \"taskNames\":[\"hotel\",\"ticket-airline\",\"rent-car\",\"scenic-spot\"]
   }"
echo
echo
echo

#       \"channel\":\"mychannel\",
#       \"peer\":\"peer0.org1.example.com\",



#[\"hotel\",\"ticket-airline\",\"rent-car\",\"scenic-spot\"],

#echo
#echo
#echo "============ test response strategy ============"
#echo
#curl -s -X POST \
#  "http://localhost:$port/response_strategy" \
#  -H "authorization: Bearer $org_token" \
#  -H "content-type: application/json" \
#  -d "{
#       \"channel\":\"mychannel\",
#       \"chaincode\":\"mycc\",
#       \"peer\":\"peer0.org1.example.com\",
#       \"roundIndex\":\"1\",
#       \"requestStrategyPath\":\"/home/perry/SCAS/strategy/request-strategy-u2.json\",
#       \"responseStrategyPath\":\"/home/perry/SCAS/strategy/response-strategy-s2.json\"
#   }"
#echo
#echo
#echo




echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
