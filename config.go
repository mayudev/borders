package main

import "fmt"

type Config struct {
	Port   uint   `env:"SERVE_PORT" default:"8080"`
	DBFile string `env:"DB_FILE" default:"./borders.db"`
}

func (c *Config) Addr() string {
	return fmt.Sprintf(":%d", c.Port)
}
