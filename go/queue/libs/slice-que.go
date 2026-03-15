package libs

type SQue struct {
	Que []interface{}
}

func NewSQue() *SQue {
	return new(SQue)
}

func (sq *SQue) Push(data interface{}) {
	sq.Que = append(sq.Que, data)
}

func (sq *SQue) Peek() (interface{}, bool) {
	if sq.IsEmpty() {
		return nil, false
	}
	return sq.Que[0], true
}

func (sq *SQue) Pop() (interface{}, bool) {
	if sq.IsEmpty() {
		return nil, false
	}
	firstVal := sq.Que[0]
	sq.Que = sq.Que[1:]
	return firstVal, true
}

func (sq *SQue) IsEmpty() bool {
	return len(sq.Que) == 0
}
