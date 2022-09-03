package queue

type Queue interface {
	Add(e any) bool

	Offer(e any) bool

	Remove() any

	Poll() any

	Peek() any
}
