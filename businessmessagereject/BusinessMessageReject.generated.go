package businessmessagereject

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// BusinessMessageReject is the fix50sp2 BusinessMessageReject type, MsgType = j.
type BusinessMessageReject struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a BusinessMessageReject from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) BusinessMessageReject {
	return BusinessMessageReject{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m BusinessMessageReject) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a BusinessMessageReject initialized with the required fields for BusinessMessageReject.
func New(refmsgtype field.RefMsgTypeField, businessrejectreason field.BusinessRejectReasonField) (m BusinessMessageReject) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("j"))
	m.Set(refmsgtype)
	m.Set(businessrejectreason)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg BusinessMessageReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "j", r
}

// SetRefSeqNum sets RefSeqNum, Tag 45.
func (m BusinessMessageReject) SetRefSeqNum(v int) {
	m.Set(field.NewRefSeqNum(v))
}

// SetText sets Text, Tag 58.
func (m BusinessMessageReject) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetRefMsgType sets RefMsgType, Tag 372.
func (m BusinessMessageReject) SetRefMsgType(v string) {
	m.Set(field.NewRefMsgType(v))
}

// SetBusinessRejectReason sets BusinessRejectReason, Tag 380.
func (m BusinessMessageReject) SetBusinessRejectReason(v enum.BusinessRejectReason) {
	m.Set(field.NewBusinessRejectReason(v))
}

// GetRefSeqNum gets RefSeqNum, Tag 45.
func (m BusinessMessageReject) GetRefSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.RefSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58.
func (m BusinessMessageReject) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRefMsgType gets RefMsgType, Tag 372.
func (m BusinessMessageReject) GetRefMsgType() (v string, err quickfix.MessageRejectError) {
	var f field.RefMsgTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBusinessRejectReason gets BusinessRejectReason, Tag 380.
func (m BusinessMessageReject) GetBusinessRejectReason() (v enum.BusinessRejectReason, err quickfix.MessageRejectError) {
	var f field.BusinessRejectReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRefSeqNum returns true if RefSeqNum is present, Tag 45.
func (m BusinessMessageReject) HasRefSeqNum() bool {
	return m.Has(tag.RefSeqNum)
}

// HasText returns true if Text is present, Tag 58.
func (m BusinessMessageReject) HasText() bool {
	return m.Has(tag.Text)
}

// HasRefMsgType returns true if RefMsgType is present, Tag 372.
func (m BusinessMessageReject) HasRefMsgType() bool {
	return m.Has(tag.RefMsgType)
}

// HasBusinessRejectReason returns true if BusinessRejectReason is present, Tag 380.
func (m BusinessMessageReject) HasBusinessRejectReason() bool {
	return m.Has(tag.BusinessRejectReason)
}
