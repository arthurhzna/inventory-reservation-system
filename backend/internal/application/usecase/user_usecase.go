package usecase

import (
	"context"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/dto/request"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/dto/response"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/entity"

	repositoryinterface "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/repository"

	securitydomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/security"

	servicedomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/service"

	errordomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/error"

	enumdomain "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/enum"
)

type UserUseCaseInterface interface {
	Register(
		ctx context.Context,
		req *request.RegisterUserRequest,
	) (*response.RegisterResponse, error)

	Login(
		ctx context.Context,
		req *request.LoginUserRequest,
	) (*response.LoginResponse, error)
}

type UserUseCase struct {
	uow repositoryinterface.UnitOfWork

	passwordHasher securitydomain.PasswordHasher
	tokenService   securitydomain.TokenService
	uuidGenerator  servicedomain.UUIDGenerator
}

func NewUserUseCase(
	uow repositoryinterface.UnitOfWork,
	uuidGenerator servicedomain.UUIDGenerator,
) UserUseCaseInterface {
	return &UserUseCase{
		uow:           uow,
		uuidGenerator: uuidGenerator,
	}
}

func (u *UserUseCase) Register(
	ctx context.Context,
	req *request.RegisterUserRequest,
) (*response.RegisterResponse, error) {

	roleID, ok := enumdomain.RoleNameToID[req.Role]
	if !ok {
		return nil, errordomain.ErrInvalidRole
	}

	var registeredUser *entity.User

	err := u.uow.WithTransaction(
		ctx,
		func(txUow repositoryinterface.UnitOfWork) error {

			existingUser, err := txUow.
				UserRepository().
				FindByEmail(
					ctx,
					req.Email,
				)

			if err != nil {
				return err
			}

			if existingUser != nil {
				return errordomain.ErrEmailAlreadyExist
			}

			hashedPassword, err := u.passwordHasher.Hash(
				req.Password,
			)

			if err != nil {
				return err
			}

			user := &entity.User{
				UUID:     u.uuidGenerator.New(),
				Name:     req.Name,
				Email:    req.Email,
				Password: hashedPassword,
				RoleID:   roleID,
			}

			err = txUow.
				UserRepository().
				Create(
					ctx,
					user,
				)

			if err != nil {
				return err
			}

			registeredUser = user

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return &response.RegisterResponse{
		User: response.UserResponse{
			UUID:  registeredUser.UUID,
			Name:  registeredUser.Name,
			Email: registeredUser.Email,
			Role:  enumdomain.RoleIDToName[registeredUser.RoleID],
		},
	}, nil
}

func (u *UserUseCase) Login(
	ctx context.Context,
	req *request.LoginUserRequest,
) (*response.LoginResponse, error) {

	user, err := u.uow.
		UserRepository().
		FindByEmail(
			ctx,
			req.Email,
		)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errordomain.ErrEmailNotFound
	}

	isValid := u.passwordHasher.Check(
		req.Password,
		user.Password,
	)

	if !isValid {
		return nil, errordomain.ErrInvalidCredential
	}

	token, err := u.tokenService.Generate(
		user.ID,
		user.RoleID,
	)

	if err != nil {
		return nil, err
	}

	return &response.LoginResponse{
		User: response.UserResponse{
			UUID:  user.UUID,
			Name:  user.Name,
			Email: user.Email,
			Role:  enumdomain.RoleIDToName[user.RoleID],
		},
		Token: token,
	}, nil
}
