package main
import 	"fmt"

///---------------PriorityQueue--------------------

type PriorityQueue struct {
	Queue queue
}

func (receiver *PriorityQueue) Len() int {
	return receiver.Queue.len()
}

func (receiver *PriorityQueue) Enqueue(element interface{}, priority int) {
	receiver.Queue.enqueue(element, priority)
}

func (receiver *PriorityQueue) Dequeue() (interface{}, error) {
	value, err := receiver.Queue.dequeue()
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (receiver *PriorityQueue) First() *node {
	return receiver.Queue.first()
}

func (receiver *PriorityQueue) Last() *node {
	return receiver.Queue.last()
}

//--------------------QUEUE------------------------------
type queue struct {
	firstEl *node
	size int
	lastId int
}

type node struct {
	next *node
	prev *node
	value interface{}
	priority int
	id int
}

func (receiver *queue) len()  int {
	return receiver.size
}

func (receiver *queue) first() *node {
	return receiver.firstEl
}

func (receiver *queue) last() *node {
	if receiver.firstEl == nil {
		return nil
	}
	return receiver.firstEl.prev
}

func (receiver *queue) enqueue (item interface{}, priority int)  {
	defer func(){
		receiver.size++
		receiver.lastId++
	}()

	if receiver.firstEl == nil {
		receiver.firstEl = &node {
			value:   item,
			priority: priority,
			id: receiver.lastId+1,
		}
		receiver.firstEl.next = receiver.firstEl
		receiver.firstEl.prev = receiver.firstEl
		return
	}
	
	lastItem := receiver.firstEl.prev
	receiver.firstEl.prev = &node {
		value:   item,
		priority: priority,
		id: receiver.lastId+1,
	}
	receiver.firstEl.prev.next = receiver.firstEl
	receiver.firstEl.prev.prev = lastItem
    lastItem.next = receiver.firstEl.prev
}

func (receiver *queue) dequeue() (interface{}, error) {

	if receiver.firstEl == nil {
		return nil, fmt.Errorf("no elements in Queue")
	}

	if receiver.size == 1 {
		pr := receiver.firstEl.value
		receiver.firstEl = nil
		receiver.size = 0
		return pr, nil
	}

	maxElId := receiver.firstEl.id
	maxElPriority := receiver.firstEl.priority
    maxQueue := receiver.firstEl
	current := receiver.firstEl.next

	for current.id != receiver.firstEl.id{
		if  (maxElPriority < current.priority) || (maxElPriority == current.priority && maxElId > current.id) {
			maxElPriority = current.priority
			maxElId = current.id
			maxQueue = current
		}
		current = current.next
	}

	receiver.size--


	maxQueue.next.prev, maxQueue.prev.next = maxQueue.prev, maxQueue.next
	if receiver.firstEl == maxQueue {
		receiver.firstEl = maxQueue.next
	}

	mvalue := maxQueue.value
	//fmt.Println(receiver.firstEl)
	maxQueue = nil
	return mvalue, nil
}

func main()	{

	t := PriorityQueue{}
	t.Enqueue(123, 0)
	t.Enqueue(456, 1)
	t.Enqueue(789, 0)
	t.Enqueue(1230, 0)
	t.Enqueue(4506, 1)
	t.Enqueue(7890, 0)


	for i:=0;i<6;i++ {
		fmt.Println(t.Dequeue())
	}
}
