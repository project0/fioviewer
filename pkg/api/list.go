package api

import (
  "net/http"

  gin "gopkg.in/gin-gonic/gin.v1"

  "github.com/project0/fioviewer/pkg/fiolog"
)

// ResponseList is an array of available log files.
type ResponseList []fiolog.Log

// Get a list of logs.
func listGet(c *gin.Context) {
  var result ResponseList
  result, err := fiolog.FioPath.Scan()

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
    return
  }

  c.JSON(http.StatusOK, result)
}
