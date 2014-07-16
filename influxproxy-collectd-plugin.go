package main

import (
	"github.com/influxdb/influxdb-go"
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

func (f Functions) Run(in string) []influxdb.Series {
	out := new([]influxdb.Series)
	return *out
}

func main() {
	f := Functions{}
	p, _ := plugin.NewPlugin()
	p.Run(f)
}
