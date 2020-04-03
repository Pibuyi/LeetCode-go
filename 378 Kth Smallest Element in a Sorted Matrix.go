type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } //最大堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func kthSmallest(matrix [][]int, k int) int {
    h := &IntHeap{}
    heap.Init(h)
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            heap.Push(h, matrix[i][j])
            if h.Len() > k {                
                heap.Pop(h)
            }
        }
    }
    // 接口断言转换类型为int， 否则提交出错
    return heap.Pop(h).(int)
}
