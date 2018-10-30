package apiteamamerica

import (
	//"bytes"
	"encoding/xml"
	//"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"io/ioutil"
	"net/http"
	"strings"
)

//Taclient is struct to hold authentificate parameters to URL
type Taclient struct {
	Username string //user name
	Password string //pass
	URL      string //client url fetc
}

//Connect is to start coonecting and returns
func (taClient *Taclient) Connect() (string, error) {

	return taClient.URL, nil
}

//ListCities returns a list with cities
func (taClient *Taclient) ListCities(r *http.Request) ([]byte, error) {

	xparam := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.wso2.org/php/xsd">
	<soapenv:Header/>
	<soapenv:Body>
	   <xsd:ListCities>
		  <xsd:UserName>test</xsd:UserName>
		  <xsd:Password>12345</xsd:Password>
	   </xsd:ListCities>
	</soapenv:Body>
 </soapenv:Envelope>`

	//must be a Google Api Engine Context
	ctx := appengine.NewContext(r)

	//get *http.Client from GAE urlfetch package
	client := urlfetch.Client(ctx)

	//invoke a POST call with xparam
	resp, err := client.Post(taClient.URL, "text/xml; charset=utf-8", strings.NewReader(xparam))

	if err != nil {
		return []byte("Error"), err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	/*

		//testing decode xml
		var envelope XmlEnvelope
		// we unmarshal our byteArray which contains our specific data
		xml.Unmarshal(body, &envelope)

		//	return buf.String(), err
		return fmt.Sprintln(envelope.Body), err
	*/
	return body, err
}

/* XML Response for fault
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
	<soapenv:Body>
		<soapenv:Fault>
			<faultcode>soapenv:Server</faultcode>
			<faultstring> Error - Login Invalid</faultstring>
			<detail/>
		</soapenv:Fault>
	</soapenv:Body>
</soapenv:Envelope>
*/

//XMLEnvelope is generic
type XMLEnvelope struct {
	XMLName xml.Name  `xml:"Envelope"`
	Body    []XMLBody `xml:"Body"`
}

//XMLBody is generic
type XMLBody struct {
	XMLName xml.Name `xml:"Body"`
	content string
}
