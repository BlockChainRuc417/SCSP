#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000

jim_org1_token=$(cat ../token/jim_org1.token)
tim_org2_token=$(cat ../token/tim_org2.token)

qunar_org1_token=$(cat ../token/qunar_org1.token)
ali_org2_token=$(cat ../token/ali_org2.token)
ctrip_org2_token=$(cat ../token/ctrip_org2.token)
dbiir_org2_token=$(cat ../token/dbiir_org2.token)

starttime=$(date +%s)

echo "export port=$port;"
#echo "export jim_org1_token=$jim_org1_token;"

echo
echo "============ set strategy file ============"
echo
curl -s -X POST \
  "http://localhost:$port/set_strategy" \
  -H "authorization: Bearer $jim_org1_token" \
  -H "content-type: application/json" \
  -d "{
        \"role\":\"requester\",
        \"strategyFile\":\"request-strategy-jim.json\"
}"
echo
echo
echo
echo "============ set strategy file ============"
echo
curl -s -X POST \
  "http://localhost:$port/set_strategy" \
  -H "authorization: Bearer $tim_org2_token" \
  -H "content-type: application/json" \
  -d "{
        \"role\":\"requester\",
        \"strategyFile\":\"request-strategy-tim.json\"
}"
echo
echo
echo
echo
echo "============ set strategy file ============"
echo
curl -s -X POST \
  "http://localhost:$port/set_strategy" \
  -H "authorization: Bearer $qunar_org1_token" \
  -H "content-type: application/json" \
  -d "{
        \"role\":\"provider\",
        \"strategyFile\":\"response-strategy-qunar.json\"
}"
echo
echo
echo
echo "============ set strategy file ============"
echo
curl -s -X POST \
  "http://localhost:$port/set_strategy" \
  -H "authorization: Bearer $ctrip_org2_token" \
  -H "content-type: application/json" \
  -d "{
        \"role\":\"provider\",
        \"strategyFile\":\"response-strategy-ctrip.json\"
}"
echo
echo
echo
echo "============ set strategy file ============"
echo
curl -s -X POST \
  "http://localhost:$port/set_strategy" \
  -H "authorization: Bearer $ali_org2_token" \
  -H "content-type: application/json" \
  -d "{
        \"role\":\"provider\",
        \"strategyFile\":\"response-strategy-ali.json\"
}"
echo
echo
echo
echo "============ set strategy file ============"
echo
curl -s -X POST \
  "http://localhost:$port/set_strategy" \
  -H "authorization: Bearer $dbiir_org2_token" \
  -H "content-type: application/json" \
  -d "{
        \"role\":\"provider\",
        \"strategyFile\":\"response-strategy-dbiir.json\"
}"
echo
echo



echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
