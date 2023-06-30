package operators

type IOperator interface {
	GetName() string
	Match(variable string, obj interface{}) bool
}