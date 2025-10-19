package util

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn/ctxwatch"
)

type ColumnInfo struct {
	ColumnName     string
	Datatype       string
	Is_nullable    string
	Column_default string
}

func GetSchema(ctx context.Context, srcConn *pgx.Conn, tableName string) []ColumnInfo {
	query := `
    SELECT column_name, data_type, is_nullable, column_default 
    FROM information_schema.columns
    WHERE table_name = $1
    ORDER BY ordinal_position
`
	fmt.Println("no P")

	rows, err := srcConn.Query(ctx, query, os.Getenv("SRC_TABLE"))
	HandleError(err)
	defer rows.Close()
	var r []ColumnInfo
	c := 0
	fmt.Println(" Columns:")
	for rows.Next() {
		var columninfo ColumnInfo
		c++
		err = rows.Scan(&columninfo.ColumnName, &columninfo.Datatype, &columninfo.Is_nullable, &columninfo.Column_default)

		fmt.Println(columninfo)
		r = append(r, columninfo)

	}

	HandleError(rows.Err())
	return r


func CompareSchema(ctx context.Context,conn1 pgx.Conn,conn2 pgx.Conn,tablename string)string){
	

	
}
