package main

import (
	"github.com/emicklei/go-restful/v3"
	"sync"
)

// Creating User define type of interface type that can holds any types of data in it
type ITem interface {
}
type Stackk struct {
	item []ITem
	mux  sync.Mutex
}

// it always returns recently inserted elemet in the stack
func (stk *Stackk) Top() ITem {
	stk.mux.Lock()
	defer stk.mux.Unlock()
	if len(stk.item) == 0 {
		return nil
	}
	//it returns last inserted elements in stack
	return stk.item[len(stk.item)-1]
}

// Push add new item of existing/stack
func (stk *Stackk) Push(item Item, req *restful.Request, res *restful.Response) {
	stk.mux.Lock()
	defer stk.mux.Unlock()
	stk.item = append(stk.item, item)

}

// Pop remove only top elements from stack
func (stk *Stackk) Pop(req *restful.Request, res *restful.Request) ITem {
	s := stk.item
	req.ReadEntity(&s)
	//res.WEntity()
	if len(stk.item) == 0 {
		return nil
	}

	//checking the length of Item array and string that length in slice
	ppItem := stk.item[len(stk.item)-1]
	stk.item = stk.item[:len(stk.item)-1]

	return ppItem

}

// IsEmpty returns whether the stack is empty or not
func (stk *Stackk) IsEmpty() bool {
	return len(stk.item) == 0

}
func (stk *Stackk) Size() int {

	return len(stk.item)
}
func (stk *Stackk) Clear() {
	stk.item = nil
}
func main() {

}
