package main

import (
	"fmt"
	"strings"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}


func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	if website.Domain == "" {
		err := fmt.Errorf("domain name is empty")
		chErr <- err
		ch <- RowData{}
		return
	}
	if website.Valid == false {
		err := fmt.Errorf("domain not valid")
		chErr <- err
		ch <- RowData{}
		return
	}
	if website.RefIPs == -1 {
		err := fmt.Errorf("domain RefIPs not valid")
		chErr <- err
		ch <- RowData{}
		return
	}

	tld := strings.Split(website.Domain, ".")
	
	if len(tld) == 3 {
		website.TLD = "." + tld[2]
		if tld[2] == "com" {
			website.IDN_TLD = ".co.id"
		}else if tld[2] == "org" {
			website.IDN_TLD = ".org.id"
		}else if tld[2] == "gov" {
			website.IDN_TLD = ".go.id"
		}else {
			website.IDN_TLD = website.TLD
		}
	}

	if len(tld) == 2 {
		website.TLD = "." +tld[1]
		if tld[1] == "com" {
			website.IDN_TLD = ".co.id"
		}else if tld[1] == "org" {
			website.IDN_TLD = ".org.id"
		}else if tld[1] == "gov" {
			website.IDN_TLD = ".go.id"
		} else {
			website.IDN_TLD = website.TLD
		}
	}

	ch <- website
	chErr <- nil

}

var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)
	out := []RowData{}

	for _, website := range data {
		go FuncProcessGetTLD(website, ch, errCh)

	}

	for i := 0; i < len(data); i++ {
		err := <-errCh
		c := <-ch
		if err !=nil {
			return nil, err
		} else if c.TLD == TLD {
			out = append(out, c)
		}
	}

	return out, nil
}

func main() {

}
