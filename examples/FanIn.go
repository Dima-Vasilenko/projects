package main

import (
	"time"   
    "fmt"
    "sync"
)

func source(id int, ch chan<- int){
    defer close(ch)
    for i := id; i < id + 5; i++{
        ch <- i
        time.Sleep(500 * time.Millisecond)
    }
}

func fanIn(channels ...<-chan int) <-chan int {
    
    out := make(chan int)
    var wg sync.WaitGroup

    for _, channel := range channels {
        wg.Add(1)
        go func(chann <-chan int){
            defer wg.Done()
            for ch := range chann {
                out <- ch
            }
        }(channel)
    }

    go func(){
        wg.Wait()
        close(out)
    }()

    return out
}

func main() {
    
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)

    go source(1, ch1)
    go source(11, ch2)
    go source(21, ch3)

    for value := range fanIn(ch1, ch2, ch3){
        fmt.Println(value)
    }
}
