package app

import (
	query "coding-test-be/repository/postgres_init"
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func (x *App) initDB(ctx context.Context) (err error) {
	var (
		host     = os.Getenv("MYSQL_HOST")
		port     = os.Getenv("MYSQL_PORT")
		user     = os.Getenv("MYSQL_USER")
		password = os.Getenv("MYSQL_PASSWORD")
		dbname   = os.Getenv("MYSQL_DBNAME")
	)

	apiUrl := os.Getenv("API_URL")
	if apiUrl != "127.0.0.1" {
		split := strings.Split(apiUrl, "//")
		split2 := strings.Split(split[1], ":")
		host = split2[0]
	}

	//sqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	sqlconn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	x.DB, err = sql.Open("mysql", sqlconn)
	if err != nil {
		fmt.Println(sqlconn)
		return err
	}

	x.DB.SetConnMaxLifetime(time.Second * 14400)

	return x.DB.PingContext(ctx)
}

func (x *App) initSchema(ctx context.Context) (err error) {
	_, err = x.DB.Exec(query.CreateSchema)
	if err != nil {
		fmt.Println("===INI?===")
		return err
	}

	return nil
}

func (x *App) initTable(ctx context.Context) (err error) {
	queryactivity := `CREATE TABLE IF NOT EXISTS challenge_2_be.activities(
		activity_id bigint AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',
		created_at timestamp NOT NULL DEFAULT now(),
		updated_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',
		updated_at timestamp NOT NULL DEFAULT now(),
		deleted_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',
		deleted_at timestamp NOT NULL DEFAULT now()
	);`

	querytodo := `CREATE TABLE IF NOT EXISTS challenge_2_be.todos(
		todo_id bigint AUTO_INCREMENT PRIMARY KEY,
		activity_group_id int8 NOT NULL,
		title VARCHAR(255) NOT NULL,
		priority VARCHAR(255) NOT NULL,
		is_active BOOLEAN NOT NULL DEFAULT true,
		created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',
		created_at timestamp NOT NULL DEFAULT now(),
		updated_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',
		updated_at timestamp NOT NULL DEFAULT now(),
		deleted_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',
		deleted_at timestamp NOT NULL DEFAULT now(),
		CONSTRAINT user_fk_products FOREIGN KEY (activity_group_id) REFERENCES challenge_2_be.activities (activity_id)
	);`

	_, err = x.DB.Exec(queryactivity)
	if err != nil {
		return err
	}

	_, err = x.DB.Exec(querytodo)
	if err != nil {
		return err
	}

	return nil
}
