package main

import "container/list"

type Object struct {
	ID string
}

type Pool interface {
	Borrow() (*Object, error)
	Return(*Object) error
}

type Allocate func() (Object, error)

type implementation struct {
	Size      int
	SizeLimit int
	Allocate  Allocate
	Objects   map[Object]bool
	FreeList  *list.List
}

func New(initSize int, limit int, alloc Allocate) (*Pool, error) {
	p := &implementation{
		Size:      initSize,
		SizeLimit: limit,
		Allocate:  alloc,
		Objects:   nil,
		FreeList:  list.New(),
	}

	for i := 0; i < initSize; i++ {
		obj, err := p.Allocate()
		if err != nil {
			return nil, err
		}

		p.FreeList.PushFront(&obj)

		p.Objects[obj] = true
	}

	return nil
}

func (p *implementation) Borrow() (*Object, error) {
	elem := p.FreeList.Front()
	obj := p.FreeList.Remove(elem)

	p.Objects[obj] = false

	return obj, nil
}

func (p *implementation) Return(ref *Object) error {
	list.PushBack(ref)

	p.Objects[*ref] = true

	return nil
}

func main() {
	p := New(0, 0, func() (Object, error) {
		return []string{}
	})

	s, _ := p.Borrow()
	s = append(s, "string")
	_ = p.Return(s)
}
