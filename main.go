package main

import (
	"log"
	"script-exporter/cmd"
	"script-exporter/collector"
	"script-exporter/global"
	"script-exporter/logs"
	"script-exporter/pkg"
	"script-exporter/router"
	"time"
)

var Log *logs.Log

func init() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("cmd.Execute error:%v", err)
	}
	conf := logs.LogConf{
		Level:       global.Level,
		AdapterName: global.AdapterName,
	}
	Log = logs.InitLog(conf)
}

func main() {
	// 额外的指标
	// global.Registry.MustRegister(prometheus.NewGoCollector())
	// global.Registry.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))

	// 定时更新指标
	go func() {
		// 创建并注册collector
		myCollector := collector.NewCollector([]string{}, []int{})
		global.Registry.MustRegister(myCollector)

		for {
			// 读取配置更新指标
			resKey, resValue := pkg.ReadFile(global.Respath)
			myCollector.Update(resKey, resValue)
			// fmt.Printf("%v  flush metrics\n", time.Now().Format("2006-01-02 15:04:05"))
			// Log.Info("Flush Metrics")
			time.Sleep(time.Duration(global.Cost) * time.Second)
		}
	}()

	// http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	// http.ListenAndServe(":8080", nil)
	//暴露指标
	r := router.SetupRouter(global.Registry, Log)
	if err := r.Run(global.Port); err != nil {
		log.Fatalf("failed to start the server:%v", err)
	}
}
