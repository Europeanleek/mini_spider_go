package loader

import (
	"fmt"
	"testing"
)

func TestSeedLoad(t *testing.T) {
	TestCase := []struct {
		seedPath string
		err      error
	}{
		{
			seedPath: "./url_test_1.data",
			err:      nil,
		},
		{
			seedPath: "/url_test_2.data",
			err:      fmt.Errorf("error"),
		},
	}
	for index, cases := range TestCase {
		_, error := SeedLoad(cases.seedPath)
		if cases.err == nil {
			if error != nil {
				t.Errorf("TestSeedLoad[%d] failed, %s\n", index, error.Error())
			}
		} else {
			if error == nil {
				t.Errorf("TestSeedLoad[%d] failed, %s\n", index, error.Error())
			}
		}
	}
}
