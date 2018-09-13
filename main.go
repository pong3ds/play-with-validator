package main

import (
	"database/sql"
	"fmt"

	"github.com/3dsinteractive/deepcopier"
	"github.com/3dsinteractive/govalidator"
)

// Post is post
type Post struct {
	Title    string `valid:"alphanum,required"`
	Message  string `valid:"duck,ascii"`
	AuthorIP string `valid:"ipv4"`
	Date     string `valid:"-"`
}

func validateSimple() {
	post := &Post{
		Title:    "My Example Post",
		Message:  "duck",
		AuthorIP: "123.234.54.3",
	}

	// Add your own struct validation tags
	govalidator.TagMap["duck"] = govalidator.Validator(func(str string) bool {
		return str == "duck"
	})

	result, err := govalidator.ValidateStruct(post)
	if err != nil {
		println(fmt.Sprintf("%+v", ErrorToJson(err)))
	}
	println(result)
}

// User model
type User struct {
	// Basic string field
	Name string
	// Deepcopier supports https://golang.org/pkg/database/sql/driver/#Valuer
	Email       sql.NullString
	Description string
}

// MethodThatTakesContext is the tranformation method
func (u *User) MethodThatTakesContext(ctx map[string]interface{}) string {
	return ctx["hello"].(string)
}

// Method2ThatTakesContext is the transformation method
func (u *User) Method2ThatTakesContext(ctx map[string]interface{}) string {
	return ctx["hello2"].(string)
}

// UserResource model
type UserResource struct {
	DisplayName             string `deepcopier:"field:Name"`
	Description             string `deepcopier:"force"`
	SkipMe                  string `deepcopier:"skip"`
	MethodThatTakesContext  string `deepcopier:"context"`
	Method2ThatTakesContext string `deepcopier:"context"`
	Email                   string `deepcopier:"force"`
}

func transformObject() {
	user := &User{
		Name: "gilles232",
		Email: sql.NullString{
			Valid:  true,
			String: "gilles@example.com",
		},
		Description: "Description haha",
	}

	resource := &UserResource{}

	deepcopier.Copy(user).
		WithContext(map[string]interface{}{"hello": "Hello", "hello2": "Hello2"}).
		To(resource)

	fmt.Println(resource.DisplayName)
	fmt.Println(resource.Email)
	fmt.Println(resource.MethodThatTakesContext)
	fmt.Println(resource.Method2ThatTakesContext)
	fmt.Println(resource.Description)
}

func main() {
	validateSimple()
	// transformObject()
}
