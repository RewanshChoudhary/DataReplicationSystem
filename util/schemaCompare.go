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

func GetSchema() ([]ColumnInfo,[]ColumnInfo) {
	query := `
    SELECT column_name, data_type, is_nullable, column_default 
FROM information_schema.columns
		WHERE table_name = $1
    ORDER BY ordinal_position
`
	fmt.Println("no P")
	srcConn,destConn, err := Init_Db()
	ctx := context.Background()

	srcCols, err := srcConn.Query(ctx, query, os.Getenv("SRC_TABLE"))
	HandleError(err)
	defer srcCols.Close()
	var src,dest []ColumnInfo


	
	for srcCols.Next(){
		var columninfo ColumnInfo
	
		err = srcCols.Scan(&columninfo.ColumnName, &columninfo.Datatype, &columninfo.Is_nullable, &columninfo.Column_default)
    HandleError(err)
		
		fmt.Println(columninfo)
		src = append(src, columninfo)

	}

  
destCols,err:=destConn.Query(ctx,query,os.Getenv("DEST_TABLE"))
  defer desCols.Close()

HandleError(err)

for destCols.Next(){
	var columninfo ColumnInfo
	err=destCols.Scan(&columninfo.ColumnName, &columninfo.Datatype, &columninfo.Is_nullable, &columninfo.Column_default)
    HandleError(err)
		dest =append(dest,columninfo)



 }

 for i :=range src {
	 sourceCols:=src[i]
	 destiantionCols:=dest[i]

	 if(compareSchema(sourceCols,destiantionCols)){
		  fmt.Println("The schemas are equal with same sequence")

		 
	 }else {
		 fmt.Println("The schemas are not equal but can contain different sequences of same columns ")

	 }




 }
  


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
	destCols := destConn.Query(ctx, query, os.Getenv("DEST_TABLE"))

	for 


}

func CompareColumns(a, b ColumnInfo) bool {
	return strings.EqualFold(a.ColumnName, b.ColumnName) &&
		strings.EqualFold(a.Datatype, b.Datatype) &&
		a.IsNullable == b.IsNullable &&
		((a.ColumnDefault == nil && b.ColumnDefault == nil) ||
			(a.ColumnDefault != nil && b.ColumnDefault != nil && *a.ColumnDefault == *b.ColumnDefault))
}
