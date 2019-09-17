#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000
cc_name=task

jim_org1_token=$(cat ../token/jim_org1.token)
tim_org2_token=$(cat ../token/tim_org2.token)
qunar_org1_token=$(cat ../token/qunar_org1.token)
ali_org2_token=$(cat ../token/ali_org2.token)
ctrip_org2_token=$(cat ../token/ctrip_org2.token)
dbiir_org2_token=$(cat ../token/dbiir_org2.token)

starttime=$(date +%s)

echo "export port=$port;"
#echo "export org_token=$org_token;"



for((i=1;i<=1;i++));
do
echo
echo
echo "============ service ============"
echo
curl -s -X POST \
  "http://localhost:$port/service" \
  -H "authorization: Bearer $jim_org1_token" \
  -H "content-type: application/json" \
  -d "{
       \"chaincode\":\"$cc_name\",
       \"args\":[\"1829291076423779329\",\"Ctrip\",\"2.0\",\"http://10.77.20.101:8080/ctrip-airline/airlines\",\"get\"]
   }"
echo
echo
done

#\"args\":[\"1803088066542830592\",\"Ctrip\",\"2.0\",\"http://10.77.20.101:8080/ctrip-airline-1.0/airlines/\",\"post\",\"hello!\"]


echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
