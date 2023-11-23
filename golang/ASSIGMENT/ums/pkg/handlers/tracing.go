package handlers

import (
	"ums/pkg/constants"

	"github.com/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

/*
Author: Arijit Nayak
*/
// TraceEveryRequest traces all the requests received in the application level to the specific
// monitor system.
func (u *UserHandlerClient) traceEveryRequest(c *gin.Context) tracer.Span {
	reqSctx, traceErr := tracer.Extract(tracer.HTTPHeadersCarrier(c.Request.Header))
	var rootSpan tracer.Span
	if traceErr != nil {
		rootSpan, _ = tracer.StartSpanFromContext(c.Request.Context(), "http.request", tracer.ServiceName(constants.SERVICE_NAME), tracer.ResourceName(c.Request.URL.Path))
	} else {
		rootSpan, _ = tracer.StartSpanFromContext(c.Request.Context(), "http.request", tracer.ServiceName(constants.SERVICE_NAME), tracer.ResourceName(c.Request.URL.Path), tracer.ChildOf(reqSctx))
	}
	if err := c.Request.ParseForm(); err == nil {
		for k, v := range c.Request.Form {
			rootSpan.SetTag(k, v)
		}
	}
	return rootSpan
}
