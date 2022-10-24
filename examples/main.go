package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/adhocore/chin"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	demo()
	demoWg()
}

func demo() {
	fmt.Print("demo without waitgroup:\n")

	s := chin.New()
	go s.Start()

	longTask()
	s.Stop()
}

func demoWg() {
	var wg sync.WaitGroup
	fmt.Print("\ndemo with waitgroup:\n")

	s := chin.New().WithWait(&wg)
	go s.Start()

	longTaskWg(&wg)
	s.Stop()
	wg.Wait()
}

func longTask() {
	n := rand.Intn(3) + 1
	fmt.Printf("[task] sleep %ds\n", n)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Print("finished task\n")
}

func longTaskWg(wg *sync.WaitGroup) {
	wg.Add(1)
	longTask()
	wg.Done()
}
