package main

import "fmt"


type AnalytycalDataService interface {
	SendXmlData()
}

type XmlDocument struct {
}
 
func (d XmlDocument) SendXmlData() {
	fmt.Println("Sendind XML")
}


type JsonDocument struct {		
}

func (d JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

type JsonDocumentAdapter struct {
	JsonDocument *JsonDocument
}

func(adapter JsonDocumentAdapter) SendXmlData() {
	adapter.JsonDocument.ConvertToXml()
	fmt.Println("Sendind XML converted")
}


func main() {
	var jsonDoc JsonDocument
	
	adapter := JsonDocumentAdapter{
		JsonDocument: &jsonDoc,
	}
	adapter.SendXmlData()
}
