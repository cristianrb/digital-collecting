db_migrate:
	migrate -path db/migration -database "postgresql://dc_admin:dc_password@localhost:5432/digital_collecting?sslmode=disable" -verbose up

mock:
	mockgen -destination mocks/db/store.go -package mockdb dc-backend/internal/storage ItemStorage
	mockgen -destination mocks/token/jwt_token.go -package mocktoken dc-backend/internal/token JWTValidator
