# Tson
Tson used to merge JSON content from a json template and a struct type object instance.

#### Prototype
1. Use `tson` that defined in struct's tag string to tag json template info help tson to merge content from struct type object 
2. Or use `$.filedNanme` that defined in json template string help tson to merge content from struct type object
3. tson `Parse(...)/ParseFrom(...)` will build a AST tree that contain json segment content
4. tson use parsed `AST` tree merge give struct type object instance content then write result to target io.Writer
5. json template can cached in memory or key-value storage such as Redis

#### Usage

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/alimy/tson"
)

type Issue struct {
	Version string `json:"version"`
	Bio     string `json:"bio"`
}

type Data struct {
	Id     string `json:"id" tson:"id"`
	Name   string `json:"name" tson:"name"`
	Age    uint8  `json:"age" tson:"age"`
	Issue  *Issue `json:"issue"`
	Random []int  `json:"random"`
}

type Info struct {
	Name  string
	Alias string
}

func main() {
	// merge json content from struct type template
	srcData := &Data{
		Age: 0,
		Issue: &Issue{
			Version: "v0.1.0",
			Bio:     "happy in codding",
		},
		Random: []int{1, 2, 3, 4, 5},
	}
	if jsonTmpl, err := tson.New("json").Parse(srcData); err == nil {
		buf := bytes.NewBuffer([1024]byte{}[:])
		jsonTmpl.Execute(buf, &Data{
			Id:   "1006013",
			Name: "Michael Li",
		})
		fmt.Println(buf.String())
	}

	// merge json content from text string template
	jsonTmplStr := `{
            "name": $.Name,
            "alias": $.Alias,
            "other": {
                "ping": "pong",
       	    },
    	}`
	if jsonTmpl, err := tson.New("tson").ParseFrom(jsonTmplStr); err == nil {
		buf := bytes.NewBuffer([1024]byte{}[:])
		jsonTmpl.Execute(buf, &Info{
			Name:  "Micheal Li",
			Alias: "Alimy",
		})
		fmt.Println(buf.String())
	}
}
// Output:
//  {
//      "id": "1006013",
//      "name": "Michael Li",
//      "age": 0,
//      "issue": {
//          "version": "v0.1.0",
//          "bio": "happy in codding",
//      },
//      "random": [1, 2, 3, 4, 5],
//  }
//  {
//      "name": "Micheal Li",
//      "alias": "Alimy",
//      "other": {
//          "ping": "pong",
//      },
//  }

```

#### Status
Project is just a prototype now will add logic implements later.
