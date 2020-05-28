package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
//    "strconv"
    "strings"
)

type WordPos struct {
    pos, len int
}

type GridIter interface {
    get(n int) byte
}

type HorizIter struct {
    line *string 
}

type VertIter struct {
    lines *[]string
    colN int
}

func (i HorizIter) get(n int) byte {
    return (*i.line)[n]
}

func (i VertIter) get(n int) byte {
    return (*i.lines)[n][i.colN]
}




func crosswordPuzzle(crossword []string, words string) []string {
    return crossword
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    var crossword []string

    for i := 0; i < 10; i++ {
        crosswordItem := readLine(reader)
        crossword = append(crossword, crosswordItem)
    }

    words := readLine(reader)

    result := crosswordPuzzle(crossword, words)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%s", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

