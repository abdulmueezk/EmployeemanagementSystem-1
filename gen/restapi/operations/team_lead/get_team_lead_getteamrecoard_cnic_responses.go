// Code generated by go-swagger; DO NOT EDIT.

package team_lead

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"EmployeemanagementSystem/gen/models"
)

// GetTeamLeadGetteamrecoardCnicOKCode is the HTTP code returned for type GetTeamLeadGetteamrecoardCnicOK
const GetTeamLeadGetteamrecoardCnicOKCode int = 200

/*GetTeamLeadGetteamrecoardCnicOK successful operation

swagger:response getTeamLeadGetteamrecoardCnicOK
*/
type GetTeamLeadGetteamrecoardCnicOK struct {

	/*
	  In: Body
	*/
	Payload *models.Teamleademployeeofficial `json:"body,omitempty"`
}

// NewGetTeamLeadGetteamrecoardCnicOK creates GetTeamLeadGetteamrecoardCnicOK with default headers values
func NewGetTeamLeadGetteamrecoardCnicOK() *GetTeamLeadGetteamrecoardCnicOK {

	return &GetTeamLeadGetteamrecoardCnicOK{}
}

// WithPayload adds the payload to the get team lead getteamrecoard cnic o k response
func (o *GetTeamLeadGetteamrecoardCnicOK) WithPayload(payload *models.Teamleademployeeofficial) *GetTeamLeadGetteamrecoardCnicOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get team lead getteamrecoard cnic o k response
func (o *GetTeamLeadGetteamrecoardCnicOK) SetPayload(payload *models.Teamleademployeeofficial) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamLeadGetteamrecoardCnicOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTeamLeadGetteamrecoardCnicBadRequestCode is the HTTP code returned for type GetTeamLeadGetteamrecoardCnicBadRequest
const GetTeamLeadGetteamrecoardCnicBadRequestCode int = 400

/*GetTeamLeadGetteamrecoardCnicBadRequest Invalid CNIC

swagger:response getTeamLeadGetteamrecoardCnicBadRequest
*/
type GetTeamLeadGetteamrecoardCnicBadRequest struct {
}

// NewGetTeamLeadGetteamrecoardCnicBadRequest creates GetTeamLeadGetteamrecoardCnicBadRequest with default headers values
func NewGetTeamLeadGetteamrecoardCnicBadRequest() *GetTeamLeadGetteamrecoardCnicBadRequest {

	return &GetTeamLeadGetteamrecoardCnicBadRequest{}
}

// WriteResponse to the client
func (o *GetTeamLeadGetteamrecoardCnicBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
