package tree

import (
	"errors"
	"strconv"
)

func (T *Tree) Encode() []string {
	if T == nil {
		return nil
	}
	result := make([]string, 0)
	var queue []*Tree
	queue = append(queue, T)

	for len(queue) > 0 {
		currentNode := queue[0]
		if currentNode == nil {
			result = append(result, "nil")
		} else {
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			} else {
				queue = append(queue, nil)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			} else {
				queue = append(queue, nil)
			}
			result = append(result, strconv.Itoa(currentNode.Value))
		}
		queue = queue[1:]
	}
	for i := len(result) - 1; result[i] == "nil"; i-- {
		result = result[:len(result)-1]
	}
	return result
}

func Decode(data []string) (*Tree, error) {
	if len(data) == 0 {
		return nil, nil
	}
	value, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, errors.New("first element is incorrect")
	}
	left, errL := Helper(data, 1)
	if errL != nil {
		return nil, errL
	}
	right, errR := Helper(data, 2)
	if errR != nil {
		return nil, errR
	}
	return &Tree{
		value,
		left,
		right,
	}, nil
}

func Helper(data []string, idx int) (*Tree, error) {
	if idx >= len(data) {
		return nil, nil
	}
	if value, err := strconv.Atoi(data[idx]); err == nil {
		left, errL := Helper(data, idx*2+1)
		if errL != nil {
			return nil, errL
		}
		right, errR := Helper(data, idx*2+2)
		if errR != nil {
			return nil, errR
		}
		return &Tree{
			value,
			left,
			right,
		}, nil
	} else if data[idx] == "nil" {
		if idx*2+1 >= len(data) && idx*2+2 >= len(data) {
			return nil, nil
		} else {
			return nil, errors.New("non-nil node is child of nil")
		}
	} else {
		return nil, errors.New("element is not nil nor integer")
	}
}
