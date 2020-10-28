package dgip_test

import (
	"testing"

	dgip "github.com/dennigogo/go-interface-parser"
	"github.com/stretchr/testify/assert"
)

func Test_JSON(t *testing.T) {

	j := dgip.JSON()

	t.Run(`dst=nil && src=nil`, func(t *testing.T) {
		result, check := j.Merge(nil, nil)
		assert.Equal(t, check, false)
		assert.Nil(t, result)

		var dstMAP map[string]interface{}
		resultMAP, check := j.Merge(dstMAP, nil)
		assert.Equal(t, check, false)
		assert.Nil(t, resultMAP)

		var srcMAP map[string]interface{}
		resultMAP, check = j.Merge(dstMAP, srcMAP)
		assert.Equal(t, check, true)
		assert.IsType(t, srcMAP, resultMAP)

		var dstARRAY []interface{}
		resultARRAY, check := j.Merge(dstARRAY, nil)
		assert.Equal(t, check, false)
		assert.Nil(t, resultARRAY)

		var srcARRAY []interface{}
		resultARRAY, check = j.Merge(dstARRAY, srcARRAY)
		assert.Equal(t, check, true)
		assert.IsType(t, srcARRAY, resultARRAY)
	})

	t.Run(`dst=nil && src!=nil`, func(t *testing.T) {
		srcMAP := map[string]interface{}{
			`test`: `testSRC`,
		}

		resultMAP, check := j.Merge(nil, srcMAP)
		assert.Equal(t, check, true)
		assert.Equal(t, resultMAP, srcMAP)

		var dst map[string]interface{}
		resultMAP, check = j.Merge(dst, srcMAP)
		assert.Equal(t, check, true)
		assert.Equal(t, resultMAP, srcMAP)

		srcARRAY := []interface{}{
			`testSRC`,
		}

		resultARRAY, check := j.Merge(nil, srcARRAY)
		assert.Equal(t, check, true)
		assert.Equal(t, resultARRAY, srcARRAY)

		var dstARRAY []interface{}
		resultARRAY, check = j.Merge(dstARRAY, srcARRAY)
		assert.Equal(t, check, true)
		assert.IsType(t, resultARRAY, srcARRAY)
	})

	t.Run(`dst!=nil && src==nil`, func(t *testing.T) {
		dstMAP := map[string]interface{}{
			`test`: `testDST`,
		}

		resultMAP, check := j.Merge(dstMAP, nil)
		assert.Equal(t, check, false)
		assert.Equal(t, resultMAP, dstMAP)

		var src map[string]interface{}
		resultMAP, check = j.Merge(dstMAP, src)
		assert.Equal(t, check, true)
		assert.Equal(t, resultMAP, dstMAP)

		dstARRAY := []interface{}{
			`testDST`,
		}

		resultARRAY, check := j.Merge(dstARRAY, nil)
		assert.Equal(t, check, false)
		assert.Equal(t, resultARRAY, dstARRAY)

		var srcARRAY []interface{}
		resultARRAY, check = j.Merge(dstARRAY, srcARRAY)
		assert.Equal(t, check, true)
		assert.IsType(t, resultARRAY, dstARRAY)
	})

	t.Run(`dst!=nil && src!=nil`, func(t *testing.T) {
		dstMAP := map[string]interface{}{
			`test`:    `testDST`,
			`testDST`: `DST`,
		}
		srcMAP := map[string]interface{}{
			`test`:    `testSRC`,
			`testSRC`: `SRC`,
		}
		resultMAP, check := j.Merge(dstMAP, srcMAP)
		assert.Equal(t, check, true)
		assert.Equal(t, resultMAP, map[string]interface{}{
			`test`:    `testSRC`,
			`testDST`: `DST`,
			`testSRC`: `SRC`,
		})

		dstARRAY := []interface{}{
			`testDST`,
		}
		srcARRAY := []interface{}{
			`testSRC`,
		}
		resultARRAY, check := j.Merge(dstARRAY, srcARRAY)
		assert.Equal(t, check, true)
		assert.Equal(t, resultARRAY, []interface{}{
			`testSRC`,
		})

		resultMAP, check = j.Merge(dstMAP, srcARRAY)
		assert.Equal(t, check, false)
		assert.Equal(t, resultMAP, srcARRAY)

		resultARRAY, check = j.Merge(dstARRAY, srcMAP)
		assert.Equal(t, check, false)
		assert.Equal(t, resultARRAY, srcMAP)

		dstMAP = map[string]interface{}{
			`test`:    `testDST`,
			`testDST`: `DST`,
			`test2`: map[string]interface{}{
				`test2-2`: `test2-2`,
			},
			`test3`: map[string]interface{}{
				`test3-1`: map[string]interface{}{
					`test3-1`: `test3-1`,
					`test3-2`: `test3-2`,
				},
			},
			`test4`: map[string]interface{}{
				`test4-1`: []interface{}{
					`test4-1`,
					`test4-3`,
				},
			},
		}
		srcMAP = map[string]interface{}{
			`test`: `testSRC`,
			`testSRC`: []interface{}{
				`SRC1`,
			},
			`test2`: map[string]interface{}{
				`test2-2`: []interface{}{
					`test2-2`,
				},
			},
			`test3`: map[string]interface{}{
				`test3-1`: map[string]interface{}{
					`test3-2`: `test3-3`,
					`test3-3`: `test3-4`,
				},
			},
			`test4`: map[string]interface{}{
				`test4-1`: []interface{}{
					`test4-2`,
					`test4-4`,
				},
			},
		}
		resultMAP, check = j.Merge(dstMAP, srcMAP)
		assert.Equal(t, check, true)
		assert.Equal(t, resultMAP, map[string]interface{}{
			`test`:    `testSRC`,
			`testDST`: `DST`,
			`test2`: map[string]interface{}{
				`test2-2`: []interface{}{
					`test2-2`,
				},
			},
			`test3`: map[string]interface{}{
				`test3-1`: map[string]interface{}{
					`test3-1`: `test3-1`,
					`test3-2`: `test3-3`,
					`test3-3`: `test3-4`,
				},
			},
			`testSRC`: []interface{}{
				`SRC1`,
			},
			`test4`: map[string]interface{}{
				`test4-1`: []interface{}{
					`test4-2`,
					`test4-4`,
				},
			},
		})

		dstMAP = map[string]interface{}{
			`test`: []interface{}{
				`testDST`,
			},
			`testDST`: `DST`,
			`test2`: map[string]interface{}{
				`test2-2`: `test2-2`,
			},
		}
		srcMAP = map[string]interface{}{
			`test`: `testSRC`,
			`testSRC`: []interface{}{
				`SRC1`,
			},
			`test2`: map[string]interface{}{
				`test2-2`: []interface{}{
					`test2-2`,
				},
			},
		}
		resultMAP, check = j.Merge(dstMAP, srcMAP)
		assert.Equal(t, check, true)
		assert.Equal(t, resultMAP, map[string]interface{}{
			`test`:    `testSRC`,
			`testDST`: `DST`,
			`test2`: map[string]interface{}{
				`test2-2`: []interface{}{
					`test2-2`,
				},
			},
			`testSRC`: []interface{}{
				`SRC1`,
			},
		})
	})
}
