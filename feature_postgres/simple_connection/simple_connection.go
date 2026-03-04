package simple_connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

// "postgres://YourUserName:YourPassword@YourHostName:5432/YourDatabaseName"
func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("CONN_STRING")
	return pgx.Connect(ctx, connString)
}
