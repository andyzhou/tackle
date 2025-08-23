package base

import (
	"errors"
	"fmt"
	"github.com/andyzhou/tackle/conf"
	"github.com/andyzhou/tackle/define"
)

/*
 * db base opt face
 */

//face info
type Base struct {
}

//open sqlite db instance
func (f *Base) OpenDB(dbFileName string) (*SqlLite, error) {
	//check
	if dbFileName == "" {
		return nil, errors.New("invalid parameter")
	}

	//get main conf
	mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
	privatePath := mainConf.PrivatePath
	dbPath := fmt.Sprintf("%v/%v", privatePath, define.StorageOfDB)

	//get db path
	dbFilePath := fmt.Sprintf("%v/%v", dbPath, dbFileName)

	//init sqlite db instance
	db := NewSqlLite()
	err := db.OpenDBFile(dbFilePath)
	if err != nil {
		return nil, err
	}
	return db, nil
}