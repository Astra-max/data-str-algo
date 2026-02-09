package stk


var StackList []interface{}

func Push(data interface{}) {
	StackList = append(StackList, data)
}

func Peek() interface{} {
	return StackList[len(StackList)-1]
}

func Pop() (interface{}, bool) {
	if Len() == 0 {
		return nil, false
	}

	val := StackList[len(StackList)]
	StackList = StackList[:len(StackList)-1]
	return val, true
}

func Len() int {
		return len(StackList)
}
