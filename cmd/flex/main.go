package main

import (
	"flex"
	"flex/pkg/rfunc"
	"fmt"
	"log"
	"time"
)

func main() {
	//pType()
	//pValue()
	//pFunc()
	pStruct()
}

func pStruct() {
	s := flex.Struct(User{})
	fmt.Println("fields count:", s.FieldCount())
	fmt.Println("fields:")
	for _, field := range s.FieldsRow() {
		fmt.Println(field)
		fmt.Println(flex.Field(field).Tag("ef"))
	}

	fmt.Println("\nmethods count:", s.MethodCount())
	fmt.Println("methods:")
	for _, method := range s.Methods() {
		fmt.Println(method)
		fm := rfunc.NewMethod(method.Func)
		fmt.Println("arguments:", fm.Arguments())
		fmt.Println("returns:", fm.Returns())

		result, err := fm.Call(User{}, "message")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("result:", result)
	}
}

func pType() {
	obj := flex.Object(User{})
	t := flex.Type(obj.Type())

	fmt.Println("is primitive:", t.IsPrimitive())
	fmt.Println("is reference:", t.IsReference())
}

func pValue() {
	//
}

func pFunc() {
	f := flex.Func(plus)
	fmt.Println("arguments count:", f.ArgumentsCount())
	fmt.Println("return count:", f.ReturnCount())
	fmt.Println("arguments:", f.Arguments())
	fmt.Println("returns:", f.Returns())

	results, err := f.Call(1, 2)
	if err != nil {
		fmt.Println("i am here #1")
		log.Fatal(err)
	}

	fmt.Println("results:", results)
}

func plus(a, b int) int {
	return a + b
}

type Entity struct {
	ID        string    `db:"id" ef:"primary,nn"`
	CreatedAt time.Time `db:"create_at" ef:"nn"`
	UpdatedAt time.Time `db:"updated_at" ef:"nn"`

	SomeBool bool
}

type User struct {
	Entity

	Username string `db:"username"`
	Password string `db:"password"`

	LastName   *string `db:"last_name"`
	FirstName  *string `db:"first_name"`
	MiddleName *string `db:"middle_name"`

	Age       int
	BirthDate time.Time `db:"birth_date"`

	isResident bool
}

func (user User) Send(message string) error {
	return nil
}
