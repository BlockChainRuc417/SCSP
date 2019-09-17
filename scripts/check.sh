#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000
cc_name=task

admin_gfe_token=$(cat ../token/admin_gfe.token)

starttime=$(date +%s)

echo
echo
echo "============ check ============"
echo
agreementJson=$(curl -s -X POST \
  "http://localhost:$port/check" \
  -H "authorization: Bearer $admin_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"channel\":\"softwarechannel\",
        \"chaincode\":\"$cc_name\",
        \"peer\":\"peer0.fabric.gfe.com\",
        \"taskId\":\"1782222264730125312\"
}")
echo $agreementJson
echo

echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
