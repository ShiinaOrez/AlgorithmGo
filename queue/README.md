## AlgorithmGo.queue

-----

This package defines the **queue** data structure, as a kind of linear table, queue's rule is FIFO(First In First Out).

In this package, I build two types of queue by links and channel as **Queue** and **ChannelQueue**. I just provide you with some simple method like: **Push**, **Pop**, **Peek**, **Empty**. If you want create more new method, please fork my code and edit it.

-----

## Package Information:

### Package: queue

#### Type: Queue
#### Type: QueueNode
#### Mehtod: Queue.Push(element eleType) Size int;
#### Mehtod: Queue.Pop();
#### Method: Queue.Peek() element eleType;
#### Method: Queue.Empty() empty bool;

#### Type: ChannelQueue
#### Mehtod: ChannelQueue.Push(element eleType) Size int;
#### Mehtod: ChannelQueue.Pop() eleType;
#### Method: ChannelQueue.Empty() empty bool;
#### Method: ChannelQueue.New() *ChannelQueue;

-----

## Some Details:

This package just defined a simple queue, please add new method by your self.

Because of Go language dosen't support generic programming for now, so eleType needs you define it by yourself at typedefs package.

-----

## In the End:

### Contributor: ShiinaOrez
