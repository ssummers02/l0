package repository

import (
	"errors"
	"fmt"
	"l0/internal/domain/entity"
	"sync"
)

var ErrCacheNotFound = errors.New("not found cache")

type ordersCache struct {
	mutex sync.RWMutex
	data  map[string]entity.Order
}

func newOrdersCache() ordersCache {
	return ordersCache{
		data: make(map[string]entity.Order),
	}
}

func (r *ordersCache) setOrders(orders ...entity.Order) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, v := range orders {
		r.data[v.OrderUID] = v
	}
}

func (r *ordersCache) len() int {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return len(r.data)
}

func (r *ordersCache) find(id string) (entity.Order, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	res, ok := r.data[id]
	if ok {
		return res, nil
	}

	return entity.Order{}, fmt.Errorf("%w: %s", ErrCacheNotFound, id)
}
