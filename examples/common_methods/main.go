package main

import (
	"github.com/ikarpovich/go-bitrix/client"
	"github.com/ikarpovich/go-bitrix/types"
	"log"
)

func main() {
	c, err := client.NewEnvClientWithOauth()

	if err != nil {
		log.Fatalf("Can't create client: %s", err)
	}

	resp, err := c.Methods(&types.MethodsRequest{})

	if err != nil {
		log.Fatalf("Request error: %s", err)
	}

	log.Print(string(resp.Body()))

}