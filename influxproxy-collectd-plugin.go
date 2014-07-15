package main

import (
	"time"
	"github.com/influxproxy/influxproxy/plugin"
)

func main() {
	p, _ := plugin.NewPlugin()
	p.Run()
	time.Sleep(30 * time.Second)
}
