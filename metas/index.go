package metas

import (
	"fmt"
	"github.com/daiguadaidai/haechi/utils"
	"strings"
)

type Index struct {
	Name        string   `json:"name" form:"name"`
	ColumnNames []string `json:"column_names" form:"column_names"`
}

func NewIndex(name string) *Index {
	return &Index{
		Name:        name,
		ColumnNames: make([]string, 0, 1),
	}
}

func (this *Index) AddColumnName(name string) {
	this.ColumnNames = append(this.ColumnNames, name)
}

func (this *Index) GetNamesStr() string {
	columnNamesStr := utils.WarpStrs(this.ColumnNames, "`%s`", ", ")

	return fmt.Sprintf("(%s)", columnNamesStr)
}

func (this *Index) GetMetaStr() (string, error) {
	items := make([]string, 0, 3)
	items = append(items, "INDEX")

	if strings.TrimSpace(this.Name) == "" {
		return "", fmt.Errorf("索引名称不能为空")
	}
	items = append(items, this.Name)

	items = append(items, this.GetNamesStr())

	return strings.Join(items, " "), nil
}

type PrimaryIndex struct {
	Index
}

func NewPrimaryIndex(name string) *PrimaryIndex {
	return &PrimaryIndex{
		Index: *NewIndex(name),
	}
}

func (this *PrimaryIndex) GetMetaStr() (string, error) {
	items := make([]string, 0, 3)
	items = append(items, "PRIMARY KEY")

	if strings.TrimSpace(this.Name) != "" {
		items = append(items, this.Name)
	}

	items = append(items, this.GetNamesStr())

	return strings.Join(items, " "), nil
}

type UniqueIndex struct {
	Index
}

func NewUniqueIndex(name string) *UniqueIndex {
	return &UniqueIndex{
		Index: *NewIndex(name),
	}
}

func (this *UniqueIndex) GetMetaStr() (string, error) {
	items := make([]string, 0, 3)
	items = append(items, "UNIQUE INDEX")

	if strings.TrimSpace(this.Name) == "" {
		return "", fmt.Errorf("唯一索引名称不能为空")
	}
	items = append(items, this.Name)

	items = append(items, this.GetNamesStr())

	return strings.Join(items, " "), nil
}
