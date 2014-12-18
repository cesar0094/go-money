package money

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

type Price struct {
	XMLName  struct{} `xml:"Price" json:"-"`
	Amount   Amount   `xml:"amount" json:"amount"`
	Currency string   `xml:"currency" json:"currency"`
}

func TestParse(t *testing.T) {
	assertEqual := func(val interface{}, exp interface{}) {
		if val != exp {
			t.Errorf("Expected %v, got %v.", exp, val)
		}
	}

	testStr := "12"
	testAmount, err := Parse(testStr)

	assertEqual(err, nil)
	assertEqual(testAmount, Amount(1200))

	testStr = "12.3"
	testAmount, err = Parse(testStr)

	assertEqual(err, nil)
	assertEqual(testAmount, Amount(1230))

	testStr = "12.34"
	testAmount, err = Parse(testStr)

	assertEqual(err, nil)
	assertEqual(testAmount, Amount(1234))

	testStr = "12.345"
	testAmount, err = Parse(testStr)

	assertEqual(err, nil)
	assertEqual(testAmount, Amount(1234))

	testStr = "12.3456"
	testAmount, err = Parse(testStr)

	assertEqual(err, nil)
	assertEqual(testAmount, Amount(1234))

	testStr = "12."
	testAmount, err = Parse(testStr)

	assertEqual(err, nil)
	assertEqual(testAmount, Amount(1200))

	testStr = ".34"
	testAmount, err = Parse(testStr)

	assertEqual(err, nil)
	assertEqual(testAmount, Amount(34))

	// Testing garbage

	testStr = "10.asd"
	testAmount, err = Parse(testStr)

	if err == nil {
		t.Errorf("'%s' was parsed correctly to: '%v'", testStr, testAmount)
	}
	assertEqual(testAmount, Amount(0))

	testStr = "asd.00"
	testAmount, err = Parse(testStr)

	if err == nil {
		t.Errorf("'%s' was parsed correctly to: '%v'", testStr, testAmount)
	}
	assertEqual(testAmount, Amount(0))

	testStr = "as.asd"
	testAmount, err = Parse(testStr)

	if err == nil {
		t.Errorf("'%s' was parsed correctly to: '%v'", testStr, testAmount)
	}
	assertEqual(testAmount, Amount(0))

	testStr = "."
	testAmount, err = Parse(testStr)

	if err == nil {
		t.Errorf("'%s' was parsed correctly to: '%v'", testStr, testAmount)
	}
	assertEqual(testAmount, Amount(0))

	testStr = ""
	testAmount, err = Parse(testStr)

	if err == nil {
		t.Errorf("'%s' was parsed correctly to: '%v'", testStr, testAmount)
	}
	assertEqual(testAmount, Amount(0))

	testStr = ".a"
	testAmount, err = Parse(testStr)

	if err == nil {
		t.Errorf("'%s' was parsed correctly to: '%v'", testStr, testAmount)
	}
	assertEqual(testAmount, Amount(0))

}

func TestXml(t *testing.T) {
	assertEqual := func(val interface{}, exp interface{}) {
		if val != exp {
			t.Errorf("Expected %v, got %v.", exp, val)
		}
	}
	xmlStr := "<Price><amount>12.34</amount><currency>USD</currency></Price>"
	price := Price{
		Amount:   1234,
		Currency: "USD",
	}
	xmlBytes, err := xml.Marshal(price)
	assertEqual(err, nil)
	assertEqual(string(xmlBytes), xmlStr)

	price2 := Price{}
	err = xml.Unmarshal([]byte(xmlStr), &price2)
	assertEqual(err, nil)
	assertEqual(price2, price)
}

func TestJson(t *testing.T) {
	assertEqual := func(val interface{}, exp interface{}) {
		if val != exp {
			t.Errorf("Expected %v, got %v.", exp, val)
		}
	}
	jsonStr := `{"amount":12.34,"currency":"USD"}`
	price := Price{
		Amount:   1234,
		Currency: "USD",
	}
	jsonBytes, err := json.Marshal(price)
	assertEqual(err, nil)
	assertEqual(string(jsonBytes), jsonStr)

	price2 := Price{}
	err = json.Unmarshal([]byte(jsonStr), &price2)
	assertEqual(err, nil)
	assertEqual(price2, price)
}

func TestFloat(t *testing.T) {
	assertEqual := func(val interface{}, exp interface{}) {
		if val != exp {
			t.Errorf("Expected %v, got %v.", exp, val)
		}
	}

	testAmount := Amount(1234)
	testFloat := testAmount.Float()
	assertEqual(testFloat, 12.34)

	testAmount = Amount(34)
	testFloat = testAmount.Float()
	assertEqual(testFloat, 0.34)

	testAmount = Amount(1200)
	testFloat = testAmount.Float()
	assertEqual(testFloat, 12.00)

}

func TestInt(t *testing.T) {
	assertEqual := func(val interface{}, exp interface{}) {
		if val != exp {
			t.Errorf("Expected %v, got %v.", exp, val)
		}
	}

	testAmount := Amount(1234)
	testInt := testAmount.Int()
	assertEqual(testInt, 12)

	testAmount = Amount(34)
	testInt = testAmount.Int()
	assertEqual(testInt, 0)

	testAmount = Amount(1200)
	testInt = testAmount.Int()
	assertEqual(testInt, 12)
}
