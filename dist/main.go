package main

import (  
  "fmt" 
  b "bufio"
)

import (
  "os"
)

func main() {
  rd := b.NewReader(os.Stdin)
  fmt.Println("Enter name:")
  name, _ := rd.ReadString('\n')
  fmt.Printf("Hello %v", name)
}
