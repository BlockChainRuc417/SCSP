#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

jq --version > /dev/null 2>&1
if [ $? -ne 0 ]; then
	echo "Please Install 'jq' https://stedolan.github.io/jq/ to execute this script"
	echo
	exit 1
fi

starttime=$(date +%s)

LANGUAGE=go
CC_SRC_PATH=github.com/task
CC_NAME=task

echo "Language is: $LANGUAGE"
echo "The source path of chaincode is: $CC_SRC_PATH"
echo "The name of chaincode is: $CC_NAME"
echo
echo
echo "POST request Enroll on Org1  ..."
echo
ADMIN_ORG1_TOKEN=$(curl -s -X POST \
  http://localhost:4000/admins \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'username=admin_cc_gfe&orgname=Gfe')
echo $ADMIN_ORG1_TOKEN
ADMIN_ORG1_TOKEN=$(echo $ADMIN_ORG1_TOKEN | jq ".token" | sed "s/\"//g")
echo $ADMIN_ORG1_TOKEN>../token/admin_org1.token
echo
echo "ORG1 token is $ADMIN_ORG1_TOKEN"
echo
echo
echo "POST request Enroll on Org2  ..."
echo
ADMIN_ORG2_TOKEN=$(curl -s -X POST \
  http://localhost:4000/admins \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'username=admin_cc_deke&orgname=Deke')
echo $ADMIN_ORG2_TOKEN
ADMIN_ORG2_TOKEN=$(echo $ADMIN_ORG2_TOKEN| jq ".token" | sed "s/\"//g")
echo $ADMIN_ORG2_TOKEN>../token/admin_org2.token
echo
echo "ORG1 token is $ADMIN_ORG1_TOKEN"
echo
sleep 3
echo
echo "POST request Create channel  ..."
echo
curl -s -X POST \
  http://localhost:4000/channels \
  -H "authorization: Bearer $ADMIN_ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"channelName":"softwarechannel",
	"channelConfigPath":"../artifacts/channel/softwarechannel.tx"
}'
echo
echo
sleep 3
echo "POST request Join channel on Org1"
echo
curl -s -X POST \
  http://localhost:4000/channels/softwarechannel/peers \
  -H "authorization: Bearer $ADMIN_ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.fabric.gfe.com","peer1.fabric.gfe.com"]
}'
echo
echo

echo "POST request Join channel on Org2"
echo
curl -s -X POST \
  http://localhost:4000/channels/softwarechannel/peers \
  -H "authorization: Bearer $ADMIN_ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.fabric.deke.com","peer1.fabric.deke.com"]
}'
echo
echo

echo "POST Install chaincode on Org1"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ADMIN_ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.fabric.gfe.com\",\"peer1.fabric.gfe.com\"],
	\"chaincodeName\":\"$CC_NAME\",
	\"chaincodePath\":\"$CC_SRC_PATH\",
	\"chaincodeType\": \"$LANGUAGE\",
	\"chaincodeVersion\":\"v0\"
}"
echo
echo

echo "POST Install chaincode on Org1"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ADMIN_ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.fabric.gfe.com\",\"peer1.fabric.gfe.com\"],
	\"chaincodeName\":\"currency\",
	\"chaincodePath\":\"github.com/currency\",
	\"chaincodeType\": \"$LANGUAGE\",
	\"chaincodeVersion\":\"v0\"
}"
echo
echo



echo "POST Install chaincode on Org2"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ADMIN_ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.fabric.deke.com\",\"peer1.fabric.deke.com\"],
	\"chaincodeName\":\"$CC_NAME\",
	\"chaincodePath\":\"$CC_SRC_PATH\",
	\"chaincodeType\": \"$LANGUAGE\",
	\"chaincodeVersion\":\"v0\"
}"
echo
echo

echo "POST Install chaincode on Org2"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ADMIN_ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.fabric.deke.com\",\"peer1.fabric.deke.com\"],
	\"chaincodeName\":\"currency\",
	\"chaincodePath\":\"github.com/currency\",
	\"chaincodeType\": \"$LANGUAGE\",
	\"chaincodeVersion\":\"v0\"
}"
echo
echo

echo "POST instantiate chaincode on peer1 of Org1"
echo
curl -s -X POST \
  http://localhost:4000/channels/softwarechannel/chaincodes \
  -H "authorization: Bearer $ADMIN_ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.fabric.gfe.com\"],
	\"chaincodeName\":\"currency\",
	\"chaincodeVersion\":\"v0\",
	\"chaincodeType\": \"$LANGUAGE\",
	\"args\":[]
}"
echo
echo

echo "POST instantiate chaincode on peer1 of Org1"
echo
curl -s -X POST \
  http://localhost:4000/channels/softwarechannel/chaincodes \
  -H "authorization: Bearer $ADMIN_ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.fabric.gfe.com\"],
	\"chaincodeName\":\"$CC_NAME\",
	\"chaincodeVersion\":\"v0\",
	\"chaincodeType\": \"$LANGUAGE\",
	\"args\":[]
}"
echo
echo

echo "GET query Installed chaincodes"
echo
curl -s -X GET \
  "http://localhost:4000/chaincodes?peer=peer0.fabric.gfe.com" \
  -H "authorization: Bearer $ADMIN_ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo

echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
