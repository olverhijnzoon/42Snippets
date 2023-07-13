package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type myConfig struct {
	Setting1 string `json:"setting1"`
	Setting2 int    `json:"setting2"`
}

func watchConfig(cli *clientv3.Client, config *myConfig, done chan bool) {
	watchChan := cli.Watch(context.Background(), "config")
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			log.Printf("Config changed: %s\n", event.Kv.Value)
			err := json.Unmarshal(event.Kv.Value, config)
			if err != nil {
				log.Printf("Failed to unmarshal config: %s\n", err)
			} else {
				log.Printf("New config: %+v\n", config)
				done <- true
				return
			}
		}
	}
}

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## Golang ETCD")

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// Read the initial configuration from a JSON file
	configBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Set the initial configuration in etcd
	_, err = cli.Put(context.Background(), "config", string(configBytes))
	if err != nil {
		log.Fatal(err)
	}

	config := &myConfig{}
	err = json.Unmarshal(configBytes, config)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Initial config: %+v\n", config)

	done := make(chan bool)
	go watchConfig(cli, config, done)

	// Wait for a change to the configuration
	<-done
}
