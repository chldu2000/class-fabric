#!/bin/bash
if [ $# != 3 ]
then
    echo "参数有误!"
    echo "正确的格式: ./createUser OrgName Username Password"
    echo "示例: ./createUser Consumer MyUserName MyPassword"
    echo "OrgName: Agency, Manufacturer, Supplier or Consumer"
    exit 1
elif [ $1 != "Agency" -a $1 != "Consumer" -a $1 != "Manufacturer" -a $1 != "Supplier" ]
then
    echo "组织名称只能为 Agency, Manufacturer, Supplier 或 Consumer"
fi

export PATH=${HOME}/go/src/github.com/hyperledger/fabric-samples/bin:$PATH
portNum="7"
if [ ${1} == "Agency" ]
then
    portNum="5"
elif [ ${1} == "Consumer" ]
then
    portNum="6"
elif [ ${1} == "Manufacturer" ]
then
    portNum="7"
elif [ ${1} == "Supplier" ]
then
    portNum="8"
fi
echo $portNum

mkdir -p organizations/peerOrganizations/org${1}.medicine.com/

export FABRIC_CA_CLIENT_HOME=${HOME}/go/src/github.com/hyperledger/fabric-samples/test-network-new/organizations/peerOrganizations/org${1}.medicine.com/

set -x
fabric-ca-client enroll -u https://admin:adminpw@localhost:${portNum}054 --caname ca-org${1} --tls.certfiles "${HOME}/go/src/github.com/hyperledger/fabric-samples/test-network-new/organizations/peerOrganizations/org${1}.medicine.com/tlsca/tlsca.org${1}.medicine.com-cert.pem"
{ set +x; } 2>/dev/null

echo "NodeOUs:
Enable: true
ClientOUIdentifier:
  Certificate: cacerts/localhost-${portNum}054-ca-org${1}.pem  OrganizationalUnitIdentifier: client
PeerOUIdentifier:
  Certificate: cacerts/localhost-${portNum}054-ca-org${1}.pem
  OrganizationalUnitIdentifier: peer
AdminOUIdentifier:
  Certificate: cacerts/localhost-${portNum}054-ca-org${1}.pem
  OrganizationalUnitIdentifier: admin
OrdererOUIdentifier:
  Certificate: cacerts/localhost-${portNum}054-ca-org${1}.pem
  OrganizationalUnitIdentifier: orderer" > "${HOME}/go/src/github.com/hyperledger/fabric-samples/test-network-new/organizations/peerOrganizations/org${1}.medicine.com/msp/config.yaml"

set -x
fabric-ca-client register --caname ca-org${1} --id.name ${2} --id.secret ${3} --id.type client --tls.certfiles "${HOME}/go/src/github.com/hyperledger/fabric-samples/test-network-new/organizations/peerOrganizations/org${1}.medicine.com/tlsca/tlsca.org${1}.medicine.com-cert.pem"
{ set +x; } 2>/dev/null

set -x
fabric-ca-client enroll -u https://${2}:${3}@localhost:${portNum}054 --caname ca-org${1} -M "${HOME}/go/src/github.com/hyperledger/fabric-samples/test-network-new/organizations/peerOrganizations/org${1}.medicine.com/users/${2}@org${1}.medicine.com/msp" --tls.certfiles "${HOME}/go/src/github.com/hyperledger/fabric-samples/test-network-new/organizations/peerOrganizations/org${1}.medicine.com/tlsca/tlsca.org${1}.medicine.com-cert.pem"
{ set +x; } 2>/dev/null

cp "${HOME}/go/src/github.com/hyperledger/fabric-samples/test-network-new/organizations/peerOrganizations/org${1}.medicine.com/msp/config.yaml" "${HOME}/go/src/github.com/hyperledger/fabric-samples/test-network-new/organizations/peerOrganizations/org${1}.medicine.com/users/${2}@org${1}.medicine.com/msp/config.yaml"
