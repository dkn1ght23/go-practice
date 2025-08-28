package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, "step", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go task("🍵 Making Tea")
	go task("📧 Checking Emails")

	time.Sleep(4 * time.Second)
}
