package main

import (
    "fmt"
    "sync"
)

var once sync.Once

type logger2 struct {
}

var singleInstance2 *logger2

func getInstanceSync() *logger2 {
    if singleInstance2 == nil {
        once.Do(
            func() {
                fmt.Println("Creating singleInstance2 instance now.")
                singleInstance2 = &logger2{}
            })
    }

    return singleInstance2
}