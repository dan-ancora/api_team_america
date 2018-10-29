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
func (client *Taclient) Connect() (string, error) {

	return client.URL, nil
}

//ListCities returns a list with cities
func (tac *Taclient) ListCities(r *http.Request) ([]byte, error) {

	xparam := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.wso2.org/php/xsd">
	<soapenv:Header/>
	<soapenv:Body>
	   <xsd:ListCities>
		  <xsd:UserName>XMLSMAY</xsd:UserName>
		  <xsd:Password>M3WgnuOV</xsd:Password>
	   </xsd:ListCities>
	</soapenv:Body>
 </soapenv:Envelope>`

	ctx := appengine.NewContext(r)

	client := urlfetch.Client(ctx)
	resp, err := client.Post(tac.URL, "text/xml", strings.NewReader(xparam))

	if err != nil {
		return []byte("Error"), err
	}

	body, err := ioutil.ReadAll(resp.Body)

	/*
		//testing decode xml
		var envelope XmlEnvelope
		// we unmarshal our byteArray which contains our
		// xmlFiles content into 'users' which we defined above
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

//Cities List strcut for decoding XML
type XmlEnvelope struct {
	XMLName xml.Name  `xml:"soapenv:Envelope"`
	Body    []XmlBody `xml:"soapenv:Body"`
}

//type body
type XmlBody struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	content string
}
