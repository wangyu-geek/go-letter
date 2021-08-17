package models

import (
	"encoding/json"
	"leeter/pkg/db"
	"leeter/pkg/errors"
)

type DemoTable struct {
	Id	int
	Hello string `orm:"type(text);description(文本)"` 
}

func (t *TTSRecord) TableName() string {
	return "demo_table"
}

func (t *DemoTable) Insert() (int64, error) {
	o := db.GetORM()
	return o.Insert(t)
}
