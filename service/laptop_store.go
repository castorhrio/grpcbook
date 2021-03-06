package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/castorhrio/grpcpcbook/grpcbook"
	"github.com/jinzhu/copier"
)

type LaptopStore interface {
	Save(laptop *grpcbook.Laptop) error
	Find(id string) (*grpcbook.Laptop, error)
}

type InMemoryLaptopStore struct {
	mutex sync.Mutex
	data  map[string]*grpcbook.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*grpcbook.Laptop),
	}
}

var ErrAlreadyExists = errors.New("record already exists")

func (store *InMemoryLaptopStore) Save(laptop *grpcbook.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	other := &grpcbook.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data:%w", err)
	}

	store.data[other.Id] = other
	return nil
}

func (store *InMemoryLaptopStore) Find(id string) (*grpcbook.Laptop, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	other := &grpcbook.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data:%w", err)
	}

	return other, nil
}
