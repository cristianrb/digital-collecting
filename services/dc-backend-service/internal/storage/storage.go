package storage

import "dc-backend/pkg/types"

type ItemStorage interface {
	GetAllItems(offset, limit int) (types.Items, error)
}
