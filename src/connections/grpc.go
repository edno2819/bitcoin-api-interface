package connections

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RPCRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RPCResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *RPCError       `json:"error"`
	ID     int             `json:"id"`
}

type RPCInterface struct {
	url            string
	rpcUser        string
	rpcPassword    string
	jsonRPCVersion string
	id             int
}

func executeRequest(req *http.Request) (*RPCResponse, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error in resquest HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %v", err)
	}

	// Deserializa a resposta JSON-RPC
	var rpcResponse RPCResponse
	if err := json.Unmarshal(body, &rpcResponse); err != nil {
		return nil, fmt.Errorf("Erro deserializing JSON-RPC response: %v", err)
	}

	// verify if there is error in the response
	if rpcResponse.Error != nil {
		return nil, fmt.Errorf("Bitcoin Core Error: %v", rpcResponse.Error.Message)
	}

	return &rpcResponse, nil
}

func formatPrint(text string, data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf("Erro while formatting JSON: %v\n\n", err)
		return
	}

	fmt.Printf("%s %v\n", text, string(jsonData))
}

func NewRPCInterface(host string, port string, rpcUser string, rpcPassword string, isHttps bool) *RPCInterface {
	method := "http"
	if isHttps {
		method = "https"
	}
	url := fmt.Sprintf("%v://%s:%s", method, host, port)

	return &RPCInterface{
		url,
		rpcUser,
		rpcPassword,
		"1.0",
		1,
	}
}

func (rpcI *RPCInterface) executeMethod(method string, params interface{}) (*RPCResponse, error) {
	rpcRequest := &RPCRequest{
		rpcI.jsonRPCVersion, method, params, rpcI.id,
	}
	requestBody, err := json.Marshal(rpcRequest)
	if err != nil {
		return nil, fmt.Errorf("Error while serializing JSON-RPC: %v", err)
	}

	req, err := http.NewRequest("POST", rpcI.url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("Error while creating HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(rpcI.rpcUser, rpcI.rpcPassword)

	return executeRequest(req)
}

func (rpcI *RPCInterface) GetBlockchainInfo() {
	response, err := rpcI.executeMethod("getblockchaininfo", []interface{}{})
	if err != nil {
		log.Fatalf("Error calling getblockchaininfo: %v", err)
	}

	formatPrint("Blockchain Info:", response.Result)
}

func (rpcI *RPCInterface) GetWalletBalance() {
	response, err := rpcI.executeMethod("getbalance", []interface{}{})
	if err != nil {
		fmt.Errorf("Error calling getbalance: %v", err)
		return
	}
	formatPrint("Blockchain Wallet:", response.Result)
}

func (rpcI *RPCInterface) CreateWallet(walletName string) {
	response, err := rpcI.executeMethod("createwallet", []interface{}{walletName})
	if err != nil {
		log.Fatalf("Error calling createwallet: %v", err)
		return
	}
	formatPrint("Wallet Created:", response.Result)
}

func (rpcI *RPCInterface) GetWalletInfo() {
	response, err := rpcI.executeMethod("getwalletinfo", []interface{}{})
	if err != nil {
		log.Fatalf("Error calling getwalletinfo: %v", err)
		return
	}
	formatPrint("Infos about Wallet:", response.Result)
}

func (rpcI *RPCInterface) DumpWallet(filePath string) {
	response, err := rpcI.executeMethod("dumpwallet", []interface{}{filePath})
	if err != nil {
		log.Fatalf("Error calling dumpwallet: %v", err)
		return
	}
	formatPrint("Wallet Dumped to:", response.Result)
}

func (rpcI *RPCInterface) EstimateRawFee(confTarget int) {
	response, err := rpcI.executeMethod("estimaterawfee", []interface{}{confTarget})
	if err != nil {
		log.Fatalf("Error calling estimaterawfee: %v", err)
		return
	}
	formatPrint("Estimate Raw Fee:", response.Result)
}

func (rpcI *RPCInterface) ListDescriptors(includePrivate bool) {
	response, err := rpcI.executeMethod("listdescriptors", []interface{}{includePrivate})
	if err != nil {
		log.Fatalf("Error calling listdescriptors: %v", err)
		return
	}
	formatPrint("List of Descriptors:", response.Result)
}
