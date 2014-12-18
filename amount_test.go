package money

import (
	"testing"
)

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
