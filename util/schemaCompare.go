package util

import (
	"context"
	"fmt"
	"os"
)

type ColumnInfo struct {
	ColumnName     string
	Datatype       string
	Is_nullable    string
	Column_default string
}

func GetSchema() []ColumnInfo {
	query := `
    SELECT column_name, data_type, is_nullable, column_default 
    FROM information_schema.columns
		WHERE table_name = $1
    ORDER BY ordinal_position
`
	fmt.Println("no P")
	srcConn, _, err := Init_Db()
	ctx := context.Background()

	rows, err := srcConn.Query(ctx, query, os.Getenv("SRC_TABLE"))
	HandleError(err)
	defer rows.Close()
	var r []ColumnInfo

	fmt.Println(" Columns:")
	for rows.Next() {
		var columninfo ColumnInfo

		err = rows.Scan(&columninfo.ColumnName, &columninfo.Datatype, &columninfo.Is_nullable, &columninfo.Column_default)

		fmt.Println(columninfo)
		r = append(r, columninfo)

	}

	HandleError(rows.Err())
	return r
}

func compareSchema() {
	query := `
    SELECT column_name, data_type, is_nullable, column_default 
    FROM information_schema.columns
		WHERE table_name = $1
    ORDER BY ordinal_position
`

	srcConn, destConn, err := Init_Db()
	if err != nil {
		HandleError(err)
	}
	srcCols := srcConn.Query(ctx, query, os.Getenv("SRC_TABLE"))
	destCols := destConn.Query(ctx, query, os.GetEnv("DEST_TABLE"))

}
