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
		return interface{}, false
	}
	return sq.Que[0], true
}

func (sq *SQue) Pop() (interface{}, bool) {
	if sq.IsEmpty() {
		return interface{}, false
	}
	firstVal := sq.SQue[0]
	sq.SQue = sq.SQue[1:]
	return firstVal, true
}

func (sq *SQue) IsEmpty() bool {
	return len(sq.Que) == 0
}
