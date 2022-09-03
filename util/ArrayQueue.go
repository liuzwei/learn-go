package util

type ArrayQueue struct {
	ary  []any
	size int
}

func ArrayQueueInstance() *ArrayQueue {

	return &ArrayQueue{
		ary:  []any{},
		size: 0,
	}
}

func (queue ArrayQueue) Add(e any) bool {

	return false
}

func (queue ArrayQueue) Offer(e any) bool {

	return false
}

func (queue ArrayQueue) Remove() any {

	return nil
}

func (queue ArrayQueue) Poll() any {

	return nil
}

func (queue ArrayQueue) Peek() any {

	return nil
}
