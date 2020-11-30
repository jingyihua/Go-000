package api

/*
	error转换成前端显示用语
*/

import (
	"database/sql"
	"github.com/pkg/errors"
	"week02/comDef"
)

const (
	NotFindAboutUserInfo = "没有找到相关用户信息"
	InputDataIsBadData   = "输入信息错误"
	SurpriseError        = "恭喜发现宝藏, 来当我测试吧"
)

func ErrTranslate(err error) string {
	rootErr := errors.Cause(err)
	switch {
	case errors.Is(rootErr, sql.ErrNoRows):
		return NotFindAboutUserInfo
	case errors.Is(rootErr, comDef.SqlFindParamErr):
		return InputDataIsBadData
	default:
		return SurpriseError
	}
}
