package rtda

type Frame struct {
	lower        *Frame        // 链表模拟栈帧
	localVars    LocalVars     // 保存局部变量指针
	operandStack *OperandStack // 保存操作数栈指针
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}