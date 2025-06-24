package ship

type ShipService interface {
	Describe(shipID uint64) (*Ship, error)
	List(cursor uint64, limit uint64) ([]Ship, error)
	Create(Ship) (uint64, error)
	Update(ShipID uint64, ship Ship) error
	Remove(ShipID uint64) (bool, error)
}

type DummyShipService struct{}

func NewShipService() *DummyShipService {
	return &DummyShipService{}
}

func (d DummyShipService) Describe(shipID uint64) (*Ship, error) {
	for _, v := range allEntities {
		if v.Id == shipID {
			return &v, nil
		}
	}
	return nil, nil
}

func (d DummyShipService) List(cursor uint64, limit uint64) ([]Ship, error) {
	return allEntities, nil
}

func (d DummyShipService) Create(s *Ship) (uint64, error) {
	allEntities = append(allEntities, Ship{
		Id:    s.Id,
		Title: s.Title,
	})
	return s.Id, nil
}

func (d DummyShipService) Update(shipId uint64, ship Ship) error {
	for i, v := range allEntities {
		if v.Id == shipId {
			allEntities[i].Title = ship.Title
		}
	}
	return nil
}

func (d DummyShipService) Remove(shipId uint64) (bool, error) {
	for i, v := range allEntities {
		if v.Id == shipId {
			allEntities = append(allEntities[:i], allEntities[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

//func (d DummyShipService) Help() () {
//	return true, nil
//}
