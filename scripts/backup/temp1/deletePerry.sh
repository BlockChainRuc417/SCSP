#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000
cc_name=task

admin_org1_token=$(cat ../token/admin_org1.token)
admin_org2_token=$(cat ../token/admin_org2.token)

jim_org1_token=$(cat ../token/jim_org1.token)
tim_org2_token=$(cat ../token/tim_org2.token)
perry_org2_token=$(cat ../token/perry_org2.token)

starttime=$(date +%s)

echo "export port=$port;"
echo "export org_token=$admin_org1_token;"



echo
echo "============ delete task ============"
echo
curl -s -X POST \
  "http://localhost:$port/delete_task" \
  -H "authorization: Bearer $admin_org2_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.deke.com\",
        \"taskName\":\"hotel\",
        \"requesterName\":\"Perry\"
}"
echo

echo
echo "============ delete task ============"
echo
curl -s -X POST \
  "http://localhost:$port/delete_task" \
  -H "authorization: Bearer $admin_org2_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.deke.com\",
        \"taskName\":\"ticket-airline\",
        \"requesterName\":\"Perry\"
}"
echo

echo
echo "============ delete task ============"
echo
curl -s -X POST \
  "http://localhost:$port/delete_task" \
  -H "authorization: Bearer $admin_org2_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.deke.com\",
        \"taskName\":\"rent-car\",
        \"requesterName\":\"Perry\"
}"
echo

echo
echo "============ delete task ============"
echo
curl -s -X POST \
  "http://localhost:$port/delete_task" \
  -H "authorization: Bearer $admin_org2_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.deke.com\",
        \"taskName\":\"scenic-spot\",
        \"requesterName\":\"Perry\"
}"
echo



echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
