package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
)

type Pairs struct {
	Id  int
	Chs string
	En  string
}

func init() {
	orm.RegisterDataBase("default", "sqlite3", "pairs.db")
	orm.RegisterModel(new(Pairs))
	orm.RunSyncdb("default", false, true)
}
