package fiolog

// Aggregation aggregates multiple metrics to a new one.
type Aggregation interface {
  // convert multiple metrics to a new one
  convert(*LogMetrics) LogMetric
}

// AggregationMin implements Aggregation.
// Selects the lowest value.
type AggregationMin struct{}

func (a AggregationMin) convert(logs *LogMetrics) LogMetric {
  newLog := (*logs)[0]

  for _, e := range *logs {
    if newLog.Value > e.Value {
      newLog = e
    }
  }

  return *newLog
}

// AggregationMax implements Aggregation.
// Selects the highest value.
type AggregationMax struct{}

func (a AggregationMax) convert(logs *LogMetrics) LogMetric {
  newLog := (*logs)[0]

  for _, e := range *logs {
    if newLog.Value < e.Value {
      newLog = e
    }
  }

  return *newLog
}

// AggregationAvg implements Aggregation.
// Calculate an average value over all metrics.
type AggregationAvg struct{}

func (a AggregationAvg) convert(logs *LogMetrics) LogMetric {

  var totalVal uint
  newLog := (*logs)[0]

  for _, e := range *logs {
    totalVal += e.Value
  }

  newLog.Value = totalVal / uint(len(*logs))
  return *newLog
}
