package copier

import "github.com/jinzhu/copier"

func MustCopy(toValue, fromValue interface{}) {
	err := copier.Copy(toValue, fromValue)
	if err != nil {
		panic(err)
	}
}
