package main

import (  
  "fmt" 
)

import (
  "os"
)

func main() {
  var name string
  fmt.Println("Enter name:")
  fmt.Scanln(&name)
  fmt.Println(name)
  fmt.Println(os.Getwd)
}
