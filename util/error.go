package util

import "fmt"

func HandleError(err error) {
	if err != nil {
		fmt.Errorf("The error %s  occurred when ", err)
		panic(err)

	}

}
func FormatError(err error) {
	if err != nil {
		fmt.Errorf("The err %s occurred when  ", err)
	}
}
