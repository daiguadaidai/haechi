package metas

import (
	"fmt"
	"testing"
)

func Test_Index(t *testing.T) {
	index := NewIndex("idx_aaa")
	index.AddColumnName("a1")
	index.AddColumnName("a2")
	index.AddColumnName("a3")

	indexStr, err := index.GetMetaStr()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(indexStr)
}

func Test_PrimaryIndex(t *testing.T) {
	index := NewPrimaryIndex("")
	index.AddColumnName("a1")
	index.AddColumnName("a2")
	index.AddColumnName("a3")

	indexStr, err := index.GetMetaStr()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(indexStr)
}

func Test_UniqueIndex(t *testing.T) {
	index := NewUniqueIndex("udx_aaa")
	index.AddColumnName("a1")
	index.AddColumnName("a2")
	index.AddColumnName("a3")

	indexStr, err := index.GetMetaStr()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(indexStr)
}
