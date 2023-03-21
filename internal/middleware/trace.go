package middleware

import (
	"github.com/gin-gonic/gin"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// Trace 拦截器
func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//jaeger配置
		cfg := jaegercfg.Configuration{
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1, //全部采样
			},
			Reporter: &jaegercfg.ReporterConfig{
				//当span发送到服务器时要不要打日志
				LogSpans:           true,
				LocalAgentHostPort: "192.168.10.130:6831",
			},
			ServiceName: "gin-grpc",
		}
		//创建jaeger
		tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
		if err != nil {
			panic(err)
		}
		defer closer.Close()
		//最开始的span，以url开始
		startSpan := tracer.StartSpan(ctx.Request.URL.Path)
		defer startSpan.Finish()
		//将tradcer和span存放到gin.context中
		ctx.Set("tracer", tracer)
		ctx.Set("parentSpan", startSpan)
		ctx.Next()
	}
}
