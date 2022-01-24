package main

import (
	"fmt"
)

type Base struct {
	Key   interface{}
	Value interface{}
}

type Store struct {
	db []Base
	//mu *sync.RWMutex
}

func NewStore() *Store {
	basetore := &Store{db: []Base{}}
	return basetore
}

func (s *Store) Append(key interface{}, value interface{}) {
	s.db = append(s.db, Base{Key: key, Value: value})
}

func (s *Store) Get(key interface{}) (interface{}, bool) {

	//mux.Rlock()

	for _, elem := range s.db {

		if elem.Key == key {
			return elem.Value, true
		}
	}

	//mux.RUnlock()
	return 0, false
}

func (s *Store) Get_index(key interface{}) (int, bool) {

	for index, elem := range s.db {
		if elem.Key == key {
			return index, true
		}
	}

	return 0, false
}

func (s *Store) Delete(key interface{}) bool {

	i, ok := s.Get_index(key)
	if !ok {
		return false
	}
	m := s.db

	m[i] = m[len(m)-1]
	s.db = m[:len(m)-1]
	return true
}

func (s *Store) Update(key interface{}, value interface{}) bool {

	i, ok := s.Get_index(key)
	if !ok {
		return false
	}

	m := &s.db[i]
	m.Value = value

	return true
}

func main() {

	//mux := &sync.RWMutex{}

	// Check Base structure
	fmt.Println(Base{Key: "one", Value: 123})
	// Make new db structure
	one := NewStore()

	// Check put method
	one.Append("one", 123)
	one.Append("two", 222)

	// Check get method
	v, _ := one.Get("one")
	fmt.Println("Get method ", v)

	// Check update method
	one.Update("one", 112)
	fmt.Println("Update ", one)

	// Check delete method
	v = one.Delete("one")
	fmt.Println("delete ", one)

}
