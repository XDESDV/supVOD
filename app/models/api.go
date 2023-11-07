package models

// WSResponse est le format standardisé de la réponse.
//   - Meta         : *Entête pré-formaté d'une réponse retournant des données.*
//   - Data         : *Donnée ou liste de données retournée(s).*
type WSResponse struct {
	Meta MetaResponse `json:"meta"`
	Data interface{}  `json:"data"`
}

// MetaResponse est une entête d'une réponse valide
//   - ObjectName  : *Information retourné au front lui permettant de savoir quel format il reçoit.*
//   - TotalCount  : *Nombre total d'enregistrement que la demande peut retourner.*
//   - Offset      : *Position de départ de la liste des enregistrements retournés au Front.*
//   - Count       : *Nombre d'enregistrement retourné au Front.*
type MetaResponse struct {
	ObjectName string `json:"object_name"`
	TotalCount int    `json:"total_count"`
	Offset     int    `json:"offSet"`
	Count      int    `json:"count"`
}

// MessageTypes est un tableau de message type retourné au Front
// *	+ OK                  : *200*
// *	+ Created             : *201*
//   - NotModified         : *304*
//
// !	+ BadRequest          : *400*
// !	+ Unauthorized        : *401*
// !	+ PaymentRequired     : *402*
// !	- Forbidden           : *403*
// !	* NotFound            : *404*
// !	+ MethodNotAllowed    : *405*
// ! + Conflict	 	 	  :*409*
// !	+ InternalServerError : *500*
type MessageTypes struct {
	OK                  string
	Created             string
	NotModified         string
	BadRequest          string
	Unauthorized        string
	PaymentRequired     string
	Forbidden           string
	NotFound            string
	Conflict            string
	MethodNotAllowed    string
	InternalServerError string
}

// BasicResponse est une réponse basic.
//   - Status : *http status*
//   - MessageType : *message typé pour le Front au format I18N*
//   - Message : *message réponse*
type BasicResponse struct {
	Status      int    `json:"status"`
	MessageType string `json:"messageType"`
	Message     string `json:"message"`
}

// Success is a basic response of 2xx.
func Success(status int, messageType string, msg string) *BasicResponse {
	return &BasicResponse{Status: status, MessageType: messageType, Message: msg}
}

// Redirection is a basic response of 3xx.
func Redirection(status int, messageType string, msg string) *BasicResponse {
	return &BasicResponse{Status: status, MessageType: messageType, Message: msg}
}

// KnownError is a basic response of know errors.
func KnownError(status int, messageType string, err error) *BasicResponse {
	return &BasicResponse{Status: status, MessageType: messageType, Message: err.Error()}
}

// UnknownError is a basic response of unknown errors.
func UnknownError(status int, err error) *BasicResponse {
	return KnownError(status, "error.unknown", err)
}
