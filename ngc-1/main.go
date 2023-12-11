package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type User struct {
	Name     string `required:"true" minLen:"5" maxLen:"10"`
	Age      int    `required:"true" max:"100" min:"5"`
	Password string `required:"true"`
	Username string `required:"true"`
	Email    string `required:"true" email:"true"`
}

func main() {
	newUser := User{
		Name:     "jonibo",
		Password: "12345",
		Username: "joni",
		Email:    "jaja@gmail.com",
	}

	err := validateStruct(newUser)
	fmt.Println(err)
}

func validateStruct(s interface{}) error {
	typeVar := reflect.TypeOf(s)
	for i := 0; i < typeVar.NumField(); i++ {
		field := typeVar.Field(i)
		v := reflect.ValueOf(s)

		// required
		if field.Tag.Get("required") == "true" {
			if v.Field(i).Interface() == "" || v.Field(i).IsZero() {
				return fmt.Errorf("%v is required", field.Name)
			}
		}

		// email
		if field.Tag.Get("email") == "true" {
			re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
			if !re.MatchString(reflect.ValueOf(s).Field(i).String()) {
				return fmt.Errorf("%v format is invalid", field.Name)
			}
		}

		switch v.Field(i).Kind() {
		case reflect.Int:
			x := reflect.ValueOf(s).Field(i).Int()

			// check max
			if field.Tag.Get("max") != "" {
				max, err := strconv.Atoi(field.Tag.Get("max"))
				if err != nil || max <= 0 {
					return err
				}
				if x > int64(max) {
					return fmt.Errorf("%v is higher than max", field.Name)
				}
			}

			// check min
			if field.Tag.Get("min") != "" {
				min, err := strconv.Atoi(field.Tag.Get("min"))
				if err != nil || min <= 0 {
					return err
				}
				if x < int64(min) {
					return fmt.Errorf("%v is lower than min", field.Name)
				}
			}
		case reflect.String:
			str := reflect.ValueOf(s).Field(i).String()

			// check max length
			if field.Tag.Get("min") != "" {
				maxLen, err := strconv.Atoi(field.Tag.Get("maxLen"))
				if err != nil || maxLen <= 0 {
					return err
				}
				if len(str) > maxLen {
					return fmt.Errorf("%v length is higher than max", field.Name)
				}
			}

			// check min length
			if field.Tag.Get("min") != "" {
				minLen, err := strconv.Atoi(field.Tag.Get("minLen"))
				if err != nil || minLen <= 0 {
					return err
				}
				if len(str) < minLen {
					return fmt.Errorf("%v length is lower than min", field.Name)
				}
			}
		}

	}
	return nil
}
