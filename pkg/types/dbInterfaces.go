package types

type IDBTransaction interface {
	Commit() error
	Rollback() error
	SavePoint(name string) error
	RollbackTo(name string) error
}

type IBaseRepository[T any] interface {
	BeginTransaction() IDBTransaction
	WithTransaction(queryTx IDBTransaction) T
}
