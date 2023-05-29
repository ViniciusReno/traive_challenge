package repository

// func TestCreateTransactions(t *testing.T) {
// 	repo := &mock.MockRepository{}
// 	transactions := []models.Transaction{}

// 	repo.On("CreateTransactions", transactions).Return(nil)

// 	err := repo.CreateTransactions(transactions)

// 	repo.AssertCalled(t, "CreateTransactions", transactions)

// 	if err != nil {
// 		t.Errorf("Expected error: nil, Received error: %v", err)
// 	}
// }

// func TestListTransactionsByID(t *testing.T) {
// 	repo := &mock.MockRepository{}
// 	id := "123"
// 	expectedTransactions := []models.Transaction{}

// 	repo.On("ListTransactionsByID", id).Return(expectedTransactions, nil)

// 	transactions, err := repo.ListTransactionsByID(id)

// 	repo.AssertCalled(t, "ListTransactionsByID", id)

// 	if err != nil {
// 		t.Errorf("Expected error: nil, Received error: %v", err)
// 	}

// 	if len(transactions) != len(expectedTransactions) {
// 		t.Errorf("Expected transactions count: %d, Received transactions count: %d", len(expectedTransactions), len(transactions))
// 	}
// }

// func TestListTransactions(t *testing.T) {
// 	repo := &mock.MockRepository{}
// 	page := 1
// 	limit := 10
// 	filters := map[string]interface{}{}

// 	expectedTransactions := []models.TransactionResponse{}

// 	repo.On("ListTransactions", page, limit, filters).Return(expectedTransactions, nil)

// 	transactions, err := repo.ListTransactions(page, limit, filters)

// 	repo.AssertCalled(t, "ListTransactions", page, limit, filters)

// 	if err != nil {
// 		t.Errorf("Expected error: nil, Received error: %v", err)
// 	}

// 	if len(transactions) != len(expectedTransactions) {
// 		t.Errorf("Expected transactions count: %d, Received transactions count: %d", len(expectedTransactions), len(transactions))
// 	}
// }
