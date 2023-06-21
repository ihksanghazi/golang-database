package golang_database

import (
	"context"
	"fmt"
	"testing"
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