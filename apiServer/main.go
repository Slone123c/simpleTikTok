// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cold-runner/simpleTikTok/apiServer/cmd"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/spf13/viper"
)

func main() {
	cmd.Init()

	tracer, cfg := tracing.NewServerTracer()
	h := server.New(
		server.WithHostPorts(viper.GetString("etcd.api-service-address")),
		server.WithHandleMethodNotAllowed(true), // coordinate with NoMethod
		server.WithMaxRequestBodySize(10000000000),
		tracer,
	)
	log.Infow("Starting API-Server at :", "address",
		viper.GetString("etcd.api-service-addr"))
	h.Use(tracing.ServerMiddleware(cfg))
	register(h)
	h.Spin()
}
