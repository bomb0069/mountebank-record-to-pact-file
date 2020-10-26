package converter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Convert(consumer string) {

	jsonFile, err := os.Open("./input/todo-1.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened First Todo")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var mounteBank Mountebank

	pact := new(Pact)

	json.Unmarshal(byteValue, &mounteBank)

	pact.Consumer.Name = consumer
	pact.Provider.Name = mounteBank.Imposters[0].Name

	pact.Interactions = make([]Interaction, 0)

	for _, imposter := range mounteBank.Imposters {

		for _, stub := range imposter.Stubs {

			interaction := new(Interaction)

			interaction.Request.FromMap(stub.ToMap())

			interaction.Response.Status = stub.Responses[0].Is.StatusCode
			interaction.Response.Headers = stub.Responses[0].Is.GetHeaderFromList("Content-Type", "X-Job-Id", "X-Request-Id", "X-Roundtrip")

			var bodyJson map[string]interface{}
			json.Unmarshal([]byte(stub.Responses[0].Is.Body), &bodyJson)

			interaction.Response.Body = bodyJson

			pact.Interactions = append(pact.Interactions, *interaction)
		}
	}
	pact.Metadata.PactSpecification.Version = "2.0.0"

	file, _ := json.MarshalIndent(pact, "", "  ")

	_ = ioutil.WriteFile("./output/actual/todo-1.json", file, 0644)

}
