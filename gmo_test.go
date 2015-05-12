package gmo_test

import (
	"os"
	"testing"

	"github.com/theplant/gmo"
)

var Client *gmo.GMO

func init() {
	Client = gmo.New(os.Getenv("SiteID"), os.Getenv("SitePass"))
}

func TestRegisterMember(t *testing.T) {
	Client.RegisterMember("12345", "jinzhu")
	Client.SearchMember("12345")
}
