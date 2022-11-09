package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Config struct {
	startTimeout time.Duration
	stopTimeout  time.Duration
	maxThread    int
}

type StartTimeoutType time.Duration
type StopTimeoutType time.Duration

func NewConfig() Config {
	return Config{
		startTimeout: 10 * time.Second,
		stopTimeout:  10 * time.Second,
		maxThread:    runtime.NumCPU(),
	}
}

func StartTimeout(config Config) StartTimeoutType {
	return StartTimeoutType(config.startTimeout)
}

func StopTimeout(config Config) StopTimeoutType {
	return StopTimeoutType(config.stopTimeout)
}

func (c *Config) MaxThread() int {
	return c.maxThread
}

func (c *Config) MaxIdleConnsPerHost() int {
	return 10
}

func (c *Config) RequestTimeout() time.Duration {
	return 5 * time.Second
}

func (c *Config) WorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wd
}

func (c *Config) DatabaseFile() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(wd, "database.db")
}
