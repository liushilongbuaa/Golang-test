package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"strings"
)

type vpc struct {
	Id          string `json:"id"`
	VpcId       string `json:"vpc_id"`
	Age         int    `json:"age"`
	SubnetCount int    `json:"subnet_count"`
}

func main() {
	vpc1 := &vpc{"i-abcd", "vpc-abcd", 9, 2}
	vpc2 := &vpc{"i-hahah", "vpc-hahah", 6, 2}
	a := []*vpc{vpc1, vpc2}

	err := FormatListOutput([]string{"VpcId", "Id", "SubnetCount"}, a)
	if err != nil {
		fmt.Println(err)
	}
}

func FormatListOutput(header []string, params interface{}) (err error) {
	header = listtocamal(header)
	table := table{}
	table.setHead(header)

	bt, err := json.Marshal(params)
	if err != nil {
		return err
	}
	mapSlice := []map[string]interface{}{}
	decoder := json.NewDecoder(bytes.NewBuffer(bt))
	decoder.UseNumber()
	err = decoder.Decode(&mapSlice)
	if err != nil {
		return err
	}
	fmt.Println(mapSlice)
	for _, rowMap := range mapSlice {
		row := []string{}
		for _, h := range header {
			if i, ok := rowMap[camaltounix(h)]; ok {
				if iS, ok := i.(string); ok {
					row = append(row, iS)
				} else if iI, ok := i.(json.Number); ok {
					row = append(row, iI.String())
				} else {
					panic("item not string or int")
				}
			} else {
				panic("header not found")
			}
		}
		table.input(row)
	}

	table.printTable()
	return nil
}
func camaltounix(i string) string {
	buff := bytes.Buffer{}
	for _, v := range i {
		if v >= 'A' && v <= 'Z' {
			buff.WriteRune('_')
			buff.WriteRune(v + 32)
		} else {
			buff.WriteRune(v)
		}
	}
	return buff.String()[1:]
}

func listtocamal(i []string) (ret []string) {
	tocamal := func(i string) (ret string) {
		var tem []byte
		up := true
		for _, v := range i {
			if v == '_' {
				up = true
			} else if up {
				if v >= 'a' && v <= 'z' {
					v -= 32
				}
				tem = append(tem, byte(v))
				up = false
			} else {
				tem = append(tem, byte(v))
			}
		}
		return string(tem)
	}
	for _, v := range i {
		ret = append(ret, tocamal(v))
	}
	return
}

type table struct {
	header []string
	width  []int
	raw    [][]string
}

func (t *table) setHead(h []string) {
	t.header = h
	t.width = make([]int, len(h))
	for i, v := range h {
		if t.width[i] < len(v)+2 {
			t.width[i] = len(v) + 2
		}
	}
}
func (t *table) input(r []string) {
	t.raw = append(t.raw, r)
	for i, v := range r {
		if t.width[i] < len(v)+2 {
			t.width[i] = len(v) + 2
		}
	}
}
func (t *table) printTable() {
	for _, v := range t.width {
		fmt.Print("+", strings.Repeat("-", v))
	}
	fmt.Println("+")
	for i, v := range t.header {
		fmt.Print("|", strings.Repeat(" ", (t.width[i]-len(v))/2),
			v,
			strings.Repeat(" ", t.width[i]-(t.width[i]-len(v))/2-len(v)))
	}
	fmt.Println("|")
	for _, v := range t.width {
		fmt.Print("+", strings.Repeat("-", v))
	}
	fmt.Println("+")

	for _, raw := range t.raw {
		for i, column := range raw {
			fmt.Print("|", strings.Repeat(" ", (t.width[i]-len(column))/2),
				column,
				strings.Repeat(" ", t.width[i]-(t.width[i]-len(column))/2-len(column)))
		}
		fmt.Println("|")
	}

	for _, v := range t.width {
		fmt.Print("+", strings.Repeat("-", v))
	}
	fmt.Println("+")
}
