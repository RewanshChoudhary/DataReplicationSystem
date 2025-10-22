package util

import (
	"context"
	"fmt"
	"os"
	"strings"
)

type ColumnInfo struct {
	ColumnName     string
	Datatype       string
	Is_nullable    string
	Column_default string
}

func GetSchema() ([]ColumnInfo, []ColumnInfo) {
	query := `
    SELECT column_name, data_type, is_nullable, column_default 
FROM information_schema.columns
		WHERE table_name = $1
    ORDER BY ordinal_position
`
	fmt.Println("no P")
	srcConn, destConn, err := Init_Db()
	ctx := context.Background()

	srcCols, err := srcConn.Query(ctx, query, os.Getenv("SRC_TABLE"))
	HandleError(err, "While querying src Db")
	defer srcCols.Close()
	var src, dest []ColumnInfo

	for srcCols.Next() {
		var columninfo ColumnInfo

		err = srcCols.Scan(&columninfo.ColumnName, &columninfo.Datatype, &columninfo.Is_nullable, &columninfo.Column_default)
		HandleError(err, "while storing src col details")

		fmt.Println(columninfo)
		src = append(src, columninfo)

	}

	destCols, err := destConn.Query(ctx, query, os.Getenv("DEST_TABLE"))
	defer destCols.Close()

	HandleError(err, "While destiantion querying ")

	for destCols.Next() {
		var columninfo ColumnInfo
		err = destCols.Scan(&columninfo.ColumnName, &columninfo.Datatype, &columninfo.Is_nullable, &columninfo.Column_default)
		HandleError(err, "While destiantion column storing")
		dest = append(dest, columninfo)

	}
	return src, dest
}

func compareSchema() {
	src, dest := GetSchema()
	equal := true

	for i := range src {
		sourceCols := src[i]
		destiantionCols := dest[i]
		if !CompareColumns(sourceCols, destiantionCols) {
			equal = false

		}
	}
	if !equal {
		fmt.Println("The tables are not equal")
		fmt.Println(src)

		fmt.Println(dest)

	} else {
		fmt.Println("The table  schema are equal")

	}

}
func CompareColumns(a, b ColumnInfo) bool {
	return strings.EqualFold(a.ColumnName, b.ColumnName) &&
		strings.EqualFold(a.Datatype, b.Datatype) &&
		a.Is_nullable == b.Is_nullable &&
		a.Column_default == b.Column_default
}
