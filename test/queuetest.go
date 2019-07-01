package algorithmtest

import (
	"fmt"
	"github.com/ShiinaOrez/AlgorithmGo/queue"
)

func QueueTest() {
	fmt.Println("------", "Queue Test:")
	classicQueue := new(queue.Queue)
	for i:=0; i<10; i++ {
		fmt.Printf("%d ", i)
		classicQueue.Push(queue.Type(i))
	}
	fmt.Println()
	for !classicQueue.Empty() {
		fmt.Printf("%d ", classicQueue.Peek())
		classicQueue.Pop()
	}
	fmt.Println()

	fmt.Println("------", "ChannelQueue Test:")
	channelQueue := new(queue.ChannelQueue)
	channelQueue.Inter = make(chan queue.Type, 1)
	for i:=0; i<10; i++ {
		channelQueue.Push(queue.Type(i))
	}
	for !channelQueue.Empty() {
		fmt.Printf("%d ", channelQueue.Pop())
	}
	fmt.Println()
}