package main

import (
    "fmt"
)

func main() {

    // fmt.Println("Initialising the goInstanceLock method")
    // for i := 0; i < 5; i++ {
    //     go getInstanceLock()
    // }

    fmt.Println("Initialising the goInstanceSync method")
    for i := 0; i < 5; i++ {
        go getInstanceSync()
    }

    // Scanln is similar to Scan, but stops scanning at a newline and
    // after the final item there must be a newline or EOF.
    fmt.Scanln()
}