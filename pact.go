package converter

func (request *Request) FromMap(requestMap map[string]string) {
	request.Method = requestMap["method"]
	request.Path = requestMap["path"]
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
	Method string `json:"method"`
	Path   string `json:"path"`
}

type PactSpecification struct {
	Version string `json:"version"`
}
