#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000
cc_name=task

admin_org1_token=$(cat ../token/admin_org1.token)
jim_org1_token=$(cat ../token/jim_org1.token)
tim_org2_token=$(cat ../token/tim_org2.token)

starttime=$(date +%s)

echo "export port=$port;"
echo "export org_token=$admin_org1_token;"



echo
echo "============ delete task ============"
echo
curl -s -X POST \
  "http://localhost:$port/delete_task" \
  -H "authorization: Bearer $admin_org1_token" \
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
  -H "authorization: Bearer $admin_org1_token" \
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
  -H "authorization: Bearer $admin_org1_token" \
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
  -H "authorization: Bearer $admin_org1_token" \
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
