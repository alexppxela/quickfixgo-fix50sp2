package executionreport

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// ExecutionReport is the fix50sp2 ExecutionReport type, MsgType = 8.
type ExecutionReport struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a ExecutionReport from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) ExecutionReport {
	return ExecutionReport{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m ExecutionReport) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a ExecutionReport initialized with the required fields for ExecutionReport.
func New(orderid field.OrderIDField, execid field.ExecIDField, exectype field.ExecTypeField, ordstatus field.OrdStatusField, side field.SideField, leavesqty field.LeavesQtyField, cumqty field.CumQtyField) (m ExecutionReport) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("8"))
	m.Set(orderid)
	m.Set(execid)
	m.Set(exectype)
	m.Set(ordstatus)
	m.Set(side)
	m.Set(leavesqty)
	m.Set(cumqty)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg ExecutionReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "8", r
}

// SetAccount sets Account, Tag 1.
func (m ExecutionReport) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

// SetAvgPx sets AvgPx, Tag 6.
func (m ExecutionReport) SetAvgPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewAvgPx(value, scale))
}

// SetClOrdID sets ClOrdID, Tag 11.
func (m ExecutionReport) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

// SetCumQty sets CumQty, Tag 14.
func (m ExecutionReport) SetCumQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCumQty(value, scale))
}

// SetCurrency sets Currency, Tag 15.
func (m ExecutionReport) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// SetExecID sets ExecID, Tag 17.
func (m ExecutionReport) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22.
func (m ExecutionReport) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetLastPx sets LastPx, Tag 31.
func (m ExecutionReport) SetLastPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastPx(value, scale))
}

// SetLastQty sets LastQty, Tag 32.
func (m ExecutionReport) SetLastQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastQty(value, scale))
}

// SetOrderID sets OrderID, Tag 37.
func (m ExecutionReport) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

// SetOrderQty sets OrderQty, Tag 38.
func (m ExecutionReport) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

// SetOrdStatus sets OrdStatus, Tag 39.
func (m ExecutionReport) SetOrdStatus(v enum.OrdStatus) {
	m.Set(field.NewOrdStatus(v))
}

// SetOrdType sets OrdType, Tag 40.
func (m ExecutionReport) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

// SetOrigClOrdID sets OrigClOrdID, Tag 41.
func (m ExecutionReport) SetOrigClOrdID(v string) {
	m.Set(field.NewOrigClOrdID(v))
}

// SetPrice sets Price, Tag 44.
func (m ExecutionReport) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m ExecutionReport) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSide sets Side, Tag 54.
func (m ExecutionReport) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m ExecutionReport) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetText sets Text, Tag 58.
func (m ExecutionReport) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTimeInForce sets TimeInForce, Tag 59.
func (m ExecutionReport) SetTimeInForce(v enum.TimeInForce) {
	m.Set(field.NewTimeInForce(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m ExecutionReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetOrdRejReason sets OrdRejReason, Tag 103.
func (m ExecutionReport) SetOrdRejReason(v enum.OrdRejReason) {
	m.Set(field.NewOrdRejReason(v))
}

// SetExecType sets ExecType, Tag 150.
func (m ExecutionReport) SetExecType(v enum.ExecType) {
	m.Set(field.NewExecType(v))
}

// SetLeavesQty sets LeavesQty, Tag 151.
func (m ExecutionReport) SetLeavesQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLeavesQty(value, scale))
}

// SetSecondaryOrderID sets SecondaryOrderID, Tag 198.
func (m ExecutionReport) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m ExecutionReport) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetMassStatusReqID sets MassStatusReqID, Tag 584.
func (m ExecutionReport) SetMassStatusReqID(v string) {
	m.Set(field.NewMassStatusReqID(v))
}

// SetOrdStatusReqID sets OrdStatusReqID, Tag 790.
func (m ExecutionReport) SetOrdStatusReqID(v string) {
	m.Set(field.NewOrdStatusReqID(v))
}

// SetTrdMatchID sets TrdMatchID, Tag 880.
func (m ExecutionReport) SetTrdMatchID(v string) {
	m.Set(field.NewTrdMatchID(v))
}

// SetTotNumReports sets TotNumReports, Tag 911.
func (m ExecutionReport) SetTotNumReports(v int) {
	m.Set(field.NewTotNumReports(v))
}

// SetLastRptRequested sets LastRptRequested, Tag 912.
func (m ExecutionReport) SetLastRptRequested(v bool) {
	m.Set(field.NewLastRptRequested(v))
}

// SetQuoteMsgID sets QuoteMsgID, Tag 1166.
func (m ExecutionReport) SetQuoteMsgID(v string) {
	m.Set(field.NewQuoteMsgID(v))
}

// SetOrderOrigination sets OrderOrigination, Tag 1724.
func (m ExecutionReport) SetOrderOrigination(v enum.OrderOrigination) {
	m.Set(field.NewOrderOrigination(v))
}

// GetAccount gets Account, Tag 1.
func (m ExecutionReport) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAvgPx gets AvgPx, Tag 6.
func (m ExecutionReport) GetAvgPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AvgPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetClOrdID gets ClOrdID, Tag 11.
func (m ExecutionReport) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCumQty gets CumQty, Tag 14.
func (m ExecutionReport) GetCumQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CumQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCurrency gets Currency, Tag 15.
func (m ExecutionReport) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExecID gets ExecID, Tag 17.
func (m ExecutionReport) GetExecID() (v string, err quickfix.MessageRejectError) {
	var f field.ExecIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m ExecutionReport) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLastPx gets LastPx, Tag 31.
func (m ExecutionReport) GetLastPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLastQty gets LastQty, Tag 32.
func (m ExecutionReport) GetLastQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderID gets OrderID, Tag 37.
func (m ExecutionReport) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderQty gets OrderQty, Tag 38.
func (m ExecutionReport) GetOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrdStatus gets OrdStatus, Tag 39.
func (m ExecutionReport) GetOrdStatus() (v enum.OrdStatus, err quickfix.MessageRejectError) {
	var f field.OrdStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrdType gets OrdType, Tag 40.
func (m ExecutionReport) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrigClOrdID gets OrigClOrdID, Tag 41.
func (m ExecutionReport) GetOrigClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.OrigClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPrice gets Price, Tag 44.
func (m ExecutionReport) GetPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m ExecutionReport) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSide gets Side, Tag 54.
func (m ExecutionReport) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m ExecutionReport) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58.
func (m ExecutionReport) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTimeInForce gets TimeInForce, Tag 59.
func (m ExecutionReport) GetTimeInForce() (v enum.TimeInForce, err quickfix.MessageRejectError) {
	var f field.TimeInForceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m ExecutionReport) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrdRejReason gets OrdRejReason, Tag 103.
func (m ExecutionReport) GetOrdRejReason() (v enum.OrdRejReason, err quickfix.MessageRejectError) {
	var f field.OrdRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExecType gets ExecType, Tag 150.
func (m ExecutionReport) GetExecType() (v enum.ExecType, err quickfix.MessageRejectError) {
	var f field.ExecTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLeavesQty gets LeavesQty, Tag 151.
func (m ExecutionReport) GetLeavesQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LeavesQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecondaryOrderID gets SecondaryOrderID, Tag 198.
func (m ExecutionReport) GetSecondaryOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryOrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m ExecutionReport) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetMassStatusReqID gets MassStatusReqID, Tag 584.
func (m ExecutionReport) GetMassStatusReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MassStatusReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrdStatusReqID gets OrdStatusReqID, Tag 790.
func (m ExecutionReport) GetOrdStatusReqID() (v string, err quickfix.MessageRejectError) {
	var f field.OrdStatusReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTrdMatchID gets TrdMatchID, Tag 880.
func (m ExecutionReport) GetTrdMatchID() (v string, err quickfix.MessageRejectError) {
	var f field.TrdMatchIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotNumReports gets TotNumReports, Tag 911.
func (m ExecutionReport) GetTotNumReports() (v int, err quickfix.MessageRejectError) {
	var f field.TotNumReportsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLastRptRequested gets LastRptRequested, Tag 912.
func (m ExecutionReport) GetLastRptRequested() (v bool, err quickfix.MessageRejectError) {
	var f field.LastRptRequestedField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteMsgID gets QuoteMsgID, Tag 1166.
func (m ExecutionReport) GetQuoteMsgID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteMsgIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderOrigination gets OrderOrigination, Tag 1724.
func (m ExecutionReport) GetOrderOrigination() (v enum.OrderOrigination, err quickfix.MessageRejectError) {
	var f field.OrderOriginationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAccount returns true if Account is present, Tag 1.
func (m ExecutionReport) HasAccount() bool {
	return m.Has(tag.Account)
}

// HasAvgPx returns true if AvgPx is present, Tag 6.
func (m ExecutionReport) HasAvgPx() bool {
	return m.Has(tag.AvgPx)
}

// HasClOrdID returns true if ClOrdID is present, Tag 11.
func (m ExecutionReport) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

// HasCumQty returns true if CumQty is present, Tag 14.
func (m ExecutionReport) HasCumQty() bool {
	return m.Has(tag.CumQty)
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m ExecutionReport) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasExecID returns true if ExecID is present, Tag 17.
func (m ExecutionReport) HasExecID() bool {
	return m.Has(tag.ExecID)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m ExecutionReport) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasLastPx returns true if LastPx is present, Tag 31.
func (m ExecutionReport) HasLastPx() bool {
	return m.Has(tag.LastPx)
}

// HasLastQty returns true if LastQty is present, Tag 32.
func (m ExecutionReport) HasLastQty() bool {
	return m.Has(tag.LastQty)
}

// HasOrderID returns true if OrderID is present, Tag 37.
func (m ExecutionReport) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

// HasOrderQty returns true if OrderQty is present, Tag 38.
func (m ExecutionReport) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

// HasOrdStatus returns true if OrdStatus is present, Tag 39.
func (m ExecutionReport) HasOrdStatus() bool {
	return m.Has(tag.OrdStatus)
}

// HasOrdType returns true if OrdType is present, Tag 40.
func (m ExecutionReport) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

// HasOrigClOrdID returns true if OrigClOrdID is present, Tag 41.
func (m ExecutionReport) HasOrigClOrdID() bool {
	return m.Has(tag.OrigClOrdID)
}

// HasPrice returns true if Price is present, Tag 44.
func (m ExecutionReport) HasPrice() bool {
	return m.Has(tag.Price)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m ExecutionReport) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSide returns true if Side is present, Tag 54.
func (m ExecutionReport) HasSide() bool {
	return m.Has(tag.Side)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m ExecutionReport) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasText returns true if Text is present, Tag 58.
func (m ExecutionReport) HasText() bool {
	return m.Has(tag.Text)
}

// HasTimeInForce returns true if TimeInForce is present, Tag 59.
func (m ExecutionReport) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m ExecutionReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasOrdRejReason returns true if OrdRejReason is present, Tag 103.
func (m ExecutionReport) HasOrdRejReason() bool {
	return m.Has(tag.OrdRejReason)
}

// HasExecType returns true if ExecType is present, Tag 150.
func (m ExecutionReport) HasExecType() bool {
	return m.Has(tag.ExecType)
}

// HasLeavesQty returns true if LeavesQty is present, Tag 151.
func (m ExecutionReport) HasLeavesQty() bool {
	return m.Has(tag.LeavesQty)
}

// HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198.
func (m ExecutionReport) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m ExecutionReport) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasMassStatusReqID returns true if MassStatusReqID is present, Tag 584.
func (m ExecutionReport) HasMassStatusReqID() bool {
	return m.Has(tag.MassStatusReqID)
}

// HasOrdStatusReqID returns true if OrdStatusReqID is present, Tag 790.
func (m ExecutionReport) HasOrdStatusReqID() bool {
	return m.Has(tag.OrdStatusReqID)
}

// HasTrdMatchID returns true if TrdMatchID is present, Tag 880.
func (m ExecutionReport) HasTrdMatchID() bool {
	return m.Has(tag.TrdMatchID)
}

// HasTotNumReports returns true if TotNumReports is present, Tag 911.
func (m ExecutionReport) HasTotNumReports() bool {
	return m.Has(tag.TotNumReports)
}

// HasLastRptRequested returns true if LastRptRequested is present, Tag 912.
func (m ExecutionReport) HasLastRptRequested() bool {
	return m.Has(tag.LastRptRequested)
}

// HasQuoteMsgID returns true if QuoteMsgID is present, Tag 1166.
func (m ExecutionReport) HasQuoteMsgID() bool {
	return m.Has(tag.QuoteMsgID)
}

// HasOrderOrigination returns true if OrderOrigination is present, Tag 1724.
func (m ExecutionReport) HasOrderOrigination() bool {
	return m.Has(tag.OrderOrigination)
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
