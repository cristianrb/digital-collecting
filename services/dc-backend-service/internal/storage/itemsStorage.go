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

func NewItemStorage(db *pgx.Conn) *ItemStorageImpl {
	return &ItemStorageImpl{
		db: db,
	}
}

func (i *ItemStorageImpl) GetAllItems(offset, limit int) (types.Items, error) {
	rows, err := i.db.Query(context.Background(), db.GetItemsWithPagination, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := types.Items{}
	for rows.Next() {
		var item *types.Item
		if err := mapRowToItem(rows, item); err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (i *ItemStorageImpl) GetItemById(id int64) (*types.Item, error) {
	row := i.db.QueryRow(context.Background(), db.GetItemById, id)
	var item types.Item
	err := mapRowToItem(row, &item)

	return &item, err
}

func (i *ItemStorageImpl) AddItemToUser(itemId, userId int64) error {
	_, err := i.db.Exec(context.Background(), db.AddItemToUser, userId, itemId)
	return err
}

func (i *ItemStorageImpl) CheckUserHasItem(itemId, userId int64) (int, error) {
	var counter int
	err := i.db.QueryRow(context.Background(), db.UserHasItem, userId, itemId).Scan(&counter)
	if err != nil {
		return -1, err
	}

	return counter, nil
}

func (i *ItemStorageImpl) GetItemsByUserId(userId int64, offset, limit int) (types.Items, error) {
	rows, err := i.db.Query(context.Background(), db.GetUserItems, userId, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := types.Items{}
	for rows.Next() {
		var item types.Item
		if err := mapRowsToItem(rows, &item); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func mapRowToItem(row pgx.Row, item *types.Item) error {
	err := row.Scan(&item.Id, &item.Name, &item.Description, &item.Image, &item.Price)
	return err
}

func mapRowsToItem(row pgx.Rows, item *types.Item) error {
	err := row.Scan(&item.Id, &item.Name, &item.Description, &item.Image, &item.Price)
	return err
}
