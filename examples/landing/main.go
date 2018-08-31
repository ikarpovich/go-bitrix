package main

import (
	"github.com/ikarpovich/go-bitrix/client"
	"github.com/ikarpovich/go-bitrix/types/landing"
	"log"
)

func main() {
	c, err := client.NewEnvClientWithWebhookAuth()

	if err != nil {
		log.Fatalf("Can't create client: %s", err)
	}

	c.SetInsecureSSL(true)
	c.SetDebug(true)

	resp, err := c.LandingRepoRegister(&landing.RepoRegisterRequest{
		Code: "test_block",
		Fields: landing.BlockFields{
			Name: "Test Block",
			Content: "<div><It works!/div>",
			Sections: "menu",
		},
		Manifest: landing.BlockManifest{},
	})

	if err != nil {
		log.Fatalf("Request error: %s", err)
	}

	log.Print(resp)
}