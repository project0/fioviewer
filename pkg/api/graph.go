package api

import (
  "fmt"
  "net/http"

  gin "gopkg.in/gin-gonic/gin.v1"

  "github.com/project0/fioviewer/pkg/fiolog"
)

// RequestGraph is the body payload of an graph request.
type RequestGraph struct {
  Files             []string `json:"files" binding:"required"`
  MaxDataPoints     uint `json:"maxDataPoints"`
  AggregationMethod string `json:"aggregation"`
  Range             fiolog.Range `json:"range"`
}

// ResponseGraph returns all data for an graph/chart.
// ChartJS compatible.
type ResponseGraph struct {
  DataSets []ResponseGraphDataSet `json:"datasets"`
  Logs     []fiolog.Log `json:"logs"`
}

// ResponseGraphDataSet represent the dataset of metrics of an log file.
type ResponseGraphDataSet struct {
  Label string `json:"label"`
  Data  []*fiolog.LogMetric `json:"data"`
}

// Get graph data for various files.
func graphGet(c *gin.Context) {

  var request RequestGraph
  if c.BindJSON(&request) != nil {
    c.JSON(http.StatusBadRequest, gin.H{"status": "wrong format of request"})
    return
  }

  aggr, err := getAggregationMethod(request.AggregationMethod)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
    return
  }

  dataSets := []ResponseGraphDataSet{}
  logInfos := make([]fiolog.Log, len(request.Files))

  // Process each log file.
  for idx, file := range request.Files {
    metrics, logInfo, err := fiolog.ParseLogFile(file, aggr, request.MaxDataPoints, request.Range)

    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
    }

    logInfos[idx] = logInfo

    // Each logfile can supply up to three metric directions.
    for direction := range metrics {
      if len(metrics[direction]) == 0 {
        continue
      }

      dataSets = append(dataSets, ResponseGraphDataSet{
        // format label for the chart
        Label: logInfo.Name + " - " + fiolog.DirectionTypes[direction],
        Data:  metrics[direction],
      })
    }
  }
  c.JSON(http.StatusOK, ResponseGraph{DataSets: dataSets, Logs: logInfos})
}

// returns an aggregation method
func getAggregationMethod(method string) (fiolog.Aggregation, error) {
  switch method {
  case "max":
    return fiolog.AggregationMax{}, nil
  case "min":
    return fiolog.AggregationMin{}, nil
  case "":
    // if not present use avg
    fallthrough
  case "avg":
    return fiolog.AggregationAvg{}, nil
  default:
    return fiolog.AggregationAvg{}, fmt.Errorf("aggregation method '%s' is not supported", method)
  }
}
