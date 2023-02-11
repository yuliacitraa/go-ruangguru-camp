package invoice

import "errors"

type InvoiceStatus string

const (
	PAID   InvoiceStatus = "paid"
	UNPAID InvoiceStatus = "unpaid"
)
// Finance invoice
type FinanceInvoice struct {
	Date     string
	Status   InvoiceStatus // status: "paid", "unpaid"
	Approved bool
	Details  []Detail
}

type Detail struct {
	Description string
	Total       int
}

func (fi FinanceInvoice) RecordInvoice() (InvoiceData, error) {
	if fi.Date ==""{
		return InvoiceData{}, errors.New("invoice date is empty")
	}

	if len(fi.Details) ==0{
		return InvoiceData{}, errors.New("invoice details is empty")
	}

	if fi.Status == "" || fi.Status != PAID && fi.Status != UNPAID {
		return InvoiceData{}, errors.New("invoice status is not valid")
	}

	total := 0
	for _, v := range fi.Details {
		if v.Total <= 0 {
			return InvoiceData{}, errors.New("total price is not valid")
		}
		total+=v.Total		
	}

	out := InvoiceData{
		Date: ChangeDate(fi.Date),
		TotalInvoice: float64(total),
		Departemen: Finance,
	}

	return out, nil // TODO: replace this
}




