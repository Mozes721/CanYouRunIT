package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("\\")

	info := new(SysInfo)

	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024

	fmt.Printf("%+v\n", info)

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

	makeMap := make(map[string]GameNum)
	makeMap[str] = GameNum{}

	j, err := json.Marshal(makeMap)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(j))
	}

	if err := json.Unmarshal(body, &makeMap); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	min_req := makeMap[str].Data.PcRequirements.Minimum
	// max_req := makeMap[str].Data.PcRequirements.Recommended
	// getGameMinRequirements(makeMap)
	// fmt.Printf(PrettyPrint(min_req))
	fmt.Println(min_req)
	output := PrettyPrint(min_req)

	GameDataAdd(output)

	// s := strings.Split(output, "Processor:")[1]
	// val := s[1]
	// replaced := replacer.Replace((s))
	// fmt.Println(replaced)

}
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func getFieldInteger(e, field string) int {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

func GameDataAdd(output string) {
	// m := make(map[string]string)
	var GameData GameData
	fmt.Println("passed")
	replacer := strings.NewReplacer(`"`, "", "br", "", "\\", "", "/", " ", "u003e", "", "u003c", "", "strong", "", "li", "", "ul", "")
	fmt.Println(output)
	s := strings.Split(output, "Sound Card:")[1]
	sound := replacer.Replace((s))
	// m["Sound Card"] = sound

	st := strings.Split(output, "Storage:")[1]
	storage := replacer.Replace((st))
	// m["Storage"] = storage

	nt := strings.Split(output, "Network:")[1]
	network := replacer.Replace((nt))
	// m["Network"] = network

	gr := strings.Split(output, "Graphics:")[1]
	graphics := replacer.Replace((gr))
	// m["Graphics"] = graphics

	mem := strings.Split(output, "Memory:")[1]
	memory := replacer.Replace((mem))
	// m["Memory"] = memory

	pr := strings.Split(output, "Processor:")[1]
	processor := replacer.Replace(pr)
	// m["Processor"] = processor

	GameData.SoundCard = sound
	GameData.Storage = storage
	GameData.Network = network
	GameData.Graphics = graphics
	GameData.Memory = memory
	GameData.Processor = processor

	fmt.Println(getFieldString(GameData, GameData.Network))
	fmt.Println(GameData)
}
