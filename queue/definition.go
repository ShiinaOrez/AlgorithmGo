package queue

import "github.com/ShiinaOrez/AlgorithmGo/typedefs"

type Type typedefs.QueueType

type QueueNode struct {
	Next  *QueueNode
	Value Type
}

type Queue struct {
	Head *QueueNode
	Tail *QueueNode
	Size int
}

type ChannelQueue struct {
	Inter chan Type
}
