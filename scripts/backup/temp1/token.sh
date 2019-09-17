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

starttime=$(date +%s)

echo "export port=$port;"

echo
echo "============ query object ============"
echo
JIM_ORG1_TOKEN=$(curl -s -X POST \
  "http://localhost:$port/newtoken" \
  -H "authorization: Bearer $jim_org1_token");
echo $JIM_ORG1_TOKEN
JIM_ORG1_TOKEN=$(echo $JIM_ORG1_TOKEN | jq ".token" | sed "s/\"//g")
echo $JIM_ORG1_TOKEN>../token/jim_org1.token


echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
