#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

port=4000
cc_name=task

jim_gfe_token=$(cat ../token/jim_gfe.token)
tim_deke_token=$(cat ../token/tim_deke.token)
perry_deke_token=$(cat ../token/perry_deke.token)

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
        \"strategyFile\":\"request-strategy-Jim.json\"
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
        \"strategyFile\":\"request-strategy-Tim.json\"
}"
echo
echo
echo
echo "============ set strategy file ============"
echo
curl -s -X POST \
  "http://localhost:$port/set_strategy" \
  -H "authorization: Bearer $perry_deke_token" \
  -H "content-type: application/json" \
  -d "{
        \"role\":\"requester\",
        \"strategyFile\":\"request-strategy-Perry.json\"
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
        \"strategyFile\":\"response-strategy-Qunar.json\"
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
        \"strategyFile\":\"response-strategy-Ctrip.json\"
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
        \"strategyFile\":\"response-strategy-Ali.json\"
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
        \"strategyFile\":\"response-strategy-Dbiir.json\"
}"
echo
echo


echo
echo
echo "============ test task-listener ============"
echo
curl -s -X POST \
  "http://localhost:$port/add_listener" \
  -H "authorization: Bearer $ali_deke_token" \
  -H "content-type: application/json" \
  -d "{
       \"chaincode\":\"$cc_name\",
       \"taskName\":\"ticket-airline\"
   }"
echo
echo
echo
echo "============ test task-listener ============"
echo
curl -s -X POST \
  "http://localhost:$port/add_listener" \
  -H "authorization: Bearer $ctrip_deke_token" \
  -H "content-type: application/json" \
  -d "{
       \"chaincode\":\"$cc_name\",
       \"taskName\":\"ticket-airline\"
   }"
echo
echo
echo
echo
echo "============ test task-listener ============"
echo
curl -s -X POST \
  "http://localhost:$port/add_listener" \
  -H "authorization: Bearer $qunar_gfe_token" \
  -H "content-type: application/json" \
  -d "{
       \"chaincode\":\"$cc_name\",
       \"taskName\":\"ticket-airline\"
   }"
echo
echo
echo
echo
echo "============ test task-listener ============"
echo
curl -s -X POST \
  "http://localhost:$port/add_listener" \
  -H "authorization: Bearer $dbiir_deke_token" \
  -H "content-type: application/json" \
  -d "{
       \"chaincode\":\"$cc_name\",
       \"taskName\":\"ticket-airline\"
   }"
echo
echo
echo
echo
echo
echo
echo "============ test task-listener ============"
echo
curl -s -X POST \
  "http://localhost:$port/add_listener" \
  -H "authorization: Bearer $qunar_gfe_token" \
  -H "content-type: application/json" \
  -d "{
       \"chaincode\":\"$cc_name\",
       \"taskName\":\"hotel\"
   }"
echo
echo
echo
echo
echo
echo "============ test task-listener ============"
echo
curl -s -X POST \
  "http://localhost:$port/add_listener" \
  -H "authorization: Bearer $ali_deke_token" \
  -H "content-type: application/json" \
  -d "{
       \"chaincode\":\"$cc_name\",
       \"taskName\":\"hotel\"
   }"
echo
echo
echo
echo
echo
echo "============ test task-listener ============"
echo
curl -s -X POST \
  "http://localhost:$port/add_listener" \
  -H "authorization: Bearer $ali_deke_token" \
  -H "content-type: application/json" \
  -d "{
       \"chaincode\":\"$cc_name\",
       \"taskName\":\"rent-car\"
   }"
echo
echo
echo
echo
echo
echo "============ test task-listener ============"
echo
curl -s -X POST \
  "http://localhost:$port/add_listener" \
  -H "authorization: Bearer $dbiir_deke_token" \
  -H "content-type: application/json" \
  -d "{
       \"chaincode\":\"$cc_name\",
       \"taskName\":\"rent-car\"
   }"
echo
echo
echo
echo
echo
echo "============ test task-listener ============"
echo
curl -s -X POST \
  "http://localhost:$port/add_listener" \
  -H "authorization: Bearer $dbiir_deke_token" \
  -H "content-type: application/json" \
  -d "{
       \"chaincode\":\"$cc_name\",
       \"taskName\":\"scenic-spot\"
   }"
echo
echo
echo
echo
echo
echo "============ test task-listener ============"
echo
curl -s -X POST \
  "http://localhost:$port/add_listener" \
  -H "authorization: Bearer $qunar_gfe_token" \
  -H "content-type: application/json" \
  -d "{
       \"chaincode\":\"$cc_name\",
       \"taskName\":\"scenic-spot\"
   }"
echo
echo
echo

echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
