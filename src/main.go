package main 

import (
    "fmt"

    "model"
)

func main() {
    model.MustLoadConfig()
    fmt.Print("Doet Het")
}