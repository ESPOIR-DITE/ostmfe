package booking

type Booking struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phoneNumber"`
	Organisation string `json:"organisation"`
	Language     string `json:"language"`
	Purpose      string `json:"purpose"`
	Country      string `json:"country"`
	Province     string `json:"province"`
	Message      string `json:"message"`
	Date         string `json:"date"`
	IsFirstVisit bool   `json:"isFirstVisit"`
	Need         string `json:"need"`
}

type BookingAddress struct {
	AddressId   string `json:"addressId"`
	BookingId   string `json:"bookingId"`
	Description string `json:"description"`
}

type BookingTransport struct {
	TransportId string `json:"transportId"`
	BookingId   string `json:"bookingId"`
	Description string `json:"description"`
}

type BookingLanguage struct {
	LanguageId  string `json:"languageId"`
	BookId      string `json:"bookId"`
	Description string `json:"description"`
}

type BookType struct {
	Id          string `json:"id"`
	TypeName    string `json:"type_name"`
	Description string `json:"description"`
}
