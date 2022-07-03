package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	// hostStat, _ := host.Info()
	// cpuStat, _ := cpu.Info()
	// vmStat, _ := mem.VirtualMemory()
	// diskStat, _ := disk.Usage("\\")

	// info := new(SysInfo)

	// info.Hostname = hostStat.Hostname
	// info.Platform = hostStat.Platform
	// info.CPU = cpuStat[0].ModelName
	// info.RAM = vmStat.Total / 1024 / 1024
	// info.Disk = diskStat.Total / 1024 / 1024

	// fmt.Printf("%+v\n", info)
	// fmt.Println(os.Args[1:])

	all_games := "https://api.steampowered.com/ISteamApps/GetAppList/v2/"

	// game_search :="https://store.steampowered.com/api/appdetails?appids="
	gameid := api_games(all_games)

	if gameid == 0 {
		fmt.Println("Game not Found")
	} else {
		checkGame(gameid)
	}
}
func api_games(url string) int {
	// get all games url
	resp, err := http.Get(url)
	var gameID int
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var allGames allGames
	// unmarshall struct to allGames
	if err := json.Unmarshal(body, &allGames); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	// user input
	fmt.Println("Enter your game of choice:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	read_line := strings.TrimRight(input, "\r\n")
	// check for game
	for _, v := range allGames.Applist.Apps {
		if v.Name == read_line {
			gameID = v.Appid
			fmt.Println("Game ID found: \n ")
			fmt.Println(gameID)
		}
	}
	return gameID

}

func checkGame(gameID int) {
	// get game details
	str := strconv.Itoa(gameID)
	game_search := "https://store.steampowered.com/api/appdetails?appids=" + str

	resp, err := http.Get(game_search)
	// fmt.Println(PrettyPrint(resp))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	jsonData := []byte(`
			"ref": *
	`)

	type gameData struct {
		GameNum `json:"appid"`
	}

	appJson := `{"GameNum":%v}`, strconv
	var gameInfo gameData

	// _ = json.Unmarshal([]byte(str), &game)

	// fmt.Println(v)
	// if err != nil {
	// 	fmt.Println("Unable to convert the JSON string to a struct")
	// } else {
	// 	// print the post
	// 	fmt.Println(PrettyPrint(game))
	// }

	if err := json.Unmarshal(body, &gameInfo); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(PrettyPrint(gameInfo))
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
