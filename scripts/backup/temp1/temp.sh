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

#echo
#echo "GET query Channels"
#echo
#channelInfo=$(curl -s -X GET \
#  "http://localhost:4000/channels?peer=peer0.org1.example.com" \
#  -H "authorization: Bearer $admin_org1_token" \
#  -H "content-type: application/json")
#echo $channelInfo
#echo

#echo "GET query Installed chaincodes"
#echo
#curl -s -X GET \
#  "http://localhost:4000/installed_chaincodes?channel=mychannel&peer=peer0.org1.example.com" \
#  -H "authorization: Bearer $admin_org1_token" \
#  -H "content-type: application/json"
#echo
#echo

#echo "GET query Instantiated chaincodes"
#echo
#curl -s -X GET \
#  "http://localhost:4000/instantiated_chaincodes?channel=mychannel&peer=peer0.org2.example.com" \
#  -H "authorization: Bearer $admin_org2_token" \
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
#  -H "authorization: Bearer $admin_org1_token"
#echo

echo
echo "============ query object ============"
echo
curl -s -X POST \
  "http://localhost:$port/getagreement" \
  -H "authorization: Bearer $jim_org1_token" \
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
  -H "authorization: Bearer $jim_org1_token" \
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
  -H "authorization: Bearer $jim_org1_token" \
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
  -H "authorization: Bearer $jim_org1_token" \
  -H "content-type: application/json" \
  -d "{
        \"chaincode\":\"$cc_name\",
        \"taskname\":\"scenic-spot\"
}"
echo


echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
