package main

import (
	"errors"
	"fmt"
	"log"
)

type NumberList struct {
	Value int
	next  *NumberList
}

func NewNumberList(value int) *NumberList {
	return &NumberList{
		Value: value,
	}
}

func (l *NumberList) GetValueByIndex(idx int) (int, error) {
	list := l
	for idx > 0 {
		if list.next == nil && idx > 0 {
			return 0, errors.New("Wrong index")
		}
		list = list.next
		idx--
	}
	return list.Value, nil
}

func (l *NumberList) Add(value int) error {
	list := l
	for list.next != nil {
		list = list.next
	}
	newList := NewNumberList(value)
	list.next = newList
	return nil
}

func (l *NumberList) GetAllValues() []int {
	values := make([]int, 0)
	list := l
	for {
		values = append(values, list.Value)
		if list.next == nil {
			break
		}
		list = list.next
	}
	return values
}

func (l *NumberList) RemoveByValue(val int) error {
	list := l
	lprev := func() *NumberList {
		return nil
	}()
	for list.Value != val {
		lprev = list
		list = list.next

		if list == nil {
			return errors.New("value not find")
		}
	}
	if list == l {
		l = l.next
		return nil
	}
	lprev.next = list.next

	return nil
}

func (l *NumberList) Len() int {
	list := l
	len := 0
	for list.next != nil {
		len++
		list = list.next
	}
	return len
}

func main() {
	l := NewNumberList(0)

	for i := 10; i >= 0; i-- {
		if err := l.Add(i); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(l.GetAllValues())
	fmt.Println(l.GetValueByIndex(0))
	fmt.Println(l.GetValueByIndex(1))
	fmt.Println(l.GetValueByIndex(19))
	fmt.Println("Len:", l.Len())
	if err := l.RemoveByValue(3); err != nil {
		log.Fatal(err)
	}
	fmt.Println(l.GetAllValues())
	fmt.Println("Len:", l.Len())
}
