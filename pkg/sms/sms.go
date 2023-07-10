package sms

import (
	"sync"

	"github.com/liqian-spec/practice/pkg/config"
)

type Message struct {
	Template string
	Data     map[string]string
	Content  string
}

type SMS struct {
	Driver Driver
}

var once sync.Once

var internalSMS *SMS

func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})

	return internalSMS
}

func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
