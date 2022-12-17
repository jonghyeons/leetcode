type MyStack []int

func (ms *MyStack) Empty() bool {
	return len(*ms) == 0
}

func (ms *MyStack) Push(x int) {
	*ms = append(*ms, x)
}

func (ms *MyStack) Peek() int {
	return (*ms)[len(*ms)-1]
}

func (ms *MyStack) Pop() int {
	last := (*ms)[len(*ms)-1]
	*ms = (*ms)[:len(*ms)-1]
	return last
}

type MyQueue struct {
	stack1 MyStack
	stack2 MyStack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.stack1.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.stack2.Empty() {
		for !this.stack1.Empty() {
			last := this.stack1.Pop()
			this.stack2.Push(last)
		}
	}
	return this.stack2.Pop()
}

func (this *MyQueue) Peek() int {
	if this.stack2.Empty() {
		for !this.stack1.Empty() {
			last := this.stack1.Pop()
			this.stack2.Push(last)
		}
	}
	return this.stack2.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.stack1.Empty() && this.stack2.Empty()
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */