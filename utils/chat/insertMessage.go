package chat

import (
	"context"

	"github.com/furre-dev/homelaan-go-backend/graph/model"
	"github.com/jackc/pgx/v5"
)

func InsertMessageToPg(ctx context.Context, db *pgx.Conn, message model.Message) (*model.Message, error) {
	return nil, nil
}
