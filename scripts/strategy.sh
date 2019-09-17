#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000

jim_gfe_token=$(cat ../token/jim_gfe.token)
tim_deke_token=$(cat ../token/tim_deke.token)

qunar_gfe_token=$(cat ../token/qunar_gfe.token)
ali_deke_token=$(cat ../token/ali_deke.token)
ctrip_deke_token=$(cat ../token/ctrip_deke.token)
dbiir_deke_token=$(cat ../token/dbiir_deke.token)

starttime=$(date +%s)

echo "export port=$port;"
#echo "export jim_gfe_token=$jim_gfe_token;"

echo
echo "============ set strategy file ============"
echo
curl -s -X POST \
  "http://localhost:$port/set_strategy" \
  -H "authorization: Bearer $jim_gfe_token" \
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
  -H "authorization: Bearer $tim_deke_token" \
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
  -H "authorization: Bearer $qunar_gfe_token" \
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
  -H "authorization: Bearer $ctrip_deke_token" \
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
  -H "authorization: Bearer $ali_deke_token" \
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
  -H "authorization: Bearer $dbiir_deke_token" \
  -H "content-type: application/json" \
  -d "{
        \"role\":\"provider\",
        \"strategyFile\":\"response-strategy-dbiir.json\"
}"
echo
echo



echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
