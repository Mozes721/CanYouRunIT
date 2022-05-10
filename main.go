package main

import (
	"fmt"
<<<<<<< HEAD
	"io/ioutil"
	"log"
	"net/http"

=======
>>>>>>> e60f04387b646b56e4a25c134680274fa8f18b36
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/v3/mem"
)

// SysInfo saves the basic system information
type SysInfo struct {
	Hostname string `bson:hostname`
	Platform string `bson:platform`
	CPU      string `bson:cpu`
	RAM      uint64 `bson:ram`
	Disk     uint64 `bson:disk`
}

// func system() {
// 	hostStat, _ := host.Info()
// }

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

	api_games(all_games)
}

func api_games(url string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
