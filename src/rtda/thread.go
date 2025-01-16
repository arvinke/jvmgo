package rtda

type thread struct {
	pc    int
	stack *Stack
}

func NewThread() *thread {
	return &thread{
		stack: newStack(1024),
	}
}
func (self *thread) PC() int      { return self.pc }
func (self *thread) SetPC(pc int) { self.pc = pc }
func (self *thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *thread) CurrentFrame() *Frame {
	return self.stack.top()
}
