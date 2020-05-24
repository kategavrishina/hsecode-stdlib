package maxqueue

//go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"

import (
	"errors"
	"github.com/cheekybits/genny/generic"
)

type ValueType generic.Number

type MaxQueue struct {
	// contains filtered or unexported fields
	in     []ValueType
	out    []ValueType
	maxIn  []ValueType
	maxOut []ValueType
}

func New() *MaxQueue {
	return &MaxQueue{}
}

func (q *MaxQueue) Max() (ValueType, error) {
	if len(q.in) == 0 && len(q.out) == 0 {
		return 0, errors.New("the queue is empty")
	}

	i := q.maxIn
	o := q.maxOut

	if len(i) == 0 && len(o) == 0 {
		return 0, errors.New("the queue is empty")
	} else if len(i) == 0 && len(o) != 0 {
		return o[len(o)-1], nil
	} else if len(o) == 0 && len(i) != 0 {
		return i[len(i)-1], nil
	}

	if o[len(o)-1] > i[len(i)-1] {
		return o[len(o)-1], nil
	} else if i[len(i)-1] > o[len(o)-1] {
		return i[len(i)-1], nil
	} else {
		return i[len(i)-1], nil
	}
}

func (q *MaxQueue) Pop() (ValueType, error) {
	if len(q.in) == 0 && len(q.out) == 0 {
		return 0, errors.New("the queue is empty")
	}
	if len(q.out) == 0 {
		for i := range q.in {
			value := q.in[len(q.in)-1-i]
			q.out = append(q.out, value)
			if len(q.maxOut) == 0 {
				q.maxOut = append(q.maxOut, value)
			} else if value > q.maxOut[len(q.maxOut)-1] {
				q.maxOut = append(q.maxOut, value)
			} else {
				q.maxOut = append(q.maxOut, q.maxOut[len(q.maxOut)-1])
			}
		}
		q.in = []ValueType{}
		q.maxIn = []ValueType{}
	}

	q.maxOut = q.maxOut[:len(q.maxOut)-1]

	last := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	if len(q.out) == 0 {
		q.maxOut = []ValueType{}
	}

	return last, nil
}

func (q *MaxQueue) Push(value ValueType) {
	q.in = append(q.in, value)
	if len(q.maxIn) == 0 {
		q.maxIn = append(q.maxIn, value)
	} else if value > q.maxIn[len(q.maxIn)-1] {
		q.maxIn = append(q.maxIn, value)
	} else {
		q.maxIn = append(q.maxIn, q.maxIn[len(q.maxIn)-1])
	}
}
