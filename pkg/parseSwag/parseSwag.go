package parseSwag

import (
	"encoding/json"
	"fmt"
	_ "github.com/general252/goempty/docs"
	"github.com/swaggo/swag"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type StructSwagger struct {
	Swagger     string                                       `json:"swagger"`
	Info        interface{}                                  `json:"info"`
	BasePath    string                                       `json:"basePath"`
	Paths       map[string]map[string]map[string]interface{} `json:"paths"`
	Definitions interface{}                                  `json:"definitions"`
}

type StructSwagger2 struct {
	Swagger     string      `json:"swagger"`
	Info        interface{} `json:"info"`
	BasePath    string      `json:"basePath"`
	Paths       int         `json:"paths"`
	Definitions interface{} `json:"definitions"`
}

func ParseSwag() {
	doc, err := swag.ReadDoc()
	if err != nil {
		fmt.Println(err)
		return
	}

	var swagMap StructSwagger
	_ = json.Unmarshal([]byte(doc), &swagMap)

	type Object struct {
		Summary string
		Path    string
		PathObj interface{}
	}
	var objList []Object
	for path, pathObj := range swagMap.Paths {
		for _, obj := range pathObj {
			var s = obj["summary"].(string)

			objList = append(objList, Object{
				Summary: s,
				Path:    path,
				PathObj: pathObj,
			})
		}
	}

	for k := range swagMap.Paths {
		delete(swagMap.Paths, k)
	}

	var newSwagMap = &StructSwagger2{
		Swagger:     swagMap.Swagger,
		Info:        swagMap.Info,
		BasePath:    swagMap.BasePath,
		Paths:       969696969696,
		Definitions: swagMap.Definitions,
	}

	sort.Slice(objList, func(i, j int) bool {
		return objList[i].Summary < objList[j].Summary
	})

	var str = ""
	for i := 0; i < len(objList); i++ {
		obj := objList[i]
		data, _ := json.MarshalIndent(obj.PathObj, "", "    ")

		str += fmt.Sprintf(" \"%s\" : %s", obj.Path, string(data))
		if i != len(objList)-1 {
			str += ","
		}
	}

	data, _ := json.MarshalIndent(newSwagMap, "", "    ")
	strJson := string(data)
	strJson = strings.ReplaceAll(strJson, "969696969696", fmt.Sprintf("{ %s }", str))

	_ = ioutil.WriteFile("docs/swagger2.json", []byte(strJson), os.ModePerm)
	//fmt.Printf("%v", newSwagMap)
}
