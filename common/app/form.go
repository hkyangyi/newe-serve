package app

import (
	"newe-serve/common/nelog"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func BindAndValid(c *gin.Context, form interface{}) error {
	err := c.Bind(form)
	if err != nil {
		nelog.Error(err)
		return err
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		nelog.Error(err)
		return err
	}
	if !check {
		MarkErrors(valid.Errors)
		return err
	}

	return nil
}

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		nelog.Error(err.Key, err.Message)
	}
	return
}
