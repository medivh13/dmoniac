package dmoniac

import "time"

type GetRiwayatModel struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	UserName  string    `db:"name"`
	Stadium   int64     `db:"stadium"`
	CreatedAt time.Time `db:"created_at"`
}
