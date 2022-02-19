package adresses

import "strings"

// TypeOfAddress checks if address is valid and returns its type
func TypeOfAddress(address string) string {
	lowerCaseAddress := strings.ToLower(address)

	validTypes := []string{"street", "avenue", "road", "highway"}

	adddressFirstWord := strings.Split(lowerCaseAddress, " ")[0]

	isAddressValid := false

	for _, addressType := range validTypes {
		if addressType == adddressFirstWord {
			isAddressValid = true
		}
	}

	if isAddressValid {
		return strings.Title(adddressFirstWord)
	}

	return "Invalid Address"
}
