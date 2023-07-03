package operators

type Condition interface {
	Compare() bool
}