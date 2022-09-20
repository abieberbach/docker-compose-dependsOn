package waiting

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/sijms/go-ora/v2"
	"strings"
	"time"
)

func NewDatabaseCheck(waitDBSelectKey, databaseDriver, databaseUrl, databaseUser, databasePassword string) CheckTask {
	return func(timeout time.Duration) bool {
		query, value := splitKeyWithValue(waitDBSelectKey)
		db, err := sql.Open(databaseDriver, enrichUrlWithUserAndPassword(databaseUrl, databaseUser, databasePassword))
		fmt.Printf("Checking SQL on server (%v) %v = %v: ", enrichUrlWithUserAndPassword(databaseUrl, databaseUser, "****"), query, value)
		if err != nil {
			fmt.Printf("error=\"%v\"\n", err)
			return false
		}
		defer db.Close()
		var result string
		err = db.QueryRow(query).Scan(&result)
		if err != nil {
			fmt.Printf("error=\"%v\"\n", err)
			return false
		}
		checkResult := result == value
		fmt.Printf("current value=\"%v\" (result: %v)\n", result, checkResult)
		return checkResult

	}
}

func enrichUrlWithUserAndPassword(url string, user string, password string) string {
	return strings.ReplaceAll(strings.ReplaceAll(url, "#user", user), "#password", password)
}

func splitKeyWithValue(keyWithValue string) (key string, value string) {
	lastIndex := strings.LastIndex(keyWithValue, "=")
	if lastIndex == -1 {
		key = keyWithValue
		value = ""
	} else {
		key = keyWithValue[:lastIndex]
		value = keyWithValue[lastIndex+1:]
	}
	return
}
