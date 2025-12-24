package main

import (
  "fmt"
  "os"
)

// Что выведет программа?

// Объяснить внутреннее устройство интерфейсов
// и их отличие от пустых интерфейсов.

func Foo() error {
  var err *os.PathError = nil
  return err
}

func main() {
	// возвращем error - интерфейс (хранит 2 значения)
  err := Foo()
	
	// nil
  fmt.Println(err)
	
	// выведется false потому это интерфейс (error)
	// а интерфейс равен nil только если оба его значения равны nil
	// В нашем случае будет примерно так:
	// 	Интерфейс error = {
	//     tab:  *itab -> информация о типе *os.PathError  // НЕ nil!
	//     data: nil   -> само значение                    // nil
	// }
  fmt.Println(err == nil)
}