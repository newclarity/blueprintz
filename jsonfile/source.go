package jsonfile

import "blueprintz/global"

type Sourcer interface {
	GetCustom() global.Urls
	GetCommercial() global.Urls
	GetOpenSource() global.Urls
}

type Source struct {
	Custom     global.Urls `json:"custom,omitempty"`
	Commercial global.Urls `json:"commercial,omitempty"`
	OpenSource global.Urls `json:"opensource,omitempty"`
}

func NewSourceFromSourcer(sourcer Sourcer) *Source {
	return &Source{
		Custom:     sourcer.GetCustom(),
		Commercial: sourcer.GetCommercial(),
		OpenSource: sourcer.GetOpenSource(),
	}
}
