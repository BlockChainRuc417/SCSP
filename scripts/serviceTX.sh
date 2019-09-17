#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000
cc_name=task

admin_gfe_token=$(cat ../token/admin_gfe.token)

starttime=$(date +%s)

echo "export port=$port;"
echo "export org_token=$admin_gfe_token;"



echo
echo "============ query object ============"
echo
curl -s -X POST \
  "http://localhost:$port/query" \
  -H "authorization: Bearer $admin_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"fcn\":\"queryServiceTxByTaskId\",
        \"args\":[\"1842622374865273856\"]
}"
echo

echo
echo "============ query object ============"
echo
curl -s -X POST \
  "http://localhost:$port/query" \
  -H "authorization: Bearer $admin_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"currency\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"fcn\":\"queryPayTxByTaskId\",
        \"args\":[\"1842622374865273856\"]
}"
echo

echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
