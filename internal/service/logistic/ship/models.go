package ship

type Ship struct {
	Id    uint64
	Title string
}

var allEntities = make([]Ship, 10)
