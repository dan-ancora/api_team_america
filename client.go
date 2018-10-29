package apiteamamerica

//Taclient is struct to hold authentificate parameters to URL
type Taclient struct {
	Username string //user name
	Password string //pass
	URL      string //client url fetc
}

//Connect is to start coonecting and returns
func (*Taclient) Connect() (string, error) {
	return "Ok", nil
}
