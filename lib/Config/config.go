package Config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type (
	OptionStruct struct {
		Token      string `json:"token"`
		Prefix     string `json:"prefix"`
		IP         string `json:"ip"`
		Status     string `json:"status"`
		StatusCh   bool   `json:"statuschannel"`
		StatusChID string `json:"statuschannelid"`
	}
)

var (
	Option_Example = OptionStruct{
		Token:      "null",
		Prefix:     ".",
		IP:         "192.168.0.1:27015",
		Status:     "null",
		StatusCh:   true,
		StatusChID: "null",
	}

	Option *OptionStruct
)

func Init() string {
	a, err := ioutil.ReadFile("config/bot.json")

	if err != nil {
		CreateExample()
		return "Change config/bot.json"
	}

	err = json.Unmarshal(a, &Option)

	if err != nil {
		log.Panic(err)
	}

	return "Succes!"
}

func CreateExample() {
	err := os.Mkdir("config", 0755)
	if err != nil {
		log.Fatal(err)
	}

	a, err := json.MarshalIndent(Option_Example, " ", " ")
	if err != nil {
		log.Panic(err)
	}

	ioutil.WriteFile("config/bot.json", a, 0644)
	ioutil.WriteFile("config/bot_example.json", a, 0644)
}
