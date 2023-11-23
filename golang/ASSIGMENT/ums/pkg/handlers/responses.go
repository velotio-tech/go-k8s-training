package handlers

import (
	"ums/pkg/exception"

	"github.com/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

/*
Author: Arijit Nayak
*/
// serveRes structure error message for client side.
// It handles all the cases of an error case of any API request.
func (u *UserHandlerClient) serveRes(c *gin.Context, statusCode int, res map[string]interface{}, excp *exception.Exception, rootSpan ddtrace.Span) {
	c.JSON(statusCode, res)

	if statusCode >= 300 {
		rootSpan.Finish(tracer.WithError(excp.Err))
	}
}
