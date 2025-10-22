package util

import (
	"context"
	"fmt"
)

func ListOfTables() {
	ctx := context.Background()
	query := `SELECT table_name FROM information_schema.tables WHERE table_schema='public'`

	srcConn, destConn, err := Init_Db()
	HandleError(err, "WHile accessing DB list of tables")
	if srcConn != nil {
		rows, err := srcConn.Query(ctx, query)

		HandleError(err, "While accessing rows of the source table ")
		defer rows.Close()

		for rows.Next() {
			var tableName string
			err := rows.Scan(&tableName)

			HandleError(err, "WHile copying table name to the variable of the source table")
			fmt.Println(tableName)

		}

	}
	if destConn != nil {
		rows, err := destConn.Query(ctx, query)

		HandleError(err, "While accessing rows of the dest table ")
		defer rows.Close()

		for rows.Next() {
			var tableName string
			err := rows.Scan(&tableName)

			HandleError(err, "WHile copying table name to the variable of the dest table")

			fmt.Println(tableName)

		}

	}

}
