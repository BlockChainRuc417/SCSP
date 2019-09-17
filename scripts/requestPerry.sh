#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000
cc_name=task

jim_gfe_token=$(cat ../token/jim_gfe.token)
tim_deke_token=$(cat ../token/tim_deke.token)
perry_deke_token=$(cat ../token/perry_deke.token)

qunar_gfe_token=$(cat ../token/qunar_gfe.token)
ali_deke_token=$(cat ../token/ali_deke.token)
ctrip_deke_token=$(cat ../token/ctrip_deke.token)
dbiir_deke_token=$(cat ../token/dbiir_deke.token)

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
  -H "authorization: Bearer $perry_deke_token" \
  -H "content-type: application/json" \
  -d "{
       \"chaincode\":\"$cc_name\",
       \"taskNames\":[\"hotel\",\"ticket-airline\",\"rent-car\",\"scenic-spot\"]
   }"
echo
echo
echo

#       \"channel\":\"mychannel\",
#       \"peer\":\"peer0.gfe.example.com\",



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
#       \"peer\":\"peer0.gfe.example.com\",
#       \"roundIndex\":\"1\",
#       \"requestStrategyPath\":\"/home/perry/SCAS/strategy/request-strategy-u2.json\",
#       \"responseStrategyPath\":\"/home/perry/SCAS/strategy/response-strategy-s2.json\"
#   }"
#echo
#echo
#echo




echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
