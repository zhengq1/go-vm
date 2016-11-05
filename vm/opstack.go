package vm

import "container/list"

type OpStack2 struct {
     list *list.List
}

func NewOpStack2() *OpStack2 {
	var stack OpStack2
	stack.list = new(list.List)
	return &stack;
}

func  (stack *OpStack2) Push (data interface{} ) {

	stack.list.PushBack(data)
}

func  (stack *OpStack2) Peek() interface{}{
	return stack.list.Back()
}


func  (stack *OpStack2) Pop() interface{}{
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}


func  (stack *OpStack2) Count() int{
	return stack.list.Len()
}



