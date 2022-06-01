package frequency

import (
	"encoding/json"
	"log"

	"github.com/registrobr/zxcvbn-go/data"
	"github.com/registrobr/zxcvbn-go/data/ptbr"
)

// List holds a frequency list
type List struct {
	Name string
	List []string
}

// Lists holds all the frequency list in a map
var Lists = make(map[string]List)

func init() {
	maleFilePath := getAsset("data/MaleNames.json")
	femaleFilePath := getAsset("data/FemaleNames.json")
	surnameFilePath := getAsset("data/Surnames.json")
	englishFilePath := getAsset("data/English.json")
	passwordsFilePath := getAsset("data/Passwords.json")

	Lists["MaleNames"] = getStringListFromAsset(maleFilePath, "MaleNames")
	Lists["FemaleNames"] = getStringListFromAsset(femaleFilePath, "FemaleNames")
	Lists["Surname"] = getStringListFromAsset(surnameFilePath, "Surname")
	Lists["English"] = getStringListFromAsset(englishFilePath, "English")
	Lists["Passwords"] = getStringListFromAsset(passwordsFilePath, "Passwords")

	//portuguese data
	ptBrCommonWords := getPtBrAsset("data/commonWords.json")
	ptBrFirstNames := getPtBrAsset("data/firstnames.json")
	ptBrLastNames := getPtBrAsset("data/lastnames.json")
	ptBrWikipedia := getPtBrAsset("data/wikipedia.json")

	Lists["CommonWords_ptbr"] = getStringListFromAsset(ptBrCommonWords, "CommonWords_ptbr")
	Lists["FirstNames_ptbr"] = getStringListFromAsset(ptBrFirstNames, "FirstNames_ptbr")
	Lists["LastNames_ptbr"] = getStringListFromAsset(ptBrLastNames, "LastNames_ptbr")
	Lists["Wikipedia_ptbr"] = getStringListFromAsset(ptBrWikipedia, "Wikipedia_ptbr")

}
func getAsset(name string) []byte {
	data, err := data.Asset(name)
	if err != nil {
		panic("Error getting asset " + name)
	}

	return data
}
func getStringListFromAsset(data []byte, name string) List {

	var tempList List
	err := json.Unmarshal(data, &tempList)
	if err != nil {
		log.Fatal(err)
	}
	tempList.Name = name
	return tempList
}

func getPtBrAsset(name string) []byte {
	data, err := ptbr.Asset(name)
	if err != nil {
		panic("Error getting asset " + name)
	}

	return data
}
