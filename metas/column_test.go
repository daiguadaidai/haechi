package metas

import (
	"fmt"
	"testing"
)

func TestColumn_GetMetaStr(t *testing.T) {
	column_01 := &Column{
		Name:             "col_01",
		Type:             "text",
		TypeLen:          10,
		TypeDecimal:      10,
		TypeValues:       []string{"1", "2", "3", "4", "5"},
		HaveDefaultValue: true,
		DefaultNull:      true,
		DefaultValue:     "aaa",
		Charset:          "utf8mb4",
		Collate:          "utf8mb4_unicode_ci",
	}

	if field, err := column_01.GetMetaStr(); err != nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(field)
	}

	column_02 := &Column{
		Name:             "col_02",
		Type:             "int",
		TypeLen:          10,
		TypeDecimal:      10,
		TypeValues:       []string{"1", "2", "3", "4", "5"},
		HaveDefaultValue: true,
		DefaultNull:      true,
		DefaultValue:     "aaa",
		Charset:          "utf8mb4",
		Collate:          "utf8mb4_unicode_ci",
	}
	if field, err := column_02.GetMetaStr(); err != nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(field)
	}

	column_03 := &Column{
		Name:             "col_03",
		Type:             "set",
		TypeLen:          10,
		TypeDecimal:      10,
		TypeValues:       []string{"1", "2", "3", "4", "5"},
		HaveDefaultValue: true,
		DefaultNull:      true,
		DefaultValue:     "aaa",
		Charset:          "utf8mb4",
		Collate:          "utf8mb4_unicode_ci",
	}
	if field, err := column_03.GetMetaStr(); err != nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(field)
	}

	column_04 := &Column{
		Name:             "col_04",
		Type:             "DECIMAL",
		TypeLen:          10,
		TypeDecimal:      10,
		TypeValues:       []string{"1", "2", "3", "4", "5"},
		HaveDefaultValue: true,
		DefaultNull:      true,
		DefaultValue:     "aaa",
		Charset:          "utf8mb4",
		Collate:          "utf8mb4_unicode_ci",
	}
	if field, err := column_04.GetMetaStr(); err != nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(field)
	}

	column_05 := &Column{
		Name:               "col_05",
		Type:               "DATETIME",
		TypeLen:            6,
		TypeDecimal:        10,
		TypeValues:         []string{"1", "2", "3", "4", "5"},
		HaveDefaultValue:   true,
		DefaultNull:        false,
		DefaultValue:       "CURRENT_TIMESTAMP(6)",
		DefaultValueIsFunc: true,
		OnUpdateValue:      "CURRENT_TIMESTAMP(6)",
		Charset:            "utf8mb4",
		Collate:            "utf8mb4_unicode_ci",
	}
	if field, err := column_05.GetMetaStr(); err != nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(field)
	}

	column_06 := &Column{
		Name:             "col_06",
		Type:             "DECIMAL",
		TypeLen:          10,
		TypeDecimal:      10,
		TypeValues:       []string{"1", "2", "3", "4", "5"},
		HaveDefaultValue: true,
		DefaultNull:      false,
		DefaultValue:     "aaa",
		Charset:          "utf8mb4",
		Collate:          "utf8mb4_unicode_ci",
	}
	if field, err := column_06.GetMetaStr(); err != nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(field)
	}
}
