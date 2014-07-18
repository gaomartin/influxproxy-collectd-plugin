package main

import (
	"github.com/influxdb/influxdb-go"
)

type Dataset struct {
	Values         []float64 `json:"values"`
	Dstypes        []string  `json:"dstypes"`
	Dsnames        []string  `json:"dsnames"`
	Time           float64   `json:"time"`
	Interval       float32   `json:"interval"`
	Host           string    `json:"host"`
	Plugin         string    `json:"plugin"`
	PluginInstance string    `json:"plugin_instance"`
	Type           string    `json:"type"`
	TypeInstance   string    `json:"type_instance"`
}

func (ds *Dataset) GetAsSeries(influxdbPrefix string) []*influxdb.Series {
	var series []*influxdb.Series

	name := ds.Host + "." + ds.Plugin
	if ds.PluginInstance != "" {
		name = name + "." + ds.PluginInstance
	}
	name = name + "." + ds.Type
	if ds.TypeInstance != "" {
		name = name + "." + ds.TypeInstance
	}

	if influxdbPrefix != "" {
		name = influxdbPrefix + "." + name
	}

	for index, z := range ds.Dstypes {
		if z == "counter" || z == "gauge" {
			var n = name + "." + ds.Dsnames[index]
			out := &influxdb.Series{
				Name:    n,
				Columns: []string{"time", "value"},
				Points: [][]interface{}{
					[]interface{}{ds.Time, ds.Values[index]},
				},
			}

			series = append(series, out)
		}
	}

	return series
}
