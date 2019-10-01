package suite

import (
	"encoding/json"
	"io/ioutil"
)

var JSON_FILE_SMALL = "../data/small.json"
var JSON_FILE_LARGE = "../data/large.json"
var JSON_FILE_GEODATA = "../data/canada.json"

type SmallPayload struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Public   string `json:"public"`
	Paginate struct {
		Default int `json:"default"`
		Max     int `json:"max"`
	} `json:"paginate"`
	Mongodb string `json:"mongodb"`
}

type LargePayload struct {
	AreaNames                map[string]string `json:"areaNames"`
	AudienceSubCategoryNames map[string]string `json:"audienceSubCategoryNames"`
	Events                   map[string]struct {
		Description interface{} `json:"description"`
		ID          int         `json:"id"`
		Logo        interface{} `json:"logo"`
		Name        string      `json:"name"`
		SubTopicIds []int       `json:"subTopicIds"`
		SubjectCode interface{} `json:"subjectCode"`
		Subtitle    interface{} `json:"subtitle"`
		TopicIds    []int       `json:"topicIds"`
	} `json:"events"`
	Performances []struct {
		EventID int         `json:"eventId"`
		ID      int         `json:"id"`
		Logo    interface{} `json:"logo"`
		Name    interface{} `json:"name"`
		Prices  []struct {
			Amount                int `json:"amount"`
			AudienceSubCategoryID int `json:"audienceSubCategoryId"`
			SeatCategoryID        int `json:"seatCategoryId"`
		} `json:"prices"`
		SeatCategories []struct {
			Areas []struct {
				AreaID   int           `json:"areaId"`
				BlockIds []interface{} `json:"blockIds"`
			} `json:"areas"`
			SeatCategoryID int `json:"seatCategoryId"`
		} `json:"seatCategories"`
		SeatMapImage interface{} `json:"seatMapImage"`
		Start        int64       `json:"start"`
		VenueCode    string      `json:"venueCode"`
	} `json:"performances"`
	SeatCategoryNames map[string]string `json:"seatCategoryNames"`
	SubTopicNames     map[string]string `json:"subTopicNames"`
	TopicNames        map[string]string `json:"topicNames"`
	TopicSubTopics    map[string]string `json:"topicSubTopics"`
	VenueNames        map[string]string `json:"venueNames"`
}

type GeoDataPayload struct {
	Type     string `json:"type"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			Name string `json:"name"`
		} `json:"properties"`
		Geometry struct {
			Type        string        `json:"type"`
			Coordinates [][][]float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

func ReadFile(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	return bytes
}

func GenerateObjectFromFile(filename string, result interface{}) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(bytes, &result)
}
