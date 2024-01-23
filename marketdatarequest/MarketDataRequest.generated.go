package marketdatarequest

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// MarketDataRequest is the fix50sp2 MarketDataRequest type, MsgType = V.
type MarketDataRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a MarketDataRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) MarketDataRequest {
	return MarketDataRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m MarketDataRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a MarketDataRequest initialized with the required fields for MarketDataRequest.
func New(mdreqid field.MDReqIDField, subscriptionrequesttype field.SubscriptionRequestTypeField, marketdepth field.MarketDepthField) (m MarketDataRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("V"))
	m.Set(mdreqid)
	m.Set(subscriptionrequesttype)
	m.Set(marketdepth)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg MarketDataRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "V", r
}

// SetNoRelatedSym sets NoRelatedSym, Tag 146.
func (m MarketDataRequest) SetNoRelatedSym(f NoRelatedSymRepeatingGroup) {
	m.SetGroup(f)
}

// SetMDReqID sets MDReqID, Tag 262.
func (m MarketDataRequest) SetMDReqID(v string) {
	m.Set(field.NewMDReqID(v))
}

// SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263.
func (m MarketDataRequest) SetSubscriptionRequestType(v enum.SubscriptionRequestType) {
	m.Set(field.NewSubscriptionRequestType(v))
}

// SetMarketDepth sets MarketDepth, Tag 264.
func (m MarketDataRequest) SetMarketDepth(v int) {
	m.Set(field.NewMarketDepth(v))
}

// SetMDUpdateType sets MDUpdateType, Tag 265.
func (m MarketDataRequest) SetMDUpdateType(v enum.MDUpdateType) {
	m.Set(field.NewMDUpdateType(v))
}

// SetNoMDEntryTypes sets NoMDEntryTypes, Tag 267.
func (m MarketDataRequest) SetNoMDEntryTypes(f NoMDEntryTypesRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m MarketDataRequest) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetNoRelatedSym gets NoRelatedSym, Tag 146.
func (m MarketDataRequest) GetNoRelatedSym() (NoRelatedSymRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedSymRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetMDReqID gets MDReqID, Tag 262.
func (m MarketDataRequest) GetMDReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MDReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263.
func (m MarketDataRequest) GetSubscriptionRequestType() (v enum.SubscriptionRequestType, err quickfix.MessageRejectError) {
	var f field.SubscriptionRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketDepth gets MarketDepth, Tag 264.
func (m MarketDataRequest) GetMarketDepth() (v int, err quickfix.MessageRejectError) {
	var f field.MarketDepthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDUpdateType gets MDUpdateType, Tag 265.
func (m MarketDataRequest) GetMDUpdateType() (v enum.MDUpdateType, err quickfix.MessageRejectError) {
	var f field.MDUpdateTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoMDEntryTypes gets NoMDEntryTypes, Tag 267.
func (m MarketDataRequest) GetNoMDEntryTypes() (NoMDEntryTypesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMDEntryTypesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m MarketDataRequest) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasNoRelatedSym returns true if NoRelatedSym is present, Tag 146.
func (m MarketDataRequest) HasNoRelatedSym() bool {
	return m.Has(tag.NoRelatedSym)
}

// HasMDReqID returns true if MDReqID is present, Tag 262.
func (m MarketDataRequest) HasMDReqID() bool {
	return m.Has(tag.MDReqID)
}

// HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263.
func (m MarketDataRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

// HasMarketDepth returns true if MarketDepth is present, Tag 264.
func (m MarketDataRequest) HasMarketDepth() bool {
	return m.Has(tag.MarketDepth)
}

// HasMDUpdateType returns true if MDUpdateType is present, Tag 265.
func (m MarketDataRequest) HasMDUpdateType() bool {
	return m.Has(tag.MDUpdateType)
}

// HasNoMDEntryTypes returns true if NoMDEntryTypes is present, Tag 267.
func (m MarketDataRequest) HasNoMDEntryTypes() bool {
	return m.Has(tag.NoMDEntryTypes)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m MarketDataRequest) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
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

// NoMDEntryTypes is a repeating group element, Tag 267.
type NoMDEntryTypes struct {
	*quickfix.Group
}

// SetMDEntryType sets MDEntryType, Tag 269.
func (m NoMDEntryTypes) SetMDEntryType(v enum.MDEntryType) {
	m.Set(field.NewMDEntryType(v))
}

// GetMDEntryType gets MDEntryType, Tag 269.
func (m NoMDEntryTypes) GetMDEntryType() (v enum.MDEntryType, err quickfix.MessageRejectError) {
	var f field.MDEntryTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasMDEntryType returns true if MDEntryType is present, Tag 269.
func (m NoMDEntryTypes) HasMDEntryType() bool {
	return m.Has(tag.MDEntryType)
}

// NoMDEntryTypesRepeatingGroup is a repeating group, Tag 267.
type NoMDEntryTypesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoMDEntryTypesRepeatingGroup returns an initialized, NoMDEntryTypesRepeatingGroup.
func NewNoMDEntryTypesRepeatingGroup() NoMDEntryTypesRepeatingGroup {
	return NoMDEntryTypesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoMDEntryTypes,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.MDEntryType),
			},
		),
	}
}

// Add create and append a new NoMDEntryTypes to this group.
func (m NoMDEntryTypesRepeatingGroup) Add() NoMDEntryTypes {
	g := m.RepeatingGroup.Add()
	return NoMDEntryTypes{g}
}

// Get returns the ith NoMDEntryTypes in the NoMDEntryTypesRepeatinGroup.
func (m NoMDEntryTypesRepeatingGroup) Get(i int) NoMDEntryTypes {
	return NoMDEntryTypes{m.RepeatingGroup.Get(i)}
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
