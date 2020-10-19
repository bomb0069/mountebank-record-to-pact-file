package converter

import "fmt"

func (stub Stub) ToMap() map[string]string {
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

func (is Is) GetHeader(filter string) map[string]string {
	returnMap := make(map[string]string)
	for key, element := range is.Headers {
		if key == filter {
			returnMap[key] = element
		}
	}
	return returnMap
}

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
		Is        Is `json:"is"`
		Behaviors struct {
			Wait int `json:"wait"`
		} `json:"_behaviors"`
	} `json:"responses"`
}

type Is struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
	Mode       string            `json:"_mode"`
}
