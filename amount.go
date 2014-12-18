package money

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type Amount int

func (amount Amount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	elem := xml.StartElement{
		Name: xml.Name{Space: "", Local: start.Name.Local},
		Attr: []xml.Attr{},
	}
	if err := e.EncodeElement(float64(amount)/100, elem); err != nil {
		return err
	}
	return nil
}

func (amount *Amount) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err == io.EOF { // found end of element
			break
		}
		if err != nil {
			return err
		}
		switch element := token.(type) {
		case xml.CharData:
			a, err := Parse(string(element))
			if err != nil {
				return err
			}
			*amount = a
		}
	}
	return nil
}

func Parse(amountStr string) (Amount, error) {
	// Make sure it is a valid float first
	_, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return Amount(0), err
	}

	index := strings.Index(amountStr, ".")
	pow := 2
	amountStrLen := len(amountStr)
	if index != -1 {
		amountStr = strings.Replace(amountStr, ".", "", 1)
		index = amountStrLen - index
		pow = 3 - index
	}

	a, err := strconv.Atoi(amountStr)
	if err != nil {
		return Amount(0), err
	}

	a = int(float64(a) * (math.Pow(float64(10), float64(pow))))
	amount := (Amount)(a)
	return amount, nil
}

func (amount Amount) MarshalJSON() ([]byte, error) {
	amountInt := int(amount)
	return []byte(fmt.Sprintf("%.2f", float64(amountInt)/100)), nil
}

func (amount *Amount) UnmarshalJSON(bytes []byte) error {
	a, err := Parse(string(bytes))
	if err != nil {
		return err
	}
	*amount = a
	return nil
}

func (amount *Amount) Float() float64 {
	amountInt := int(*amount)
	return float64(amountInt) / 100
}

// NOTE: this function truncates decimals
func (amount *Amount) Int() int {
	amountInt := int(*amount)
	return amountInt / 100
}
