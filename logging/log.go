package main

import (
	"fmt"
	"log"
	"time"
)

func runAndWaitLog() int {
	time.Sleep(time.Second * 1)
	return 10
}

func main() {
	for i := 0; i < 5; i++ {
		a := runAndWaitLog()

		log.Println("runAndWait finished...")
		log.Println("result", a))
	}
	fmt.Println("done")
}
