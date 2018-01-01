package models

import (
	"bytes"
	"strconv"
)

type Data struct {
	Proof        int64          `json:"proof-of-work"`
	Transactions []*Transaction `json:"transactions"`
}

// Returns a string representation of the data.
func (d *Data) String() string {
	var buf bytes.Buffer

	buf.WriteString(strconv.FormatInt(d.Proof, 10))
	for _, txn := range d.Transactions {
		buf.WriteString(txn.String())
	}

	return buf.String()
}
