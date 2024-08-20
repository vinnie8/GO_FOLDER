package logic

import(
	"encoding/json"
	"strings"
)

func ToursFromJson(content string) []Tour{
	
	tours:=make([]Tour,0,20)
	
	decoder:= json.NewDecoder(strings.NewReader(content))
	_, err:=decoder.Token()
	
	if err!=nil{
		panic(err)
	}

	var tour Tour
	
	for decoder.More(){
		err:=decoder.Decode(&tour)
		
		if err!=nil{
			panic(err)
		}
		tours = append(tours, tour)
	}

	return tours
}

type Tour struct{
	Name, Price string
}