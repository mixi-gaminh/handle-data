package main

import (
	"fmt"

	"github.com/mixi-gaminh/handle-data/jsonparse"
)

type Teststruct struct {
	Str  string `json:"string"`
	List []int  `json:"list"`
}

func main() {
	data := map[string]interface{}{
		"string": nil,
		"list":   []int{1234, 3124},
		"mapstring": map[string]interface{}{
			"test": "134",
		},
	}

	// bodyBytes := []byte(`{
	// 	"string": null,
	// 	"list":   [1234, 3124]
	// }`)

	// data2 := make(map[string]interface{})
	// var data3 interface{}
	// _ = json.Unmarshal(bodyBytes, &data2)
	// _ = json.Unmarshal(bodyBytes, &data3)

	// st := new(Teststruct)
	// err := convert.ToStruct(data2, st)
	// fmt.Println("debug", err)
	// if err == nil {
	// 	fmt.Println("struct:", st.List)
	// 	fmt.Println("struct:", st.Str)
	// }

	// var a interface{} = 1234
	// var b = "1234"
	// var c = "1234a"
	// var d interface{} = nil
	// var e = ""
	// var f = 10.11111111

	// fmt.Print("To Float: ")
	// fmt.Println(convert.ToFloat64(b))

	fmt.Print("Get string: ")
	fmt.Println(jsonparse.GetSlice(data, "list"))
	// fmt.Print("To slice map string: ")
	// fmt.Println(convert.ToMapString(data["mapstring"]))
	// fmt.Print("To slice map string: ")
	// fmt.Println(convert.ToMapString(data3))

	// fmt.Print("To slice string: ")
	// fmt.Println(convert.ToSliceString(data2["list"]))
	// fmt.Print("To slice map string: ")
	// fmt.Println(convert.ToSliceMapString(data2["list"]))

	// fmt.Print("To int a: ")
	// fmt.Println(convert.ToInt(a))
	// fmt.Print("To int b: ")
	// fmt.Println(convert.ToInt(b))
	// fmt.Print("To int c: ")
	// fmt.Println(convert.ToInt(c))
	// fmt.Print("To int d: ")
	// fmt.Println(convert.ToInt(d))
	// fmt.Print("To int e: ")
	// fmt.Println(convert.ToInt(e))
	// fmt.Print("To int f: ")
	// fmt.Println(convert.ToInt(f))
}
