package DatabaseFunctions

import (
	"context"
	"fmt"

	"github.com/RewanshChoudhary/DataReplicationSystem/config"
	"github.com/RewanshChoudhary/DataReplicationSystem/util"
	"github.com/jackc/pgx/v5"
)

func Init_Db() (*pgx.Conn, *pgx.Conn) {
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

	util.GetSchema(ctx, srcConn, "person_details")

	destConn, err := pgx.Connect(ctx, destDSN)
	util.HandleError(err)
	fmt.Println("✅ Destination DB connected")

	return srcConn, destConn

}
