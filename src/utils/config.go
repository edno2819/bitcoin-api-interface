package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type RPCConfig struct {
	Host        string
	Port        string
	RPCUser     string
	RPCPassword string
	IsHTTPS     bool
}

type Config struct {
	NodeType       string
	WalletVersion  string
	RPCConfig      RPCConfig
	ServerPingTime time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		NodeType:       "mainnet",
		WalletVersion:  "2.0",
		ServerPingTime: 10 * time.Second,
		RPCConfig: RPCConfig{
			Host:        "127.0.0.1",
			Port:        "8332",
			RPCUser:     "user",
			RPCPassword: "password",
			IsHTTPS:     false,
		},
	}
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	config := DefaultConfig()

	if value, exists := os.LookupEnv("NODE_TYPE"); exists && value != "" {
		config.NodeType = value
	}

	if value, exists := os.LookupEnv("WALLET_VERSION"); exists && value != "" {
		config.WalletVersion = value
	}

	if value, exists := os.LookupEnv("RPC_HOST"); exists && value != "" {
		config.RPCConfig.Host = value
	}

	if value, exists := os.LookupEnv("RPC_PORT"); exists && value != "" {
		config.RPCConfig.Port = value
	}

	if value, exists := os.LookupEnv("RPC_USER"); exists && value != "" {
		config.RPCConfig.RPCUser = value
	}

	if value, exists := os.LookupEnv("RPC_PASSWORD"); exists && value != "" {
		config.RPCConfig.RPCPassword = value
	}

	if value, exists := os.LookupEnv("RPC_IS_HTTPS"); exists && value != "" {
		isHTTPS, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		config.RPCConfig.IsHTTPS = isHTTPS
	}

	if value, exists := os.LookupEnv("SERVER_PING_TIME"); exists && value != "" {
		duration, err := time.ParseDuration(value)
		if err != nil {
			return nil, err
		}
		config.ServerPingTime = duration
	}

	return config, nil
}
