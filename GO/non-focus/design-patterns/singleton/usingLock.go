package main

import (
    "fmt"
    "sync"
)

var lock = &sync.Mutex{}

type logger1 struct {
}

var singleInstance1 *logger1

func getInstanceLock() *logger1 {
    if singleInstance1 == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance1 == nil {
            fmt.Println("Creating singleInstance1 instance now.")
            singleInstance1 = &logger1{}
        } else {
            fmt.Println("singleInstance1 instance already created.")
        }
    } else {
        fmt.Println("singleInstance1 instance already created.")
    }

    return singleInstance1
}