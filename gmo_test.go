package gmo_test

import (
	"os"
)

var Client *GMO

func init() {
	Client = gmo.New(os.Getenv("SiteID"), os.Getenv("SitePass"))
}
