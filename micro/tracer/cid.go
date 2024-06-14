package tracer

import (
	id "github.com/jinmukeji/go-pkg/v2/id-gen"
)

func NewCid() string {
	cid := id.NewXid()
	return cid
}
