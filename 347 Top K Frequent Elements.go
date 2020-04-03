type Node struct {
	val   int
	times int
}
type NodeHeap []*Node
func (h NodeHeap) Len() int { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].times < h[j].times } //最小堆
func (h NodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int, 0)
    for _, v := range nums {
        m[v] += 1
    }

    h := &NodeHeap{}
    for val, times := range m {
        heap.Push(h, &Node{
            val: val,
            times: times,
        })
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    res := make([]int, 0, k)
    for i := 0; i < k; i++ {
        res = append(res, heap.Pop(h).(*Node).val)
    }
    return res
}
