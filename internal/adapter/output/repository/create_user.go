package repository

import (
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"go.uber.org/zap"
	"time"
)

func (ur *userRepository) CreateUser(userDomain domain.UserDomain) *rest_err.RestErr {
	userEntity := converter.ConvertUserDomainToEntity(userDomain)
	userEntity.CreatedAt = time.Now()

	if err := ur.db.Create(&userEntity).Error; err != nil {
		ur.logger.Error("erro ao realizar a consulta CreateUser",
			zap.String("err", err.Error()),
		)

		return rest_err.NewInternalServerError("unable create user")
	}

	return nil
}
