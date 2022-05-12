# An example using Context

One very important factor of Context as we've seen is the ability to contain valuable data for
specific requets. The ability to correlate those events across a system is one of the principles
behind distributed tracing. To find these correlations, distributed systems need to be able to
collect, store, and transfer metadata referred to as context.

Using a context we can now identify this particular set of calls, for example the current span, and traces,
and some extra metadata we might want to add to it. This context is propagated and bundled accross services via HTTP headers. This is called **Propagation** - which if I get lucky and understand it, we'll continue to explore in a follow-up

# Sample application - Fibonnacci

This mini example demonstrates how you can instrument an application using Go Open Telemetry in the context of a single-user application, using the **context.Background()**. This are my annotations around the example provided.


## Traces

A trace is a Type of OpenTelemetry to identify work being done by a service. It is a record
of the connections between various participants processing a transaction, often through client/server
requests.

Each part of the work that service performs is represented by a **span**. Spans are defined with
relations to one another. The root span is the only span without a parent. All other spans have a
parent relationship to other spans in the Trace.

In our case, our Trace would look like this:

```
Run
├── Poll
└── Write
    └── Fibonacci
```

## Exporters
The SDK connects Telemetry from the OpenTelemetry API to exporters. Exporters are packages that allow telemetry data to be emitted somewhere - either to the console (which is what we’re doing here), or to a remote system or collector for further analysis and/or enrichment.

Famous exporters include:
* Prometheus
* Jaeger
* Zipkin


##  Resources
Since we need to identify from which service, or which service instance is the trace coming from, Otel, uses
a resource called... You guessed it: **Resource**. This represent the entity producing the telemetry.



# Study material
* https://opentelemetry.lightstep.com/core-concepts/context-propagation/
* https://opentelemetry.io/docs/instrumentation/go/getting-started/
