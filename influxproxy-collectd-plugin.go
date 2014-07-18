package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/influxdb/influxdb-go"
	"github.com/influxproxy/influxproxy/plugin"
)

type Functions struct{}

func (f Functions) Describe() plugin.Description {
	d := plugin.Description{
		Description: "This plugin works as endpoint for CollectD, feeding the posted data into the given InfluxDB",
		Author:      "github.com/sontags",
		Version:     "0.1.0",
		Arguments: []plugin.Argument{
			{
				Name:        "prefix",
				Description: "Prefix of the series, will be separated with a '.' if given",
				Optional:    true,
			},
		},
	}
	return d
}

func (f Functions) Run(in plugin.Request) plugin.Response {

	var series []*influxdb.Series
	dec := json.NewDecoder(strings.NewReader(in.Body))
	for {
		var datasets []Dataset
		if err := dec.Decode(&datasets); err == io.EOF {
			break
		} else if err != nil {
			return plugin.Response{
				Series: nil,
				Error:  err.Error(),
			}
		}

		for _, ds := range datasets {
			series = append(series, ds.GetAsSeries(in.Query.Get("prefix"))...)
		}
	}

	return plugin.Response{
		Series: series,
		Error:  "",
	}
}

func main() {
	f := Functions{}
	p, err := plugin.NewPlugin()
	if err != nil {
		fmt.Println(err)
	} else {
		p.Run(f)
	}
}
