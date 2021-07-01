const { GrpcInstrumentation } = require('@opentelemetry/instrumentation-grpc');
const { trace } = require("@opentelemetry/api");
const grpc = require('@grpc/grpc-js');
const { registerInstrumentations } = require('@opentelemetry/instrumentation');
const { NodeTracerProvider } = require('@opentelemetry/node')
const { HttpInstrumentation } = require('@opentelemetry/instrumentation-http');
const { ConsoleSpanExporter, BatchSpanProcessor } = require('@opentelemetry/tracing');
const { CollectorTraceExporter } = require('@opentelemetry/exporter-collector-grpc');


const metadata = new grpc.Metadata();
// For instance, an API key or access token might go here.
metadata.set('x-honeycomb-dataset', 'x');
metadata.set('x-honeycomb-team', 'x');

const provider = new NodeTracerProvider({ resource: { attributes: { "service.name": "test-node-otel" } } });
const exporter = new CollectorTraceExporter({
    url: 'https://api.honeycomb.io:443',
    credentials: grpc.credentials.createSsl(),
    metadata,
});
provider.addSpanProcessor(new BatchSpanProcessor(exporter));
provider.addSpanProcessor(new BatchSpanProcessor(new ConsoleSpanExporter()));
provider.register();

registerInstrumentations({
    instrumentations: [
        new HttpInstrumentation(),
        new GrpcInstrumentation(),
    ],
});

var express = require('express');
var app = express();


var services = require('./greet_grpc_pb');
var messages = require('./greet_pb.js');
var client = new services.GreetServiceClient("localhost:50051",
    grpc.credentials.createInsecure());
app.get('/', function (req, res) {
    var greetingRequest = new messages.GreetRequest();
    var greeting = new messages.Greeting();
    greeting.setFirstname('ricardo');
    greeting.setLastname('linck');
    greetingRequest.setGreeting(greeting);
    client.greet(greetingRequest, function (err, response) {
        console.log('Greeting:', response.getResult());
    });
    res.send('Hello World!');
});

app.listen(3000, function () {
    console.log('Example app listening on port 3000!');
});
