Useful set of golang benchmarks

	go test -bench=. -benchmem

	goos: linux
	goarch: amd64
	pkg: github.com/kachan1208/golang-benchmarks
	BenchmarkIterateMapWithoutPointer-4           	 5000000	       248 ns/op	      70 B/op	       0 allocs/op
	BenchmarkIterateMapWithPointer-4              	 3000000	       379 ns/op	      66 B/op	       1 allocs/op
	BenchmarkStructEncode-4                       	 3000000	       417 ns/op	     240 B/op	       2 allocs/op
	BenchmarkInterfaceEncode-4                    	 3000000	       502 ns/op	     256 B/op	       3 allocs/op
	BenchmarkPassStructIntoEncodeFunc-4           	 3000000	       428 ns/op	     240 B/op	       2 allocs/op
	BenchmarkPassInterfaceIntoEncodeFunc-4        	 3000000	       646 ns/op	     256 B/op	       3 allocs/op
	BenchmarkArrayEncode-4                        	 1000000	      1067 ns/op	     880 B/op	       4 allocs/op
	BenchmarkSliceEncode-4                        	 1000000	      1142 ns/op	     896 B/op	       4 allocs/op
	BenchmarkPassArrayIntoEncodeFunc-4            	 1000000	      1080 ns/op	     880 B/op	       4 allocs/op
	BenchmarkPassSliceIntoEncodeFunc-4            	 1000000	      1055 ns/op	     896 B/op	       4 allocs/op
	BenchmarkPassArrayByPointerIntoEncodeFunc-4   	 1000000	      1013 ns/op	     864 B/op	       3 allocs/op
	BenchmarkPassSliceByPointerIntoEncodeFunc-4   	 1000000	      1113 ns/op	     864 B/op	       3 allocs/op
	BenchmarkMapGetKeysReflect-4                  	 5000000	       283 ns/op	     102 B/op	       1 allocs/op
	BenchmarkMapGetKeysForeach-4                  	 5000000	       257 ns/op	     114 B/op	       0 allocs/op
	PASS
	ok  	github.com/kachan1208/golang-benchmarks	20.735s
