package persist

import (
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	itemCount := 0
	go func() {
		for {
			item := <-out
			log.Printf("ItemSaver got item #%d : %v", itemCount, item)
			itemCount++
		}
	}()
	return out
}
