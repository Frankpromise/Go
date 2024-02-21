# Go CLI Application

## Import Statements:

```
package main

import (
   "bufio"
   "flag"
   "fmt"
   "log"
   "strings"
   "os"
)

```

1. main: Indicates that this is the `main` package.
2. import: Imports necessary packages from the standard library (flag, fmt, log, os, strings, bufio).

## Main Function:
```
func main() {
```

The entry point of the Go program.

## Command-Line Flags:
```
   entry := flag.String("Entry", "1", "log entry to filter for")
   flag.Parse()
```

1. Defines a command-line flag named "level" with a default value of "CRITICAL". This flag is used to filter log entries based on the 
   specified log level.
2. flag.Parse() parses the command-line arguments.

## Open and Read Log File:
```
f, err := os.Open("./log.txt")
   if err != nil {
      log.Fatal(err)
   }
   defer f.Close()
```

1. Opens the "log.txt" file for reading.
2. If there's an error (e.g., file not found), it logs the error and terminates the program.
3. defer f.Close() ensures that the file is closed when the function exits.

## Read and Filter Log Entries:

```
   bufReader := bufio.NewReader(f)
   for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
      if strings.Contains(line, *entry) {
         fmt.Println(line)
      }
   }
```

1. Creates a buffered reader (bufReader) for efficient reading from the file.
2. Iterates over each line in the file until the end is reached.
3. Checks if the line contains the specified log level (*level), and if so, prints the line to the console.
