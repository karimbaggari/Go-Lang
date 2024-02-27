package goroutine

import "fmt"

func worker(done chan bool) {
    fmt.Println("Working...")
    done <- true
}

func main() {
    done := make(chan bool)
    go worker(done)
    <-done
    fmt.Println("Work done!")
}
