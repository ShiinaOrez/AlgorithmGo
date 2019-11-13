package queue

import (
	"fmt"
	_ "fmt"
)

func (queue *ChannelQueue) Push(value Type) int {
	select {
	case queue.Inter <- value:
	default:
		length := len(queue.Inter)
		fmt.Println("[ChannelQueue]: The internal channel is full. Will be updated at next.")
		queue.Inter = func(origin chan Type) chan Type {
			newBuffer := make(chan Type, 2*length)
			close(origin)
			for i := range origin {
				newBuffer <- i
			}
			return newBuffer
		}(queue.Inter)
		queue.Inter <- value
		fmt.Printf("[ChannelQueue]: Expanding is over, now size is: %d.\n", length*2)
	}
	return len(queue.Inter)
}

func (queue *ChannelQueue) Pop() Type {
	var element Type
	select {
	case element = <-queue.Inter:
	default:
		fmt.Println("[ChannelQueue]: Queue is empty!!!!")
	}
	return element
}

func (queue *ChannelQueue) Empty() bool {
	return len(queue.Inter) == 0
}

func (queue *ChannelQueue) New() *ChannelQueue {
	newQueue := new(ChannelQueue)
	newQueue.Inter = make(chan Type, 1)
	return newQueue
}
