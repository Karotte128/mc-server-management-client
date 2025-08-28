package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/karotte128/mcsmplib"
)

// RPCResponse represents a JSON-RPC response
type RPCResponse struct {
	JSONRPC string    `json:"jsonrpc"`
	Result  any       `json:"result"`
	Error   *RPCError `json:"error,omitempty"`
	ID      int       `json:"id"`
}

// RPCError represents a JSON-RPC error
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// RPCClient manages the WebSocket connection and JSON-RPC communication
type RPCClient struct {
	conn *websocket.Conn
}

// NewRPCClient creates a new JSON-RPC client
func NewRPCClient(url string) (*RPCClient, error) {
	// Establish WebSocket connection
	wsConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %v", url, err)
	}

	return &RPCClient{
		conn: wsConn,
	}, nil
}

// Close shuts down the WebSocket connection
func (c *RPCClient) Close() error {
	return c.conn.Close()
}

// Call makes a JSON-RPC request and waits for the response
func (c *RPCClient) Call(input_request mcsmplib.Request) (*RPCResponse, error) {
	// Set request ID
	// TODO: Generate unique request ID
	reqID := 1234

	// Prepare request
	request := map[string]any{
		"jsonrpc": "2.0",
		"method":  input_request.Method,
		"params":  input_request.Params,
		"id":      reqID,
	}

	// Marshal request to JSON
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	// Send request
	err = c.conn.WriteMessage(websocket.TextMessage, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	// Read response
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	// Unmarshal response
	var response RPCResponse
	err = json.Unmarshal(message, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	// Verify response ID matches our request
	if response.ID != reqID {
		return nil, fmt.Errorf("response ID mismatch: expected %d, got %d",
			reqID, response.ID)
	}

	// Check for errors
	if response.Error != nil {
		return nil, fmt.Errorf("RPC error %d: %s", response.Error.Code,
			response.Error.Message)
	}

	return &response, nil
}
