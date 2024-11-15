package utils

import (
	"fmt"
	"time"
)

const (
	hostNodeBTC  = "1212.154.54"
	loginNodeBTC = ""
)

var (
	urlBTC = fmt.Sprintf("https://%s:%s@127.0.0.1:8332", loginNodeBTC, hostNodeBTC)
)

type RPCConfig struct {
	hostNodeBTC    string
	loginNodeBTC   string
	urlBTC         string
	ServerPingTime time.Duration `long:"server-ping-time" description:"How long the server waits on a gRPC stream with no activity before pinging the client."`
}

type Config struct {
	nodeType      string    
	walletVersion string    
	rpcConfig     RPCConfig
}

func DefaultConfig() *Config {
	return &Config{
		nodeType:      "mainnet",
		walletVersion: "2.0",
		rpcConfig : &RPCConfig{
			hostNodeBTC:    hostNodeBTC,
            loginNodeBTC:   loginNodeBTC,
            urlBTC:         urlBTC,
            ServerPingTime: 10 * time.Second,
		}
	}
}
