package main

import "fmt"

type MinHeap struct {
	arr []int
}

//inserting into the heap

func (h *MinHeap) insert(val int) {
	h.arr = append(h.arr, val)
	h.heapifyTop()
}

//heapify to maintain the heap properties

func (h *MinHeap) heapifyTop() {
	index := len(h.arr) - 1
	for h.arr[parent(index)] > h.arr[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

//extracting the min element

func (h *MinHeap) extract() {
	element := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	fmt.Println(element)
	h.heapifyBottom()
}

//heapify to maintain the heap properties

func (h *MinHeap) heapifyBottom() {
	index := 0
	for lchild(index) <= len(h.arr)-1 {
		lc := lchild(index)
		rc := rchild(index)
		if rc > len(h.arr)-1{
			if h.arr[lc] < h.arr[index] {
				h.swap(lc, index)
				index = lchild(index)
			}else{
				return
			}
		} else if h.arr[lc] < h.arr[rc] && h.arr[index] > h.arr[lc] {
			h.swap(lc, index)
			index = lchild(index)
		} else if h.arr[rc] < h.arr[lc] && h.arr[index] > h.arr[rc] {
			h.swap(rc, index)
			index = rchild(index)
		} else {
			return
		}
	}
}

func (h *MinHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
func parent(i int) int {
	return (i - 1) / 2
}
func lchild(i int) int {
	return (i * 2) + 1
}
func rchild(i int) int {
	return (i * 2) + 2
}
func main() {
	var h MinHeap
	arr := []int{1, 2, 3}
	for _, v := range arr {
		h.insert(v)
	}
	fmt.Println(h.arr)
	h.extract()
	fmt.Println(h.arr)
}
