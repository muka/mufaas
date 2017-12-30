package main

import (
	"time"
)

func main() {
	t := time.NewTicker(time.Duration(60) * time.Second)
	<-t.C
	select {}
}
