package main

import (
	"time"   
    "fmt"
    "sync"
)


func worker(id int, wg *sync.WaitGroup, jobs <-chan int, res chan<- int){
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("ID: %d, task number %d\n", id, job)

        time.Sleep(500 * time.Millisecond)

        res <- job*job
    }
}


func main() {
    
    var wg sync.WaitGroup
    jobs := make(chan int, 10)
    result := make(chan int, 10)

    for i := 0; i < 3; i++ {
        wg.Add(1)
        go worker(i, &wg, jobs, result)
    }
    
    for j := 1; j <= 10; j++ {
        jobs <- j
    }
    close(jobs)

    go func(){
        wg.Wait()
        close(result)
    }()
    
    for res := range result{
        fmt.Println("RES: ", res)
    }
}
