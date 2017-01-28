package fiolog

import (
  "testing"
  "reflect"
)

func TestAggregation(t *testing.T) {

  logs := LogMetrics{
    &LogMetric{
      Time:  0,
      Value: 200,
    },
    &LogMetric{
      Time:  10,
      Value: 100,
    },
    &LogMetric{
      Time:  30,
      Value: 20,
    },
    &LogMetric{
      Time:  40,
      Value: 500,
    },
    &LogMetric{
      Time:  50,
      Value: 120,
    },
  }

  logTests := []struct {
    aggregation Aggregation
    resultLog   LogMetric
  }{
    {
      AggregationAvg{},
      LogMetric{
        Time:  0,
        Value: 188,
      },
    },
    {
      AggregationMax{},
      LogMetric{
        Time:  40,
        Value: 500,
      },
    },
    {
      AggregationMin{},
      LogMetric{
        Time:  30,
        Value: 20,
      },
    },
  }

  for i, test := range logTests {
    result := test.aggregation.convert(&logs)

    if !reflect.DeepEqual(test.resultLog, result) {
      t.Fatalf("Test %d expected %v, got %v", i, test.resultLog, result)
    }
  }
}
