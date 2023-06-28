package states

import "context"

type IState interface {
	Run(context.Context)
}
