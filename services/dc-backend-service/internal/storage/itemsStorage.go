package storage

import (
	"context"
	"dc-backend/db"
	"dc-backend/pkg/types"

	"github.com/jackc/pgx/v5"
)

type ItemStorageImpl struct {
	db *pgx.Conn
}

func (i *ItemStorageImpl) GetAllItems(offset, limit int) (types.Items, error) {
	rows, err := i.db.Query(context.Background(), db.GetItemsWithPagination, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := types.Items{}
	for rows.Next() {
		var item types.Item
		if err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Image, &item.Price); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func NewItemStorage(db *pgx.Conn) *ItemStorageImpl {
	return &ItemStorageImpl{
		db: db,
	}
}
