# Tson
Tson used to merge JSON content from a template and an object.

#### Usage

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/alimy/tson"
)

type Data struct {
	Id   string
	Name string
	Age  uint8
}

func main() {
	json := `{
		"id": $.Id,
		"name": $.Name,
		"age": $.Age,
		"issue": {
			"version": "v0.1.0",
			"bio": "happy in codding",
		},
		"random": [1, 2, 3, 4, 5],
	}`

	if tmpl, err := tson.New("json").Parse(json); err == nil {
		buf := bytes.NewBuffer([1024]byte{}[:])
		tmpl.Execute(buf, &Data {
			Id:   "1006013",
			Name: "Michael Li",
			Age:  0,
		})
		fmt.Println(buf.String())
	}	
}
// Output:
// {
// 	"id": "1006013",
// 	"name": "Michael Li",
// 	"age": 0,
// 	"issue": {
// 	"version": "v0.1.0",
// 	"bio": "happy in codding",
// 	},
// 	"random": [1, 2, 3, 4, 5],
// }
//
```

#### Status
Project is just a prototype now will add logic implements later.
