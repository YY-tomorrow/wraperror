package dao

import (
	"github.com/pkg/errors"
	"wrapError/common"
)

func FindById(id string) (string, error) {
	var name string
	err := common.DB.QueryRow("select name from user where id = ?", id).Scan(&name)
	if err != nil {
		return name, errors.Wrapf(err, "FindById %v", id)
	}
	return name, nil
}
