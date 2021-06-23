module github.com/RicardoLinck/otel-honeycomb

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.20.0
	go.opentelemetry.io/otel v0.20.0
	go.opentelemetry.io/otel/exporters/otlp v0.20.0
	go.opentelemetry.io/otel/sdk v0.20.0
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)
