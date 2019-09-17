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
echo "export org_token=$admin_org1_token;"

echo
echo "============ query object ============"
echo
curl -s -X POST \
  "http://localhost:$port/query" \
  -H "authorization: Bearer $admin_org1_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"fcn\":\"queryByObjectType\",
        \"args\":[\"1811948164073653249\",\"request\"]
}"
echo

echo
echo "============ query object ============"
echo
curl -s -X POST \
  "http://localhost:$port/query" \
  -H "authorization: Bearer $admin_org1_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"fcn\":\"queryByObjectType\",
        \"args\":[\"1811948164073653249\",\"response\"]
}"
echo

echo
echo "============ query object ============"
echo
curl -s -X POST \
  "http://localhost:$port/query" \
  -H "authorization: Bearer $admin_org1_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"fcn\":\"queryTask\",
        \"args\":[]
}"
echo

echo
echo "============ query object ============"
echo
curl -s -X GET \
  "http://localhost:$port/get_taskmap" \
  -H "authorization: Bearer $admin_org1_token"
echo

echo
echo "============ query object ============"
echo
curl -s -X GET \
  "http://localhost:$port/get_requestmap" \
  -H "authorization: Bearer $admin_org1_token"
echo

echo
echo "============ query object ============"
echo
curl -s -X GET \
  "http://localhost:$port/get_taskjsonmap" \
  -H "authorization: Bearer $admin_org1_token"
echo

echo
echo "============ query object ============"
echo
curl -s -X POST \
  "http://localhost:$port/query" \
  -H "authorization: Bearer $admin_org1_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"fcn\":\"queryServiceTxByTaskId\",
        \"args\":[\"1829041010425463808\"]
}"
echo

echo
echo "============ query object ============"
echo
curl -s -X POST \
  "http://localhost:$port/query" \
  -H "authorization: Bearer $admin_org1_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"currency\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"fcn\":\"queryPayTxByTaskId\",
        \"args\":[\"1829041010425463808\"]
}"
echo

echo
echo "============ query object ============"
echo
curl -s -X POST \
  "http://localhost:$port/query" \
  -H "authorization: Bearer $admin_org1_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"task\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"fcn\":\"queryTaskByNameAndRequester\",
        \"args\":[\"Jim\",\"ticket-airline\"]
}"
echo

echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
