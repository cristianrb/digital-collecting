package db

const GetItemsWithPagination = `
SELECT * FROM items
OFFSET $1
LIMIT $2
`