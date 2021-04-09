package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cfg := jaegercfg.Configuration{
			ServiceName: "go-jaeger",
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans:           true,
				LocalAgentHostPort: "127.0.0.1:6831",
			},
		}
		jLogger := jaegerlog.StdLogger
		tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jLogger),
		)
		if err != nil {
			panic(err)
		}
		defer closer.Close()
		//span :=opentracing.StartSpan("go-grpc")
		//time.Sleep(time.Second)
		//defer span.Finish()

		// 单次请求 响应时间
		startSpan := tracer.StartSpan(ctx.Request.URL.Path)
		defer startSpan.Finish()
		ctx.Set("tracer", tracer)
		ctx.Set("parentSpan", startSpan)
		ctx.Next()
	}
}
