package store

import (
	mystore "bookstore/store"
	"bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func (s *MemStore) Create(b *mystore.Book) error {
	return nil
}

func (s *MemStore) Update(b *mystore.Book) error {
	return nil
}

func (s *MemStore) Get(key string) (mystore.Book, error) {
	b := mystore.Book{}
	return b, nil
}

func (s *MemStore) GetAll() ([]mystore.Book, error) {
	var list []mystore.Book
	return list, nil

}

func (s *MemStore) Delete(key string) error {

	return nil
}
