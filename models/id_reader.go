package models

type IdReader interface {
	GetID() uint
}

func GetIds[T IdReader](items []T) []uint {
	var ids []uint
	for _, item := range items {
		ids = append(ids, item.GetID())
	}
	return ids
}
