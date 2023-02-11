package invoice

import "errors"

// Warehouse invoice

type WarehouseInvoice struct {
	Date        string
	InvoiceType InvoiceTypeName
	Approved    bool
	Products    []Product
}

type InvoiceTypeName string

const (
	PURCHASE InvoiceTypeName = "purchase"
	SALES    InvoiceTypeName = "sales"
)

type Product struct {
	Name     string
	Unit     int
	Price    float64
	Discount float64
}

func (wi WarehouseInvoice) RecordInvoice() (InvoiceData, error) {
	if wi.Date ==""{
		return InvoiceData{}, errors.New("invoice date is empty")
	}

	if wi.InvoiceType == "" || wi.InvoiceType != PURCHASE && wi.InvoiceType != SALES{
		return InvoiceData{}, errors.New("invoice type is invalid")
	}

	if len(wi.Products) ==0 {
		return InvoiceData{}, errors.New("unit product is not valid")
	}

	var total float64
	for _, v := range wi.Products {
		if v.Unit <= 0  {
			return InvoiceData{}, errors.New("unit product is not valid")
		}

		if v.Price <= 0  {
			return InvoiceData{}, errors.New("price product is not valid")
		}
		total += float64(v.Unit) * float64(v.Price) * (1-v.Discount)
	}

	out := InvoiceData{
		Date: wi.Date,
		TotalInvoice: float64(total),
		Departemen: Warehouse,
	}
	
	return out, nil // TODO: replace this
}
