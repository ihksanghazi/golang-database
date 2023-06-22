package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

func TestSQLInjectionSafe(t *testing.T) {
	db:= GetConnection()
	defer db.Close()

	ctx:=context.Background()

	username:="admin' ; #"
	password:="salah"

	query:= "SELECT name FROM user WHERE name = ? AND password = ? LIMIT 1"
	rows,err:= db.QueryContext(ctx,query,username,password)
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

func TestAutoIncrement(t *testing.T) {
	db:= GetConnection();
	defer db.Close()

	ctx:= context.Background()

	email:= "sandy@gmail.com"
	comment:= "apa aja"

	query:= "INSERT INTO comments(email,comment) VALUES(?,?)"

	result,err:= db.ExecContext(ctx,query,email,comment)
	if err != nil{
		panic(err)
	}
	insertId,err:=result.LastInsertId()
	if err != nil{
		panic(err)
	}

	fmt.Println("Sukses Insert Comment with id ",insertId)

}

func TestPrepareStatment(t *testing.T) {
	db:=GetConnection()
	defer db.Close()

	ctx:=context.Background()
	query:= "INSERT INTO comments(email,comment) VALUES(?,?)"
	statement,err:=db.PrepareContext(ctx,query)
	if err != nil{
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email:="sandy"+strconv.Itoa(i)+"@.gmail.com"	
		comment:="Komentar ke " + strconv.Itoa(i)

		result,err:= statement.ExecContext(ctx,email,comment)
		if err != nil {
			panic(err)
		}

		id,err:=result.LastInsertId()
		if err != nil{
			panic(err)
		}
		fmt.Println("comment id ",id)
	}


}

func TestTransaction(t *testing.T) {
	db:= GetConnection()
	defer db.Close()

	ctx:=context.Background()
	tx,err:=db.Begin()
	if err != nil{
		panic(err)
	}
	
	// do Transaction
	query:= "INSERT INTO comments(email,comment) VALUES(?,?)"
	for i := 0; i < 10; i++ {
		email:="sandy"+strconv.Itoa(i)+"@.gmail.com"	
		comment:="Komentar ke " + strconv.Itoa(i)
		result,err:= tx.ExecContext(ctx,query,email,comment)
		if err != nil {
			panic(err)
		}

		id,err:=result.LastInsertId()
		if err != nil{
			panic(err)
		}
		fmt.Println("comment id ",id)
	}

	err=tx.Commit() // commit data masuk ke db  rollback data gagal masuk ke db
	if err != nil{
		panic(err)
	}
}