package securitylistrequest

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// SecurityListRequest is the fix50sp2 SecurityListRequest type, MsgType = x.
type SecurityListRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a SecurityListRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) SecurityListRequest {
	return SecurityListRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m SecurityListRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a SecurityListRequest initialized with the required fields for SecurityListRequest.
func New(securityreqid field.SecurityReqIDField, securitylistrequesttype field.SecurityListRequestTypeField) (m SecurityListRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("x"))
	m.Set(securityreqid)
	m.Set(securitylistrequesttype)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg SecurityListRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "x", r
}

// SetSecurityReqID sets SecurityReqID, Tag 320.
func (m SecurityListRequest) SetSecurityReqID(v string) {
	m.Set(field.NewSecurityReqID(v))
}

// SetSecurityListRequestType sets SecurityListRequestType, Tag 559.
func (m SecurityListRequest) SetSecurityListRequestType(v enum.SecurityListRequestType) {
	m.Set(field.NewSecurityListRequestType(v))
}

// GetSecurityReqID gets SecurityReqID, Tag 320.
func (m SecurityListRequest) GetSecurityReqID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityListRequestType gets SecurityListRequestType, Tag 559.
func (m SecurityListRequest) GetSecurityListRequestType() (v enum.SecurityListRequestType, err quickfix.MessageRejectError) {
	var f field.SecurityListRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSecurityReqID returns true if SecurityReqID is present, Tag 320.
func (m SecurityListRequest) HasSecurityReqID() bool {
	return m.Has(tag.SecurityReqID)
}

// HasSecurityListRequestType returns true if SecurityListRequestType is present, Tag 559.
func (m SecurityListRequest) HasSecurityListRequestType() bool {
	return m.Has(tag.SecurityListRequestType)
}
