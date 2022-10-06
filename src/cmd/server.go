package cmd

import (
	"class_main_service/src/config"
	"class_main_service/src/middlewares"
	"class_main_service/src/routes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.uber.org/zap"
)

const (
	service     = "class"
	environment = "development"
	id          = 1
)

func server() *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.GetConfig()
			if err != nil {
				zap.S().Fatalln(err)
			}

			tp, err := tracerProvider("http://localhost:14268/api/traces")
			if err != nil {
				zap.S().Error(err)
			}

			otel.SetTracerProvider(tp)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			defer func(ctx context.Context) {
				if err := tp.Shutdown(ctx); err != nil {
					zap.S().Error(err)
				}
			}(ctx)

			r := gin.Default()
			r.Use(middlewares.Cors, middlewares.Recover)

			routes.Bootstrap(r, cfg)

			r.Run(":7000")
		},
	}
}

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}
