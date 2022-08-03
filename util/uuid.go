package util

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func GenUUID() string {

	return fmt.Sprintf("%s", uuid.NewV4())
}
