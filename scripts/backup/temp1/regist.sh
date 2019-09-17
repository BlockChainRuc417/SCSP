
port=4000
cc_name=currency
#org_token=$(cat org1.token)

starttime=$(date +%s)

echo "export port=$port;"
#echo "export org_token=$org_token;"

echo
echo "================ regist ================="
echo
JIM_ORG1_TOKEN=$(curl -s -X POST \
  "http://localhost:$port/users" \
  -H "content-type: application/json" \
  -d "{
		   \"channel\":\"softwarechannel\",
		   \"peer\":\"peer0.fabric.gfe.com\",
       \"userName\":\"Jim\",
	     \"orgName\":\"Gfe\",
		   \"description\":\"travel agent\"
     }");
echo $JIM_ORG1_TOKEN
JIM_ORG1_TOKEN=$(echo $JIM_ORG1_TOKEN | jq ".token" | sed "s/\"//g")
echo $JIM_ORG1_TOKEN>../token/jim_org1.token
echo
echo
echo "================ regist ================="
echo
TIM_ORG2_TOKEN=$(curl -s -X POST \
  "http://localhost:$port/users" \
  -H "content-type: application/json" \
  -d "{
		   \"channel\":\"softwarechannel\",
		   \"peer\":\"peer0.fabric.deke.com\",
       \"userName\":\"Tim\",
	     \"orgName\":\"Deke\",
		   \"description\":\"travel agent\"
     }");
echo $TIM_ORG2_TOKEN
TIM_ORG2_TOKEN=$(echo $TIM_ORG2_TOKEN | jq ".token" | sed "s/\"//g")
echo $TIM_ORG2_TOKEN>../token/tim_org2.token
echo
echo
echo "================ regist ================="
echo
PERRY_ORG2_TOKEN=$(curl -s -X POST \
  "http://localhost:$port/users" \
  -H "content-type: application/json" \
  -d "{
		   \"channel\":\"softwarechannel\",
		   \"peer\":\"peer1.fabric.deke.com\",
       \"userName\":\"Perry\",
	     \"orgName\":\"Deke\",
		   \"description\":\"travel agent\"
     }");
echo $PERRY_ORG2_TOKEN
PERRY_ORG2_TOKEN=$(echo $PERRY_ORG2_TOKEN | jq ".token" | sed "s/\"//g")
echo $PERRY_ORG2_TOKEN>../token/perry_org2.token
echo
echo
echo "================ regist ================="
echo
QUNAR_ORG1_TOKEN=$(curl -s -X POST \
  "http://localhost:$port/users" \
  -H "content-type: application/json" \
  -d "{
		   \"channel\":\"softwarechannel\",
		   \"peer\":\"peer0.fabric.gfe.com\",
       \"userName\":\"Qunar\",
	     \"orgName\":\"Gfe\",
		   \"description\":\"qunar service\"
     }");
echo $QUNAR_ORG1_TOKEN
QUNAR_ORG1_TOKEN=$(echo $QUNAR_ORG1_TOKEN | jq ".token" | sed "s/\"//g")
echo $QUNAR_ORG1_TOKEN>../token/qunar_org1.token
echo
echo
echo "================ regist ================="
echo
CTRIP_ORG2_TOKEN=$(curl -s -X POST \
  "http://localhost:$port/users" \
  -H "content-type: application/json" \
  -d "{
		   \"channel\":\"softwarechannel\",
		   \"peer\":\"peer1.fabric.deke.com\",
       \"userName\":\"Ctrip\",
	     \"orgName\":\"Deke\",
		   \"description\":\"ctrip service\"
     }");
echo $CTRIP_ORG2_TOKEN
CTRIP_ORG2_TOKEN=$(echo $CTRIP_ORG2_TOKEN | jq ".token" | sed "s/\"//g")
echo $CTRIP_ORG2_TOKEN>../token/ctrip_org2.token
echo
echo
echo
echo "================ regist ================="
echo
ALI_ORG2_TOKEN=$(curl -s -X POST \
  "http://localhost:$port/users" \
  -H "content-type: application/json" \
  -d "{
		   \"channel\":\"softwarechannel\",
		   \"peer\":\"peer0.fabric.deke.com\",
       \"userName\":\"Ali\",
	     \"orgName\":\"Deke\",
		   \"description\":\"ali service\"
     }");
echo $ALI_ORG2_TOKEN
ALI_ORG2_TOKEN=$(echo $ALI_ORG2_TOKEN | jq ".token" | sed "s/\"//g")
echo $ALI_ORG2_TOKEN>../token/ali_org2.token
echo
echo
echo "================ regist ================="
echo
DBIIR_ORG2_TOKEN=$(curl -s -X POST \
  "http://localhost:$port/users" \
  -H "content-type: application/json" \
  -d "{
		   \"channel\":\"softwarechannel\",
		   \"peer\":\"peer1.fabric.deke.com\",
       \"userName\":\"Dbiir\",
	     \"orgName\":\"Deke\",
		   \"description\":\"dbiir service\"
     }");
echo $DBIIR_ORG2_TOKEN
DBIIR_ORG2_TOKEN=$(echo $DBIIR_ORG2_TOKEN | jq ".token" | sed "s/\"//g")
echo $DBIIR_ORG2_TOKEN>../token/dbiir_org2.token
echo
echo


echo
echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
