#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000
cc_name=currency

admin_org1_token=$(cat ../token/admin_org1.token)
admin_org2_token=$(cat ../token/admin_org2.token)
jim_org1_token=$(cat ../token/jim_org1.token)
tim_org2_token=$(cat ../token/tim_org2.token)
qunar_org1_token=$(cat ../token/qunar_org1.token)
ali_org2_token=$(cat ../token/ali_org2.token)
ctrip_org2_token=$(cat ../token/ctrip_org2.token)
dbiir_org2_token=$(cat ../token/dbiir_org2.token)

starttime=$(date +%s)

echo "export port=$port;"

echo
echo
echo "============ query payTX ============"
echo
curl -s -X POST \
  "http://localhost:$port/query" \
  -H "authorization: Bearer $admin_org1_token" \
  -H "content-type: application/json" \
  -d "{
       \"channel\":\"mychannel\",
       \"peer\":\"peer0.org1.example.com\",
       \"chaincode\":\"$cc_name\",
       \"fcn\":\"queryPayTxByTaskId\",
       \"args\":[\"1798026674009801728\"]
   }"

   echo
   echo
   echo "============ query payTX ============"
   echo
   curl -s -X POST \
     "http://localhost:$port/query" \
     -H "authorization: Bearer $admin_org1_token" \
     -H "content-type: application/json" \
     -d "{
          \"channel\":\"mychannel\",
          \"peer\":\"peer0.org1.example.com\",
          \"chaincode\":\"$cc_name\",
          \"fcn\":\"queryPayTxByPayee\",
          \"args\":[\"contract\"]
      }"


echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
