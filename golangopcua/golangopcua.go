package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## Golang OPCUA")

	endpoint := "opc.tcp://localhost:32541"

	client, err := opcua.NewClient(
		endpoint,
		opcua.SecurityMode(ua.MessageSecurityModeSignAndEncrypt),
		opcua.SecurityPolicy("Basic256Sha256"),
		//opcua.AuthAnonymous(),
		opcua.AuthUsername("test123", "test123"),
	)
	slog.Info("client", client)
	if err != nil {
		slog.Error("err", err)
	}

	// Connect to the server
	if err := client.Connect(context.Background()); err != nil {
		slog.Error("err", err)
	}

	// Read a value from the server
	nodeID := ua.NewNumericNodeID(0, 2258)
	value, err := client.Node(nodeID).Value(context.Background())
	if err != nil {
		slog.Error("err", err)
	}

	fmt.Printf("Value: %v\n", value)

	// Close the connection
	client.Close(context.Background())
}
