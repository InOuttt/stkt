package utils

import (
	"github.com/inouttt/stkt/test_2/internal/entity"
	"github.com/inouttt/stkt/test_2/pkg/utils"
)

func MapToLog(mp map[string]string) entity.Log {
	return entity.Log{
		RequestBody: utils.ConvertMapToString(mp),
	}
}
