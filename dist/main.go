package main

import (  
  "fmt" 
  "os"
)

func main() {
  var name string
  fmt.Println("Enter name:")
  fmt.Scanln(&name)
  fmt.Println(name)
  fmt.Println(os.Getwd)
}
