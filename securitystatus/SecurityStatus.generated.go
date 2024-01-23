package securitystatus

import (
	"time"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// SecurityStatus is the fix50sp2 SecurityStatus type, MsgType = f.
type SecurityStatus struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a SecurityStatus from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) SecurityStatus {
	return SecurityStatus{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m SecurityStatus) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a SecurityStatus initialized with the required fields for SecurityStatus.
func New() (m SecurityStatus) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("f"))

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg SecurityStatus, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "f", r
}

// SetCurrency sets Currency, Tag 15.
func (m SecurityStatus) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22.
func (m SecurityStatus) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m SecurityStatus) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m SecurityStatus) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m SecurityStatus) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetSecurityStatusReqID sets SecurityStatusReqID, Tag 324.
func (m SecurityStatus) SetSecurityStatusReqID(v string) {
	m.Set(field.NewSecurityStatusReqID(v))
}

// SetSecurityTradingStatus sets SecurityTradingStatus, Tag 326.
func (m SecurityStatus) SetSecurityTradingStatus(v enum.SecurityTradingStatus) {
	m.Set(field.NewSecurityTradingStatus(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m SecurityStatus) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// GetCurrency gets Currency, Tag 15.
func (m SecurityStatus) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m SecurityStatus) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m SecurityStatus) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m SecurityStatus) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m SecurityStatus) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityStatusReqID gets SecurityStatusReqID, Tag 324.
func (m SecurityStatus) GetSecurityStatusReqID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityStatusReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityTradingStatus gets SecurityTradingStatus, Tag 326.
func (m SecurityStatus) GetSecurityTradingStatus() (v enum.SecurityTradingStatus, err quickfix.MessageRejectError) {
	var f field.SecurityTradingStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m SecurityStatus) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m SecurityStatus) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m SecurityStatus) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m SecurityStatus) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m SecurityStatus) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m SecurityStatus) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasSecurityStatusReqID returns true if SecurityStatusReqID is present, Tag 324.
func (m SecurityStatus) HasSecurityStatusReqID() bool {
	return m.Has(tag.SecurityStatusReqID)
}

// HasSecurityTradingStatus returns true if SecurityTradingStatus is present, Tag 326.
func (m SecurityStatus) HasSecurityTradingStatus() bool {
	return m.Has(tag.SecurityTradingStatus)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m SecurityStatus) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}
