package utils

import (
	"fmt"
	"strings"
	"testing"
)

func Test_OriginStrs(t *testing.T) {
	datas := []interface{}{1, 2, 3, 4}
	strs := OriginStrs(datas)
	fmt.Println(strings.Join(strs, ", "))

	datas = []interface{}{"1", "a", "2", "b"}
	strs = OriginStrs(datas)
	fmt.Println(strings.Join(strs, ", "))
}
