// unit tests

package adresses

import "testing"

type testScenario struct {
	receivedAddress string
	expectedAddress string
}

func TestTypeOfAddress(t *testing.T) {
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
