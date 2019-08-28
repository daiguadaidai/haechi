package metas

import (
	"fmt"
	"testing"
)

func Test_RangePartitionItem(t *testing.T) {
	name1 := "p_01"
	values := []interface{}{0, '1', "a"}
	part1 := NewRangePartitionItem(name1, values, false)
	part1Str, err := part1.GetMetaStr()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(part1Str)

	nameMax := "p_max"
	partMax := NewRangePartitionItem(nameMax, values, true)
	partMaxStr, err := partMax.GetMetaStr()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(partMaxStr)
}

func Test_RangePartition(t *testing.T) {
	rangePartition := NewRangePartition()

	part1 := NewRangePartitionItem("p_01", []interface{}{1, 2, 3, 4}, false)
	part2 := NewRangePartitionItem("p_02", []interface{}{"1", "2", "3", "4"}, false)
	part3 := NewRangePartitionItem("p_03", []interface{}{5, 6, 7, 8}, false)
	part4 := NewRangePartitionItem("p_04", []interface{}{"a", "B", "c", "d"}, false)
	partMax := NewRangePartitionItem("p_max", []interface{}{1, 2, 3, 4}, true)

	rangePartition.AddPartitionItem(part1).AddPartitionItem(part2).AddPartitionItem(part3).AddPartitionItem(part4).AddPartitionItem(partMax)

	partStr, err := rangePartition.GetMetaStr()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(partStr)
}
