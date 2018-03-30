Useful set of golang benchmarks

    go test -bench=. -benchmem

    goos: linux
    goarch: amd64
    pkg: github.com/kachan1208/golang-benchmarks
    BenchmarkIterateMapWithoutPointer-4              5000000           268 ns/op      70 B/op          0 allocs/op
    BenchmarkIterateMapWithPointer-4                 3000000           398 ns/op      66 B/op          1 allocs/op
    BenchmarkStructEncode-4                          5000000           412 ns/op     240 B/op          2 allocs/op
    BenchmarkInterfaceEncode-4                       3000000           492 ns/op     256 B/op          3 allocs/op
    BenchmarkPassStructIntoEncodeFunc-4              5000000           394 ns/op     240 B/op          2 allocs/op
    BenchmarkPassInterfaceIntoEncodeFunc-4           3000000           480 ns/op     256 B/op          3 allocs/op
    BenchmarkArrayEncode-4                           2000000           935 ns/op     880 B/op          4 allocs/op
    BenchmarkSliceEncode-4                           2000000           945 ns/op     896 B/op          4 allocs/op
    BenchmarkPassArrayIntoEncodeFunc-4               2000000           931 ns/op     880 B/op          4 allocs/op
    BenchmarkPassSliceIntoEncodeFunc-4               2000000           943 ns/op     896 B/op          4 allocs/op
    BenchmarkPassArrayByPointerIntoEncodeFunc-4      2000000           902 ns/op     864 B/op          3 allocs/op
    BenchmarkPassSliceByPointerIntoEncodeFunc-4      2000000           961 ns/op     864 B/op          3 allocs/op
    BenchmarkMapGetKeysReflect-4                     5000000           303 ns/op     102 B/op          1 allocs/op
    BenchmarkMapGetKeysForeach-4                     5000000           260 ns/op     114 B/op          0 allocs/op
    BenchmarkIncrementInt-4                          2000000000        0.30 ns/op    0 B/op            0 allocs/op
    BenchmarkIncrementIntWithCastFromInterface-4     100000000         16.7 ns/op    8 B/op            1 allocs/op
    BenchmarkCastIntFromInterface-4                  2000000000        0.30 ns/op    0 B/op            0 allocs/op
    PASS
    ok      github.com/kachan1208/golang-benchmarks 35.369s

