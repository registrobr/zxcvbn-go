package frequency

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/registrobr/zxcvbn-go/data"
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
	ptBrCommonWords := getDataFromFile("data/pt-br/commonWords.json")
	ptBrFirstNames := getDataFromFile("data/pt-br/firstnames.json")
	ptBrLastNames := getDataFromFile("data/pt-br/lastnames.json")
	ptBrWikipedia := getDataFromFile("data/pt-br/wikipedia.json")

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

func getDataFromFile(name string) []byte {
	// Open our jsonFile
	jsonFile, err := os.Open("./data/" + name)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatalf("error: %s\n", err)
		return nil
	}
	log.Printf("Successfully Opened %s\n", name)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue

}
