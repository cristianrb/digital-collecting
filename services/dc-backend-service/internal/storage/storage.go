package storage

import "dc-backend/pkg/types"

type ItemStorage interface {
	GetAllItems(offset, limit int) (types.Items, error)
	GetItemById(id int64) (*types.Item, error)
	AddItemToUser(itemId, userId int64) error
	CheckUserHasItem(itemId, userId int64) (int, error)
	GetItemsByUserId(userId int64, offset, limit int) (types.Items, error)
}
