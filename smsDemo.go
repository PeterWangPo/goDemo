package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Context struct {
	Tel string
	Txt string
}

type providerType string

const (
	TencentProvider providerType = "tencentSms"
	AliProvider     providerType = "aliSms"
	BaiduProvider   providerType = "baiduSms"
)

var providers = [...]providerType{TencentProvider, AliProvider, BaiduProvider}

type Msg interface {
	Send(ctx *Context) error
	getProvider() providerType
}

type TencentSms struct{}

func (t *TencentSms) Send(ctx *Context) error {
	s := fmt.Sprintf("provider: %s, sms. tel:%s, txt:%s", t.getProvider(), ctx.Tel, ctx.Txt)
	fmt.Println(s)
	return nil
}

func (t *TencentSms) getProvider() providerType {
	return TencentProvider
}

type AliSms struct{}

func (t *AliSms) Send(ctx *Context) error {
	s := fmt.Sprintf("provider: %s, sms. tel:%s, txt:%s", t.getProvider(), ctx.Tel, ctx.Txt)
	fmt.Println(s)
	return nil
}

func (t *AliSms) getProvider() providerType {
	return AliProvider
}

type BaiduSms struct{}

func (t *BaiduSms) Send(ctx *Context) error {
	s := fmt.Sprintf("provider: %s, sms. tel:%s, txt:%s", t.getProvider(), ctx.Tel, ctx.Txt)
	fmt.Println(s)
	return nil
}

func (t *BaiduSms) getProvider() providerType {
	return BaiduProvider
}

type ManageInstance struct {
	T               time.Duration
	currentProvider Msg
}

func (m *ManageInstance) InitStatus() {
	m.refershStatus()
	go func() {
		for {
			select {
			case <-time.NewTimer(m.T).C:
				m.refershStatus()
			}
		}
	}()
}

func (m *ManageInstance) GetCurrentProvider() Msg {
	return m.currentProvider
}

func (m *ManageInstance) refershStatus() error {
	r := rand.New(rand.NewSource(time.Now().Unix())).Intn(len(providers))
	p := providers[r]
	switch p {
	case TencentProvider:
		m.currentProvider = &TencentSms{}
	case AliProvider:
		m.currentProvider = &AliSms{}
	case BaiduProvider:
		m.currentProvider = &BaiduSms{}
	default:
		panic(fmt.Sprintf("Unsupported provider : %s", p))
	}
	return nil
}

func main() {
	m := &ManageInstance{
		T: time.Millisecond * 300,
	}
	m.InitStatus()
	for i := 0; i < 30; i++ {
		m.GetCurrentProvider().Send(&Context{
			Tel: "18611454187",
			Txt: "this is phone",
		})
		time.Sleep(500 * time.Millisecond)
	}
}
