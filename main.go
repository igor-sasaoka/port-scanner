package main

import (
	"fmt"
	"net"
	"sort"
)

const portsNumber = 65535

func worker(portsChannel, results chan int, id int) {
    for p := range portsChannel {
        address := fmt.Sprintf("localhost:%d", p) 
        conn, err := net.Dial("tcp", address)
        if err != nil {
            results <- 0
            continue 
        }
        conn.Close()
        results <- p
    }
}

func main() {
    portsChannel := make(chan int, 100)
    results := make(chan int)
    
    var openPorts []int
   
    for i := 0; i < 100; i++ {
        go worker(portsChannel, results, i)
    }

    go func() {
        for i := 1; i <= portsNumber; i++ {
            portsChannel <- i
        }
    }()

    for i := 0; i < portsNumber; i++ {
        p := <-results
        if p != 0 { 
            openPorts = append(openPorts, p)
        }
    }

    close(portsChannel)
    close(results)

    sort.Ints(openPorts)
    fmt.Println("Open ports: ")
    fmt.Println(openPorts)
}
