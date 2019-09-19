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

#echo
#echo "GET query Channels"
#echo
#channelInfo=$(curl -s -X GET \
#  "http://localhost:4000/channels?peer=peer0.gfe.example.com" \
#  -H "authorization: Bearer $admin_gfe_token" \
#  -H "content-type: application/json")
#echo $channelInfo
#echo

#echo "GET query Installed chaincodes"
#echo
#curl -s -X GET \
#  "http://localhost:4000/installed_chaincodes?channel=mychannel&peer=peer0.gfe.example.com" \
#  -H "authorization: Bearer $admin_gfe_token" \
#  -H "content-type: application/json"
#echo
#echo

#echo "GET query Instantiated chaincodes"
#echo
#curl -s -X GET \
#  "http://localhost:4000/instantiated_chaincodes?channel=mychannel&peer=peer0.deke.example.com" \
#  -H "authorization: Bearer $admin_deke_token" \
#  -H "content-type: application/json"
#echo
#echo

#echo "GET query "
#echo
#curl -s -X GET "http://10.77.20.101:8080/ctrip-airline-1.0/airlines?abb=%E5%9B%BD%E8%88%AA"
#echo
#echo

#echo
#echo "GET query Organizations"
#echo
#curl -s -X GET \
#  "http://localhost:$port/organizations?channel=mychannel" \
#  -H "authorization: Bearer $admin_gfe_token"
#echo

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
