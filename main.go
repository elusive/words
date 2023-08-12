package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
)

func main() {
    // define boolean flag -l to count lines instead of words
    lines := flag.Bool("l", false, "Count lines")
    bytes := flag.Bool("b", false, "Count bytes")
    flag.Parse()

    // calling count fx to count words rec'd STDIN
    fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
    // a scanner is used to read text from Reader
    scanner := bufio.NewScanner(r)

    // if the flag for counting lines is false then we define
    // the scanner split type to words (default is split by lines)
    if countBytes {
        scanner.Split(bufio.ScanBytes)
    } else if !countLines {
        scanner.Split(bufio.ScanWords)
    }

    // counter
    wc := 0

    // loop on every split in scanner
    for scanner.Scan() {
        wc++
    }

    return wc
}
