package client

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestAwx(t *testing.T) {
	a, _ := (&AwxCient{}).Run(context.TODO(), nil)
	v, _ := json.Marshal(a)
	fmt.Println(string(v))
	time.Sleep(500*time.Millisecond)
}