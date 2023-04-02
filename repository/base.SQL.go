package repository

import (
	"database/sql"
	"fmt"
)

type SQLTransaction struct{}

func (SQLTransaction) EndTx(tx *sql.Tx, err error) error {
	if tx == nil {
		return fmt.Errorf("database: Invalid Transaction")
	}

	if msg := "rollback"; err != nil {
		if errR := tx.Rollback(); errR != nil {
			msg = fmt.Sprintf("failed when rollback, err :%s", err)
		}

		return fmt.Errorf("database: %s because %w", msg, err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("database: %w", err)
	}

	return nil
}
