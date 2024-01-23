package securitystatusrequest

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// SecurityStatusRequest is the fix50sp2 SecurityStatusRequest type, MsgType = e.
type SecurityStatusRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a SecurityStatusRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) SecurityStatusRequest {
	return SecurityStatusRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m SecurityStatusRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a SecurityStatusRequest initialized with the required fields for SecurityStatusRequest.
func New(securitystatusreqid field.SecurityStatusReqIDField, subscriptionrequesttype field.SubscriptionRequestTypeField) (m SecurityStatusRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("e"))
	m.Set(securitystatusreqid)
	m.Set(subscriptionrequesttype)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg SecurityStatusRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "e", r
}

// SetCurrency sets Currency, Tag 15.
func (m SecurityStatusRequest) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22.
func (m SecurityStatusRequest) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m SecurityStatusRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m SecurityStatusRequest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263.
func (m SecurityStatusRequest) SetSubscriptionRequestType(v enum.SubscriptionRequestType) {
	m.Set(field.NewSubscriptionRequestType(v))
}

// SetSecurityStatusReqID sets SecurityStatusReqID, Tag 324.
func (m SecurityStatusRequest) SetSecurityStatusReqID(v string) {
	m.Set(field.NewSecurityStatusReqID(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m SecurityStatusRequest) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// GetCurrency gets Currency, Tag 15.
func (m SecurityStatusRequest) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m SecurityStatusRequest) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m SecurityStatusRequest) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m SecurityStatusRequest) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263.
func (m SecurityStatusRequest) GetSubscriptionRequestType() (v enum.SubscriptionRequestType, err quickfix.MessageRejectError) {
	var f field.SubscriptionRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityStatusReqID gets SecurityStatusReqID, Tag 324.
func (m SecurityStatusRequest) GetSecurityStatusReqID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityStatusReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m SecurityStatusRequest) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m SecurityStatusRequest) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m SecurityStatusRequest) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m SecurityStatusRequest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m SecurityStatusRequest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263.
func (m SecurityStatusRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

// HasSecurityStatusReqID returns true if SecurityStatusReqID is present, Tag 324.
func (m SecurityStatusRequest) HasSecurityStatusReqID() bool {
	return m.Has(tag.SecurityStatusReqID)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m SecurityStatusRequest) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}
