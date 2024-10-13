package transactions

import (
	"encoding/json"
	"time"
)

type CreateTransactionDto struct {
	TransactionOn time.Time `json:"transaction_on" binding:"required"`
	PostedOn      time.Time `json:"posted_on" binding:"required"`
	Amount        uint      `json:"amount" binding:"required"`
	CategoryName  string    `json:"category" binding:"required"`
	VendorName    string    `json:"vendor" binding:"required"`
}

func (t *CreateTransactionDto) UnmarshalJSON(data []byte) error {
	type Alias CreateTransactionDto
	aux := &struct {
		TransactionOn string `json:"transaction_on"`
		PostedOn      string `json:"posted_on"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	layout := "2006-01-02"
	var err error
	t.TransactionOn, err = time.Parse(layout, aux.TransactionOn)

	if err != nil {
		return err
	}

	t.PostedOn, err = time.Parse(layout, aux.PostedOn)

	if err != nil {
		return err
	}

	return nil
}
