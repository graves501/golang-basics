package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("What is your name: ")
    name, _ := reader.ReadString('\n')
    fmt.Print("Your name is: " + name)
}
