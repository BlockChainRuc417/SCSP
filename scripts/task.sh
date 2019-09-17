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

qunar_gfe_token=$(cat ../token/qunar_gfe.token)
ali_deke_token=$(cat ../token/ali_deke.token)
ctrip_deke_token=$(cat ../token/ctrip_deke.token)
dbiir_deke_token=$(cat ../token/dbiir_deke.token)

starttime=$(date +%s)

echo "export port=$port;"
echo "export cc_name=$cc_name"


echo
echo "============ write task ============"
echo
taskJson=$(curl -s -X POST \
  "http://localhost:$port/writetask" \
  -H "authorization: Bearer $tim_deke_token" \
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
  -H "authorization: Bearer $perry_deke_token" \
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
