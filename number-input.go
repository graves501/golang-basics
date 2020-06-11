package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    inputNumber := getNumberFromUserInput()
    checkNumber(inputNumber)
}

func getNumberFromUserInput() (int64){
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a number between 1 and 10: ")
    userInput, _ := reader.ReadString('\n')
    inputNumber, err := strconv.ParseInt(strings.TrimSpace(userInput), 0, 64)

    if err != nil {
      panic(err)
    }

    return inputNumber
}

func checkNumber(number int64){
    if(number < 0 || number > 10){
      fmt.Println("Your number is not a number between 1 and 10!")
    } else {
      fmt.Printf("Your number is %d", number)
    }
}
