package main

import (
	"errors"
	"fmt"
	"log"
)

type Queu struct {
	Arr  []int
	Size int
}

//var q *Queu

func (q *Queu) Enque(ele int) {
	if q.Size == q.GetSize() {
		log.Println("Queue is Full")
	}
	//here is element getting append in queue
	q.Arr = append(q.Arr, ele)
}

// it uses to delete the elements from the queue
func (q *Queu) Dequ() int {
	if q.IsEmp() { //checking if element exists then delete from queue or not then empty queue message
		log.Println("Queue Is Empty")
		return 0
	}

	ele := q.Arr[0]       // in queue always 0th position elements get deleted from the queue
	if q.GetSize() == 1 { //if size is one then enter this loop and
		q.Arr = nil
		return ele
	}
	// it is slice and skipping the elements of 0th position
	q.Arr = q.Arr[1:] // it will print all elements excepts zeroth elements
	return ele        //it returns only deueu elements from
}

// it returns size of array
func (q *Queu) GetSize() int {
	return len(q.Arr)
}
func (q *Queu) PeeK() (int, error) {
	if q.IsEmp() {
		log.Println()
		return 0, errors.New("Empty")
	}
	//peel return always first elements from the queue
	return q.Arr[0], nil
}
func (q *Queu) IsEmp() bool {
	return len(q.Arr) == 0

}
func main() {
	que := Queu{Size: 3}
	//log.Println(que.Arr)
	que.Enque(1)
	log.Println(que.Arr)
	que.Enque(2)
	log.Println(que.Arr)
	que.Enque(3)
	log.Println(que.Arr)
	que.Enque(4)
	log.Println(que.Arr)
	e := que.Dequ()
	log.Println("Dequeue", e)
	q := que.Dequ()
	log.Println("Dequeue", q)
	//	PeeK()
	fmt.Println("Peek", que.Arr)
	que.Enque(5)
	log.Println(que.Arr)
	que.Enque(8)
	log.Println(que.Arr)
	que.Enque(9)
	log.Println(que.Arr)
	que.Enque(10)
	log.Println(que.Arr)
	que.Enque(19)
	log.Println(que.Arr)
	ee := que.Dequ()
	log.Println("Dequeue", ee)
	qq := que.Dequ()
	log.Println("Dequeue", qq)
	e1 := que.Dequ()
	log.Println("Dequeue", e1)
	q1 := que.Dequ()
	log.Println("Dequeue", q1)
	//e2 := que.Dequ()
	//log.Println("Dequeue", e2)
	//q2 := que.Dequ()
	//log.Println("Dequeue", q2)
	//e3 := que.Dequ()
	//log.Println("Dequeue", e3)
	//q3 := que.Dequ()
	//log.Println("Dequeue", q3)
	que.PeeK()
	fmt.Println("Peek", que.Arr)

}
