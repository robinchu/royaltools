package service

import (
	 "github.com/astaxie/beego/config"
	"os"
	_ "strconv"
	"royaltools/util"
	"log"
)

const jsoncontextwitharray = `[
	{
		"url": "user",
		"serviceAPI": "http://www.test.com/user"
	},
	{
		"url": "employee",
		"serviceAPI": "http://www.test.com/employee"
	}
]`

type ValidateFlashService struct {
	Width, Height float64
}

func (r *ValidateFlashService) GetConfig() []string {


	f, err := os.Create("testjsonWithArra.conf")
	if err != nil {
	}
	_, err = f.WriteString(jsoncontextwitharray)
	if err != nil {
		f.Close()
	}
	f.Close()
	defer os.Remove("testjsonWithArra.conf")
	jsonconf, err := config.NewConfig("json", util.RootDir + "/conf/validate_flash.conf")
	if err != nil {
		log.Fatal(err)
	}
	// strconv的用法
	//str1 = jsonconf.String("svn::url::test")
	//str = "name: " + strconv.Itoa(len(strs)) + " as " + strs[0] + " as " + strs[1]
	urlTest := "svn.url.test: " + jsonconf.String("svn::url::test")
	urlQa := "svn.url.qa: " + jsonconf.String("svn::url::qa")
	username := "svn.username: " + jsonconf.String("svn::user")
	password := "svn.password: " + jsonconf.String("svn::password")
	
	//strs := []string{urlTest, urlQa, username, password}
	return []string{urlTest, urlQa, username, password}
}
func (r *ValidateFlashService) Area() float64 {
	return r.Width * r.Height
}
