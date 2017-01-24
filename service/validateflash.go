package service

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

type ValidateFlashService struct {
	Width, Height float64
}

func (r *ValidateFlashService) GetConfig() string {
	jsonconf, err := config.NewConfig("json", "/www/go/src/royaltools/conf/validate_flash.conf")
	if err != nil {
		fmt.Printf("parse json error")
		return "error"
	}
	return jsonconf.String("name")
}
func (r *ValidateFlashService) Area() float64 {
	return r.Width * r.Height
}
