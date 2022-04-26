package model

import (
	"strings"

	"github.com/google/uuid"
)

//生成UUID
func GetUUID() string {
	u := uuid.New().String()
	uu := strings.Replace(u, "-", "", -1)
	return uu
}
