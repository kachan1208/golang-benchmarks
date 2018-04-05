Useful set of golang benchmarks

    go test -bench=. -benchmem
    
    goos: linux
    goarch: amd64
    pkg: github.com/kachan1208/golang-benchmarks
    BenchmarkIterateMapWithoutPointer-4             10000000           223 ns/op      70 B/op          0 allocs/op
    BenchmarkIterateMapWithPointer-4                 5000000           373 ns/op      76 B/op          1 allocs/op
    BenchmarkStructEncode-4                          5000000           361 ns/op     240 B/op          2 allocs/op
    BenchmarkInterfaceEncode-4                       3000000           460 ns/op     256 B/op          3 allocs/op
    BenchmarkPassStructIntoEncodeFunc-4              5000000           365 ns/op     240 B/op          2 allocs/op
    BenchmarkPassInterfaceIntoEncodeFunc-4           3000000           467 ns/op     256 B/op          3 allocs/op
    BenchmarkArrayEncode-4                           2000000           887 ns/op     880 B/op          4 allocs/op
    BenchmarkSliceEncode-4                           2000000           908 ns/op     896 B/op          4 allocs/op
    BenchmarkPassArrayIntoEncodeFunc-4               2000000           953 ns/op     880 B/op          4 allocs/op
    BenchmarkPassSliceIntoEncodeFunc-4               2000000           926 ns/op     896 B/op          4 allocs/op
    BenchmarkPassArrayByPointerIntoEncodeFunc-4      2000000           887 ns/op     864 B/op          3 allocs/op
    BenchmarkPassSliceByPointerIntoEncodeFunc-4      2000000           904 ns/op     864 B/op          3 allocs/op
    BenchmarkMapGetKeysReflect-4                     5000000           254 ns/op     102 B/op          1 allocs/op
    BenchmarkMapGetKeysForeach-4                    10000000           229 ns/op     113 B/op          0 allocs/op
    BenchmarkIncrementInt-4                         2000000000           0.31 ns/op        0 B/op          0 allocs/op
    BenchmarkIncrementIntWithCastFromInterface-4    100000000           16.1 ns/op         8 B/op          1 allocs/op
    BenchmarkCastIntFromInterface-4                 100000000           17.3 ns/op         8 B/op          1 allocs/op
    BenchmarkCustomHash-4                           10000000           117 ns/op      51 B/op          2 allocs/op
    BenchmarkMD5Hash-4                               5000000           389 ns/op     211 B/op          5 allocs/op
    BenchmarkFnvHash-4                              10000000           211 ns/op     112 B/op          4 allocs/op
