package DatabaseFunctions

import (
	"context"
	"fmt"
	"github.com/RewanshChoudhary/DataReplicationSystem/config"
	"github.com/RewanshChoudhary/DataReplicationSystem/util"
	"github.com/jackc/pgx/v5"
	"os"
)

func Init_Db() {
	ctx := context.Background()

	configs, err := config.LoadFile()
	util.HandleError(err)

	job := configs.Jobs[0]
	srcDSN := config.ExpandDSN(job.Source.DSN)
	destDSN := config.ExpandDSN(job.Destination.DSN)

	// 🔹 Debug: print expanded DSNs
	fmt.Println("SRC DSN:", srcDSN)
	fmt.Println("DEST DSN:", destDSN)

	srcConn, err := pgx.Connect(ctx, srcDSN)
	util.HandleError(err)
	fmt.Println("✅ Source DB connected")
	defer srcConn.Close(ctx)

	destConn, err := pgx.Connect(ctx, destDSN)
	util.HandleError(err)
	fmt.Println("✅ Destination DB connected")
	defer destConn.Close(ctx)
	query := `
    SELECT column_name
    FROM information_schema.columns
    WHERE table_name = $1
    ORDER BY ordinal_position
`

	rows, err := srcConn.Query(ctx, query, os.Getenv("SRC_TABLE"))
	util.HandleError(err)
	defer rows.Close()

	fmt.Println("🔹 Columns:")
	for rows.Next() {
		var columnName string
		err = rows.Scan(&columnName)
		util.HandleError(err)
		fmt.Println(columnName)
	}

	util.HandleError(rows.Err())
}
