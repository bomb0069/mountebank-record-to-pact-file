package converter

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func (request *Request) FromMap(requestMap map[string]interface{}) {
	request.Method = fmt.Sprintf("%v", requestMap["method"])
	request.Path = fmt.Sprintf("%v", requestMap["path"])

	var requestBodyJson map[string]interface{}
	json.Unmarshal([]byte(fmt.Sprintf("%v", requestMap["body"])), &requestBodyJson)
	request.Body = requestBodyJson

	request.Headers = make(map[string]string)

	fmt.Printf("request-header :------------------------: %v", requestMap["headers"])
	fmt.Printf("request-header-type :------------------------: %v", reflect.TypeOf(requestMap["headers"]))
	switch v := requestMap["headers"].(type) {
	case map[string]interface{}:
		for key, value := range v {
			request.Headers[key] = fmt.Sprintf("%v", value)
		}
	default:
		fmt.Printf("XXXXX %v", v)
	}
}

type Pact struct {
	Consumer struct {
		Name string `json:"name"`
	} `json:"consumer"`
	Provider struct {
		Name string `json:"name"`
	} `json:"provider"`
	Interactions []Interaction `json:"interactions"`
	Metadata     struct {
		PactSpecification PactSpecification `json:"pactSpecification"`
	} `json:"metadata"`
}

type Interaction struct {
	Description string  `json:"description"`
	Request     Request `json:"request"`
	Response    struct {
		Status  int                    `json:"status"`
		Headers map[string]string      `json:"headers"`
		Body    map[string]interface{} `json:"body"`
	} `json:"response"`
}

type Request struct {
	Method  string                 `json:"method"`
	Path    string                 `json:"path"`
	Body    map[string]interface{} `json:"body"`
	Headers map[string]string      `json:"headers"`
}

type PactSpecification struct {
	Version string `json:"version"`
}
