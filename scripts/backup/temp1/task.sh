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

qunar_org1_token=$(cat ../token/qunar_org1.token)
ali_org2_token=$(cat ../token/ali_org2.token)
ctrip_org2_token=$(cat ../token/ctrip_org2.token)
dbiir_org2_token=$(cat ../token/dbiir_org2.token)

starttime=$(date +%s)

echo "export port=$port;"
echo "export cc_name=$cc_name"


echo
echo "============ write task ============"
echo
taskJson=$(curl -s -X POST \
  "http://localhost:$port/writetask" \
  -H "authorization: Bearer $tim_org2_token" \
  -H "content-type: application/json" \
  -d "{
        \"chaincode\":\"$cc_name\",
        \"name\":\"analyze1\",
        \"description\":\"000\"
}")
echo $taskJson
taskId=$(echo $taskJson | jq ".id" | sed "s/\"//g")
echo "taskId is $taskId"

echo
echo "============ write task ============"
echo
taskJson=$(curl -s -X POST \
  "http://localhost:$port/writetask" \
  -H "authorization: Bearer $perry_org2_token" \
  -H "content-type: application/json" \
  -d "{
        \"chaincode\":\"$cc_name\",
        \"name\":\"analyze2\",
        \"description\":\"000\"
}")
echo $taskJson
taskId=$(echo $taskJson | jq ".id" | sed "s/\"//g")
echo "taskId is $taskId"

echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
