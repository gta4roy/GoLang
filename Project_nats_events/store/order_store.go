package store

import (
	pb "Project_nats_events/order"
)

type OrderStore struct {
	orderList []pb.Order
}

func (store *OrderStore) CreateOrder(order *pb.Order) error {
	store.orderList = append(store.orderList, *order)
	return nil
}

func (store *OrderStore) GetOrders() []pb.Order {
	return store.orderList
}
