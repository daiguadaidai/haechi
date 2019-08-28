package metas

import (
	"fmt"
	"github.com/daiguadaidai/haechi/utils"
	"strings"
)

type Partition struct {
	Type           string          `json:"type" form:"type"`
	ColumnNames    []string        `json:"column_names" form:"column_names"`
	ListPartition  *ListPartition  `json:"list_partition" form:"list_partition"`
	RangePartition *RangePartition `json:"range_partition" form:"range_partition"`
}

type ListPartition struct {
	ListPartitionItems []*ListPartitionItem `json:"list_partition_items" form:"list_partition_items"`
}

type ListPartitionItem struct {
}

type RangePartition struct {
	RangePartitionItems []*RangePartitionItem `json:"range_partition_items" form:"range_partition_items"`
}

func NewRangePartition() *RangePartition {
	return &RangePartition{
		RangePartitionItems: make([]*RangePartitionItem, 0, 5),
	}
}

func (this *RangePartition) AddPartitionItem(partitionItem *RangePartitionItem) *RangePartition {
	this.RangePartitionItems = append(this.RangePartitionItems, partitionItem)
	return this
}

func (this *RangePartition) GetMetaStr() (string, error) {
	parts := make([]string, 0, 3)

	for _, item := range this.RangePartitionItems {
		partStr, err := item.GetMetaStr()
		if err != nil {
			return "", err
		}
		partStr = fmt.Sprintf("    %s", partStr)
		parts = append(parts, partStr)
	}

	return strings.Join(parts, ",\n"), nil
}

type RangePartitionItem struct {
	Name       string        `json:"name" form:"name"`
	Values     []interface{} `json:"values" form:"values"`
	IsMaxValue bool          `json:"is_max_value" json:"is_max_value"`
}

func NewRangePartitionItem(name string, values []interface{}, isMaxValue bool) *RangePartitionItem {
	return &RangePartitionItem{
		Name:       name,
		Values:     values,
		IsMaxValue: isMaxValue,
	}
}

func (this *RangePartitionItem) GetMetaStr() (string, error) {
	if strings.TrimSpace(this.Name) == "" {
		return "", fmt.Errorf("Range 分区必须要有分区名. %v", this.Values)
	}

	if len(this.Values) < 1 { // 没有值
		return "", fmt.Errorf("Range 分区必须要有指定范围值. %s", this.Name)
	}

	prePart := fmt.Sprintf("PARTITION %s VALUES LESS THAN", this.Name)

	// MAXVALUE 分区
	if this.IsMaxValue {
		values := make([]string, 0, len(this.Values))
		for _, _ = range this.Values {
			values = append(values, "MAXVALUE")
		}
		valuePart := utils.WarpStrs(values, "%s", ", ")
		finalStr := fmt.Sprintf("%s (%s)", prePart, valuePart)
		return finalStr, nil
	}

	// LIST 分区
	valuePart := utils.WarpInterfaceStrs(this.Values, "%#v", ", ")
	finalStr := fmt.Sprintf("%s (%s)", prePart, valuePart)

	return finalStr, nil
}
