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
curl -s -X POST \
  "http://localhost:$port/getagreement" \
  -H "authorization: Bearer $jim_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"chaincode\":\"$cc_name\",
        \"taskname\":\"ticket-airline\"
}"
echo

echo
echo "============ query object ============"
echo
curl -s -X POST \
  "http://localhost:$port/getagreement" \
  -H "authorization: Bearer $jim_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"chaincode\":\"$cc_name\",
        \"taskname\":\"hotel\"
}"
echo

echo
echo "============ query object ============"
echo
curl -s -X POST \
  "http://localhost:$port/getagreement" \
  -H "authorization: Bearer $jim_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"chaincode\":\"$cc_name\",
        \"taskname\":\"rent-car\"
}"
echo

echo
echo "============ query object ============"
echo
curl -s -X POST \
  "http://localhost:$port/getagreement" \
  -H "authorization: Bearer $jim_gfe_token" \
  -H "content-type: application/json" \
  -d "{
        \"chaincode\":\"$cc_name\",
        \"taskname\":\"scenic-spot\"
}"
echo


echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
