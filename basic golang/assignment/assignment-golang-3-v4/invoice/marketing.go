package invoice

import (
	"errors"
	"strconv"
	"strings"
)

// Marketing invoice
type MarketingInvoice struct {
	Date        string
	StartDate   string
	EndDate     string
	PricePerDay int
	AnotherFee  int
	Approved    bool
}

func (mi MarketingInvoice) RecordInvoice() (InvoiceData, error) {
	if mi.Date ==""{
		return InvoiceData{}, errors.New("invoice date is empty")
	}

	if mi.StartDate == "" || mi.EndDate == "" {
		return InvoiceData{}, errors.New("travel date is empty")
	}

	if mi.PricePerDay <= 0 {
		return InvoiceData{}, errors.New("price per day is not valid")
	}
	
	var total float64
	splitStart := strings.Split(mi.StartDate, "/")
	splitEnd := strings.Split(mi.EndDate, "/")

	start, _ := strconv.ParseFloat(splitStart[0], 64)
	end, _ := strconv.ParseFloat(splitEnd[0], 64)

	total = (end - start+1) * float64(mi.PricePerDay) + float64(mi.AnotherFee)

	out := InvoiceData{
		Date: ChangeDate(mi.Date),
		TotalInvoice: float64(total),
		Departemen: Marketing,
	}

	return out, nil // TODO: replace this
}
