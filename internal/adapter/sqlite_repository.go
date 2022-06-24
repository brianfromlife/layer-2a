package adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/trevatk/common/database"
	"github.com/trevatk/layer-2a/internal/domain/user"
)

// sqlUser
type sqlUser struct {
	ID          int64  `db:"id"`
	Password    string `db:"password"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Phone       string `db:"phone"`
	CountryCode string `db:"country_code"`
	CreatedAt   string `db:"created_at"`
}

// SQLiteRepository ...
type SQLiteRepository struct {
	db *database.DB
	f  *user.Factory
}

// compile time interface verification
var _ user.IRepository = (*SQLiteRepository)(nil)

// NewRepository create new instance
func NewRepository(database *database.DB, factory *user.Factory) *SQLiteRepository {
	return &SQLiteRepository{db: database, f: factory}
}

// CreateUser insert and read new user from database
func (repo *SQLiteRepository) CreateUser(ctx context.Context, u *user.User) (*user.User, error) {

	sqlUser := &sqlUser{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Phone:       u.Phone,
		CountryCode: u.CountryCode,
	}

	stmt :=
		`
	INSERT INTO users (name, email, phone, country_code) VALUES (?,?,?,?);
	`

	timeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	conn, tx, err := repo.nTx(timeout)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	result, err := tx.ExecContext(timeout, stmt, sqlUser.Name, sqlUser.Email, sqlUser.Email, sqlUser.Phone, sqlUser.CountryCode)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("tx.ExecContext: %v", err)
	}

	liid, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()

		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}

		return nil, fmt.Errorf("result.LastInsertId: %v", err)
	}

	stmt =
		`
	SELECT id, name, email, country_code, phone, created_at FROM users WHERE id = ? AND deleted_at IS NULL;
	`

	row := tx.QueryRowContext(ctx, stmt, liid)
	err = row.Scan(
		&sqlUser.ID,
		&sqlUser.Name,
		&sqlUser.Email,
		&sqlUser.CountryCode,
		&sqlUser.Phone,
		&sqlUser.CreatedAt,
	)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("row.Scan: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("tx.Commit: %v", err)
	}

	domainUser, err := repo.f.UnmarshalUserFromDatabase(
		sqlUser.ID,
		sqlUser.Name,
		sqlUser.Email,
		sqlUser.CountryCode,
		sqlUser.Phone,
		sqlUser.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("f.UnmarshalUserFromDatabase: %v", err)
	}

	return domainUser, nil
}

// ReadUser read user by id
func (repo *SQLiteRepository) ReadUser(ctx context.Context, ID int64) (*user.User, error) {

	stmt :=
		`
	SELECT id, name, email, country_code, phone, created_at FROM users WHERE id = ? AND deleted_at IS NULL;
	`

	timeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	conn, tx, err := repo.nTx(timeout)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var u sqlUser
	row := tx.QueryRowContext(timeout, stmt, ID)
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.CountryCode, &u.Phone, &u.CreatedAt)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("row.Scan: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("tx.Commit: %v", err)
	}

	domainUser, err := repo.f.UnmarshalUserFromDatabase(
		u.ID,
		u.Name,
		u.Email,
		u.CountryCode,
		u.Phone,
		u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return domainUser, nil
}

// ReadAllUsers read all users from database
func (repo *SQLiteRepository) ReadAllUsers(ctx context.Context) ([]*user.User, error) {

	stmt :=
		`
	SELECT id, name, email, country_code, phone, created_at FROM users;
	`

	timeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	conn, tx, err := repo.nTx(timeout)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := tx.QueryContext(timeout, stmt)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("tx.QueryContext: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("tx.Commit: %v", err)
	}

	users := make([]*user.User, 0)

	for rows.Next() {

		var su sqlUser
		err = rows.Scan(
			&su.ID,
			&su.Name,
			&su.Email,
			&su.CountryCode,
			&su.Phone,
			&su.CreatedAt,
		)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("rows.Scan: %v", err)
		}

		user, err := repo.f.UnmarshalUserFromDatabase(
			su.ID,
			su.Name,
			su.Email,
			su.CountryCode,
			su.Phone,
			su.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("f.UnmarshalUserFromDatabase: %v", err)
		}

		users = append(users, user)
	}

	return users, nil
}

// UpdateUser update and read user record in database
func (repo *SQLiteRepository) UpdateUser(ctx context.Context, u *user.User) (*user.User, error) {

	stmt :=
		`
	UPDATE users
	SET name = ?, email = ?, country_code = ?, phone = ?, updated_at = CURRENT_TIMESTAMP
	WHERE id = ?;
	`

	timeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	conn, tx, err := repo.nTx(timeout)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	result, err := tx.ExecContext(timeout, stmt, u.Name, u.Email, u.CountryCode, u.Phone, u.ID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("tx.ExecContext: %v", err)
	}

	a, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("result.RowsAffect: %v", err)
	}

	if !(a > 0) {
		tx.Rollback()
		return nil, errors.New("no record updated")
	}

	stmt =
		`
	SELECT id, name, email, country_code, phone, created_at FROM users WHERE id = ?;
	`

	var sqlUser sqlUser
	row := tx.QueryRowContext(timeout, stmt, u.ID)
	err = row.Scan(
		&sqlUser.ID,
		&sqlUser.Name,
		&sqlUser.Email,
		&sqlUser.CountryCode,
		&sqlUser.Phone,
		&sqlUser.CreatedAt,
	)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("row.Scan: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("tx.Commit: %v", err)
	}

	domainUser, err := repo.f.UnmarshalUserFromDatabase(
		sqlUser.ID,
		sqlUser.Name,
		sqlUser.Email,
		sqlUser.CountryCode,
		sqlUser.Phone,
		sqlUser.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return domainUser, nil
}

// DeleteUser soft delete user record
func (repo *SQLiteRepository) DeleteUser(ctx context.Context, ID int64) error {

	stmt :=
		`
	UPDATE users
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE id = ?;
	`

	timeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	conn, tx, err := repo.nTx(timeout)
	if err != nil {
		return err
	}
	defer conn.Close()

	result, err := tx.ExecContext(timeout, stmt, ID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("tx.ExecContext, %v", err)
	}

	a, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("result.RowsAffected: %v", err)
	}

	if !(a > 0) {
		tx.Rollback()
		return errors.New("no record deleted")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("tx.Commit: %v", err)
	}

	return nil
}

func (repo *SQLiteRepository) nTx(ctx context.Context) (*sql.Conn, *sql.Tx, error) {

	conn, err := repo.db.Conn(ctx)
	if err != nil {
		return nil, nil, err
	}

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return conn, tx, nil
}
