package main

import (
    "time"
    "strconv"
    "os"
    "fmt"
)

func main() {
    if len(os.Args) == 1 || os.Args[1] == "-h" {
        fmt.Println("usage is:\nsleep <# of seconds to sleep>")
    } else {
        ssec, _ := strconv.Atoi(os.Args[1])
        time.Sleep(time.Duration(ssec) * time.Second)
    }
}
