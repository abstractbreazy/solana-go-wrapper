package solego

import (
	"context"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

// Client wraps Solana-sdk clients.
type Client struct {
	rpc      *rpc.Client
	wss      *ws.Client
	Accounts map[string]string
}

// New creates a new Solana-wrapper client.
func New(
	ctx context.Context,
	rpcEndpoint, wsEndpoint string,
) (*Client, error) {

	rpcClient, err := NewRPCClient(ctx, rpcEndpoint)
	if err != nil {
		return nil, err
	}

	wsClient, err := NewWSClient(ctx, wsEndpoint)
	if err != nil {
		return nil, err
	}

	return &Client{
		rpc: rpcClient,
		wss: wsClient,
	}, nil
}

// NewRPCClient creates a new Solana RPC client.
func NewRPCClient(
	ctx context.Context,
	rpcEndpont string,
) (*rpc.Client, error) {

	rpcClient := rpc.New(rpcEndpont)
	_, err := rpcClient.GetVersion(ctx)
	if err != nil {
		return nil, err
	}

	return rpcClient, nil
}

// NewWSClient creates a new Solana WS client.
func NewWSClient(
	ctx context.Context,
	wsEndpoint string,
) (*ws.Client, error) {

	wsClient, err := ws.Connect(ctx, wsEndpoint)
	if err != nil {
		return nil, err
	}

	return wsClient, nil
}
