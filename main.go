package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type SoapEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    SoapBody
}

type SoapBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Result  string   `xml:"GetCityWeatherByZIPResult"`
}

func main() {
	soapXML := `
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ws="http://ws.cdyne.com/WeatherWS/">
	   <soapenv:Header/>
	   <soapenv:Body>
		  <ws:GetCityWeatherByZIP>
			 <ws:ZIP>60187</ws:ZIP>
		  </ws:GetCityWeatherByZIP>
	   </soapenv:Body>
	</soapenv:Envelope>`

	req, err := http.NewRequest("POST", "http://wsf.cdyne.com/WeatherWS/Weather.asmx", bytes.NewBufferString(soapXML))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "http://ws.cdyne.com/WeatherWS/GetCityWeatherByZIP")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	
	var envelope SoapEnvelope
	// Nota: O unmarshal pode exigir estruturas mais complexas dependendo do WSDL real
	fmt.Println(string(body))
}