package database

import (
	"cake-store-golang-restfull-api/helper"
	"context"
	"database/sql"
)

func Migration(db *sql.DB) {
	context := context.Background()

	var query string = "CREATE TABLE IF NOT EXISTS cakes(id INT PRIMARY KEY AUTO_INCREMENT, title VARCHAR(255) NOT NULL, description TEXT NULL, rating FLOAT DEFAULT 0, image VARCHAR(191) NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP)"

	_, err := db.ExecContext(context, query)
	helper.IfError(err)
}
