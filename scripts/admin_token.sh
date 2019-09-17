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

starttime=$(date +%s)

echo "export port=$port;"

echo
echo "============ query object ============"
echo
ADMIN_GFE_TOKEN=$(curl -s -X POST \
  "http://localhost:$port/admin_newtoken" \
  -H "authorization: Bearer $admin_gfe_token")
echo $ADMIN_GFE_TOKEN
ADMIN_GFE_TOKEN=$(echo $ADMIN_GFE_TOKEN | jq ".token" | sed "s/\"//g")
echo $ADMIN_GFE_TOKEN>../token/admin_gfe.token


echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
