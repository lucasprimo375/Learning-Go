// unit tests

package addresses // only tests can use a package with a different name than the folder

import (
	"testing"
	//"tests/addresses" // can write an alias before the package name: .; no need to mention the package when calling a function from it
)

type testScenario struct {
	receivedAddress string
	expectedAddress string
}

func TestTypeOfAddress(t *testing.T) {
	//t.Parallel() // all tests with this stamente in run in parallel

	testScenarios := []testScenario{
		{"Street ABC", "Street"},
		{"Avenue Paulista", "Avenue"},
		{"Road of the Fools", "Road"},
		{"Highway of Highwater", "Highway"},
		{"Pool Form", "Invalid Address"},
		{"", "Invalid Address"},
		{"HIGHWAY CHAGAS", "Highway"},
		{"Road 31", "Road"},
		{"Avenue Town", "Avenue"},
		{"Street Leeds", "Street"},
	}

	for _, scenario := range testScenarios {
		receivedAddressType := TypeOfAddress(scenario.receivedAddress)

		if receivedAddressType != scenario.expectedAddress {
			t.Errorf("Received type %s is different than expected type %s", receivedAddressType, scenario.expectedAddress)
		}
	}
}

func TestWhatever(t *testing.T) {
	//t.Parallel()

	if 1 > 2 {
		t.Error("Broken test")
	}
}
