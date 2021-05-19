package main

import (  
  "fmt" 
  "bufio"
)

import (
  "os"
)

func main() {
  rd := bufio.NewReader(os.Stdin)
  fmt.Println("Enter name:")
  name, _ := rd.ReadString('\n')
  fmt.Printf("Hello %v", name)
}
