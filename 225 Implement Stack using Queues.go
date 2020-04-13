type MyStack struct {
    slice []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.slice = append(this.slice, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    if len(this.slice) == 0 {
        return -1
    }
    res := this.slice[len(this.slice) - 1]
    this.slice = this.slice[:len(this.slice) - 1]
    return res
}


/** Get the top element. */
func (this *MyStack) Top() int {
    if len(this.slice) == 0 {
        return -1
    }
    res := this.slice[len(this.slice) - 1]
    return res
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.slice) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
