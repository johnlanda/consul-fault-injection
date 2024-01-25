Kind = "proxy-defaults"
Name = "global"
Config = {
  protocol           = "http"
  envoy_prometheus_bind_addr = "0.0.0.0:20200"
  envoy_tracing_json = <<EOF
{
  "http":{
    "name":"envoy.tracers.zipkin",
    "typedConfig":{
      "@type":"type.googleapis.com/envoy.config.trace.v3.ZipkinConfig",
      "collector_cluster":"jaeger_9411",
      "collector_endpoint_version":"HTTP_JSON",
      "collector_endpoint":"/api/v2/spans",
      "shared_span_context":false
    }
  }
}
EOF

  envoy_extra_static_clusters_json = <<EOF
{
  "name":"jaeger_9411",
  "type":"STRICT_DNS",
  "connect_timeout":"5s",
  "load_assignment":{
    "cluster_name":"jaeger_9411",
    "endpoints":[
      {
        "lb_endpoints":[
          {
            "endpoint":{
              "address":{
                "socket_address":{
                  "address":"jaeger-collector.observability.svc.cluster.local",
                  "port_value":9411
                }
              }
            }
          }
        ]
      }
    ]
  }
}
EOF
}