package marketdataincrementalrefresh

import (
	"github.com/shopspring/decimal"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// MarketDataIncrementalRefresh is the fix50sp2 MarketDataIncrementalRefresh type, MsgType = X.
type MarketDataIncrementalRefresh struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a MarketDataIncrementalRefresh from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) MarketDataIncrementalRefresh {
	return MarketDataIncrementalRefresh{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m MarketDataIncrementalRefresh) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a MarketDataIncrementalRefresh initialized with the required fields for MarketDataIncrementalRefresh.
func New() (m MarketDataIncrementalRefresh) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("X"))

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "X", r
}

// SetMDReqID sets MDReqID, Tag 262.
func (m MarketDataIncrementalRefresh) SetMDReqID(v string) {
	m.Set(field.NewMDReqID(v))
}

// SetNoMDEntries sets NoMDEntries, Tag 268.
func (m MarketDataIncrementalRefresh) SetNoMDEntries(f NoMDEntriesRepeatingGroup) {
	m.SetGroup(f)
}

// GetMDReqID gets MDReqID, Tag 262.
func (m MarketDataIncrementalRefresh) GetMDReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MDReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoMDEntries gets NoMDEntries, Tag 268.
func (m MarketDataIncrementalRefresh) GetNoMDEntries() (NoMDEntriesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMDEntriesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasMDReqID returns true if MDReqID is present, Tag 262.
func (m MarketDataIncrementalRefresh) HasMDReqID() bool {
	return m.Has(tag.MDReqID)
}

// HasNoMDEntries returns true if NoMDEntries is present, Tag 268.
func (m MarketDataIncrementalRefresh) HasNoMDEntries() bool {
	return m.Has(tag.NoMDEntries)
}

// NoMDEntries is a repeating group element, Tag 268.
type NoMDEntries struct {
	*quickfix.Group
}

// SetMDUpdateAction sets MDUpdateAction, Tag 279.
func (m NoMDEntries) SetMDUpdateAction(v enum.MDUpdateAction) {
	m.Set(field.NewMDUpdateAction(v))
}

// SetMarketDepth sets MarketDepth, Tag 264.
func (m NoMDEntries) SetMarketDepth(v int) {
	m.Set(field.NewMarketDepth(v))
}

// SetMDEntryType sets MDEntryType, Tag 269.
func (m NoMDEntries) SetMDEntryType(v enum.MDEntryType) {
	m.Set(field.NewMDEntryType(v))
}

// SetMDEntryID sets MDEntryID, Tag 278.
func (m NoMDEntries) SetMDEntryID(v string) {
	m.Set(field.NewMDEntryID(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m NoMDEntries) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m NoMDEntries) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22.
func (m NoMDEntries) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetCurrency sets Currency, Tag 15.
func (m NoMDEntries) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// SetMDEntryPx sets MDEntryPx, Tag 270.
func (m NoMDEntries) SetMDEntryPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewMDEntryPx(value, scale))
}

// SetOrdType sets OrdType, Tag 40.
func (m NoMDEntries) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

// SetMDEntrySize sets MDEntrySize, Tag 271.
func (m NoMDEntries) SetMDEntrySize(value decimal.Decimal, scale int32) {
	m.Set(field.NewMDEntrySize(value, scale))
}

// SetMDEntryDate sets MDEntryDate, Tag 272.
func (m NoMDEntries) SetMDEntryDate(v string) {
	m.Set(field.NewMDEntryDate(v))
}

// SetMDEntryTime sets MDEntryTime, Tag 273.
func (m NoMDEntries) SetMDEntryTime(v string) {
	m.Set(field.NewMDEntryTime(v))
}

// SetTradeCondition sets TradeCondition, Tag 277.
func (m NoMDEntries) SetTradeCondition(v enum.TradeCondition) {
	m.Set(field.NewTradeCondition(v))
}

// SetOpenCloseSettlFlag sets OpenCloseSettlFlag, Tag 286.
func (m NoMDEntries) SetOpenCloseSettlFlag(v enum.OpenCloseSettlFlag) {
	m.Set(field.NewOpenCloseSettlFlag(v))
}

// SetOrderID sets OrderID, Tag 37.
func (m NoMDEntries) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

// SetQuoteEntryID sets QuoteEntryID, Tag 299.
func (m NoMDEntries) SetQuoteEntryID(v string) {
	m.Set(field.NewQuoteEntryID(v))
}

// SetTradeID sets TradeID, Tag 1003.
func (m NoMDEntries) SetTradeID(v string) {
	m.Set(field.NewTradeID(v))
}

// SetText sets Text, Tag 58.
func (m NoMDEntries) SetText(v string) {
	m.Set(field.NewText(v))
}

// GetMDUpdateAction gets MDUpdateAction, Tag 279.
func (m NoMDEntries) GetMDUpdateAction() (v enum.MDUpdateAction, err quickfix.MessageRejectError) {
	var f field.MDUpdateActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketDepth gets MarketDepth, Tag 264.
func (m NoMDEntries) GetMarketDepth() (v int, err quickfix.MessageRejectError) {
	var f field.MarketDepthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryType gets MDEntryType, Tag 269.
func (m NoMDEntries) GetMDEntryType() (v enum.MDEntryType, err quickfix.MessageRejectError) {
	var f field.MDEntryTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryID gets MDEntryID, Tag 278.
func (m NoMDEntries) GetMDEntryID() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m NoMDEntries) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m NoMDEntries) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m NoMDEntries) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCurrency gets Currency, Tag 15.
func (m NoMDEntries) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryPx gets MDEntryPx, Tag 270.
func (m NoMDEntries) GetMDEntryPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MDEntryPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrdType gets OrdType, Tag 40.
func (m NoMDEntries) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntrySize gets MDEntrySize, Tag 271.
func (m NoMDEntries) GetMDEntrySize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MDEntrySizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryDate gets MDEntryDate, Tag 272.
func (m NoMDEntries) GetMDEntryDate() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryTime gets MDEntryTime, Tag 273.
func (m NoMDEntries) GetMDEntryTime() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeCondition gets TradeCondition, Tag 277.
func (m NoMDEntries) GetTradeCondition() (v enum.TradeCondition, err quickfix.MessageRejectError) {
	var f field.TradeConditionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOpenCloseSettlFlag gets OpenCloseSettlFlag, Tag 286.
func (m NoMDEntries) GetOpenCloseSettlFlag() (v enum.OpenCloseSettlFlag, err quickfix.MessageRejectError) {
	var f field.OpenCloseSettlFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderID gets OrderID, Tag 37.
func (m NoMDEntries) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteEntryID gets QuoteEntryID, Tag 299.
func (m NoMDEntries) GetQuoteEntryID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteEntryIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeID gets TradeID, Tag 1003.
func (m NoMDEntries) GetTradeID() (v string, err quickfix.MessageRejectError) {
	var f field.TradeIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58.
func (m NoMDEntries) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasMDUpdateAction returns true if MDUpdateAction is present, Tag 279.
func (m NoMDEntries) HasMDUpdateAction() bool {
	return m.Has(tag.MDUpdateAction)
}

// HasMarketDepth returns true if MarketDepth is present, Tag 264.
func (m NoMDEntries) HasMarketDepth() bool {
	return m.Has(tag.MarketDepth)
}

// HasMDEntryType returns true if MDEntryType is present, Tag 269.
func (m NoMDEntries) HasMDEntryType() bool {
	return m.Has(tag.MDEntryType)
}

// HasMDEntryID returns true if MDEntryID is present, Tag 278.
func (m NoMDEntries) HasMDEntryID() bool {
	return m.Has(tag.MDEntryID)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m NoMDEntries) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m NoMDEntries) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m NoMDEntries) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m NoMDEntries) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasMDEntryPx returns true if MDEntryPx is present, Tag 270.
func (m NoMDEntries) HasMDEntryPx() bool {
	return m.Has(tag.MDEntryPx)
}

// HasOrdType returns true if OrdType is present, Tag 40.
func (m NoMDEntries) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

// HasMDEntrySize returns true if MDEntrySize is present, Tag 271.
func (m NoMDEntries) HasMDEntrySize() bool {
	return m.Has(tag.MDEntrySize)
}

// HasMDEntryDate returns true if MDEntryDate is present, Tag 272.
func (m NoMDEntries) HasMDEntryDate() bool {
	return m.Has(tag.MDEntryDate)
}

// HasMDEntryTime returns true if MDEntryTime is present, Tag 273.
func (m NoMDEntries) HasMDEntryTime() bool {
	return m.Has(tag.MDEntryTime)
}

// HasTradeCondition returns true if TradeCondition is present, Tag 277.
func (m NoMDEntries) HasTradeCondition() bool {
	return m.Has(tag.TradeCondition)
}

// HasOpenCloseSettlFlag returns true if OpenCloseSettlFlag is present, Tag 286.
func (m NoMDEntries) HasOpenCloseSettlFlag() bool {
	return m.Has(tag.OpenCloseSettlFlag)
}

// HasOrderID returns true if OrderID is present, Tag 37.
func (m NoMDEntries) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

// HasQuoteEntryID returns true if QuoteEntryID is present, Tag 299.
func (m NoMDEntries) HasQuoteEntryID() bool {
	return m.Has(tag.QuoteEntryID)
}

// HasTradeID returns true if TradeID is present, Tag 1003.
func (m NoMDEntries) HasTradeID() bool {
	return m.Has(tag.TradeID)
}

// HasText returns true if Text is present, Tag 58.
func (m NoMDEntries) HasText() bool {
	return m.Has(tag.Text)
}

// NoMDEntriesRepeatingGroup is a repeating group, Tag 268.
type NoMDEntriesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoMDEntriesRepeatingGroup returns an initialized, NoMDEntriesRepeatingGroup.
func NewNoMDEntriesRepeatingGroup() NoMDEntriesRepeatingGroup {
	return NoMDEntriesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoMDEntries,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.MDUpdateAction),
				quickfix.GroupElement(tag.MarketDepth),
				quickfix.GroupElement(tag.MDEntryType),
				quickfix.GroupElement(tag.MDEntryID),
				quickfix.GroupElement(tag.Symbol),
				quickfix.GroupElement(tag.SecurityID),
				quickfix.GroupElement(tag.SecurityIDSource),
				quickfix.GroupElement(tag.Currency),
				quickfix.GroupElement(tag.MDEntryPx),
				quickfix.GroupElement(tag.OrdType),
				quickfix.GroupElement(tag.MDEntrySize),
				quickfix.GroupElement(tag.MDEntryDate),
				quickfix.GroupElement(tag.MDEntryTime),
				quickfix.GroupElement(tag.TradeCondition),
				quickfix.GroupElement(tag.OpenCloseSettlFlag),
				quickfix.GroupElement(tag.OrderID),
				quickfix.GroupElement(tag.QuoteEntryID),
				quickfix.GroupElement(tag.TradeID),
				quickfix.GroupElement(tag.Text),
			},
		),
	}
}

// Add create and append a new NoMDEntries to this group.
func (m NoMDEntriesRepeatingGroup) Add() NoMDEntries {
	g := m.RepeatingGroup.Add()
	return NoMDEntries{g}
}

// Get returns the ith NoMDEntries in the NoMDEntriesRepeatinGroup.
func (m NoMDEntriesRepeatingGroup) Get(i int) NoMDEntries {
	return NoMDEntries{m.RepeatingGroup.Get(i)}
}
