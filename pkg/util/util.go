package util

import "fmt"

type Utils string

var (
	UtilsInstance Utils = "utils"
)

func GetUtils() Utils {
	return UtilsInstance
}
func (u *Utils) Name() {
	fmt.Print(u)
}
