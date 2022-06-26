package datamodel

import (
	"github.com/mt1976/mwt-go-dev/core"
	"github.com/mt1976/mwt-go-dev/logs"
)

type Lookup_Item struct {
	ID   string
	Name string
}

type FieldProperties struct {
	MsgType         string
	MsgMessage      string
	MsgLog          bool
	MsgGlyph        string
	MsgFeedBackType string
	InputType       string
	InputMessage    string
	InputLog        bool
	InputGlyph      string
	DefaultValue    string
}

func NewFieldProperties() *FieldProperties {
	return &FieldProperties{}
}

func AddFieldMessage(fP FieldProperties, typeStr string, message string, log bool, glyph string) FieldProperties {
	fP.MsgType = typeStr
	fP.MsgMessage = message
	fP.MsgLog = log
	fP.MsgGlyph = glyph
	fP.MsgFeedBackType = "form-helper"
	if fP.MsgType == core.FieldMessage_POSITIVE {
		fP.MsgFeedBackType = "valid-feedback"

	}
	if fP.MsgType == core.FieldMessage_ERROR {
		fP.MsgFeedBackType = "invalid-feedback"
	}
	if log || fP.MsgType == core.FieldMessage_ERROR {
		logs.Warning(message)
	}
	return fP
}

type PageProperties struct {
	Title    string
	Glyph    string
	MsgType  string
	MsgClass string
	MsgTitle string
	MsgBody  string
	MsgLog   bool
	MsgGlyph string
}

func NewPageProperties() *PageProperties {
	return &PageProperties{}
}

func AddPageMessage(typeStr string, class string, title string, body string, log bool, glyph string) *PageProperties {
	return &PageProperties{MsgType: typeStr, MsgClass: class, MsgTitle: title, MsgBody: body, MsgLog: log, MsgGlyph: glyph}
}
