package types

type Item struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

type ItemId struct {
	Id int64 `json:"id"`
}

type Items []Item

type Coins struct {
	Coins float64 `json:"coins"`
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Coins    int64  `json:"coins"`
}

type ApiError struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
