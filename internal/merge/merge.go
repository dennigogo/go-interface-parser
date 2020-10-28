package merge

type Interface interface {
	Merge(dst interface{}, src interface{}) (interface{}, bool)
}

type detectType uint

const (
	noneType detectType = iota
	mapType
	arrayType
)

type JSONMerge struct{}

func NewJSONMerge() Interface {
	return &JSONMerge{}
}

func (m *JSONMerge) Merge(dst interface{}, src interface{}) (interface{}, bool) {
	switch {
	case dst == nil && src != nil:
		result := src
		return result, true
	case dst == nil && src == nil, dst != nil && src == nil:
		result := dst
		return result, false
	default:
		return m.processing(dst, src)
	}
}

func (m *JSONMerge) processing(dst interface{}, src interface{}) (interface{}, bool) {
	dstType := m.typeDetect(dst)
	srcType := m.typeDetect(src)

	switch {
	case dstType == srcType && srcType == arrayType:
		dstARRAY, _ := dst.([]interface{})
		srcARRAY, _ := src.([]interface{})
		return m.mergeARRAY(dstARRAY, srcARRAY), true
	case dstType == srcType && srcType == mapType:
		dstMAP, _ := dst.(map[string]interface{})
		srcMAP, _ := src.(map[string]interface{})
		return m.mergeMAP(dstMAP, srcMAP), true
	default:
		result := src
		return result, false
	}
}

func (m *JSONMerge) typeDetect(input interface{}) detectType {
	if _, ok := input.([]interface{}); ok {
		return arrayType
	}

	if _, ok := input.(map[string]interface{}); ok {
		return mapType
	}

	return noneType
}

func (m *JSONMerge) mergeARRAY(dst []interface{}, src []interface{}) []interface{} {
	var test []interface{}
	if len(src) > len(dst) {
		test = make([]interface{}, len(src))
	} else {
		test = make([]interface{}, len(dst))
	}

	for i := range dst {
		test[i] = dst[i]
	}

	for i := range src {
		test[i] = src[i]
	}

	var result []interface{}
	for i := range test {
		if test[i] != nil {
			result = append(result, test[i])
		}
	}

	test = nil

	return result
}

func (m *JSONMerge) mergeMAP(dst map[string]interface{}, src map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for i := range dst {
		if dst[i] != nil {
			result[i] = dst[i]
		}
	}

	var test interface{}
	for i := range src {
		if _, ok := result[i]; !ok || m.typeDetect(result[i]) == noneType {
			if src[i] != nil {
				result[i] = src[i]
			} else {
				delete(result, i)
			}
			continue
		}

		test, _ = m.Merge(result[i], src[i])
		if src[i] != nil {
			result[i] = test
		}
	}

	return result
}
