package main

import "context"

type service struct {
	store OrderStore
}

func NewService(store OrderStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(ctx context.Context) {
	return nil
}
