package dataframe

import (
	"testing"
)

func TestFromCSV(t *testing.T) {
	expected := "   first_name  last_name  username\n0         Rob       Pike       rob\n1         Ken   Thompson       ken\n2      Robert  Griesemer       gri"

	df := FromCSV("dataframe_test/test.csv")

	if df.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, df.String())
	}
}

func TestParseSQL(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected ParseSQL to panic, but it did not")
		}
	}()

	panic("Test not implemented")
}

