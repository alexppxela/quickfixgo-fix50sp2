package quotecancel

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// QuoteCancel is the fix50sp2 QuoteCancel type, MsgType = Z.
type QuoteCancel struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a QuoteCancel from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) QuoteCancel {
	return QuoteCancel{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m QuoteCancel) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a QuoteCancel initialized with the required fields for QuoteCancel.
func New(quotecanceltype field.QuoteCancelTypeField) (m QuoteCancel) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("Z"))
	m.Set(quotecanceltype)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg QuoteCancel, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "Z", r
}

// SetAccount sets Account, Tag 1.
func (m QuoteCancel) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

// SetQuoteID sets QuoteID, Tag 117.
func (m QuoteCancel) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

// SetNoQuoteEntries sets NoQuoteEntries, Tag 295.
func (m QuoteCancel) SetNoQuoteEntries(f NoQuoteEntriesRepeatingGroup) {
	m.SetGroup(f)
}

// SetQuoteCancelType sets QuoteCancelType, Tag 298.
func (m QuoteCancel) SetQuoteCancelType(v enum.QuoteCancelType) {
	m.Set(field.NewQuoteCancelType(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m QuoteCancel) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetQuoteMsgID sets QuoteMsgID, Tag 1166.
func (m QuoteCancel) SetQuoteMsgID(v string) {
	m.Set(field.NewQuoteMsgID(v))
}

// GetAccount gets Account, Tag 1.
func (m QuoteCancel) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteID gets QuoteID, Tag 117.
func (m QuoteCancel) GetQuoteID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoQuoteEntries gets NoQuoteEntries, Tag 295.
func (m QuoteCancel) GetNoQuoteEntries() (NoQuoteEntriesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoQuoteEntriesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetQuoteCancelType gets QuoteCancelType, Tag 298.
func (m QuoteCancel) GetQuoteCancelType() (v enum.QuoteCancelType, err quickfix.MessageRejectError) {
	var f field.QuoteCancelTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m QuoteCancel) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetQuoteMsgID gets QuoteMsgID, Tag 1166.
func (m QuoteCancel) GetQuoteMsgID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteMsgIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAccount returns true if Account is present, Tag 1.
func (m QuoteCancel) HasAccount() bool {
	return m.Has(tag.Account)
}

// HasQuoteID returns true if QuoteID is present, Tag 117.
func (m QuoteCancel) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

// HasNoQuoteEntries returns true if NoQuoteEntries is present, Tag 295.
func (m QuoteCancel) HasNoQuoteEntries() bool {
	return m.Has(tag.NoQuoteEntries)
}

// HasQuoteCancelType returns true if QuoteCancelType is present, Tag 298.
func (m QuoteCancel) HasQuoteCancelType() bool {
	return m.Has(tag.QuoteCancelType)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m QuoteCancel) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasQuoteMsgID returns true if QuoteMsgID is present, Tag 1166.
func (m QuoteCancel) HasQuoteMsgID() bool {
	return m.Has(tag.QuoteMsgID)
}

// NoQuoteEntries is a repeating group element, Tag 295.
type NoQuoteEntries struct {
	*quickfix.Group
}

// SetSymbol sets Symbol, Tag 55.
func (m NoQuoteEntries) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m NoQuoteEntries) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22.
func (m NoQuoteEntries) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetCurrency sets Currency, Tag 15.
func (m NoQuoteEntries) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// GetSymbol gets Symbol, Tag 55.
func (m NoQuoteEntries) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m NoQuoteEntries) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m NoQuoteEntries) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCurrency gets Currency, Tag 15.
func (m NoQuoteEntries) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m NoQuoteEntries) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m NoQuoteEntries) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m NoQuoteEntries) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m NoQuoteEntries) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// NoQuoteEntriesRepeatingGroup is a repeating group, Tag 295.
type NoQuoteEntriesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoQuoteEntriesRepeatingGroup returns an initialized, NoQuoteEntriesRepeatingGroup.
func NewNoQuoteEntriesRepeatingGroup() NoQuoteEntriesRepeatingGroup {
	return NoQuoteEntriesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoQuoteEntries,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.Symbol),
				quickfix.GroupElement(tag.SecurityID),
				quickfix.GroupElement(tag.SecurityIDSource),
				quickfix.GroupElement(tag.Currency),
			},
		),
	}
}

// Add create and append a new NoQuoteEntries to this group.
func (m NoQuoteEntriesRepeatingGroup) Add() NoQuoteEntries {
	g := m.RepeatingGroup.Add()
	return NoQuoteEntries{g}
}

// Get returns the ith NoQuoteEntries in the NoQuoteEntriesRepeatinGroup.
func (m NoQuoteEntriesRepeatingGroup) Get(i int) NoQuoteEntries {
	return NoQuoteEntries{m.RepeatingGroup.Get(i)}
}

// NoPartyIDs is a repeating group element, Tag 453.
type NoPartyIDs struct {
	*quickfix.Group
}

// SetPartyID sets PartyID, Tag 448.
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

// SetPartyIDSource sets PartyIDSource, Tag 447.
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

// SetPartyRole sets PartyRole, Tag 452.
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

// SetNoPartySubIDs sets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetPartyRoleQualifier sets PartyRoleQualifier, Tag 2376.
func (m NoPartyIDs) SetPartyRoleQualifier(v enum.PartyRoleQualifier) {
	m.Set(field.NewPartyRoleQualifier(v))
}

// GetPartyID gets PartyID, Tag 448.
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyIDSource gets PartyIDSource, Tag 447.
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyRole gets PartyRole, Tag 452.
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartySubIDs gets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetPartyRoleQualifier gets PartyRoleQualifier, Tag 2376.
func (m NoPartyIDs) GetPartyRoleQualifier() (v enum.PartyRoleQualifier, err quickfix.MessageRejectError) {
	var f field.PartyRoleQualifierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartyID returns true if PartyID is present, Tag 448.
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

// HasPartyIDSource returns true if PartyIDSource is present, Tag 447.
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

// HasPartyRole returns true if PartyRole is present, Tag 452.
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

// HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802.
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

// HasPartyRoleQualifier returns true if PartyRoleQualifier is present, Tag 2376.
func (m NoPartyIDs) HasPartyRoleQualifier() bool {
	return m.Has(tag.PartyRoleQualifier)
}

// NoPartySubIDs is a repeating group element, Tag 802.
type NoPartySubIDs struct {
	*quickfix.Group
}

// SetPartySubID sets PartySubID, Tag 523.
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

// SetPartySubIDType sets PartySubIDType, Tag 803.
func (m NoPartySubIDs) SetPartySubIDType(v enum.PartySubIDType) {
	m.Set(field.NewPartySubIDType(v))
}

// GetPartySubID gets PartySubID, Tag 523.
func (m NoPartySubIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartySubIDType gets PartySubIDType, Tag 803.
func (m NoPartySubIDs) GetPartySubIDType() (v enum.PartySubIDType, err quickfix.MessageRejectError) {
	var f field.PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartySubID returns true if PartySubID is present, Tag 523.
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

// HasPartySubIDType returns true if PartySubIDType is present, Tag 803.
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

// NoPartySubIDsRepeatingGroup is a repeating group, Tag 802.
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup.
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartySubIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PartySubID),
				quickfix.GroupElement(tag.PartySubIDType),
			},
		),
	}
}

// Add create and append a new NoPartySubIDs to this group.
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

// Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup.
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoPartyIDsRepeatingGroup is a repeating group, Tag 453.
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup.
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartyIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PartyID),
				quickfix.GroupElement(tag.PartyIDSource),
				quickfix.GroupElement(tag.PartyRole),
				NewNoPartySubIDsRepeatingGroup(),
				quickfix.GroupElement(tag.PartyRoleQualifier),
			},
		),
	}
}

// Add create and append a new NoPartyIDs to this group.
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

// Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup.
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}
