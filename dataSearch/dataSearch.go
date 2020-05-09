package dataSearch

import (
	"encoding/json"
	uuid "github.com/google/uuid"
	qrcode "github.com/skip2/go-qrcode"
	"io/ioutil"
	"os"
)


type JsonObject struct {
	Path   string `json:path`
	ImgUrl string `json:imgUrl`
}

func SearchItem(url string) *JsonObject {
	items := readJsonFile("data.json")
	for _, obj := range *items {
		if obj.Path == url {
			return &JsonObject{Path: url, ImgUrl: obj.ImgUrl}
		}
	}
	return createNewItem(url)
}
func createNewItem(url string) *JsonObject {
	//create uuid and image path
	id, _ := uuid.NewRandom()
	imgStr := id.String() + ".png"

	//create QRcode
	err := createQRcode(url, "./qrimg/"+imgStr)
	if err != nil {
		panic(err)
	}

	//create new Item
	newItem := JsonObject{
		Path:   url,
		ImgUrl: id.String(),
	}
	writeJsonFile(newItem)

	return &newItem
}
func createQRcode(url string, writePath string) error {

	err := qrcode.WriteFile(url, qrcode.Medium, 256, writePath)
	return err
}

func writeJsonFile(item JsonObject) {
	items := readJsonFile("data.json")
	*items = append(*items, item) //add new item

	//items to json byte
	jsonTypeStr, _ := json.Marshal(*items)
	ioutil.WriteFile("data.json", jsonTypeStr, os.ModePerm)
}

func readJsonFile(filePath string) *[]JsonObject {
	var items []JsonObject
	jsonFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &items)
	return &items
}
