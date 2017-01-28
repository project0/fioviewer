package fiolog

const (
  logBandwith          string = "bw"
  logLatency           string = "lat"
  logCompletionLatency string = "clat"
  logSubmissionLatency string = "slat"
  logIops              string = "iops"
)

// Log represent a fio log file
type Log struct {
  Filename string `json:"filename"`
  Name     string  `json:"name"`
  Type     LogType `json:"type"`
  Duration Range `json:"duration"`
  Interval uint `json:"interval"`
}

// Range sets the start and end duration.
// Metrics which are within this time range will be added to resulting graph.
type Range struct {
  Start uint `json:"start"`
  End   uint `json:"end"`
}

// LogMetrics is a collection of metrics
type LogMetrics []*LogMetric

// LogMetric is a metric of an fio log file
type LogMetric struct {
  Time      uint `json:"x"`
  Value     uint `json:"y"`
  //Direction can be 0,1 or 2 - see DirectionTypes
  Direction uint8 `json:"-"`
  Offset    uint `json:"-"`
}

// LogType provides some additional info about the log file
type LogType struct {
  Type      string `json:"type"`
  TypeGroup string `json:"group"`
  Name      string `json:"name"`
  Unit      string `json:"unit"`
}

var logTypes = map[string]LogType{
  logBandwith: LogType{
    Type:        logBandwith,
    TypeGroup:   logBandwith,
    Name:        "Bandwith",
    Unit:        "KiB/s",
  },
  logLatency: LogType{
    Type:        logLatency,
    TypeGroup:   logLatency,
    Name:        "Latency",
    Unit:        "ms",
  },
  logCompletionLatency: LogType{
    Type:        logCompletionLatency,
    TypeGroup:   logLatency,
    Name:        "Completion Latency",
    Unit:        "ms",
  },
  logSubmissionLatency: LogType{
    Type:        logSubmissionLatency,
    TypeGroup:   logLatency,
    Name:        "Submission Latency",
    Unit:        "ms",
  },
  logIops: LogType{
    Type:        logIops,
    TypeGroup:   logIops,
    Name:        "IOPS",
    Unit:        "iops",
  },
}

// DirectionTypes represent the string of an LogMetric direction
var DirectionTypes = map[int]string{
  0: "READ",
  1: "WRITE",
  2: "TRIM",
}
