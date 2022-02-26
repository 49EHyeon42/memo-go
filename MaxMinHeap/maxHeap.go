package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, value int
	fmt.Fscanln(reader, &N)

	heap := NewHeap()

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &value)
		if value == 0 {
			if heap.size == 0 {
				fmt.Fprintln(writer, 0)
			} else {
				fmt.Fprintln(writer, heap.Pop())
			}
		} else {
			heap.Push(value)
		}
	}
}

func NewHeap() heap {
	return heap{make([]int, 1), 0}
}

type heap struct {
	linkedList []int
	size       int
}

// add linked list
func (h *heap) add(key int) {
	h.linkedList, h.size = append(h.linkedList, key), h.size+1
}

// remove linked list
func (h *heap) remove() {
	h.linkedList, h.size = h.linkedList[:len(h.linkedList)-1], h.size-1
}

func (h *heap) Push(key int) {
	h.add(key)

	child := h.size
	parent := child / 2

	for child > 1 && h.linkedList[child] > h.linkedList[parent] {
		// swap
		h.linkedList[child], h.linkedList[parent] = h.linkedList[parent], h.linkedList[child]

		child = parent
		parent = child / 2
	}
}

func (h *heap) Pop() int {
	answer := h.linkedList[1]

	h.linkedList[1] = h.linkedList[h.size]
	h.remove()

	parent := 1
	child := parent * 2

	// check the key of the two child
	if h.size >= child+1 && h.linkedList[child] <= h.linkedList[child+1] {
		child += 1
	}

	for child <= h.size && h.linkedList[parent] < h.linkedList[child] {
		// swap
		h.linkedList[parent], h.linkedList[child] = h.linkedList[child], h.linkedList[parent]

		parent = child
		child = parent * 2

		// check the key of the two child
		if h.size >= child+1 && h.linkedList[child] <= h.linkedList[child+1] {
			child += 1
		}
	}

	return answer
}
