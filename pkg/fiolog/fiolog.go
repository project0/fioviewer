// Package fiolog manages all backend operation regarding the log files.
package fiolog

import (
  "regexp"
  "strings"
  "log"
)

// FioPath handler to the path of the log files
var FioPath FioPathHandler
var regexLogFiles *regexp.Regexp

// Register and prepare the usage to the log files path
func Register(path *string) {

  FioPath = FioPathHandler(*path)
  var err error

  // prepare regexLogFiles to find log files
  keys := make([]string, 0, len(logTypes))
  for k := range logTypes {
    keys = append(keys, k)
  }

  // precompile
  regexLogFiles, err = regexp.Compile("^(.*?)_(" + strings.Join(keys, "|") + ").?(\\d+|)\\.log$")
  if err != nil {
    log.Fatal(err)
  }
}
