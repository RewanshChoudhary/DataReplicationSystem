package util

import "fmt"

func HandleError(err error, mess string) {
	if err != nil {
		fmt.Errorf("The error %w  occurred when %s", err, mess)
		panic(err)

	}
}
