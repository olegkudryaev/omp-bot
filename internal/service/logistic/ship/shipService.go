package ship

import (
	"errors"
	"fmt"
	"log"
)

type DummyShipService struct{}

func NewShipService() *DummyShipService {
	return &DummyShipService{}
}

func (d DummyShipService) Describe(shipID uint64) (*Ship, error) {
	for i := range allEntities {
		if allEntities[i].Id == shipID {
			return &allEntities[i], nil
		}
	}
	return nil, fmt.Errorf("корабль с id %d не найден", shipID)
}

func (d DummyShipService) List(cursor uint64, limit uint64) ([]Ship, error) {
	n := uint64(len(allEntities))
	if cursor >= n {
		return nil, errors.New("курсор за пределами списка")
	}

	end := cursor + limit
	if end > n {
		end = n
	}

	return allEntities[cursor:end], nil
}

func (d DummyShipService) Create(s *Ship) (uint64, error) {
	for _, ship := range allEntities {
		if ship.Id == s.Id && s.Id != 0 {
			return 0, fmt.Errorf("корабль с id %d уже существует", s.Id)
		}
	}

	newShip := Ship{
		Id:    s.Id,
		Title: s.Title,
	}
	allEntities = append(allEntities, newShip)
	return s.Id, nil
}

func (d DummyShipService) Update(shipId uint64, s Ship) error {
	if s.Title == "" {
		log.Printf("название корабля не может быть пустым")
		return errors.New("название корабля не может быть пустым")
	}

	for i, v := range allEntities {
		if v.Id == shipId {
			allEntities[i].Title = s.Title
			return nil
		}
	}
	log.Printf("корабль с id %d не найден\n", shipId)
	return fmt.Errorf("корабль с id %d не найден", shipId)
}

func (d DummyShipService) Remove(shipId uint64) (bool, error) {
	for i, v := range allEntities {
		if v.Id == shipId {
			allEntities = append(allEntities[:i], allEntities[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New(fmt.Sprintf("корабль по Id %d не найден", shipId))
}
