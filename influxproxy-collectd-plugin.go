package main

import (
	"time"

	"github.com/influxproxy/influxproxy/plugin"
)

type Functions struct{}

func (f Functions) Ping() bool {
	return true
}

func (f Functions) Describe() string {
	return "Describe"
}

func (f Functions) Run(in []*interface{}) string {
	return "Run"
}

func main() {
	f := Functions{}
	p, _ := plugin.NewPlugin()
	p.Run(f)
	time.Sleep(20 * time.Second)
}
