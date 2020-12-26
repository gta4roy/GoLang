package store

import (
	pb "Project_nats_events/order"
)

// EventStore provides CRUD operations against the collection "orders"
type EventStore struct {
	eventStoreList []pb.EventStore
}

// CreateEvent inserts the value of struct Order into collection.
func (store *EventStore) CreateEvent(order *pb.EventStore) error {
	store.eventStoreList = append(store.eventStoreList, *order)
	return nil
}

// GetEvents returns all documents from the collection.
func (store *EventStore) GetEvents() []pb.EventStore {
	return store.eventStoreList
}
