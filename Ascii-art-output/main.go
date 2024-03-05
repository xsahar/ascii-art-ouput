package main

import (
"fmt"
"os"
"strings"
)

func main() {
if len(os.Args) < 4 {
fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
fmt.Println("Example: go run . --output=<fileName.txt> something standard")
return
}

outputFile := ""
if strings.HasPrefix(os.Args[1], "--output=") {
outputFile = strings.TrimPrefix(os.Args[1], "--output=")
} else {
fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
fmt.Println("Example: go run . --output=<fileName.txt> something standard")
return
}

input := os.Args[2]
font := os.Args[3]

var file string
switch font {
case "standard":
file = "standard.txt"
case "shadow":
file = "shadow.txt"
case "thinkertoy":
file = "thinkertoy.txt"
default:
fmt.Println("Invalid font type")
return
}

data, err := os.ReadFile(file)
if err != nil {
fmt.Println("Failed to read file", err)
return
}

input = strings.Replace(input, "\\n", "\n", -1)
words := strings.Split(input, "\n")

lines := strings.Split(string(data), "\n")

output, err := os.Create(outputFile)
if err != nil {
fmt.Println("Failed to create output file", err)
return
}
defer output.Close()

for _, word := range words {
if word == "" {
fmt.Fprintln(output, "")
continue
}

printWord(word, lines, output)
}
}

func printWord(word string, lines []string, output *os.File) {
letters := strings.Split(word, "")
var ascii []int
for _, letter := range letters {
l := int([]rune(letter)[0])
ascii = append(ascii, l)
}
for j := 1; j < 9; j++ {
str := ""
for _, val := range ascii {
line := (val - 32) * 9
if line+j >= len(lines) {
fmt.Fprintln(output, "Error: Insufficient ASCII characters for the word:", word)
return
}
str += lines[line+j]
}

fmt.Fprintln(output, str)
}
}