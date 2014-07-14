package main

import (
// "encoding/json"
// "io"
// "log"
// "os"
// "strconv"
// "strings"

// "github.com/influxdb/influxdb-go"
// "github.com/influxproxy/influxproxy/plugin"
)

func main() {
	// ...
}

type Plugin struct{}

func (p *Plugin) Ping() string {
	return "Pong"
}

func (p *Plugin) Run() string {
	return "Pong"
}

func (p *Plugin) Describe() string {
	return "Pong"
}
