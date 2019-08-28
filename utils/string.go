package utils

import (
	"fmt"
	"strings"
)

func EnumStrs(datas []string) []string {
	strs := make([]string, len(datas))
	for i, data := range datas {
		strs[i] = fmt.Sprintf("%#v", data)
	}
	return strs
}

/* 将指定的items中的每个item按指定的 warp 包裹起来, 包裹后的值使用 sep分割
Param:
    items: 需要包裹的值
    warp: 包裹的值是什么. 传参如:(%s), 就会将每个item使用()包裹起来
    sep: 包裹后以什么分割
Return:
    string
*/
func WarpStrs(items []string, warp string, sep string) string {
	if len(items) == 0 {
		return ""
	}

	warpItems := make([]string, len(items))
	for i, item := range items {
		warpItems[i] = WarpStr(item, warp)
	}

	return strings.Join(warpItems, sep)
}

func WarpStr(item string, warp string) string {
	return fmt.Sprintf(warp, item)
}

func WarpInterfaceStrs(items []interface{}, warp string, sep string) string {
	if len(items) == 0 {
		return ""
	}

	warpItems := make([]string, len(items))
	for i, item := range items {
		warpItems[i] = WarpInterfaceStr(item, warp)
	}

	return strings.Join(warpItems, sep)
}

func WarpInterfaceStr(item interface{}, warp string) string {
	return fmt.Sprintf(warp, item)
}
