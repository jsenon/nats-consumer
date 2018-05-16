package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/nats-io/go-nats"
	"go.uber.org/zap"
)

// NOTE: Use tls scheme for TLS, e.g. nats-sub -s tls://demo.nats.io:4443 foo

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		logger.Error("Failed to create zap logger",
			zap.String("status", "ERROR"),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}

	urls := os.Getenv("MY_NATSBOOTSTRAP")
	showTime := os.Getenv("MY_TIMESTAMP")
	subj := os.Getenv("MY_TOPIC")
	queue := os.Getenv("MY_QUEUE")

	//nc, err = nats.Connect("tls://localhost:4443", nats.RootCAs("./configs/certs/ca.pem"))

	nc, err := nats.Connect(urls)
	if err != nil {
		logger.Error("Error nats connection:",
			zap.Error(err),
			zap.String("status", "ERROR"),
			zap.Duration("backoff", time.Second),
		)
	}
	defer nc.Close() // nolint: errcheck

	if queue != "" {
		i := 0

		_, err = nc.QueueSubscribe(subj, queue, func(msg *nats.Msg) {
			i++
			printMsg(msg, i)
		})
		if err != nil {
			logger.Error("Error nats subscription:",
				zap.Error(err),
				zap.String("status", "ERROR"),
			)
		}
		logger.Error("Subscription:",
			zap.String("topic", subj),
			zap.String("queue", queue),
			zap.Duration("backoff", time.Second),
		)
		err = nc.Flush()
		if err != nil {
			logger.Error("Error nats flush:",
				zap.Error(err),
				zap.String("status", "ERROR"),
				zap.Duration("backoff", time.Second),
			)
		}

		if err := nc.LastError(); err != nil {
			logger.Error("Error nats:",
				zap.Error(err),
				zap.String("status", "ERROR"),
				zap.Duration("backoff", time.Second),
			)
		}

		log.Printf("Listening on [%s]\n", subj)
		if showTime != "false" {
			log.SetFlags(log.LstdFlags)
		}

	} else {
		i := 0

		_, err = nc.Subscribe(subj, func(msg *nats.Msg) {
			i++
			printMsg(msg, i)
		})
		if err != nil {
			logger.Error("Error nats subscription:",
				zap.Error(err),
				zap.String("status", "ERROR"),
				zap.Duration("backoff", time.Second),
			)
		}
		err = nc.Flush()
		if err != nil {
			logger.Error("Error nats flush:",
				zap.Error(err),
				zap.String("status", "ERROR"),
				zap.Duration("backoff", time.Second),
			)
		}

		if err := nc.LastError(); err != nil {
			logger.Error("Error nats:",
				zap.Error(err),
				zap.String("status", "ERROR"),
				zap.Duration("backoff", time.Second),
			)
		}

		log.Printf("Listening on [%s]\n", subj)
		if showTime != "false" {
			log.SetFlags(log.LstdFlags)
		}

	}

	runtime.Goexit()
}
