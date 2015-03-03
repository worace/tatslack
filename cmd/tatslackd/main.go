package main

import (
	"flag"
	"github.com/worace/tatslack"
	"log"
)

func main() {
	log.SetFlags(0)

	token := flag.String("token", "", "slack API token")
	flag.Parse()

	if *token == "" {
		log.Fatal("token required")
	}

	c := tatslack.Client{
		Token: *token,
	}

	resp, err := c.ChannelHistory("C03HCH8RB")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MESSAGES")
	log.Println("========")
	for _, m := range resp.Messages {
		log.Println(m.Text)
	}
}
