package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Query map[string]interface{}

func (model Query) Find(tabel string) Query {
	var values []interface{}
	var where []string
	for k, v := range model {
		values = append(values, v)
		//MySQL Way: where = append(where, fmt.Sprintf("%s = ?", k))
		where = append(where, fmt.Sprintf(`"%s" = %s`, k, "$"+strconv.Itoa(len(values))))
	}
	string := ("SELECT name FROM users WHERE " + strings.Join(where, " AND "))
	//for testing purposes i didn't ran actual query, just print it in the console and returned JSON back
	fmt.Println(string)
	return model

}
