package securitylist

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// SecurityList is the fix50sp2 SecurityList type, MsgType = y.
type SecurityList struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a SecurityList from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) SecurityList {
	return SecurityList{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m SecurityList) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a SecurityList initialized with the required fields for SecurityList.
func New() (m SecurityList) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("y"))

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg SecurityList, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "y", r
}

// SetNoRelatedSym sets NoRelatedSym, Tag 146.
func (m SecurityList) SetNoRelatedSym(f NoRelatedSymRepeatingGroup) {
	m.SetGroup(f)
}

// SetSecurityReqID sets SecurityReqID, Tag 320.
func (m SecurityList) SetSecurityReqID(v string) {
	m.Set(field.NewSecurityReqID(v))
}

// SetSecurityResponseID sets SecurityResponseID, Tag 322.
func (m SecurityList) SetSecurityResponseID(v string) {
	m.Set(field.NewSecurityResponseID(v))
}

// SetTotNoRelatedSym sets TotNoRelatedSym, Tag 393.
func (m SecurityList) SetTotNoRelatedSym(v int) {
	m.Set(field.NewTotNoRelatedSym(v))
}

// SetSecurityRequestResult sets SecurityRequestResult, Tag 560.
func (m SecurityList) SetSecurityRequestResult(v enum.SecurityRequestResult) {
	m.Set(field.NewSecurityRequestResult(v))
}

// SetLastFragment sets LastFragment, Tag 893.
func (m SecurityList) SetLastFragment(v bool) {
	m.Set(field.NewLastFragment(v))
}

// GetNoRelatedSym gets NoRelatedSym, Tag 146.
func (m SecurityList) GetNoRelatedSym() (NoRelatedSymRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedSymRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetSecurityReqID gets SecurityReqID, Tag 320.
func (m SecurityList) GetSecurityReqID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityResponseID gets SecurityResponseID, Tag 322.
func (m SecurityList) GetSecurityResponseID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityResponseIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotNoRelatedSym gets TotNoRelatedSym, Tag 393.
func (m SecurityList) GetTotNoRelatedSym() (v int, err quickfix.MessageRejectError) {
	var f field.TotNoRelatedSymField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityRequestResult gets SecurityRequestResult, Tag 560.
func (m SecurityList) GetSecurityRequestResult() (v enum.SecurityRequestResult, err quickfix.MessageRejectError) {
	var f field.SecurityRequestResultField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLastFragment gets LastFragment, Tag 893.
func (m SecurityList) GetLastFragment() (v bool, err quickfix.MessageRejectError) {
	var f field.LastFragmentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasNoRelatedSym returns true if NoRelatedSym is present, Tag 146.
func (m SecurityList) HasNoRelatedSym() bool {
	return m.Has(tag.NoRelatedSym)
}

// HasSecurityReqID returns true if SecurityReqID is present, Tag 320.
func (m SecurityList) HasSecurityReqID() bool {
	return m.Has(tag.SecurityReqID)
}

// HasSecurityResponseID returns true if SecurityResponseID is present, Tag 322.
func (m SecurityList) HasSecurityResponseID() bool {
	return m.Has(tag.SecurityResponseID)
}

// HasTotNoRelatedSym returns true if TotNoRelatedSym is present, Tag 393.
func (m SecurityList) HasTotNoRelatedSym() bool {
	return m.Has(tag.TotNoRelatedSym)
}

// HasSecurityRequestResult returns true if SecurityRequestResult is present, Tag 560.
func (m SecurityList) HasSecurityRequestResult() bool {
	return m.Has(tag.SecurityRequestResult)
}

// HasLastFragment returns true if LastFragment is present, Tag 893.
func (m SecurityList) HasLastFragment() bool {
	return m.Has(tag.LastFragment)
}

// NoRelatedSym is a repeating group element, Tag 146.
type NoRelatedSym struct {
	*quickfix.Group
}

// SetSymbol sets Symbol, Tag 55.
func (m NoRelatedSym) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m NoRelatedSym) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22.
func (m NoRelatedSym) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetCurrency sets Currency, Tag 15.
func (m NoRelatedSym) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// GetSymbol gets Symbol, Tag 55.
func (m NoRelatedSym) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m NoRelatedSym) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m NoRelatedSym) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCurrency gets Currency, Tag 15.
func (m NoRelatedSym) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m NoRelatedSym) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m NoRelatedSym) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m NoRelatedSym) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m NoRelatedSym) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// NoRelatedSymRepeatingGroup is a repeating group, Tag 146.
type NoRelatedSymRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedSymRepeatingGroup returns an initialized, NoRelatedSymRepeatingGroup.
func NewNoRelatedSymRepeatingGroup() NoRelatedSymRepeatingGroup {
	return NoRelatedSymRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRelatedSym,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.Symbol),
				quickfix.GroupElement(tag.SecurityID),
				quickfix.GroupElement(tag.SecurityIDSource),
				quickfix.GroupElement(tag.Currency),
			},
		),
	}
}

// Add create and append a new NoRelatedSym to this group.
func (m NoRelatedSymRepeatingGroup) Add() NoRelatedSym {
	g := m.RepeatingGroup.Add()
	return NoRelatedSym{g}
}

// Get returns the ith NoRelatedSym in the NoRelatedSymRepeatinGroup.
func (m NoRelatedSymRepeatingGroup) Get(i int) NoRelatedSym {
	return NoRelatedSym{m.RepeatingGroup.Get(i)}
}
