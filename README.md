# go-interface-parser

## Merge of interfaces

```
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

result, check := dgip.JSON().Merge(dst, src)
check: true
result: map[string]interface{}{
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
}
```