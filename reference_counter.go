package main

import (
	"fmt"
	"sync"
)

type RefCountedObject struct {
	data     string
	refCount int
	lock     sync.Mutex
}

func NewRefCountedObject(data string) *RefCountedObject {
	return &RefCountedObject{
		data:     data,
		refCount: 1,
	}
}

func (r *RefCountedObject) AddRef() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.refCount++
}

func (r *RefCountedObject) Release() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.refCount--
	if r.refCount == 0 {
		fmt.Println("Object is no longer referenced and can be freed:", r.data)
	}
}

func main() {
	obj := NewRefCountedObject("TestObject")

	obj.AddRef()
	fmt.Println(obj.refCount)

	obj.Release()
	fmt.Println(obj.refCount)

	obj.Release()
}

// Output:
// 2
// 1
// Object is no longer referenced and can be freed: TestObject
