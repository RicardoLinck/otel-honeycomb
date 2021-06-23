package otel

import (
	"context"
	"log"

	"google.golang.org/grpc/credentials"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
)

func RegisterTracer() {
	ctx := context.Background()

	// Create an OTLP exporter, passing in Honeycomb credentials as environment variables.
	exp, err := otlp.NewExporter(
		ctx,
		otlpgrpc.NewDriver(
			otlpgrpc.WithEndpoint("api.honeycomb.io:443"),
			otlpgrpc.WithHeaders(map[string]string{
				"x-honeycomb-team":    "x",
				"x-honeycomb-dataset": "x",
			}),
			otlpgrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, "")),
		),
	)

	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	// Create a new tracer provider with a batch span processor and the otlp exporter.
	// Add a resource attribute service.name that identifies the service in the Honeycomb UI.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(
			sdktrace.NewBatchSpanProcessor(exp),
		),
		sdktrace.WithResource(resource.NewWithAttributes(semconv.ServiceNameKey.String("test-go-otel"))),
	)

	// // Handle this error in a sensible manner where possible
	// defer func() { _ = tp.Shutdown(ctx) }()

	// Set the Tracer Provider and the W3C Trace Context propagator as globals
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)

	log.Println("Created the otlp exporter")

	// Create a tracer instance.
	// return otel.Tracer("example/honeycomb-go")
}
