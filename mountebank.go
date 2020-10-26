package converter

import "fmt"

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

func (stub Stub) ToMap() map[string]interface{} {
	deepEqualsMap := make(map[string]interface{})

	for k := 0; k < len(stub.Predicates); k++ {
		predicate := stub.Predicates[k]
		fmt.Printf("%+v\n", predicate)
		for key, value := range predicate.DeepEquals {
			deepEqualsMap[key] = value
		}
	}
	return deepEqualsMap
}

type Stub struct {
	Predicates []struct {
		CaseSensitive bool                   `json:"caseSensitive"`
		DeepEquals    map[string]interface{} `json:"deepEquals"`
	} `json:"predicates"`
	Responses []struct {
		Is        Is `json:"is"`
		Behaviors struct {
			Wait int `json:"wait"`
		} `json:"_behaviors"`
	} `json:"responses"`
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

func (is Is) GetHeaderFromList(filter ...string) map[string]string {
	returnMap := make(map[string]string)
	for _, key := range filter {
		returnMap[key] = is.Headers[key]
	}
	return returnMap
}

type Is struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
	Mode       string            `json:"_mode"`
}
