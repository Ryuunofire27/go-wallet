package uow_test

import (
	"context"
	"errors"
	"go-wallet/internal/shared/src/application/port/generators"
	"go-wallet/internal/shared/src/application/port/unit_of_work"
	"go-wallet/internal/shared/src/domain/entity"
	"go-wallet/internal/shared/src/infrastructure/database/repository"
	infraUnitOfWork "go-wallet/internal/shared/src/infrastructure/database/unit_of_work"
	infraGenerators "go-wallet/internal/shared/src/infrastructure/generators"
	testutil "go-wallet/internal/shared/test/util"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestUser struct {
	entity.Entity[string]
	Name  string
	Email string
}

type TestUserModel struct {
	ID    string
	Name  string
	Email string
}

type TestUserRepository interface {
	Create(ctx context.Context, domain *TestUser) (*TestUser, error)
}

type GormTestUserRepository struct {
	repository.GormEntityRepository[TestUser, string, TestUserModel]
}

type TestUserRepositoryMapper struct{}

func (m TestUserRepositoryMapper) ToDomain(entity *TestUserModel) (*TestUser, error) {
	return &TestUser{
		ID:    entity.ID,
		Name:  entity.Name,
		Email: entity.Email,
	}, nil
}

func (m TestUserRepositoryMapper) ToModel(domain *TestUser) (*TestUserModel, error) {
	return &TestUserModel{
		ID:    domain.ID,
		Name:  domain.Name,
		Email: domain.Email,
	}, nil
}

type TestCreateUserUseCase struct {
	uow            unit_of_work.TransactionManager
	userRepository TestUserRepository
	uuidGenerator  generators.UUIDGenerator
	afterCreate    func() error
}

type CreateUserInput struct {
	Name  string
	Email string
}

func (uc *TestCreateUserUseCase) Execute(
	ctx context.Context,
	input CreateUserInput,
) (*TestUser, error) {
	return unit_of_work.Transaction(
		ctx,
		uc.uow,
		func(ctx context.Context) (*TestUser, error) {
			uuid, err := uc.uuidGenerator.Generate()

			if err != nil {
				return nil, err
			}

			user, err := uc.userRepository.Create(ctx, &TestUser{
				Name:  input.Name,
				Email: input.Email,
				ID:    uuid,
			})

			if err != nil {
				return nil, err
			}

			if uc.afterCreate != nil {
				err := uc.afterCreate()
				if err != nil {
					return nil, err
				}
			}

			return user, nil
		},
	)
}

func TestUnitOfWork(t *testing.T) {
	db := testutil.SetupPostgres(t)

	if err := db.AutoMigrate(&TestUserModel{}); err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	uow := infraUnitOfWork.NewGormUnitOfWork(db)

	mapper := &TestUserRepositoryMapper{}

	userRepository := &GormTestUserRepository{
		GormEntityRepository: *repository.NewGormEntityRepository(
			uow,
			mapper,
			"GormTestUserRepository",
		),
	}

	uuidGenerator := &infraGenerators.UUIDGenerator{}

	t.Run("should commit transaction", func(t *testing.T) {

		commitUseCase := &TestCreateUserUseCase{
			uow:            uow,
			userRepository: userRepository,
			uuidGenerator:  uuidGenerator,
		}
		t.Cleanup(func() {
			db.Exec("TRUNCATE TABLE test_user_models CASCADE")
		})
		input := CreateUserInput{
			Name:  "John",
			Email: "john@test.com",
		}

		user, err := commitUseCase.Execute(context.Background(), input)

		require.NoError(t, err)
		require.NotNil(t, user)
		require.NotEmpty(t, user.ID)

		var persistedUser TestUserModel

		err = db.First(&persistedUser, "id = ?", user.ID).Error

		require.NoError(t, err)
		require.Equal(t, input.Name, persistedUser.Name)
		require.Equal(t, input.Email, persistedUser.Email)

	})

	t.Run("should rollback transaction", func(t *testing.T) {

		rollbackUseCase := &TestCreateUserUseCase{
			uow:            uow,
			userRepository: userRepository,
			uuidGenerator:  uuidGenerator,
			afterCreate: func() error {
				return errors.New("Test error")
			},
		}

		_, err := rollbackUseCase.Execute(context.Background(), CreateUserInput{
			Name:  "Jane",
			Email: "jane@test.com",
		})

		require.Error(t, err)

		var count int64
		err = db.Model(&TestUserModel{}).
			Count(&count).Error

		require.NoError(t, err)
		require.Zero(t, count)

	})
}
