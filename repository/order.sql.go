// Code generated by sqlc. DO NOT EDIT.
// source: order.sql

package repository

import (
	"context"
	"time"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
  order_id, order_price, order_type,order_crypto, order_status, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING order_id, order_price, order_type, order_crypto, order_status, created_at, updated_at
`

type CreateOrderParams struct {
	OrderID     int64     `json:"order_id"`
	OrderPrice  float64   `json:"order_price"`
	OrderType   string    `json:"order_type"`
	OrderCrypto string    `json:"order_crypto"`
	OrderStatus string    `json:"order_status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.queryRow(ctx, q.createOrderStmt, createOrder,
		arg.OrderID,
		arg.OrderPrice,
		arg.OrderType,
		arg.OrderCrypto,
		arg.OrderStatus,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.OrderPrice,
		&i.OrderType,
		&i.OrderCrypto,
		&i.OrderStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders
WHERE order_id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, orderID int64) error {
	_, err := q.exec(ctx, q.deleteOrderStmt, deleteOrder, orderID)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT order_id, order_price, order_type, order_crypto, order_status, created_at, updated_at FROM orders
WHERE order_id = $1 LIMIT 1
`

func (q *Queries) GetOrder(ctx context.Context, orderID int64) (Order, error) {
	row := q.queryRow(ctx, q.getOrderStmt, getOrder, orderID)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.OrderPrice,
		&i.OrderType,
		&i.OrderCrypto,
		&i.OrderStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listOrder = `-- name: ListOrder :many
SELECT order_id, order_price, order_type, order_crypto, order_status, created_at, updated_at FROM orders
LIMIT $1
OFFSET $2
`

type ListOrderParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOrder(ctx context.Context, arg ListOrderParams) ([]Order, error) {
	rows, err := q.query(ctx, q.listOrderStmt, listOrder, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderID,
			&i.OrderPrice,
			&i.OrderType,
			&i.OrderCrypto,
			&i.OrderStatus,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
