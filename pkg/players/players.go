package players

import (
	"encoding/json"
	mcNet "github.com/Tnze/go-mc/net"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type playerListResponse struct {
	PlayerList []string `json:"player_list"`
}

//There are 3 of a max 20 players online: bar, foo, baz
var (
	rconClient       mcNet.RCONClientConn
	listPlayersRegex = regexp.MustCompile(`There are \d+ of a max \d+ players online:\s(?P<players>.+)`)
)

func init() {
	address := ""
	password := ""
	rcon, err := mcNet.DialRCON(address, password)
	if err != nil {
		panic(err)
	}
	rconClient = rcon
}

func Players(w http.ResponseWriter, r *http.Request) {
	log.Println("request received")
	switch r.Method {
	case http.MethodGet:
		response, err := getHandler()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("could not parse player list"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(response)
	case http.MethodPost:
	default:

	}
}

func getHandler() ([]byte, error) {
	err := rconClient.Cmd("list")
	if err != nil {
		log.Println("error with listing players")
	}
	response, err := rconClient.Resp()
	if err != nil {
		log.Println("error with getting response")
	}
	parsedResponse := parseListResponse(response)
	responseBytes, err := json.Marshal(parsedResponse)
	if err != nil {
		log.Println("error marshalling the json")
		return nil, err
	}
	log.Println(string(responseBytes))
	return responseBytes, nil
}

func parseListResponse(response string) playerListResponse {
	match := listPlayersRegex.FindStringSubmatch(response)
	if len(match) == 0 {
		return playerListResponse{PlayerList: []string{}}
	}
	result := match[1]
	playerList := strings.Split(result, ", ")
	return playerListResponse{
		PlayerList: playerList,
	}
}
