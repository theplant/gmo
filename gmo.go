package gmo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
)

const (
	TestEndpoint       = "https://pt01.mul-pay.jp"
	ProductionEndpoint = "https://p01.mul-pay.jp"
)

type GMO struct {
	Endpoint string
	Version  string
	SiteID   string
	SitePass string
	ShopID   string
	ShopPass string
	// Debug    bool
}

type Params map[string]string

func New(siteID, sitePass, shopID, shopPass, endPoint string) *GMO {
	return &GMO{
		Version:  "3",
		SiteID:   siteID,
		SitePass: sitePass,
		ShopID:   shopID,
		ShopPass: shopPass,
		Endpoint: endPoint,
	}
}

func (gmo *GMO) HandleSiteRequest(action string, params Params, output interface{}) error {
	values := url.Values{}
	values.Add("Version", gmo.Version)
	values.Add("SiteID", gmo.SiteID)
	values.Add("SitePass", gmo.SitePass)
	for key, value := range params {
		values[key] = []string{value}
	}
	return gmo.HandleRawRequest(action, values, output)
}

func (gmo *GMO) HandleShopRequest(action string, params Params, output interface{}) error {
	values := url.Values{}
	values.Add("Version", gmo.Version)
	values.Add("ShopID", gmo.ShopID)
	values.Add("ShopPass", gmo.ShopPass)
	for key, value := range params {
		values[key] = []string{value}
	}
	return gmo.HandleRawRequest(action, values, output)
}

func (gmo *GMO) HandleRawRequest(action string, params url.Values, output interface{}) (err error) {
	resp, err := http.PostForm(gmo.Endpoint+action, params)
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	results, err := url.ParseQuery(string(bytes))
	if err != nil {
		return
	} else if errc := results.Get("ErrCode"); errc != "" {
		err = fmt.Errorf("%v: %s", errc, results.Get("ErrInfo"))
		return
	}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	err = decoder.Decode(output, results)
	return
}
