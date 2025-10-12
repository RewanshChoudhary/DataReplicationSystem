package DatabaseFunctions

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Db_init() (*pgx.Conn,error){
	conn,err:=pgx.Connect(context.Background(),"postgres://postgres:rootrewansh@localhost:5432/postgres")

	if (err!=nil){
		fmt.Println(err)

	}

	return conn, err
	


}

