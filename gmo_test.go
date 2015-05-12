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
	Client = gmo.New(os.Getenv("SiteID"), os.Getenv("SitePass"), os.Getenv("ShopID"), os.Getenv("ShopPass"))
	Client.Debug = true
}

func checkErr(err error, t *testing.T) error {
	if err != nil {
		t.Error(err.Error())
	}
	return err
}
func TestMemberAPIs(t *testing.T) {
	userID := fmt.Sprintf("%v", time.Now().UnixNano())
	if _, err := Client.RegisterMember(userID, "jinzhu"); checkErr(err, t) == nil {
		if resp, err := Client.SearchMember(userID); checkErr(err, t) == nil {
			if resp.Get("MemberID") != userID || resp.Get("MemberName") != "jinzhu" {
				t.Errorf("User should be registered successfully")
			}
		}
	}

	if _, err := Client.UpdateMember(userID, "jinzhu-new"); checkErr(err, t) == nil {
		if resp, err := Client.SearchMember(userID); checkErr(err, t) == nil {
			if resp.Get("MemberID") != userID || resp.Get("MemberName") != "jinzhu-new" {
				t.Errorf("User should be updated successfully")
			}
		}
	}

	if _, err := Client.DeleteMember(userID); checkErr(err, t) == nil {
		if _, err := Client.SearchMember(userID); err == nil {
			t.Errorf("User should be deleted successfully")
		}
	}
}

func TestCardAPIs(t *testing.T) {
	userID := fmt.Sprintf("%v", time.Now().UnixNano())
	Client.RegisterMember(userID, "jinzhu")

	if resp, err := Client.SaveCard(userID, "4024007154567043", "1219", "jinzhu"); checkErr(err, t) == nil {
		seq := resp.Get("CardSeq")
		if _, err := Client.SearchCard(userID, seq); checkErr(err, t) != nil {
			t.Errorf("Card should be created successfully")
		}

		if _, err := Client.SaveCard(userID, "4024007154567043", "1219", "jinzhu new"); checkErr(err, t) == nil {
			if resp, err := Client.SearchCard(userID, seq); checkErr(err, t) != nil || resp.Get("HolderName") != "jinzhu new" {
				t.Errorf("Card should be updated successfully")
			}
		}

		if _, err := Client.DeleteCard(userID, seq); checkErr(err, t) == nil {
			if resp, err := Client.SearchCard(userID, seq); checkErr(err, t) != nil || resp.Get("DeleteFlag") != "1" {
				t.Errorf("Card should be deleted successfully")
			}
		}
	}
}

func TestCreateOrderWithSavedCard(t *testing.T) {
	userID := fmt.Sprintf("%v", time.Now().UnixNano())
	Client.RegisterMember(userID, "jinzhu")
	Client.SaveCard(userID, "4111111111111111", "1219", "jinzhu")
	orderID := "111111"

	if results, err := Client.EntryTran(orderID, "1000", "100"); checkErr(err, t) == nil {
		fmt.Println(Client.ExecTran(results.Get("AccessID"), results.Get("AccessPass"), orderID, userID, "0", "123"))
	} else {
		fmt.Println(results)
		t.Errorf("No error should happen when register order to GMO")
	}
}
