package sort

import "container/heap"

type BigHeap []int

func (b BigHeap) Len() int {
	return len(b)
}

func (b BigHeap) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b BigHeap) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b *BigHeap) Push(x interface{}) {
	*b = append(*b, x.(int))
}

func (b *BigHeap) Pop() interface{} {
	ret := (*b)[len(*b)-1]
	*b = (*b)[:len(*b)-1]
	return ret
}

// 维护一个大小为 K 的大顶堆
// 当遍历到一个新的元素时，需要知道这个新元素是否比堆中最大的元素更大，更大的话就把堆中最大元素去除，并将新元素添加到堆中。
// 堆顶元素就是第 k 大的元素
func findKthLargest(nums []int, k int) int {
	h := make(BigHeap, 0)
	size := 0
	heap.Init(&h)
	for i := range nums {
		if size < k {
			heap.Push(&h, nums[i])
			size++
		} else {
			// 最大堆，堆顶就是第 k 大的元素
			if nums[i] > h[0] {
				heap.Pop(&h)
				heap.Push(&h, nums[i])
			} // if>>>
		} // else>>
	} // for>

	return h[0]
}
