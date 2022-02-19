// unit tests

package adresses

import "testing"

func TestTypeOfAddress(t *testing.T) {
	testAddress := "Street Paulista"

	expectedAddressType := "Street"

	receivedAddressType := TypeOfAddress(testAddress)

	if receivedAddressType != expectedAddressType {
		t.Errorf("Received address type is different than the expected. Expected %s, received %s", expectedAddressType, receivedAddressType)
	}
}
