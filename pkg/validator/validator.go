package validator

import (
	"fmt"
	"log"
	"net/mail"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/SeyramWood/app/framework/database"
)

func Validate(i interface{}) interface{} {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	var structure = make(map[string]interface{})
	var errors = make([]string, 0)

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	if t.Kind() == reflect.Ptr {
		t = t.Elem() //Gets the type in the type pointer

	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem() //Get the value in the value address
	}

	if t.Kind() != reflect.Struct {
		log.Fatal("Please provide a struct type")
	}

	for i := 0; i < t.NumField(); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if _, ok := t.Field(i).Tag.Lookup("json"); ok {
				msg := validator(i, t, v)
				mut.Lock()
				if msg == "" {
					errors = append(errors, t.Field(i).Tag.Get("json"))
				}
				structure[t.Field(i).Tag.Get("json")] = msg
				mut.Unlock()
			}
		}(i)
	}
	wg.Wait()
	if len(errors) == t.NumField() {
		structure = make(map[string]interface{})
		errors = []string{}
		return nil
	}
	return structure

}

func validator(index int, t reflect.Type, v reflect.Value) interface{} {
	rules := strings.Split(t.Field(index).Tag.Get("validate"), "|")
	field := t.Field(index).Name
	fieldType := t.Field(index).Type.String()
	value := v.Field(index)
	formattedField := formatFieldName(field)

	for _, rule := range rules {
		if rule == "required" && fieldType == "string" && value.Len() == 0 {
			return fmt.Sprintf("The %s field is required", formattedField)
		} else if rule == "required" && fieldType == "bool" && !value.Bool() {
			return fmt.Sprintf("The %s field is required", formattedField)
		} else if fieldType == "string" && value.Len() != 0 {

			switch rule {
			case "string":
				r, _ := regexp.Compile("^[0-9a-zA-Z-+ ]+$")
				if !r.MatchString(value.String()) {
					return fmt.Sprintf("The %s must be a string.", formattedField)
				}
			case "email":
				if _, err := mail.ParseAddress(value.String()); err != nil {
					return fmt.Sprintf("The %s must be a valid email address.", formattedField)
				}
			case "email_phone":
				email, _ := regexp.Compile(`^[A-Za-z]+$`)
				phone, _ := regexp.Compile(`^0\d{9}$`)

				if email.MatchString(value.String()) || strings.Contains(value.String(), "@") {
					if _, err := mail.ParseAddress(value.String()); err != nil {
						return fmt.Sprintf("The %s must be a valid email address.", formattedField)
					}
				} else {
					if phone.MatchString(value.String()) {
						if !phone.MatchString(value.String()) {
							return fmt.Sprintf("The %s field must be a valid phone number.", formattedField)
						}
					}
					return fmt.Sprintf("The %s field must be a valid phone number.", formattedField)
				}

			case "id_card":
				r, _ := regexp.Compile(`^GHA-\d{9}-\d{1}$`)
				if !r.MatchString(value.String()) {
					return "The The ID field must be a valid Ghana Card."
				}
			case "digital_address":
				r, _ := regexp.Compile(`[A-Z]{2}-\d{3}-\d{4}$`)
				if !r.MatchString(value.String()) {
					return "The address field must be a valid digital address."
				}
			default:
				if strings.Contains(rule, ":") {
					r := strings.Split(rule, ":")
					switch r[0] {
					case "max":
						val, _ := strconv.Atoi(r[1])
						if value.Len() > val {
							return fmt.Sprintf("The %s must not be greater than %v characters", formattedField, val)
						}
					case "min":
						val, _ := strconv.Atoi(r[1])
						if value.Len() < val {
							return fmt.Sprintf("The %s must be at least %v characters", formattedField, val)
						}
					case "match":
						var val reflect.Value
						f := strings.ToUpper(string([]byte{r[1][0]})) + strings.ToLower(r[1][1:])
						for i := 0; i < t.NumField(); i++ {
							if t.Field(i).Name == f {
								val = v.Field(i)
								break
							}
						}
						if value.String() != val.String() {
							return fmt.Sprintf("The %s does not matched", "passwords")
						}
					case "unique":
						table := r[1]
						if r := isUsernameExist(value.String(), field, table); r != nil {
							return r
						}
						// case "file":
						// mimes := r[1]
						// fmt.Println(r)

					}

				}

			}
		}
	}

	return ""
}

func formatFieldName(field string) string {
	var text string
	for i := 0; i < len(field); i++ {
		c := string([]byte{field[i]})
		if c == strings.ToUpper(c) {
			if len(text) != 0 {
				text += " "
			}
			text += strings.ToLower(c)
		} else {
			text += strings.ToLower(c)
		}
	}
	return text

}

func snakeCase(field string) string {
	var text string
	for i := 0; i < len(field); i++ {
		c := string([]byte{field[i]})
		if c == strings.ToUpper(c) {
			if len(text) != 0 {
				text += "_"
			}
			text += strings.ToLower(c)
		} else {
			text += strings.ToLower(c)
		}
	}
	return text

}

func isUsernameExist(username, field, table string) interface{} {
	dbField := snakeCase(field)
	fieldName := formatFieldName(field)
	queryStr := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?", dbField, table, dbField)
	db := database.Connect()
	err := db.QueryRow(queryStr, username).Scan(&username)
	if err == nil {
		return fmt.Sprintf("The %s is already taken", fieldName)
	}
	defer db.Close()
	return nil
}
