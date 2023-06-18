package db

const GetItemsWithPagination = `
SELECT * FROM items
OFFSET $1
LIMIT $2
`

const GetItemById = `
SELECT * FROM items
WHERE id = $1
`

const AddItemToUser = `
INSERT INTO user_items(user_id, item_id) VALUES ($1, $2)
`
const UserHasItem = `
SELECT COUNT(*) FROM user_items
WHERE user_id = $1 AND item_id = $2`
