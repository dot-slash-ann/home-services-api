package transactions

import (
	"encoding/json"
	"time"
)

type UpdateTransactionDto struct {
	TransactionOn time.Time `json:"transaction_on"`
	PostedOn      time.Time `json:"posted_on"`
	Amount        uint      `json:"amount"`
	CategoryId    uint      `json:"category_id"`
}

func (t *UpdateTransactionDto) UnmarshalJSON(data []byte) error {
	type Alias UpdateTransactionDto
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
