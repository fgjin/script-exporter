package router

import (
	"script-exporter/logs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRouter(registry *prometheus.Registry, log *logs.Log) *gin.Engine {
	r := gin.New()
	//使用自定义日志中间件
	r.Use(MyLogger(log))
	r.Use(gin.Recovery())
	//访问/metrics会自动更新指标
	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(registry, promhttp.HandlerOpts{})))
	return r
}

// 自定义日志格式
func MyLogger(log *logs.Log) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		ctx.Next()
		mp := make(map[string]any)
		if raw != "" {
			path = path + "?" + raw
		}
		mp["Path"] = path
		mp["Method"] = ctx.Request.Method
		mp["ClientIP"] = ctx.ClientIP()
		mp["StatusCode"] = ctx.Writer.Status()
		mp["BodySize"] = ctx.Writer.Size()
		mp["Latency"] = time.Since(start)
		if len(ctx.Errors) > 0 || ctx.Writer.Status() >= 400 {
			mp["Errors"] = ctx.Errors.String()
			log.WithFields(mp).Error("Request processing error")
		} else {
			log.WithFields(mp).Info("Request processed successfully")
		}
	}
}
