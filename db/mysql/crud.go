package db

import (
	"github.com/miles990/golang-mysql-example/util/token"
	"github.com/miles990/golang-mysql-example/db/models"	
)

import (
	"fmt"
	"strings"
	"reflect"
	"time"
)

func getMysqlInsertStringAndArgs(tableName string, o interface{}) (string, []interface{}) {
	var colNames = []string{}
	var colValues = []string{}
	var rtnArgs = []interface{}{}
	var rtnString = "insert into " + tableName + "("
	var t = reflect.TypeOf(o)
	var v = reflect.ValueOf(o)
	var colName string
	for i := 0; i < t.NumField(); i++ {
		if colName = t.Field(i).Tag.Get("db"); colName != "" {
			if t.Field(i).Tag.Get("dbinsert") != "false" {
				colNames = append(colNames, fmt.Sprintf("`%s`", colName))
				colValues = append(colValues, "?")
				rtnArgs = append(rtnArgs, v.Field(i).Interface())
			}
		}
	}
	rtnString = rtnString + strings.Join(colNames, ",") + ") values(" + strings.Join(colValues, ",") + ")"

	return rtnString, rtnArgs
}

func ResetTable() {
	_ = dropTable()
	_ = createTable()
}

func CreateAccount(name string) (dbAccount *models.Account, err error) {

	err = ensureConnection()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	
	dbAccount = new(models.Account)
	dbAccount.UID, err = token.AccountUidWorker.NextId()
	if err != nil {		
		fmt.Println(err.Error())	
		return nil, err
	}
	
	dbAccount.CreateTimestamp = time.Now().UTC().Unix()
	dbAccount.Name = name



	var insertString string
	var insertArgs []interface{}

	insertString, insertArgs = getMysqlInsertStringAndArgs("accounts", *dbAccount)
	if _, err = connector.Exec(insertString, insertArgs...); err != nil {			
		fmt.Println(insertString, err.Error())
		return nil, err
	}
	fmt.Println("新增成功")
	
	return dbAccount, nil
}

func GetAccounts() ([]*models.Account, error) {
	var err error
	err = ensureConnection()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var accounts []*models.Account

	var sqlcommand = `select * from gamedata.accounts `

	sqlcommand += ` order by createtimestamp asc;`

	err = connector.Select(&accounts, sqlcommand)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return accounts, nil
}
