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

starttime=$(date +%s)

echo "export port=$port;"

echo
echo "============ query object ============"
echo
JIM_GFE_TOKEN=$(curl -s -X POST \
  "http://localhost:$port/newtoken" \
  -H "authorization: Bearer $jim_gfe_token")
echo $JIM_GFE_TOKEN
JIM_GFE_TOKEN=$(echo $JIM_GFE_TOKEN | jq ".token" | sed "s/\"//g")
echo $JIM_GFE_TOKEN>../token/jim_gfe.token


echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
