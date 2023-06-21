package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSQL(t *testing.T) {
	db:= GetConnection()
	defer db.Close()

	ctx:=context.Background()

	query:= "INSERT INTO customer(id,name) VALUES('azhi','AZHI')"
	_,err:= db.ExecContext(ctx,query)
	if err != nil{
		panic(err)
	}

	fmt.Println("Sukses insert a new customer")
}

func TestQuerySQL(t *testing.T) {
	db:= GetConnection()
	defer db.Close()

	ctx:=context.Background()

	query:= "SELECT id,name FROM customer"
	rows,err:= db.QueryContext(ctx,query)
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		var id,name string
		err:=rows.Scan(&id,&name)
		if err != nil{
			panic(err)
		}
		fmt.Println("id : ",id)
		fmt.Println("name : ",name)
	}
}

func TestQuerySQLComplex(t *testing.T) {
	db:= GetConnection()
	defer db.Close()

	ctx:=context.Background()

	query:= "SELECT id,name,email,balance,rating,birth_date,married,created_at FROM customer"
	rows,err:= db.QueryContext(ctx,query)
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		var id,name string
		var email sql.NullString
		var balance int
		var rating float64
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool
		err:=rows.Scan(&id,&name,&email,&balance,&rating,&birth_date,&married,&created_at)
		if err != nil{
			panic(err)
		}
		fmt.Println("====================")
		fmt.Println("id : ",id)
		fmt.Println("name : ",name)
		if email.Valid{
			fmt.Println("email : ",email.String)
		}
		fmt.Println("rating : ",rating)
		if birth_date.Valid{
			fmt.Println("birth date : ",birth_date.Time)
		}
		fmt.Println("married : ",married)
		fmt.Println("created at : ",created_at)
	}
}

func TestSQLInjection(t *testing.T) {
	db:= GetConnection()
	defer db.Close()

	ctx:=context.Background()

	username:="admin' ; #"
	password:="salah"

	query:= "SELECT name FROM user WHERE name= '"+username+"' AND password = '"+password+"' LIMIT 1"
	rows,err:= db.QueryContext(ctx,query)
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	if rows.Next(){
		var username string
		err:=rows.Scan(&username)
		if err != nil{
			panic(err)
		}
		fmt.Println("Sukses Login ",username)
	}else{
		fmt.Println("Gagal Login")
	}
}