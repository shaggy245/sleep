package main

import (
    "time"
    "strconv"
    "os"
)

func main() {
    ssec, _ := strconv.Atoi(os.Args[1])
    time.Sleep(time.Duration(ssec) * time.Second)
}
