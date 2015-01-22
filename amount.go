package money

import (
	"encoding/xml"
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
	if err := e.EncodeElement(amount.String(), elem); err != nil {
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
	return []byte(amount.String()), nil
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

func (amount *Amount) String() string {
	// stay close to how Go prints float64(0.00)
	if int(*amount) == 0 {
		return "0"
	}
	amountStr := strconv.Itoa(int(*amount))
	length := len(amountStr)
	// Add left zeros to have format "0.01" and "0.12"
	for i := 0; i < 3-length; i += 1 {
		amountStr = "0" + amountStr
	}
	length = len(amountStr)
	decimals := amountStr[length-2:]
	integers := amountStr[:length-2]
	return integers + "." + decimals
}
