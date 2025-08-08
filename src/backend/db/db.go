package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// DB 包裝資料庫連接
type DB struct {
	conn *pgx.Conn
}

// New 建立新的資料庫連接
func New() (*DB, error) {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME")))
	if err != nil {
		return nil, fmt.Errorf("無法連接到資料庫: %v", err)
	}

	return &DB{conn: conn}, nil
}

// Close 關閉資料庫連接
func (db *DB) Close() {
	if db.conn != nil {
		db.conn.Close(context.Background())
	}
}

// GetConn 取得原始連接（用於 controller 中的查詢）
func (db *DB) GetConn() *pgx.Conn {
	return db.conn
}

// TestConnection 測試資料庫連接
func (db *DB) TestConnection() error {
	if db.conn == nil {
		return fmt.Errorf("資料庫連接未初始化")
	}

	// 執行簡單的查詢來測試連接
	var result int
	err := db.conn.QueryRow(context.Background(), "SELECT 1").Scan(&result)
	if err != nil {
		return fmt.Errorf("資料庫連接測試失敗: %v", err)
	}

	if result != 1 {
		return fmt.Errorf("資料庫連接測試返回不正確的結果: %d", result)
	}

	return nil
}

// Ping 檢查資料庫是否仍然可用
func (db *DB) Ping() error {
	if db.conn == nil {
		return fmt.Errorf("資料庫連接未初始化")
	}

	return db.conn.Ping(context.Background())
}
