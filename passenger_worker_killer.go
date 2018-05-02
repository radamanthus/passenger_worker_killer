package main

import (
  "flag"
  "fmt"
  "log"
  "os/exec"
  "regexp"
)

func main() {
  var memoryLimit int
  var runMode string
  var testFilename string

  flag.IntVar(&memoryLimit, "limit", 500, "worker memory limit")
  flag.StringVar(&runMode, "mode", "dryrun", "run mode")
  flag.StringVar(&testFilename, "testFilename", "test.txt", "Test file")

  flag.Parse()

  // Review: https://bountify.co/golang-parse-stdout
  // also see: https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/
  // read input from test file if it was supplied
  if runMode == "test" && testFilename != "" {
    fmt.Println("Running in test mode.")
    fmt.Printf("Reading input from %s\n", testFilename)
  } else {

  // run passenger-memory-stats and parse the output of the command

  // run mode
  // dry run - do nothing, but print PIDs that will be terminated
  // live - actually terminate the PIDs

    out, err := exec.Command("passenger-memory-stats").Output()
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("Terminating workers that exceed the %dMB limit\n", memoryLimit)
    var rackFinder = regexp.MustCompile(`RackApp`)
    result := rackFinder.FindStringSubmatch(string(out))

    fmt.Println(result)
  }

}
