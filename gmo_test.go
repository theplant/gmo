package gmo_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/theplant/gmo"
)

var Client *gmo.GMO

func init() {
	Client = gmo.New(os.Getenv("SiteID"), os.Getenv("SitePass"))
	// Client.Debug = true
}

func TestRegisterMember(t *testing.T) {
	userID := fmt.Sprintf("%v", time.Now().UnixNano())
	if _, err := Client.RegisterMember(userID, "jinzhu"); err == nil {
		if resp, err := Client.SearchMember(userID); err == nil {
			if resp.Get("MemberID") != userID || resp.Get("MemberName") != "jinzhu" {
				t.Errorf("User should be registered successfully")
			}
		}
	}

	if _, err := Client.UpdateMember(userID, "jinzhu-new"); err == nil {
		if resp, err := Client.SearchMember(userID); err == nil {
			if resp.Get("MemberID") != userID || resp.Get("MemberName") != "jinzhu-new" {
				t.Errorf("User should be updated successfully")
			}
		}
	}

	if _, err := Client.DeleteMember(userID); err == nil {
		if _, err := Client.SearchMember(userID); err == nil {
			t.Errorf("User should be deleted successfully")
		}
	}
}
