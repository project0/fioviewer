package fiolog

import (
  "os"
  "bufio"
  "log"
  "strings"
  "strconv"
  "math"
)

func parseLine(line string) (LogMetric, error) {
  lineSplit := strings.Split(line, ",")

  time, err := strconv.ParseUint(strings.TrimSpace(lineSplit[0]), 0, 32)
  if err != nil {
    return LogMetric{}, err
  }

  value, err := strconv.ParseUint(strings.TrimSpace(lineSplit[1]), 0, 32)
  if err != nil {
    return LogMetric{}, err
  }

  direction, err := strconv.ParseUint(strings.TrimSpace(lineSplit[2]), 0, 8)
  if err != nil {
    return LogMetric{}, err
  }

  offset, err := strconv.ParseUint(strings.TrimSpace(lineSplit[3]), 0, 32)
  if err != nil {
    return LogMetric{}, err
  }

  return LogMetric{
    Time:      uint(time),
    Value:     uint(value),
    Direction: uint8(direction),
    Offset:    uint(offset),
  }, nil
}

func readLastLine(fname string) (LogMetric, error) {

  file, err := os.Open(fname)
  if err != nil {
    return LogMetric{}, err
  }
  defer file.Close()

  var finalOffset int64
  var offset int64 = -2
  b := make([]byte, 1)

  // detect length of last line
  for string(b) != "\n" {
    // offset relative to end of file
    finalOffset, _ = file.Seek(offset, 2)
    _, err = file.Read(b)
    if err != nil {
      return LogMetric{}, err
    }
    offset--
  }

  finalOffset++
  b = make([]byte, (offset * -1) - 3)
  file.ReadAt(b, finalOffset)

  return parseLine(string(b))
}

func readFirstLine(fname string) (LogMetric, error) {
  file, err := os.Open(fname)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Scan()

  if err := scanner.Err(); err != nil {
    return LogMetric{}, err
  }
  return parseLine(scanner.Text())
}

// ParseLogFile scans a fio log file and returns the metrics grouped by the direction and sorted by time.
// If maxDataPoints is to 0, all metrics are returned and no aggregation will be processed, aggregation will not be used.
func ParseLogFile(fname string, aggregation Aggregation, maxDataPoints uint, timeRange Range) ([3]LogMetrics, Log, error) {

  // set end value of range to maximum value if not present
  if timeRange.End == 0 {
    timeRange.End = math.MaxUint64
  }

  // create for each direction
  byDirection := [3]LogMetrics{
    []*LogMetric{},
    []*LogMetric{},
    []*LogMetric{},
  }

  // tmp struct to hold start and end times
  durations := [3]Range{
    {
      Start: math.MaxUint64,
      End:   0,
    },
    {
      Start: math.MaxUint64,
      End:   0,
    },
    {
      Start: math.MaxUint64,
      End:   0,
    },
  }

  file, err := os.Open(fname)
  // usually file not found
  if err != nil {
    return byDirection, Log{}, err
  }
  defer file.Close()

  loginfo, err := getLogInfo(fname)
  if err != nil {
    return byDirection, loginfo, err
  }

  // each line contains an metric
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    if err := scanner.Err(); err != nil {
      return byDirection, loginfo, err
    }

    entry, err := parseLine(scanner.Text())
    if err != nil {
      return byDirection, loginfo, err
    }

    if entry.Time >= timeRange.Start && entry.Time <= timeRange.End {
      // update duration
      s := &durations[entry.Direction].Start
      e := &durations[entry.Direction].End
      if *s > entry.Time {
        *s = entry.Time
      }
      if *e < entry.Time {
        *e = entry.Time
      }

      byDirection[entry.Direction] = append(byDirection[entry.Direction], &entry)
    }
  }

  if maxDataPoints > 0 {
    for idx := range byDirection {
      metrics := uint(len(byDirection[idx]))

      // batch only if maxDataPoints is less than amount of available metrics
      if maxDataPoints > metrics {
        continue
      }

      // not enough points for aggregation?
      // if metrics / maxDataPoints < 2 {
      //   continue
      // }

      interval := (durations[idx].End - durations[idx].Start) / maxDataPoints

      // unfortunately the metric values are not stored sequential
      // put each metric in the associated point of the graph and aggregate later by point
      cacheAggregationData := make([]LogMetrics, maxDataPoints + 1)
      for i := range byDirection[idx] {
        metric := byDirection[idx][i]
        // destination index of metric in new array
        point := (metric.Time - durations[idx].Start) / interval
        cacheAggregationData[point] = append(cacheAggregationData[point], metric)
      }

      byDirection[idx] = make(LogMetrics, len(cacheAggregationData))
      for i := range cacheAggregationData {
        metric := &cacheAggregationData[i]

        // fix out of range if datapoint has no metrics
        if len(*metric) > 0 {
          newMetric := aggregation.convert(metric)
          // set newMetric time based on the point and interval
          newMetric.Time = uint(i) * interval
          byDirection[idx][i] = &newMetric
        }
      }
    }
  }
  return byDirection, loginfo, nil
}
