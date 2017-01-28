// Package api provides an REST service for the frontend of fioviewer.
package api

import (
  "net/http"

  gin "gopkg.in/gin-gonic/gin.v1"
)

// Register api routes/handler for frontend interactions.
func Register(r *gin.Engine) {

  r.Use(corsHeader)

  // List all available graphs.
  r.GET("/list", listGet)

  // Retrieve all data and metrics for an graph/chart.
  // RFC does not forbid request bodys on get.
  r.GET("/graph", graphGet)
  // However, most client-libraries does not support this -> support POST as well.
  r.POST("/graph", graphGet)
}

// Add Cross-origin headers.
func corsHeader(c *gin.Context) {

  c.Header("Access-Control-Allow-Origin", "*")
  c.Header("Access-Control-Allow-Methods", "*")
  c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

  // Answer on option methods.
  if c.Request.Method == http.MethodOptions {
    c.Status(http.StatusOK)
  }
}
