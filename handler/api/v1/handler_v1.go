package v1

import (
	"fmt"
	setting "logger/config"
)

var (
	s = setting.LoadSettings()
)

func HandlerV1 () {
	fmt.Println("setting :", s.API_PREFIX_V1)
}