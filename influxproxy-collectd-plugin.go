package main

import (
	"github.com/influxproxy/influxproxy/plugin"
)

type Functions struct{}

func (f Functions) Ping() bool {
	return true
}

func (f Functions) Describe() plugin.Description {
	args := new([]plugin.Argument)

	arg := plugin.Argument{
		Name:        "test",
		Description: "test",
		Optional:    true,
	}

	*args = append(*args, arg)

	d := plugin.Description{
		Description: "This plugin works as endpoint for CollectD, feeding the posted data into the given InfluxDB",
		Author:      "githob.com/sontags",
		Version:     "0.1.0",
		Arguments:   *args,
	}
	return d
}

func (f Functions) Run(in []*interface{}) string {
	return "Run"
}

func main() {
	f := Functions{}
	p, _ := plugin.NewPlugin()
	p.Run(f)
}

type Description struct {
	Description string
	Author      string
	Version     string
	Arguments   []*Argument
}

type Argument struct {
	Name        string
	Description string
	Optional    bool
}
