package handlers

import (
	"encoding/json"
	"net/http"
)

// PrepareHeaders prépare les `headers` de la rèponse
func PrepareHeaders(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cache-control")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Accept", "application/json, application/xml")

	w.WriteHeader(statusCode)
}

// ReturnResponse Renvoi le message bien formaté
func ReturnResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	PrepareHeaders(w, statusCode)
	json.NewEncoder(w).Encode(message)
}

// ReturnResponse Renvoi le message bien formaté
func ReturnResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	PrepareHeaders(w, statusCode)
	json.NewEncoder(w).Encode(message)
}

// ErrorResponse is a json response of errors.
func ErrorResponse(w http.ResponseWriter, status int, messageTypes *models.MessageTypes, err error) {
	switch status {
	case http.StatusNotModified:
		KnownErrorResponse(w, status, messageTypes.NotModified, err)
	case http.StatusBadRequest:
		KnownErrorResponse(w, status, messageTypes.BadRequest, err)
	case http.StatusUnauthorized:
		KnownErrorResponse(w, status, messageTypes.Unauthorized, err)
	case http.StatusPaymentRequired:
		KnownErrorResponse(w, status, messageTypes.PaymentRequired, err)
	case http.StatusForbidden:
		KnownErrorResponse(w, status, messageTypes.Forbidden, err)
	case http.StatusNotFound:
		KnownErrorResponse(w, status, messageTypes.NotFound, err)
	case http.StatusMethodNotAllowed:
		KnownErrorResponse(w, status, messageTypes.MethodNotAllowed, err)
	case http.StatusInternalServerError:
		KnownErrorResponse(w, status, messageTypes.InternalServerError, err)
	case http.StatusConflict:
		KnownErrorResponse(w, status, messageTypes.Conflict, err)
	default:
		UnknownErrorResponse(w, status, err)
	}
}

// SuccessResponse est une réponse Json 2xx.
func SuccessResponse(w http.ResponseWriter, status int, messageType string, msg string) {
	ReturnResponse(w, status, models.Success(status, messageType, msg))
}

// RedirectionResponse est une réponse Json 3xx.
func RedirectionResponse(w http.ResponseWriter, status int, messageType string, msg string) {
	ReturnResponse(w, status, models.Redirection(status, messageType, msg))
}

// KnownErrorResponse est une réponse Json connu.
func KnownErrorResponse(w http.ResponseWriter, status int, messageType string, err error) {
	ReturnResponse(w, status, models.KnownError(status, messageType, err))
}

// UnknownErrorResponse est une réponse Json iconnu.
func UnknownErrorResponse(w http.ResponseWriter, status int, err error) {
	ReturnResponse(w, status, models.UnknownError(status, err))
}
