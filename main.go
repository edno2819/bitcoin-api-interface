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

	// Example usage:
	a.GetBlockchainInfo()
	a.EstimateRawFee(2)
	a.GetWalletBalance()
	a.CreateWallet("Wallet1")
	a.GetWalletBalance()
	a.ListDescriptors(false)
	a.DumpWallet("./my_wallet_key.txt")

}
