package main

import (
	"cnns_vika/examples"
	_ "cnns_vika/utils/im" // image utils
	_ "cnns_vika/utils/u"  // data utils
	"log"
	"time"
)

func main() {
	log.Println("Start")
	timeStart := time.Now()
	// examples.CheckConvolutional()
	examples.CheckConvLayer()
	// examples.CheckFClayer()
	// examples.CheckXOR()
	// examples.CheckOR()
	// examples.CheckAND()
	log.Println("Done in:", time.Since(timeStart))
}