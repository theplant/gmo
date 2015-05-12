package gmo

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type GMO struct {
	Endpoint string
	Version  string
	SiteID   string
	SitePass string
}

type Params map[string]string

func New(siteID string, sitePass string) *GMO {
	return &GMO{Version: "3", SiteID: siteID, SitePass: sitePass, Endpoint: "https://pt01.mul-pay.jp"}
}

func (gmo *GMO) HandleRequest(action string, params *Params) {
	values := url.Values{}
	values.Add("Version", gmo.Version)
	values.Add("SiteID", gmo.SiteID)
	values.Add("SitePass", gmo.SitePass)

	for key, value := range *params {
		values[key] = []string{value}
	}

	resp, err := http.PostForm(path.Join(gmo.Endpoint, action), values)
	fmt.Println(resp)
	fmt.Println(err)
}

func (gmo *GMO) RegisterMember(id, name string) {
	var params = Params{"MemberID": id, "MemberName": name}
	gmo.HandleRequest("/payment/SaveMember.idPass", &params)
}

func (gmo *GMO) UpdateMember(id, name string) {
	var params = Params{"MemberID": id, "MemberName": name}
	gmo.HandleRequest("/payment/UpdateMember.idPass", &params)
}

func (gmo *GMO) DeleteMember(id, name string) {
	var params = Params{"MemberID": id}
	gmo.HandleRequest("/payment/DeleteMember.idPass", &params)
}

func (gmo *GMO) SaveCard(cardno, expire, name string) {
	// /payment/SaveCard.idPass

	// CardSeq
}

func (gmo *GMO) DeleteCard(cardSeq string) {
	// /payment/DeleteCard.idPass
}

// /payment/EntryTranPaypal.idPass
