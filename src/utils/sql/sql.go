package sql

import (
	"context"
	"database/sql"
	"fmt"
	"test_dealls/src/utils/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Interface interface {
	Close() error
	Ping(ctx context.Context) error
	QueryRow(ctx context.Context, name string, query string, args ...interface{}) (*sqlx.Row, error)
	Query(ctx context.Context, name string, query string, args ...interface{}) (*sqlx.Rows, error)
	NamedQuery(ctx context.Context, name string, query string, arg interface{}) (*sqlx.Rows, error)

	NamedExec(ctx context.Context, name string, query string, args interface{}) (sql.Result, error)
	Exec(ctx context.Context, name string, query string, args ...interface{}) (sql.Result, error)
}

type TxOptions struct {
	Isolation sql.IsolationLevel
	ReadOnly  bool
}

type Config struct {
	Driver   string
	User     string
	Password string
	Host     string
	DB       string
	Port     string
}

type command struct {
	cfg Config
	db  *sqlx.DB
	log log.Interface
}

func Init(cfg Config, log log.Interface) Interface {
	cmd := &command{
		cfg: cfg,
		log: log,
	}
	db := cmd.connect()

	cmd.db = db

	return cmd
}

func (c *command) connect() *sqlx.DB {
	ctx := context.Background()
	cfg := c.cfg
	uri := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?tls=false&parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)

	// Initialize the DB connection pool
	db, err := sql.Open("mysql", uri)
	if err != nil {
		c.log.Fatal(ctx, fmt.Sprintf("Failed to initialize database connection: %v", err))
	}

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		c.log.Fatal(ctx, fmt.Sprintf("Failed to connect to database: %v", err))
	} else {
		c.log.Info(ctx, "Successfully connected to the database!")
	}

	sqlxDB := sqlx.NewDb(db, cfg.Driver)

	return sqlxDB
}

func (c *command) Close() error {
	return c.db.Close()
}

func (c *command) Ping(ctx context.Context) error {
	return c.db.PingContext(ctx)
}

// QueryRow should be avoided as it cannot be mocked using ExpectQuery
func (c *command) QueryRow(ctx context.Context, name string, query string, args ...interface{}) (*sqlx.Row, error) {
	row := c.db.QueryRowxContext(ctx, query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	return row, nil
}

func (c *command) Query(ctx context.Context, name string, query string, args ...interface{}) (*sqlx.Rows, error) {
	res, err := c.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *command) NamedQuery(ctx context.Context, name string, query string, arg interface{}) (*sqlx.Rows, error) {
	res, err := c.db.NamedQueryContext(ctx, query, arg)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *command) NamedExec(ctx context.Context, name string, query string, args interface{}) (sql.Result, error) {
	res, err := c.db.NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *command) Exec(ctx context.Context, name string, query string, args ...interface{}) (sql.Result, error) {
	return c.db.ExecContext(ctx, query, args...)
}
