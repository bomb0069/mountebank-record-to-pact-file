package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	jsonFile, err := os.Open("./input/todo-1.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened First Todo")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var mountebank Mountebank

	pact := new(Pact)

	json.Unmarshal(byteValue, &mountebank)

	pact.Consumer.Name = "Localhost"
	pact.Provider.Name = mountebank.Imposters[0].Name

	pact.Interactions = make([]Interaction, len(mountebank.Imposters))

	for i := 0; i < len(mountebank.Imposters); i++ {
		imposter := mountebank.Imposters[i]

		for j := 0; j < len(imposter.Stubs); j++ {
			stub := imposter.Stubs[j]
			interaction := new(Interaction)

			deepEqualsMap := stub.toMap()

			fmt.Printf("%+v\n", deepEqualsMap)

			interaction.Request.Method = deepEqualsMap["method"]
			interaction.Request.Path = deepEqualsMap["path"]
			pact.Interactions[j] = *interaction
		}
	}

	file, _ := json.MarshalIndent(pact, "", "  ")

	_ = ioutil.WriteFile("./output/actual/todo-1.json", file, 0644)

	defer jsonFile.Close()
}

func (stub Stub) toMap() map[string]string {
	deepEqualsMap := make(map[string]string)

	for k := 0; k < len(stub.Predicates); k++ {
		predicate := stub.Predicates[k]
		fmt.Printf("%+v\n", predicate)
		for key, value := range predicate.DeepEquals {
			deepEqualsMap[key] = value
		}
	}
	return deepEqualsMap
}

//Mountebank

type Mountebank struct {
	Imposters []Imposter `json:"imposters"`
}

type Imposters struct {
	Imposters []Imposter `json:"imposters"`
}

type Imposter struct {
	Protocol       string `json:"protocol"`
	Port           int    `json:"port"`
	Name           string `json:"name"`
	RecordRequests bool   `json:"recordRequests"`
	Stubs          []Stub `json:"stubs"`
}

type Stub struct {
	Predicates []struct {
		CaseSensitive bool              `json:"caseSensitive"`
		DeepEquals    map[string]string `json:"deepEquals"`
	} `json:"predicates"`
	Responses []struct {
		Is struct {
			StatusCode int `json:"statusCode"`
			Headers    struct {
				Date                          string `json:"Date"`
				ContentType                   string `json:"Content-Type"`
				ContentLength                 string `json:"Content-Length"`
				Connection                    string `json:"Connection"`
				SetCookie                     string `json:"Set-Cookie"`
				XPoweredBy                    string `json:"X-Powered-By"`
				XRatelimitLimit               string `json:"X-Ratelimit-Limit"`
				XRatelimitRemaining           string `json:"X-Ratelimit-Remaining"`
				XRatelimitReset               string `json:"X-Ratelimit-Reset"`
				Vary                          string `json:"Vary"`
				AccessControlAllowCredentials string `json:"Access-Control-Allow-Credentials"`
				CacheControl                  string `json:"Cache-Control"`
				Pragma                        string `json:"Pragma"`
				Expires                       string `json:"Expires"`
				XContentTypeOptions           string `json:"X-Content-Type-Options"`
				Etag                          string `json:"Etag"`
				Via                           string `json:"Via"`
				CFCacheStatus                 string `json:"CF-Cache-Status"`
				Age                           string `json:"Age"`
				AcceptRanges                  string `json:"Accept-Ranges"`
				CfRequestID                   string `json:"cf-request-id"`
				ExpectCT                      string `json:"Expect-CT"`
				ReportTo                      string `json:"Report-To"`
				NEL                           string `json:"NEL"`
				Server                        string `json:"Server"`
				CFRAY                         string `json:"CF-RAY"`
			} `json:"headers"`
			Body string `json:"body"`
			Mode string `json:"_mode"`
		} `json:"is"`
		Behaviors struct {
			Wait int `json:"wait"`
		} `json:"_behaviors"`
	} `json:"responses"`
}

// Pact
type Pact struct {
	Consumer struct {
		Name string `json:"name"`
	} `json:"consumer"`
	Provider struct {
		Name string `json:"name"`
	} `json:"provider"`
	Interactions []Interaction `json:"interactions"`
	Metadata     struct {
		PactSpecification struct {
			Version string `json:"version"`
		} `json:"pactSpecification"`
	} `json:"metadata"`
}

type Interaction struct {
	Description string `json:"description"`
	Request     struct {
		Method string `json:"method"`
		Path   string `json:"path"`
	} `json:"request"`
	Response struct {
		Status  int `json:"status"`
		Headers struct {
			ContentType string `json:"Content-Type"`
		} `json:"headers"`
		Body struct {
			UserID    int    `json:"userId"`
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		} `json:"body"`
	} `json:"response"`
}
