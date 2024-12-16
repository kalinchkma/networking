package interfaces

import "gnja_server/internal/database"

type Configuration struct {
	DB         *database.Queries
	JWT_SECRET string
}
