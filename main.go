package dgip

import "github.com/dennigogo/go-interface-parser/internal/merge"

type Interface interface {
	Merge(dst interface{}, src interface{}) (interface{}, bool)
}

type dgJR struct {
	merge merge.Interface
}

func JSON() Interface {
	return &dgJR{
		merge: merge.NewJSONMerge(),
	}
}

func (j *dgJR) Merge(dst interface{}, src interface{}) (interface{}, bool) {
	return j.merge.Merge(dst, src)
}
