package models

// func TestParseOperationType(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected error
// 	}{
// 		{"credit", nil},
// 		{"debit", nil},
// 		{"pix", nil},
// 		{"invalid", errors.New("invalid operation type")},
// 	}

// 	for _, test := range tests {
// 		err := ParseOperationType(test.input)

// 		if (err == nil && test.expected != nil) || (err != nil && test.expected == nil) {
// 			t.Errorf("Expected error: %v, Received error: %v", test.expected, err)
// 		} else if err != nil && test.expected != nil && err.Error() != test.expected.Error() {
// 			t.Errorf("Expected error message: %s, Received error message: %s", test.expected.Error(), err.Error())
// 		}
// 	}
// }

// func TestParseOrigin(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected error
// 	}{
// 		{"desktop-web", nil},
// 		{"mobile-android", nil},
// 		{"mobile-ios", nil},
// 		{"invalid", errors.New("invalid origin")},
// 	}

// 	for _, test := range tests {
// 		err := ParseOrigin(test.input)

// 		if (err == nil && test.expected != nil) || (err != nil && test.expected == nil) {
// 			t.Errorf("Expected error: %v, Received error: %v", test.expected, err)
// 		} else if err != nil && test.expected != nil && err.Error() != test.expected.Error() {
// 			t.Errorf("Expected error message: %s, Received error message: %s", test.expected.Error(), err.Error())
// 		}
// 	}
// }
