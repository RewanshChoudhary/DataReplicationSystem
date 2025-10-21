package util

import (
	"context"
	"fmt"

	"github.com/RewanshChoudhary/DataReplicationSystem/config"

	"github.com/jackc/pgx/v5"
)

func Init_Db() (*pgx.Conn, *pgx.Conn, error) {
	ctx := context.Background()

	configs, err := config.LoadFile()
	if err != nil {
		return nil, nil, err
	}

	job := configs.Jobs[0]
	srcDSN := config.ExpandDSN(job.Source.DSN)
	destDSN := config.ExpandDSN(job.Destination.DSN)

	// 🔹 Debug: print expanded DSNs
	fmt.Println("SRC DSN:", srcDSN)
	fmt.Println("DEST DSN:", destDSN)

	srcConn, err := pgx.Connect(ctx, srcDSN)
	if err != nil {
		return nil, nil, err

	}
	fmt.Println(" Source DB connected")

	destConn, err := pgx.Connect(ctx, destDSN)
	HandleError(err)
	fmt.Println(" Destination DB connected")

	return srcConn, destConn, nil

}
