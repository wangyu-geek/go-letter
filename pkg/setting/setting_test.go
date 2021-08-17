package setting

import (
	"log"
	"testing"
)

func Test_setAppPath(t *testing.T) {
	//appPath := setAppPath()
	//log.Println(appPath)
	appPath := GetAppPath()
	log.Println(appPath)
}
