package main

import (
	"log"

	"github.com/VishvaNavanjana/Docker-Go-REST-API/database"
	"github.com/gofiber/fiber/v2"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"

	// "github.com/gofiber/contrib/otelfiber"
)



const (
	service     = "traces-mvp-rest-api"
	environment = "production"
	id          = 1
)

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(

		tracesdk.WithSyncer(exp),

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




func main() {
	// Create a new TracerProvider that will send spans to the Jaeger
	// collector running on localhost.
	tp, err := tracerProvider("http://172.17.0.1:14268/api/traces")

	//handle err
	if err != nil {
		log.Fatal(err)
	}
	
	otel.SetTracerProvider(tp)


	database.ConnectDb()
	app := fiber.New()

	// use otel middleware for automatic instrumentation
	// app.Use(otelfiber.Middleware("traces-mvp-rest-api", otelfiber.WithTracerProvider(tp)))

	setupRoutes(app)

	app.Listen(":3000")
}