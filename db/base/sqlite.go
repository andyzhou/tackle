package base

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

/*
 * sql lite db interface
 */

//sql lite info
type SqlLite struct {
	dbFile string
	db *sql.DB
}

//construct
func NewSqlLite() *SqlLite {
	this := &SqlLite{
	}
	return this
}

//open db file
func (s *SqlLite) OpenDBFile(dbFile string) error {
	//check
	if dbFile == "" {
		return errors.New("invalid parameter")
	}

	//init db
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Println("SqlLite, open db failed, err:", err.Error())
		return err
	}

	//sync db
	s.db = db
	return nil
}

//close db
func (s *SqlLite) Close() {
	if s.db != nil {
		s.db.Close()
		s.db = nil
	}
}

//execute
func (s *SqlLite) Execute(
	sql string,
	args []interface{}) (int64, int64, error) {
	var (
		lastInsertId, effectRows int64
		err error
	)
	//check
	if sql == "" || s.db == nil {
		return lastInsertId, effectRows, errors.New("invalid parameter")
	}
	//exec sql
	result, err := s.db.Exec(sql, args...)
	if err != nil {
		return lastInsertId, effectRows, err
	}
	lastInsertId, err = result.LastInsertId()
	effectRows, err = result.RowsAffected()
	if err != nil {
		return lastInsertId, effectRows, err
	}

	return lastInsertId, effectRows, nil
}

//query
func (s *SqlLite) Query(
	sql string,
	args []interface{}) ([]map[string]interface{}, error) {
	var (
		integerVal int
		colSize, i int
		err error
		//tempStr string
		tempSlice = make([]interface{}, 0)
		results = make([]map[string]interface{}, 0)
	)
	if sql == "" || s.db == nil {
		return nil, errors.New("invalid parameter")
	}
	rows, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	//init temp slice
	colSize = len(cols)
	for i = 0; i < colSize; i++ {
		tempSlice = append(tempSlice, new([]byte))
	}

	for rows.Next() {
		//process single row record
		err = rows.Scan(tempSlice...)
		i = 0
		tempMap := make(map[string]interface{})
		for _, col := range cols {
			//tempStr = ""
			switch v := tempSlice[i].(type) {
			case *[]uint8:
				{
					//check integer or string
					integerVal, err = strconv.Atoi(string(*v))
					if err == nil {
						//integer success
						tempMap[col] = integerVal
					}else{
						//string value
						tempMap[col] = string(*v)
					}
				}
			default:
				tempMap[col] = v
			}
			//tempMap[col] = tempSlice[i]
			i++
		}
		results = append(results, tempMap)
	}

	//clear temp variable
	tempSlice = tempSlice[:0]
	return results, nil
}
