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
		return nil, fmt.Errorf("erro ao fazer a requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler a resposta: %v", err)
	}

	// Deserializa a resposta JSON-RPC
	var rpcResponse RPCResponse
	if err := json.Unmarshal(body, &rpcResponse); err != nil {
		return nil, fmt.Errorf("erro ao desserializar resposta JSON-RPC: %v", err)
	}

	// verify if there is error in the response
	if rpcResponse.Error != nil {
		return nil, fmt.Errorf("erro do Bitcoin Core: %v", rpcResponse.Error.Message)
	}

	return &rpcResponse, nil
}

func formatPrint(text string, data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf("Erro ao formatar JSON: %v", err)
	}

	fmt.Printf("%s %v", text, string(jsonData))
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
		return nil, fmt.Errorf("erro ao serializar JSON-RPC: %v", err)
	}

	req, err := http.NewRequest("POST", rpcI.url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar a requisição HTTP: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(rpcI.rpcUser, rpcI.rpcPassword)

	return executeRequest(req)
}

func (rpcI *RPCInterface) GetBlockchainInfo() {
	response, err := rpcI.executeMethod("getblockchaininfo", []interface{}{})
	if err != nil {
		log.Fatalf("Erro ao chamar getblockchaininfo: %v", err)
	}

	formatPrint("Blockchain Info:", response.Result)
}

func (rpcI *RPCInterface) GetWalletBalance() {
	response, err := rpcI.executeMethod("getbalance", []interface{}{})
	if err != nil {
		log.Fatalf("Erro ao chamar getbalance: %v", err)
	}
	formatPrint("Blockchain Wallet:", response.Result)
}
