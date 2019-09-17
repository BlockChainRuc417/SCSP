#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000
cc_name=currency

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

echo
echo
echo "============ query payTX ============"
echo
curl -s -X POST \
  "http://localhost:$port/query" \
  -H "authorization: Bearer $admin_gfe_token" \
  -H "content-type: application/json" \
  -d "{
       \"channel\":\"mychannel\",
       \"peer\":\"peer0.gfe.example.com\",
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
     -H "authorization: Bearer $admin_gfe_token" \
     -H "content-type: application/json" \
     -d "{
          \"channel\":\"mychannel\",
          \"peer\":\"peer0.gfe.example.com\",
          \"chaincode\":\"$cc_name\",
          \"fcn\":\"queryPayTxByPayee\",
          \"args\":[\"contract\"]
      }"


echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
