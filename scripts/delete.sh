#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000
cc_name=task

admin_gfe_token=$(cat ../token/admin_gfe.token)
jim_gfe_token=$(cat ../token/jim_gfe.token)
tim_deke_token=$(cat ../token/tim_deke.token)

starttime=$(date +%s)

echo "export port=$port;"
echo "export org_token=$admin_gfe_token;"



echo
echo "============ delete task ============"
echo
curl -s -X POST \
  "http://localhost:$port/delete_task" \
  -H "authorization: Bearer $admin_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"taskName\":\"hotel\",
        \"requesterName\":\"Jim\"
}"
echo

echo
echo "============ delete task ============"
echo
curl -s -X POST \
  "http://localhost:$port/delete_task" \
  -H "authorization: Bearer $admin_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"taskName\":\"ticket-airline\",
        \"requesterName\":\"Jim\"
}"
echo

echo
echo "============ delete task ============"
echo
curl -s -X POST \
  "http://localhost:$port/delete_task" \
  -H "authorization: Bearer $admin_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"taskName\":\"rent-car\",
        \"requesterName\":\"Jim\"
}"
echo

echo
echo "============ delete task ============"
echo
curl -s -X POST \
  "http://localhost:$port/delete_task" \
  -H "authorization: Bearer $admin_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"taskName\":\"scenic-spot\",
        \"requesterName\":\"Jim\"
}"
echo



echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
