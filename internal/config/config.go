package config

import (
	"flag"
	"time"
)

type Config struct {
	Producers   int
	Consumers   int
	RPS         int
	Duration    time.Duration
	DSN         string
	ChannelSize int
}

func Load() Config {
	producers := flag.Int("producers", 1, "Number of producers")
	consumers := flag.Int("consumers", 1, "Number of consumers")
	rps := flag.Int("rps", 10, "Messages per second")
	duration := flag.Int("duration", 10, "Run duration in seconds")
	flag.Parse()

	return Config{
		Producers:   *producers,
		Consumers:   *consumers,
		RPS:         *rps,
		Duration:    time.Duration(*duration) * time.Second,
		DSN:         "root:root@123@tcp(localhost:3306)/producer_consumer", // Hardcoded for local dev
		ChannelSize: 1000,
	}
}
