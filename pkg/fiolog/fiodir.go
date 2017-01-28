package fiolog

import (
  "os"
  "strings"
  "fmt"
  "path/filepath"
)

// FioPathHandler provides handy mechanism to fio log paths
type FioPathHandler string

// Scan for available fio log files
func (dir FioPathHandler) Scan() ([]Log, error) {
  handler := walkHandler{
    logs: &[]Log{},
  }
  err := filepath.Walk(dir.String(), handler.walk)
  return *handler.logs, err
}

// String returns the filepath of the logs
func (dir FioPathHandler) String() string {
  return string(dir)
}

type walkHandler struct {
  logs *[]Log
}

// walk implements the WalkFun of the filepath and verifies and adds fio logs
func (h walkHandler) walk(path string, f os.FileInfo, err error) error {
  if f.IsDir() {
    return nil
  }

  l, err := getLogInfo(path)

  // if log file of fio, append to list
  if err == nil {
    *h.logs = append(*h.logs, l)
  }

  return nil
}

func getLogInfo(path string) (Log, error) {
  match := regexLogFiles.FindStringSubmatch(filepath.Base(path))
  if len(match) <= 1 {
    return Log{}, fmt.Errorf("filename %s not matched pattern", filepath.Base(path))
  }

  // validate first and last line of the file
  firstEntry, err := readFirstLine(path)
  if err != nil {
    return Log{}, err
  }

  lastEntry, err := readLastLine(path)
  if err != nil {
    return Log{}, err
  }

  logType := logTypes[match[2]]

  // Beautify name for labels
  name := match[1]
  // remove path prefix of given fio filepath
  subPath := strings.Replace(path, FioPath.String(), "", 1)
  // remove filename
  subPath = strings.Replace(subPath, filepath.Base(path), "", 1)
  name = subPath + name

  // append job index if present
  if match[3] != "" {
    name += " (" + match[3] + ")"
  }
  name += " (" + logType.Type + ")"

  return Log{
    Filename: path,
    Name:     name,
    Type:     logType,
    Duration: Range{
      Start: firstEntry.Value,
      End:   lastEntry.Value,
    },
  }, nil
}
