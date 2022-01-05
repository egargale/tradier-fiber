// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"

	"github.com/jackc/pgtype"
)

type Account struct {
	AccountID      int32     `json:"account_id"`
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	AccountNumber  string    `json:"account_number"`
	Classification string    `json:"classification"`
	DateCreated    time.Time `json:"date_created"`
	DayTrader      bool      `json:"day_trader"`
	OptionLevel    int32     `json:"option_level"`
	Status         string    `json:"status"`
	Type           string    `json:"type"`
	CreatedAt      time.Time `json:"created_at"`
}

type Position struct {
	PositionID   int32          `json:"position_id"`
	ProfileID    sql.NullString `json:"profile_id"`
	CostBasis    pgtype.Numeric `json:"cost_basis"`
	DateAcquired sql.NullTime   `json:"date_acquired"`
	ID           sql.NullInt32  `json:"id"`
	Quantity     sql.NullInt32  `json:"quantity"`
	Symbol       sql.NullString `json:"symbol"`
}

type PositionsMaster struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Classification string         `json:"classification"`
	Symbol         sql.NullString `json:"symbol"`
	Quantity       sql.NullInt32  `json:"quantity"`
	CostBasis      pgtype.Numeric `json:"cost_basis"`
	DateAcquired   sql.NullTime   `json:"date_acquired"`
}

type Todo struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Completed sql.NullBool `json:"completed"`
}