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

	fmt.Println("Your PC stats")
	fmt.Println(info)

	all_games := "https://api.steampowered.com/ISteamApps/GetAppList/v2/"
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
		fmt.Printf("Data has been recived as a %T \n", string(j))
	}

	if err := json.Unmarshal(body, &makeMap); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	min_req := makeMap[str].Data.PcRequirements.Minimum
	output := PrettyPrint(min_req)
	GameDataAdd(output)
	fmt.Println("Would you like to see the game price and Metacritic score? \n ")
	fmt.Println("Answer with yes or no:")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "yes" {
			metacritic := makeMap[str].Data.Metacritic.Score
			getScore(metacritic)
			price := makeMap[str].Data.PriceOverview.Final
			getPrice(price)
			break
		} else {
			fmt.Println("Exiting...")
			break
		}
	}
	// getPriceScore(makeMap)

}
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func GameDataAdd(output string) {
	replacer := strings.NewReplacer(`"`, "", "br", "", "\\", "", "/", " ", "u003e", "", "u003c", "", "strong", "", "li", "", "ul", "")
	s := replacer.Replace((output))
	result := strings.Split(s, ":")
	fmt.Println("Ignore class=bb_")
	for i := range result {
		fmt.Println(result[i])
	}
}

func getScore(output int) {
	fmt.Println(output)
}

func getPrice(output int) {
	fmt.Println(output)
}
