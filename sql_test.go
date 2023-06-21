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

	query:= "INSERT INTO customer(id,name) VALUES('sandy','SANDY')"
	_,err:= db.ExecContext(ctx,query)
	if err != nil{
		panic(err)
	}

	fmt.Println("Sukses insert a new customer")
}