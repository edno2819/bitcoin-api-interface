package main

import (
	"bitcoin-api-interface/src/connections"
	"bitcoin-api-interface/src/utils"
	"fmt"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		fmt.Errorf("Error loading config:")
		panic(err)
	}
	rpcConfig := &config.RPCConfig
	a := connections.NewRPCInterface(rpcConfig.Host, rpcConfig.Port, rpcConfig.RPCUser, rpcConfig.RPCPassword, rpcConfig.IsHTTPS)
	a.GetBlockchainInfo()
	a.GetWalletBalance()
}
