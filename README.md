= GMO

GMO is a Go client library for the GMO Payment Platform

== Usage

```go
Client = gmo.New(os.Getenv("SiteID"), os.Getenv("SitePass"), os.Getenv("ShopID"), os.Getenv("ShopPass"))

// member related APIs
Client.RegisterMember("123", "jinzhu")
Client.UpdateMember("123", "jinzhu new")
Client.DeleteMember("123")
Client.SearchMember("123")

// cards related APIs
Client.SaveCard(memberID, "4024007154567043", "1219", "jinzhu")
Client.SearchCard(memberID, cardSeq)
Client.DeleteCard(memberID, cardSeq)

// Transaction related APIs
Client.EntryTran(orderID, orderAmount, orderTax)
Client.ExecTranPaypal(AccessID, AccessPass, orderID, memberID, "0", "123")

// Paypal related APIs
Client.EntryTranPaypal(orderID, orderAmount, orderTax, "USD")
Client.ExecTranPaypal(AccessID, AccessPass, orderID, orderDescription, redirectURL)

// Other SiteAPIs
Client.HandleSiteRequest(action string, params *Params)

// Other Shop APIs
Client.HandleShopRequest(action string, params *Params)
```

== Author

ThePlant, MIT License
