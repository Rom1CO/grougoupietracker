package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Players struct {
	Count        int `json:"count"`
	CountTotal   int `json:"count_total"`
	Page         int `json:"page"`
	PageTotal    int `json:"page_total"`
	ItemsPerPage int `json:"items_per_page"`
	Items        []struct {
		ID                    int           `json:"id"`
		ResourceID            interface{}   `json:"resource_id"`
		Name                  string        `json:"name"`
		Age                   int           `json:"age"`
		ResourceBaseID        interface{}   `json:"resource_base_id"`
		FutBinID              interface{}   `json:"fut_bin_id"`
		FutWizID              interface{}   `json:"fut_wiz_id"`
		FirstName             interface{}   `json:"first_name"`
		LastName              interface{}   `json:"last_name"`
		CommonName            string        `json:"common_name"`
		Height                int           `json:"height"`
		Weight                int           `json:"weight"`
		BirthDate             string        `json:"birth_date"`
		League                int           `json:"league"`
		Nation                int           `json:"nation"`
		Club                  int           `json:"club"`
		Rarity                int           `json:"rarity"`
		Traits                []interface{} `json:"traits"`
		Specialities          []interface{} `json:"specialities"`
		Position              string        `json:"position"`
		SkillMoves            int           `json:"skill_moves"`
		WeakFoot              int           `json:"weak_foot"`
		Foot                  string        `json:"foot"`
		AttackWorkRate        string        `json:"attack_work_rate"`
		DefenseWorkRate       string        `json:"defense_work_rate"`
		TotalStats            interface{}   `json:"total_stats"`
		TotalStatsInGame      interface{}   `json:"total_stats_in_game"`
		Rating                int           `json:"rating"`
		RatingAverage         interface{}   `json:"rating_average"`
		Pace                  int           `json:"pace"`
		Shooting              int           `json:"shooting"`
		Passing               int           `json:"passing"`
		Dribbling             int           `json:"dribbling"`
		Defending             int           `json:"defending"`
		Physicality           int           `json:"physicality"`
		PaceAttributes        interface{}   `json:"pace_attributes"`
		ShootingAttributes    interface{}   `json:"shooting_attributes"`
		PassingAttributes     interface{}   `json:"passing_attributes"`
		DribblingAttributes   interface{}   `json:"dribbling_attributes"`
		DefendingAttributes   interface{}   `json:"defending_attributes"`
		PhysicalityAttributes interface{}   `json:"physicality_attributes"`
		GoalkeeperAttributes  interface{}   `json:"goalkeeper_attributes"`
	} `json:"items"`
}

type Singleplayer struct {
	Item struct {
		ID                    int           `json:"id"`
		ResourceID            interface{}   `json:"resource_id"`
		Name                  string        `json:"name"`
		Age                   int           `json:"age"`
		ResourceBaseID        interface{}   `json:"resource_base_id"`
		FutBinID              interface{}   `json:"fut_bin_id"`
		FutWizID              interface{}   `json:"fut_wiz_id"`
		FirstName             interface{}   `json:"first_name"`
		LastName              interface{}   `json:"last_name"`
		CommonName            string        `json:"common_name"`
		Height                int           `json:"height"`
		Weight                int           `json:"weight"`
		BirthDate             string        `json:"birth_date"`
		League                int           `json:"league"`
		Nation                int           `json:"nation"`
		Club                  int           `json:"club"`
		Rarity                int           `json:"rarity"`
		Traits                []interface{} `json:"traits"`
		Specialities          []interface{} `json:"specialities"`
		Position              string        `json:"position"`
		SkillMoves            int           `json:"skill_moves"`
		WeakFoot              int           `json:"weak_foot"`
		Foot                  string        `json:"foot"`
		AttackWorkRate        string        `json:"attack_work_rate"`
		DefenseWorkRate       string        `json:"defense_work_rate"`
		TotalStats            interface{}   `json:"total_stats"`
		TotalStatsInGame      interface{}   `json:"total_stats_in_game"`
		Rating                int           `json:"rating"`
		RatingAverage         interface{}   `json:"rating_average"`
		Pace                  int           `json:"pace"`
		Shooting              int           `json:"shooting"`
		Passing               int           `json:"passing"`
		Dribbling             int           `json:"dribbling"`
		Defending             int           `json:"defending"`
		Physicality           int           `json:"physicality"`
		PaceAttributes        interface{}   `json:"pace_attributes"`
		ShootingAttributes    interface{}   `json:"shooting_attributes"`
		PassingAttributes     interface{}   `json:"passing_attributes"`
		DribblingAttributes   interface{}   `json:"dribbling_attributes"`
		DefendingAttributes   interface{}   `json:"defending_attributes"`
		PhysicalityAttributes interface{}   `json:"physicality_attributes"`
		GoalkeeperAttributes  interface{}   `json:"goalkeeper_attributes"`
	} `json:"item"`
}

func add(a int) int { return a + 1 }

func main() {
	directory := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", directory))
	http.HandleFunc("/home/", homeHandler)
	http.HandleFunc("/player/", playerHandler)
	http.ListenAndServe(":9999", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	httpClient := http.Client{
		Timeout: time.Second * 8,
	}
	page := strings.ReplaceAll(r.URL.Path, "/home/", "")
	lien := "https://futdb.app/api/players?page=" + page
	req, err := http.NewRequest(http.MethodGet, lien, nil)
	req.Header.Set("X-AUTH-TOKEN", "75becef4-5d0b-4dc3-b599-1ba5cc611cf2")

	if err != nil {
		log.Fatalln(err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	response := Players{}

	jsonErr := json.Unmarshal(body, &response)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	response.Page = add(response.Page)
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, response)

}

func playerHandler(w http.ResponseWriter, r *http.Request) {
	httpClient := http.Client{
		Timeout: time.Second * 8,
	}
	id := strings.ReplaceAll(r.URL.Path, "/player/", "")
	url := "https://futdb.app/api/players/" + id
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("X-AUTH-TOKEN", "75becef4-5d0b-4dc3-b599-1ba5cc611cf2")
	if err != nil {
		log.Fatalln(err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	response := Singleplayer{}

	jsonErr := json.Unmarshal(body, &response)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	tmpl := template.Must(template.ParseFiles("player.html"))
	tmpl.Execute(w, response)

}
