package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RPCRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

type RPCResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *RPCError       `json:"error"`
	ID     int             `json:"id"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func callBitcoinRPC(method string, params interface{}) (*RPCResponse, error) {
	url := "http://127.0.0.1:18332"
	rpcUser := "seuusuário"
	rpcPassword := "suasenha"

	rpcRequest := RPCRequest{
		Jsonrpc: "1.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}
	requestBody, err := json.Marshal(rpcRequest)
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar JSON-RPC: %v", err)
	}

	// Cria uma nova solicitação HTTP
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar a requisição HTTP: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(rpcUser, rpcPassword)

	// Executa a requisição HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Lê e interpreta a resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler a resposta: %v", err)
	}

	// Deserializa a resposta JSON-RPC
	var rpcResponse RPCResponse
	if err := json.Unmarshal(body, &rpcResponse); err != nil {
		return nil, fmt.Errorf("erro ao desserializar resposta JSON-RPC: %v", err)
	}

	// Verifica se houve erro no resultado JSON-RPC
	if rpcResponse.Error != nil {
		return nil, fmt.Errorf("erro do Bitcoin Core: %v", rpcResponse.Error.Message)
	}

	return &rpcResponse, nil
}

func getBlockchainInfo() {
	response, err := callBitcoinRPC("getblockchaininfo", []interface{}{})
	if err != nil {
		log.Fatalf("Erro ao chamar getblockchaininfo: %v", err)
	}

	fmt.Println("Blockchain Info:", string(response.Result))
}

func main() {
	getBlockchainInfo()
}
