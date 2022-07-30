package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"markettracker.com/tracker/internal/domain"
)

type AssetDTO struct {
	Id    string    `db:"id"`
	Date  time.Time `db:"created_at"`
	Price float32   `db:"price"`
}

type PostgresqlConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Dbname          string
	SslMode         string
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
}

type Postgresql struct {
	table string
	conn  *sqlx.DB
}

func NewPostgresql(table string, config PostgresqlConfig) (*Postgresql, error) {
	if config.SslMode == "" {
		config.SslMode = "disable"
	}
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Dbname, config.SslMode,
	)
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	return &Postgresql{
		conn:  db,
		table: table,
	}, nil
}

func (r *Postgresql) Save(asset domain.Asset) error {
	sentence := fmt.Sprintf("INSERT INTO %s (id, created_at, price) VALUES (:id, :created_at, :price)", r.table)
	row, err := r.conn.NamedQuery(sentence, AssetDTO{
		Id:    asset.ID(),
		Date:  asset.Date(),
		Price: asset.Float32Price(),
	})
	log.Println(row)
	return err
}
