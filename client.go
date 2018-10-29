package apiteamamerica

import (
	//"bytes"
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
		return []byte("Eroare"), err
	}

	body, err := ioutil.ReadAll(resp.Body)

	//	return buf.String(), err
	return body, err
}
