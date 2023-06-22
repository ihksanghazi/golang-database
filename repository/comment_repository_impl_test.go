package repository

import (
	"context"
	"fmt"
	"golang_database"
	"golang_database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository:= NewCommentRepository(golang_database.GetConnection())
	
	ctx:= context.Background()
	comment:= entity.Comment{
		Id: 0,
		Email: "repositor@test.com",
		Comment: "Test Repository",
	}
	result,err:=commentRepository.Insert(ctx,comment)
	if err != nil{
		panic(err)
	}
	
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository:= NewCommentRepository(golang_database.GetConnection())
	ctx:= context.Background()
	comment,err:= commentRepository.FindById(ctx,22)
	if err != nil{
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository:= NewCommentRepository(golang_database.GetConnection())
	ctx:= context.Background()
	comments,err:= commentRepository.FindAll(ctx)
	if err != nil{
		panic(err)
	}
	for _, comment:= range comments{
		fmt.Println(comment)
	}

}