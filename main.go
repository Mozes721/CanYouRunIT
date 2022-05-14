package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/v3/mem"
)

var game_arg string

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
	// fmt.Println(os.Args[1:])

	all_games := "https://api.steampowered.com/ISteamApps/GetAppList/v2/"

	// game_search :="https://store.steampowered.com/api/appdetails?appids="

	api_games(all_games)
}

func api_games(url string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var allGames allGames

	if err := json.Unmarshal(body, &allGames); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(PrettyPrint(allGames))
}

func PrettyPrint(i interface{}) string {
	// game_arg = os.Args[1:]
	// s, _ := json.MarshalIndent(i, "", "\t")
	// println(game_arg)
	return string(game_arg)
}
