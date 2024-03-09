package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"github.com/rodrigobsimon/togodo/internal/usecase"
	_ "github.com/go-sql-driver/mysql"
)

type PersistentImplementationMySQL struct {
	User string
	Password string
	Host string
	Port int
	Database string
}

func (pi PersistentImplementationMySQL) SaveTask(t usecase.Task) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", pi.User, pi.Password, pi.Host, pi.Port, pi.Database))
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	query := "INSERT INTO `task` (`description`, `priority`, `status`) VALUES(?, ?, ?);"
	result, err := db.ExecContext(context.Background(), query, t.Description, t.Priority, t.Status)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (pi PersistentImplementationMySQL) ListTask() ([]usecase.Task, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", pi.User, pi.Password, pi.Host, pi.Port, pi.Database))
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	rows, err := db.Query("SELECT * FROM task ORDER BY priority DESC;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ts []usecase.Task

	for rows.Next() {
		var t usecase.Task
		var id int
		if err := rows.Scan(&id, &t.Description, &t.Priority, &t.Status); err != nil {
			return ts, err
		}
		ts = append(ts, t)
	}

	if err = rows.Err(); err != nil {
		return ts, err
	}

	return ts, nil
}
