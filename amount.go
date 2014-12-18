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
			a, err := GetAmountFromBytes(element)
			if err != nil {
				return err
			}
			*amount = *a
		}
	}
	return nil
}

func Parse(bytes []byte) (*Amount, error) {
	amountStr := string(bytes)
	index := strings.Index(amountStr, ".")
	pow := 2
	amountStrLen := len(amountStr)
	if index != -1 {
		amountStr = strings.Replace(amountStr, ".", "", 1)
		pow = amountStrLen - index - 3
	}

	a, err := strconv.Atoi(amountStr)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	fmt.Println(pow)

	a = a * int(math.Pow(float64(10), float64(pow)))
	amount := (Amount)(a)
	return &amount, nil
}

func (amount Amount) MarshalJSON() ([]byte, error) {
	amountInt := int(amount)
	return []byte(fmt.Sprintf("%.2f", float64(amountInt)/100)), nil
}

func (amount *Amount) UnmarshalJSON(bytes []byte) error {
	a, err := GetAmountFromBytes(bytes)
	if err != nil {
		return err
	}
	*amount = *a
	return nil
}

func (amount *Amount) ToFloat() float64 {
	amountInt := int(*amount)
	return float64(amountInt) / 100
}
