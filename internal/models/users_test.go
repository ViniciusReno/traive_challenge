package models

// func TestUser_Validate(t *testing.T) {
// 	tests := []struct {
// 		user     User
// 		expected error
// 	}{
// 		{
// 			user: User{
// 				FirstName: "John",
// 				LastName:  "Doe",
// 				Password:  "Password123!",
// 				Email:     "john.doe@example.com",
// 			},
// 			expected: nil,
// 		},
// 		{
// 			user: User{
// 				FirstName: "John",
// 				LastName:  "Doe",
// 				Password:  "Password",
// 				Email:     "john.doe@example.com",
// 			},
// 			expected: fmt.Errorf("password must contain at least one number"),
// 		},
// 	}

// 	for _, test := range tests {
// 		err := test.user.Validate()

// 		if (err == nil && test.expected != nil) || (err != nil && test.expected == nil) {
// 			t.Errorf("Expected error: %v, Received error: %v", test.expected, err)
// 		} else if err != nil && test.expected != nil && err.Error() != test.expected.Error() {
// 			t.Errorf("Expected error message: %s, Received error message: %s", test.expected.Error(), err.Error())
// 		}
// 	}
// }
