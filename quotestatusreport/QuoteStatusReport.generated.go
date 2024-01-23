package quotestatusreport

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// QuoteStatusReport is the fix50sp2 QuoteStatusReport type, MsgType = AI.
type QuoteStatusReport struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a QuoteStatusReport from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) QuoteStatusReport {
	return QuoteStatusReport{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m QuoteStatusReport) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a QuoteStatusReport initialized with the required fields for QuoteStatusReport.
func New() (m QuoteStatusReport) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("AI"))

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg QuoteStatusReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "AI", r
}

// SetAccount sets Account, Tag 1.
func (m QuoteStatusReport) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

// SetCurrency sets Currency, Tag 15.
func (m QuoteStatusReport) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22.
func (m QuoteStatusReport) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetOrderQty sets OrderQty, Tag 38.
func (m QuoteStatusReport) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

// SetOrdType sets OrdType, Tag 40.
func (m QuoteStatusReport) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

// SetPrice sets Price, Tag 44.
func (m QuoteStatusReport) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m QuoteStatusReport) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSide sets Side, Tag 54.
func (m QuoteStatusReport) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m QuoteStatusReport) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetText sets Text, Tag 58.
func (m QuoteStatusReport) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m QuoteStatusReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetQuoteID sets QuoteID, Tag 117.
func (m QuoteStatusReport) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

// SetBidPx sets BidPx, Tag 132.
func (m QuoteStatusReport) SetBidPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidPx(value, scale))
}

// SetOfferPx sets OfferPx, Tag 133.
func (m QuoteStatusReport) SetOfferPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferPx(value, scale))
}

// SetBidSize sets BidSize, Tag 134.
func (m QuoteStatusReport) SetBidSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidSize(value, scale))
}

// SetOfferSize sets OfferSize, Tag 135.
func (m QuoteStatusReport) SetOfferSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferSize(value, scale))
}

// SetQuoteStatus sets QuoteStatus, Tag 297.
func (m QuoteStatusReport) SetQuoteStatus(v enum.QuoteStatus) {
	m.Set(field.NewQuoteStatus(v))
}

// SetQuoteCancelType sets QuoteCancelType, Tag 298.
func (m QuoteStatusReport) SetQuoteCancelType(v enum.QuoteCancelType) {
	m.Set(field.NewQuoteCancelType(v))
}

// SetQuoteRejectReason sets QuoteRejectReason, Tag 300.
func (m QuoteStatusReport) SetQuoteRejectReason(v enum.QuoteRejectReason) {
	m.Set(field.NewQuoteRejectReason(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m QuoteStatusReport) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetQuoteMsgID sets QuoteMsgID, Tag 1166.
func (m QuoteStatusReport) SetQuoteMsgID(v string) {
	m.Set(field.NewQuoteMsgID(v))
}

// SetBidMDEntryID sets BidMDEntryID, Tag 1745.
func (m QuoteStatusReport) SetBidMDEntryID(v string) {
	m.Set(field.NewBidMDEntryID(v))
}

// SetOfferMDEntryID sets OfferMDEntryID, Tag 1746.
func (m QuoteStatusReport) SetOfferMDEntryID(v string) {
	m.Set(field.NewOfferMDEntryID(v))
}

// SetBidQuoteID sets BidQuoteID, Tag 1747.
func (m QuoteStatusReport) SetBidQuoteID(v string) {
	m.Set(field.NewBidQuoteID(v))
}

// SetOfferQuoteID sets OfferQuoteID, Tag 1748.
func (m QuoteStatusReport) SetOfferQuoteID(v string) {
	m.Set(field.NewOfferQuoteID(v))
}

// GetAccount gets Account, Tag 1.
func (m QuoteStatusReport) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCurrency gets Currency, Tag 15.
func (m QuoteStatusReport) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m QuoteStatusReport) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderQty gets OrderQty, Tag 38.
func (m QuoteStatusReport) GetOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrdType gets OrdType, Tag 40.
func (m QuoteStatusReport) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPrice gets Price, Tag 44.
func (m QuoteStatusReport) GetPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m QuoteStatusReport) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSide gets Side, Tag 54.
func (m QuoteStatusReport) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m QuoteStatusReport) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58.
func (m QuoteStatusReport) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m QuoteStatusReport) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteID gets QuoteID, Tag 117.
func (m QuoteStatusReport) GetQuoteID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBidPx gets BidPx, Tag 132.
func (m QuoteStatusReport) GetBidPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOfferPx gets OfferPx, Tag 133.
func (m QuoteStatusReport) GetOfferPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBidSize gets BidSize, Tag 134.
func (m QuoteStatusReport) GetBidSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOfferSize gets OfferSize, Tag 135.
func (m QuoteStatusReport) GetOfferSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteStatus gets QuoteStatus, Tag 297.
func (m QuoteStatusReport) GetQuoteStatus() (v enum.QuoteStatus, err quickfix.MessageRejectError) {
	var f field.QuoteStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteCancelType gets QuoteCancelType, Tag 298.
func (m QuoteStatusReport) GetQuoteCancelType() (v enum.QuoteCancelType, err quickfix.MessageRejectError) {
	var f field.QuoteCancelTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteRejectReason gets QuoteRejectReason, Tag 300.
func (m QuoteStatusReport) GetQuoteRejectReason() (v enum.QuoteRejectReason, err quickfix.MessageRejectError) {
	var f field.QuoteRejectReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m QuoteStatusReport) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetQuoteMsgID gets QuoteMsgID, Tag 1166.
func (m QuoteStatusReport) GetQuoteMsgID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteMsgIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBidMDEntryID gets BidMDEntryID, Tag 1745.
func (m QuoteStatusReport) GetBidMDEntryID() (v string, err quickfix.MessageRejectError) {
	var f field.BidMDEntryIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOfferMDEntryID gets OfferMDEntryID, Tag 1746.
func (m QuoteStatusReport) GetOfferMDEntryID() (v string, err quickfix.MessageRejectError) {
	var f field.OfferMDEntryIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBidQuoteID gets BidQuoteID, Tag 1747.
func (m QuoteStatusReport) GetBidQuoteID() (v string, err quickfix.MessageRejectError) {
	var f field.BidQuoteIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOfferQuoteID gets OfferQuoteID, Tag 1748.
func (m QuoteStatusReport) GetOfferQuoteID() (v string, err quickfix.MessageRejectError) {
	var f field.OfferQuoteIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAccount returns true if Account is present, Tag 1.
func (m QuoteStatusReport) HasAccount() bool {
	return m.Has(tag.Account)
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m QuoteStatusReport) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m QuoteStatusReport) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasOrderQty returns true if OrderQty is present, Tag 38.
func (m QuoteStatusReport) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

// HasOrdType returns true if OrdType is present, Tag 40.
func (m QuoteStatusReport) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

// HasPrice returns true if Price is present, Tag 44.
func (m QuoteStatusReport) HasPrice() bool {
	return m.Has(tag.Price)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m QuoteStatusReport) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSide returns true if Side is present, Tag 54.
func (m QuoteStatusReport) HasSide() bool {
	return m.Has(tag.Side)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m QuoteStatusReport) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasText returns true if Text is present, Tag 58.
func (m QuoteStatusReport) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m QuoteStatusReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasQuoteID returns true if QuoteID is present, Tag 117.
func (m QuoteStatusReport) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

// HasBidPx returns true if BidPx is present, Tag 132.
func (m QuoteStatusReport) HasBidPx() bool {
	return m.Has(tag.BidPx)
}

// HasOfferPx returns true if OfferPx is present, Tag 133.
func (m QuoteStatusReport) HasOfferPx() bool {
	return m.Has(tag.OfferPx)
}

// HasBidSize returns true if BidSize is present, Tag 134.
func (m QuoteStatusReport) HasBidSize() bool {
	return m.Has(tag.BidSize)
}

// HasOfferSize returns true if OfferSize is present, Tag 135.
func (m QuoteStatusReport) HasOfferSize() bool {
	return m.Has(tag.OfferSize)
}

// HasQuoteStatus returns true if QuoteStatus is present, Tag 297.
func (m QuoteStatusReport) HasQuoteStatus() bool {
	return m.Has(tag.QuoteStatus)
}

// HasQuoteCancelType returns true if QuoteCancelType is present, Tag 298.
func (m QuoteStatusReport) HasQuoteCancelType() bool {
	return m.Has(tag.QuoteCancelType)
}

// HasQuoteRejectReason returns true if QuoteRejectReason is present, Tag 300.
func (m QuoteStatusReport) HasQuoteRejectReason() bool {
	return m.Has(tag.QuoteRejectReason)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m QuoteStatusReport) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasQuoteMsgID returns true if QuoteMsgID is present, Tag 1166.
func (m QuoteStatusReport) HasQuoteMsgID() bool {
	return m.Has(tag.QuoteMsgID)
}

// HasBidMDEntryID returns true if BidMDEntryID is present, Tag 1745.
func (m QuoteStatusReport) HasBidMDEntryID() bool {
	return m.Has(tag.BidMDEntryID)
}

// HasOfferMDEntryID returns true if OfferMDEntryID is present, Tag 1746.
func (m QuoteStatusReport) HasOfferMDEntryID() bool {
	return m.Has(tag.OfferMDEntryID)
}

// HasBidQuoteID returns true if BidQuoteID is present, Tag 1747.
func (m QuoteStatusReport) HasBidQuoteID() bool {
	return m.Has(tag.BidQuoteID)
}

// HasOfferQuoteID returns true if OfferQuoteID is present, Tag 1748.
func (m QuoteStatusReport) HasOfferQuoteID() bool {
	return m.Has(tag.OfferQuoteID)
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
