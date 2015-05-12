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
	ShopID   string
	ShopPass string
	Debug    bool
}

type Params map[string]string

func New(siteID, sitePass, shopID, shopPass string) *GMO {
	return &GMO{Version: "3", SiteID: siteID, SitePass: sitePass, ShopID: shopID, ShopPass: shopPass, Endpoint: "https://pt01.mul-pay.jp", Debug: false}
}

func (gmo *GMO) HandleSiteRequest(action string, params *Params) (url.Values, error) {
	values := url.Values{}
	values.Add("Version", gmo.Version)
	values.Add("SiteID", gmo.SiteID)
	values.Add("SitePass", gmo.SitePass)

	for key, value := range *params {
		values[key] = []string{value}
	}
	return gmo.HandleRawRequest(action, values)
}

func (gmo *GMO) HandleShopRequest(action string, params *Params) (url.Values, error) {
	values := url.Values{}
	values.Add("Version", gmo.Version)
	values.Add("ShopID", gmo.ShopID)
	values.Add("ShopPass", gmo.ShopPass)

	for key, value := range *params {
		values[key] = []string{value}
	}
	return gmo.HandleRawRequest(action, values)
}

func (gmo *GMO) HandleRawRequest(action string, values url.Values) (url.Values, error) {
	var (
		resp    *http.Response
		err     error
		results = url.Values{}
	)

	defer func() {
		if err != nil && gmo.Debug {
			fmt.Printf("%v\t%v\nGot Error: %v\n\n", action, values, err.Error())
		}
	}()

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
	return gmo.HandleSiteRequest("/payment/SaveMember.idPass", &params)
}

func (gmo *GMO) UpdateMember(id, name string) (url.Values, error) {
	var params = Params{"MemberID": id, "MemberName": name}
	return gmo.HandleSiteRequest("/payment/UpdateMember.idPass", &params)
}

func (gmo *GMO) SearchMember(id string) (url.Values, error) {
	var params = Params{"MemberID": id}
	return gmo.HandleSiteRequest("/payment/SearchMember.idPass", &params)
}

func (gmo *GMO) DeleteMember(id string) (url.Values, error) {
	var params = Params{"MemberID": id}
	return gmo.HandleSiteRequest("/payment/DeleteMember.idPass", &params)
}

func (gmo *GMO) SaveCard(memberID, cardNo, expire, holderName string) (url.Values, error) {
	var params = Params{"MemberID": memberID, "CardNo": cardNo, "Expire": expire, "HolderName": holderName, "SeqMode": "1"}
	return gmo.HandleSiteRequest("/payment/SaveCard.idPass", &params)
}

func (gmo *GMO) SearchCard(memberID, cardSeq string) (url.Values, error) {
	var params = Params{"MemberID": memberID, "CardSeq": cardSeq, "SeqMode": "1"}
	return gmo.HandleSiteRequest("/payment/SearchCard.idPass", &params)
}

func (gmo *GMO) DeleteCard(memberID, cardSeq string) (url.Values, error) {
	var params = Params{"MemberID": memberID, "CardSeq": cardSeq, "SeqMode": "1"}
	return gmo.HandleSiteRequest("/payment/DeleteCard.idPass", &params)
}

func (gmo *GMO) EntryTran(orderID, amount, tax string) (url.Values, error) {
	var params = Params{"OrderID": orderID, "JobCd": "CAPTURE", "Amount": amount, "Tax": tax}
	return gmo.HandleShopRequest("/payment/EntryTran.idPass", &params)
}

func (gmo *GMO) ExecTran(accessID, accessPass, orderID, memberID, cardSeq, securityCode string) (url.Values, error) {
	var params = Params{"AccessID": accessID, "AccessPass": accessPass, "OrderID": orderID, "MemberID": memberID, "CardSeq": cardSeq, "SecurityCode": securityCode, "SeqMode": "1"}
	return gmo.HandleSiteRequest("/payment/ExecTran.idPass", &params)
}
