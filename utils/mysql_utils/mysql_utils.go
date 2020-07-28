package mysql_utils

import (
	"github.com/alvarezcarlos/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

func parseError(err error) *errors.RestErr{
	//sqlErr, ok := err.(*mysql.MySQLError)
	//
	//if !ok {
	//
	//}
}