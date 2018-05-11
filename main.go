package main

import (
	"log"
	"os"
	"runtime"

	"github.com/nats-io/go-nats"
)

// NOTE: Use tls scheme for TLS, e.g. nats-sub -s tls://demo.nats.io:4443 foo

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}

func main() {

	urls := os.Getenv("MY_NATSBOOTSTRAP")
	showTime := os.Getenv("MY_TIMESTAMP")
	subj := os.Getenv("MY_TOPIC")

	nc, err := nats.Connect(urls)
	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}

	i := 0

	nc.Subscribe(subj, func(msg *nats.Msg) {
		i += 1
		printMsg(msg, i)
	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]\n", subj)
	if showTime != "false" {
		log.SetFlags(log.LstdFlags)
	}

	runtime.Goexit()
}
