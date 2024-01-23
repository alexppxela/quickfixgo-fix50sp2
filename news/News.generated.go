package news

import (
	"time"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// News is the fix50sp2 News type, MsgType = B.
type News struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a News from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) News {
	return News{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m News) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a News initialized with the required fields for News.
func New(headline field.HeadlineField) (m News) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("B"))
	m.Set(headline)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg News, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "B", r
}

// SetNoLinesOfText sets NoLinesOfText, Tag 33.
func (m News) SetNoLinesOfText(f NoLinesOfTextRepeatingGroup) {
	m.SetGroup(f)
}

// SetOrigTime sets OrigTime, Tag 42.
func (m News) SetOrigTime(v time.Time) {
	m.Set(field.NewOrigTime(v))
}

// SetUrgency sets Urgency, Tag 61.
func (m News) SetUrgency(v enum.Urgency) {
	m.Set(field.NewUrgency(v))
}

// SetNoRelatedSym sets NoRelatedSym, Tag 146.
func (m News) SetNoRelatedSym(f NoRelatedSymRepeatingGroup) {
	m.SetGroup(f)
}

// SetHeadline sets Headline, Tag 148.
func (m News) SetHeadline(v string) {
	m.Set(field.NewHeadline(v))
}

// SetURLLink sets URLLink, Tag 149.
func (m News) SetURLLink(v string) {
	m.Set(field.NewURLLink(v))
}

// SetMarketID sets MarketID, Tag 1301.
func (m News) SetMarketID(v string) {
	m.Set(field.NewMarketID(v))
}

// GetNoLinesOfText gets NoLinesOfText, Tag 33.
func (m News) GetNoLinesOfText() (NoLinesOfTextRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLinesOfTextRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetOrigTime gets OrigTime, Tag 42.
func (m News) GetOrigTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.OrigTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUrgency gets Urgency, Tag 61.
func (m News) GetUrgency() (v enum.Urgency, err quickfix.MessageRejectError) {
	var f field.UrgencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRelatedSym gets NoRelatedSym, Tag 146.
func (m News) GetNoRelatedSym() (NoRelatedSymRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedSymRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetHeadline gets Headline, Tag 148.
func (m News) GetHeadline() (v string, err quickfix.MessageRejectError) {
	var f field.HeadlineField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetURLLink gets URLLink, Tag 149.
func (m News) GetURLLink() (v string, err quickfix.MessageRejectError) {
	var f field.URLLinkField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketID gets MarketID, Tag 1301.
func (m News) GetMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasNoLinesOfText returns true if NoLinesOfText is present, Tag 33.
func (m News) HasNoLinesOfText() bool {
	return m.Has(tag.NoLinesOfText)
}

// HasOrigTime returns true if OrigTime is present, Tag 42.
func (m News) HasOrigTime() bool {
	return m.Has(tag.OrigTime)
}

// HasUrgency returns true if Urgency is present, Tag 61.
func (m News) HasUrgency() bool {
	return m.Has(tag.Urgency)
}

// HasNoRelatedSym returns true if NoRelatedSym is present, Tag 146.
func (m News) HasNoRelatedSym() bool {
	return m.Has(tag.NoRelatedSym)
}

// HasHeadline returns true if Headline is present, Tag 148.
func (m News) HasHeadline() bool {
	return m.Has(tag.Headline)
}

// HasURLLink returns true if URLLink is present, Tag 149.
func (m News) HasURLLink() bool {
	return m.Has(tag.URLLink)
}

// HasMarketID returns true if MarketID is present, Tag 1301.
func (m News) HasMarketID() bool {
	return m.Has(tag.MarketID)
}

// NoLinesOfText is a repeating group element, Tag 33.
type NoLinesOfText struct {
	*quickfix.Group
}

// SetText sets Text, Tag 58.
func (m NoLinesOfText) SetText(v string) {
	m.Set(field.NewText(v))
}

// GetText gets Text, Tag 58.
func (m NoLinesOfText) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasText returns true if Text is present, Tag 58.
func (m NoLinesOfText) HasText() bool {
	return m.Has(tag.Text)
}

// NoLinesOfTextRepeatingGroup is a repeating group, Tag 33.
type NoLinesOfTextRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoLinesOfTextRepeatingGroup returns an initialized, NoLinesOfTextRepeatingGroup.
func NewNoLinesOfTextRepeatingGroup() NoLinesOfTextRepeatingGroup {
	return NoLinesOfTextRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoLinesOfText,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.Text),
			},
		),
	}
}

// Add create and append a new NoLinesOfText to this group.
func (m NoLinesOfTextRepeatingGroup) Add() NoLinesOfText {
	g := m.RepeatingGroup.Add()
	return NoLinesOfText{g}
}

// Get returns the ith NoLinesOfText in the NoLinesOfTextRepeatinGroup.
func (m NoLinesOfTextRepeatingGroup) Get(i int) NoLinesOfText {
	return NoLinesOfText{m.RepeatingGroup.Get(i)}
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
