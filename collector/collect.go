package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

// 实现Collector接口
type MyCollector struct {
	desc       *prometheus.Desc
	labelValue []string
	resValue   []int
}

func NewCollector(labelValue []string, resValue []int) *MyCollector {
	return &MyCollector{
		desc: prometheus.NewDesc(
			"script_execute_res",
			"Check Scripts Result",
			[]string{"scriptName"},
			map[string]string{"custom": "script"},
		),
		labelValue: labelValue,
		resValue:   resValue,
	}
}

func (c *MyCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.desc
}

func (c *MyCollector) Collect(ch chan<- prometheus.Metric) {
	for k, v := range c.resValue {
		ch <- prometheus.MustNewConstMetric(c.desc, prometheus.GaugeValue, float64(v), c.labelValue[k])
	}
}

func (c *MyCollector) Update(labelValue []string, resValue []int) {
	c.labelValue = labelValue
	c.resValue = resValue
}
