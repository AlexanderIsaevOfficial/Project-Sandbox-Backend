package utils

//
import (
	"encoding/json"
	types "gameback_v1/types"
	"log"
	"os"
)

func ReadConfig(path string) (cnf types.Config) {

	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("ERROR - READ CONFIG", err)
	}

	err = json.Unmarshal(b, &cnf)
	if err != nil {
		log.Fatalln("ERROR - PARSE CONFIG", err)
	}
	return
}

func ConvertCategoryDBRequestToMainResponse(in []types.CategoryDBRequest) (out types.MainResponse, err error) {

	var items []types.Item
	var CR types.CategoryResponse
	var idPrew int

	for i, v := range in {
		//fmt.Println(items)
		if i > 1 {
			if idPrew == v.Id {
				//fmt.Println(types.Item{Id: v.I.Id, Title: v.I.Title, Logo: v.I.Logo})
				items = append(items, types.Item{Id: v.I.Id, Title: v.I.Title, Logo: v.I.Logo})
			} else {
				CR.Items = append(CR.Items, items...)
				out.Cat = append(out.Cat, CR)
				items = nil
			}
		} else {
			//fmt.Println(types.Item{Id: v.I.Id, Title: v.I.Title, Logo: v.I.Logo})
			items = append(items, types.Item{Id: v.I.Id, Title: v.I.Title, Logo: v.I.Logo})
		}
		CR.Id = v.Id
		CR.Title = v.Title

		idPrew = v.Id

	}

	CR.Items = append(CR.Items, items...)
	out.Cat = append(out.Cat, CR)
	clear(items)

	return out, nil
}
