package http_client

import (
	"fmt"

	"go.opentelemetry.io/otel/trace"
)

const (
	// W3CTraceparentHeader is the standard W3C Trace-Context HTTP
	// header for trace propagation.
	W3CTraceparentHeader = "traceparent"

	// TracestateHeader is the standard W3C Trace-Context HTTP header
	// for vendor-specific trace propagation.
	TracestateHeader = "tracestate"
)

func PopulateTraceparentHeadersFromOtelContext(ctx trace.SpanContext) (headers map[string]string) {
	headers = make(map[string]string)
	headers[W3CTraceparentHeader] = fmt.Sprintf("00-%s-%s-%s", ctx.TraceID().String(), ctx.SpanID().String(), ctx.TraceFlags().String())
	if tracestate := ctx.TraceState().String(); tracestate != "" {
		headers[TracestateHeader] = tracestate
	}

	return
}
