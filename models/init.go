package models

import (
	"github.com/beego/beego/v2/client/orm"
)

func init() {
	orm.RegisterModel(new(DemoTable))
}
