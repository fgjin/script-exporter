自定义exporter，用于收集脚本执行结果，需要统一将脚本执行结果写到yaml文件里
实现Collector接口--注册collector--/metrics暴露指标，时间、metrics请求触发指标更新

修改prometheus的配置
scrape_configs:
  - job_name: scripts
    static_configs:
      - targets: ["localhost:9090"]