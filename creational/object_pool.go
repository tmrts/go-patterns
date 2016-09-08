package main

import (
	"container/list"
	"errors"
	"log"
)

var (
	ErrNoMoreObject       = errors.New("no more object")
	ErrNotEnoughPoolSpace = errors.New("not enough pool space")
)

type Object struct {
	ID string
}

type Pool interface {
	Borrow() (*Object, error)
	Return(*Object) error
}

type Allocate func() (*Object, error)

type implementation struct {
	Size     int
	Allocate Allocate
	FreeList *list.List
}

func New(initSize int, alloc Allocate) (Pool, error) {
	p := &implementation{
		Size:     initSize,
		Allocate: alloc,
		FreeList: list.New(),
	}

	for i := 0; i < initSize; i++ {
		obj, err := p.Allocate()
		if err != nil {
			return nil, err
		}

		p.FreeList.PushFront(obj)
	}

	return p, nil
}

func (p *implementation) Borrow() (*Object, error) {
	elem := p.FreeList.Front()
	if elem == nil {
		return nil, ErrNoMoreObject
	}

	obj := p.FreeList.Remove(elem)

	o := obj.(*Object)

	return o, nil
}

func (p *implementation) Return(ref *Object) error {
	if p.FreeList.Len() == p.Size {
		return ErrNotEnoughPoolSpace
	}
	p.FreeList.PushBack(ref)
	return nil
}

func main() {
	const poolSize = 3

	p, _ := New(poolSize, func() (*Object, error) {
		return &Object{}, nil
	})

	ob, err := p.Borrow()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("borrow a object from pool: %#v\n", *ob)

	for i := 0; i < poolSize-1; i++ {
		o, err := p.Borrow()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("borrow a object from pool: %#v\n", *o)
	}

	_, err = p.Borrow()
	if err.Error() != ErrNoMoreObject.Error() {
		log.Fatalf("expect: %v\n", ErrNoMoreObject)
	}

	p.Return(ob)
	ob1, err := p.Borrow()
	if err != nil {
		log.Fatal(err)
	}

	if ob != ob1 {
		log.Fatal("expect the same object")
	}
}
