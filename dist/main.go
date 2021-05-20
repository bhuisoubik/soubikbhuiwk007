package main

import (  
  brainfuck "github.com/soubikbhuiwk007/brainfuck/bf"
)

import "fmt"


func main() {
  fmt.Println("Starting..")
  brainfuck.Prase(`
    ++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.
  `)
  fmt.Println("Completed..")
}
