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
perry_deke_token=$(cat ../token/perry_deke.token)

starttime=$(date +%s)

echo "export port=$port;"
echo "export org_token=$admin_gfe_token;"



echo
echo "============ delete task ============"
echo
curl -s -X POST \
  "http://localhost:$port/delete_task" \
  -H "authorization: Bearer $admin_deke_token" \
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
  -H "authorization: Bearer $admin_deke_token" \
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
  -H "authorization: Bearer $admin_deke_token" \
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
  -H "authorization: Bearer $admin_deke_token" \
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
