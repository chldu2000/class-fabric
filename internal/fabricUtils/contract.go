package fabricUtils

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"log"
	"os"
	"path/filepath"
)

func GetContract(userId, org, contractName string) (*gateway.Contract, error) {
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}
	ccpPath := filepath.Join(
		"..",
		"fabric-samples",
		"test-network-new",
		"organizations",
		"peerOrganizations",
		"org"+org+".medicine.com",
		"connection-org"+org+".yaml",
	)

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		return nil, err
	}
	walletLabel := userId + "@" + org
	log.Println("连接到网关……")
	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, walletLabel),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	log.Println("连接到通道……")
	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	log.Println("获取合约……")
	contract := network.GetContract(contractName)
	log.Println(contract.Name())
	return contract, nil
}
