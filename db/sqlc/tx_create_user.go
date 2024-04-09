package db

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

type CreateUserTxResult struct {
	User User
}

// func (store *SQLStore) CreateUserTx(
// 	ctx context.Context,
// 	arg CreateUserTxParams,
// ) (CreateUserTxResult, error) {
// 	var reult CreateUserTxResult
// }
