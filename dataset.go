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

// [
// {"values":[0,0.01,0.05],"dstypes":["gauge","gauge","gauge"],"dsnames":["shortterm","midterm","longterm"],"time":1405941780.747,"interval":10.000,"host":"raspi","plugin":"load","plugin_instance":"","type":"load","type_instance":""},
// {"values":[4.41139e+07],"dstypes":["gauge"],"dsnames":["value"],"time":1405941780.763,"interval":10.000,"host":"raspi","plugin":"memory","plugin_instance":"","type":"memory","type_instance":"used"},
// {"values":[5.21093e+07],"dstypes":["gauge"],"dsnames":["value"],"time":1405941780.770,"interval":10.000,"host":"raspi","plugin":"memory","plugin_instance":"","type":"memory","type_instance":"buffered"},
// {"values":[6.04897e+07],"dstypes":["gauge"],"dsnames":["value"],"time":1405941780.770,"interval":10.000,"host":"raspi","plugin":"memory","plugin_instance":"","type":"memory","type_instance":"cached"},
// {"values":[3.02793e+08],"dstypes":["gauge"],"dsnames":["value"],"time":1405941780.770,"interval":10.000,"host":"raspi","plugin":"memory","plugin_instance":"","type":"memory","type_instance":"free"}
// ]

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
