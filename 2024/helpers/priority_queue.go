package helpers

type Item struct {
	value    interface{}
	priority int
}

type PriorityQueue []Item

func (pq *PriorityQueue) Enqueue(val interface{}, prio int) {
	item := Item{value: val, priority: prio}
	*pq = append(*pq, item)
	pq.heapifyUp(len(*pq) - 1)
}

func (pq *PriorityQueue) Dequeue() interface{} {
	n := len(*pq)
	item := (*pq)[0]
	(*pq)[0] = (*pq)[n-1]
	*pq = (*pq)[:n-1]
	pq.heapifyDown(0, len(*pq))

	return item.value
}

func (pq *PriorityQueue) heapifyUp(index int) {
	if index == 0 {
		return
	}

	parent := (index - 1) / 2
	if (*pq)[index].priority > (*pq)[parent].priority {
		return
	}

	(*pq)[index], (*pq)[parent] = (*pq)[parent], (*pq)[index]
	pq.heapifyUp(parent)
}

func (pq *PriorityQueue) heapifyDown(index int, size int) {
	left := (2 * index) + 1
	right := (2 * index) + 2

	smallest := index
	if left < size && (*pq)[left].priority < (*pq)[smallest].priority {
		smallest = left
	}
	if right < size && (*pq)[right].priority < (*pq)[smallest].priority {
		smallest = right
	}

	if smallest == index {
		return
	}

	(*pq)[index], (*pq)[smallest] = (*pq)[smallest], (*pq)[index]
	pq.heapifyDown(smallest, size)
}
