package main

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
)

type SimpleHandler struct{}

var logger *slog.Logger

func (h *SimpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received %s request for %s", r.Method, r.URL.String())

	if r.Method == http.MethodConnect {
		handleConnect(w, r)
	} else {
		handleHTTP(w, r)
	}
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		logger.Error("Error during RoundTrip: %v", err)
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	logger.Info("Response Status: %s", resp.Status)
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func handleConnect(w http.ResponseWriter, r *http.Request) {
	logger.Info("Handling CONNECT method for %s", r.URL.String())
	targetConn, err := net.Dial("tcp", r.Host)
	if err != nil {
		logger.Error("Error during CONNECT Dial: %v", err)
		http.Error(w, "Server failed to connect", http.StatusServiceUnavailable)
		return
	}
	logger.Info("Successfully dialed target %s", r.Host)

	// Inform the client that a connection has been established
	w.WriteHeader(http.StatusOK)

	clientConn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		logger.Error("Error during Hijack: %v", err)
		targetConn.Close() // Close the connection to the target
		return
	}

	// Start tunneling: Copy data between the client and target in both directions
	go func() {
		defer targetConn.Close()
		defer clientConn.Close()
		_, err := io.Copy(targetConn, clientConn)
		if err != nil {
			logger.Error("Error copying from client to target: %v", err)
		}
	}()

	go func() {
		defer clientConn.Close()
		defer targetConn.Close()
		io.Copy(clientConn, targetConn)
	}()
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func main() {
	fmt.Println("# 42Snippets")
	fmt.Println("## Golang Proxy")

	/*
		This code is a simple HTTP and HTTPS proxy written in Go. Upon receiving a request, the proxy logs the details and checks if the request method is CONNECT (for HTTPS) or a standard HTTP request. For HTTP requests, the proxy forwards them using the http.DefaultTransport and returns the response to the client. For HTTPS requests, it establishes a TCP connection with the target server and tunnels the client's connection to the server, allowing them to communicate directly.
	*/

	handlerOptions := slog.HandlerOptions{AddSource: true}
	logger = slog.New(slog.NewTextHandler(os.Stdout, &handlerOptions))

	server := &http.Server{
		Addr:    ":8080",
		Handler: &SimpleHandler{},
	}

	logger.Info("Starting proxy on :8080...")
	err := server.ListenAndServe()
	if err != nil {
		logger.Error("Error starting proxy: %v", err)
	}
}
