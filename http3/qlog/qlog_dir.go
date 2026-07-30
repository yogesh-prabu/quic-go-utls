package qlog

import (
	"context"

	"github.com/yogesh-prabu/quic-go-utls"
	"github.com/yogesh-prabu/quic-go-utls/qlog"
	"github.com/yogesh-prabu/quic-go-utls/qlogwriter"
)

const EventSchema = "urn:ietf:params:qlog:events:http3-12"

func DefaultConnectionTracer(ctx context.Context, isClient bool, connID quic.ConnectionID) qlogwriter.Trace {
	return qlog.DefaultConnectionTracerWithSchemas(ctx, isClient, connID, []string{qlog.EventSchema, EventSchema})
}
