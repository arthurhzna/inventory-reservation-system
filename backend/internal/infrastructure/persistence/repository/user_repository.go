package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/entity"

	repositoryiface "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/repository"

	dbtx "github.com/arthurhzna/inventory-reservation-system/backend/internal/infrastructure/persistence/dbtx"
)

type userRepository struct {
	db dbtx.DBTX
}

func NewUserRepository(
	db dbtx.DBTX,
) repositoryiface.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(
	ctx context.Context,
	user *entity.User,
) error {

	query := `
		INSERT INTO users (
			uuid,
			name,
			email,
			password,
			role_id,
			created_at,
			updated_at
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		)
		RETURNING id
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		user.UUID,
		user.Name,
		user.Email,
		user.Password,
		user.RoleID,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(
		&user.ID,
	)
}

func (r *userRepository) FindByUUID(
	ctx context.Context,
	uuid string,
) (*entity.User, error) {

	query := `
		SELECT
			id,
			uuid,
			name,
			email,
			password,
			role_id,
			created_at,
			updated_at
		FROM users
		WHERE uuid = $1
	`

	var user entity.User

	err := r.db.QueryRowxContext(
		ctx,
		query,
		uuid,
	).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*entity.User, error) {

	query := `
		SELECT
			id,
			uuid,
			name,
			email,
			password,
			role_id,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`

	var user entity.User

	err := r.db.QueryRowxContext(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(
	ctx context.Context,
	user *entity.User,
) error {

	query := `
		UPDATE users
		SET
			name = $1,
			email = $2,
			password = $3,
			role_id = $4,
			updated_at = $5
		WHERE uuid = $6
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		user.RoleID,
		user.UpdatedAt,
		user.UUID,
	)

	return err
}

func (r *userRepository) DeleteByUUID(
	ctx context.Context,
	uuid string,
) error {

	query := `
		DELETE FROM users
		WHERE uuid = $1
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		uuid,
	)

	return err
}
