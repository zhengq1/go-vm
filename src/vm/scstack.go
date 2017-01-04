package vm


type ScStack struct {
	Element []*ScriptContext
}

func NewScStack() *ScStack {
	var stack ScStack
	e := make([]*ScriptContext,0)
	stack.Element = e
	return &stack;
}

func  (s *ScStack) Push (data *ScriptContext ) {
	//stack.list.PushBack(data)
	s.Element = append(s.Element,data)
}

func  (s *ScStack) Peek() *ScriptContext{
	//return stack.list.Back()
	len := len(s.Element)
	if(len == 0) {return nil}

	return s.Element[len-1]
}

func  (s *ScStack) Pop() *ScriptContext{
	len := len(s.Element)
	if(len == 0) {return nil}

	e := s.Element[len-1]
	s.Element = s.Element[:len-1]

	return e
}

func  (s *ScStack) Count() int{
	return len(s.Element)
}



