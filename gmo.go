package gmo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type GMO struct {
	Endpoint string
	Version  string
	SiteID   string
	SitePass string
	Debug    bool
}

type Params map[string]string

func New(siteID string, sitePass string) *GMO {
	return &GMO{Version: "3", SiteID: siteID, SitePass: sitePass, Endpoint: "https://pt01.mul-pay.jp", Debug: false}
}

func (gmo *GMO) HandleRequest(action string, params *Params) (url.Values, error) {
	var (
		resp            *http.Response
		err             error
		results, values = url.Values{}, url.Values{}
	)

	defer func() {
		if err != nil && gmo.Debug {
			fmt.Printf("%v\t%v\nGot Error: %v\n\n", action, values, err.Error())
		}
	}()

	values.Add("Version", gmo.Version)
	values.Add("SiteID", gmo.SiteID)
	values.Add("SitePass", gmo.SitePass)

	for key, value := range *params {
		values[key] = []string{value}
	}

	if resp, err = http.PostForm(gmo.Endpoint+action, values); err == nil {
		var bytes []byte
		if bytes, err = ioutil.ReadAll(resp.Body); err == nil {
			if results, err = url.ParseQuery(string(bytes)); err == nil {
				if errStr := results.Get("ErrCode"); errStr == "" {
					return results, nil
				} else {
					err = fmt.Errorf("error code: %v", errStr)
				}
			}
		}
	}
	return results, err
}

func (gmo *GMO) RegisterMember(id, name string) (url.Values, error) {
	var params = Params{"MemberID": id, "MemberName": name}
	return gmo.HandleRequest("/payment/SaveMember.idPass", &params)
}

func (gmo *GMO) UpdateMember(id, name string) (url.Values, error) {
	var params = Params{"MemberID": id, "MemberName": name}
	return gmo.HandleRequest("/payment/UpdateMember.idPass", &params)
}

func (gmo *GMO) SearchMember(id string) (url.Values, error) {
	var params = Params{"MemberID": id}
	return gmo.HandleRequest("/payment/SearchMember.idPass", &params)
}

func (gmo *GMO) DeleteMember(id string) (url.Values, error) {
	var params = Params{"MemberID": id}
	return gmo.HandleRequest("/payment/DeleteMember.idPass", &params)
}

func (gmo *GMO) SaveCard(memberID, cardNo, expire, holderName string) (url.Values, error) {
	var params = Params{"MemberID": memberID, "CardNo": cardNo, "Expire": expire, "HolderName": holderName, "SeqMode": "1"}
	return gmo.HandleRequest("/payment/SaveCard.idPass", &params)
}

func (gmo *GMO) SearchCard(memberID, cardSeq string) (url.Values, error) {
	var params = Params{"MemberID": memberID, "CardSeq": cardSeq, "SeqMode": "1"}
	return gmo.HandleRequest("/payment/SearchCard.idPass", &params)
}

func (gmo *GMO) DeleteCard(memberID, cardSeq string) (url.Values, error) {
	var params = Params{"MemberID": memberID, "CardSeq": cardSeq, "SeqMode": "1"}
	return gmo.HandleRequest("/payment/DeleteCard.idPass", &params)
}

// /payment/EntryTranPaypal.idPass
