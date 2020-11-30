package service

import (
	"github.com/pkg/errors"
	"week02/comDef"
	"week02/dao"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (serve *Service) GetUserAge(id int) (int, error) {
	if id < 0 {
		return 0, errors.Wrapf(comDef.SqlFindParamErr, "bad id: %d", id)
	}

	age, err := dao.GetUserAge(id)
	if err != nil {
		return 0, err
	}
	return age, nil
}
