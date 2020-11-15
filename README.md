Useful set of golang benchmarks

    go test -bench=. -benchmem
    goos: linux
    goarch: amd64
    BenchmarkIterateMapWithoutPointer-8            	 6023126	       223 ns/op      61 B/op	       0 allocs/op
    BenchmarkIterateMapWithPointer-8               	 4191570	       360 ns/op	      89 B/op	       1 allocs/op
    BenchmarkStructEncode-8                        	 5033884	       225 ns/op	      96 B/op	       2 allocs/op
    BenchmarkInterfaceEncode-8                     	 3573952	       338 ns/op	     112 B/op	       3 allocs/op
    BenchmarkPassStructIntoEncodeFunc-8            	 4538497	       238 ns/op	      96 B/op	       2 allocs/op
    BenchmarkPassInterfaceIntoEncodeFunc-8         	 3834438	       310 ns/op	     112 B/op	       3 allocs/op
    BenchmarkArrayEncode-8                         	 2732959	       435 ns/op	     256 B/op	       2 allocs/op
    BenchmarkSliceEncode-8                         	 2543551	       456 ns/op	     272 B/op	       2 allocs/op
    BenchmarkPassArrayIntoEncodeFunc-8             	 2782016	       435 ns/op	     256 B/op	       2 allocs/op
    BenchmarkPassSliceIntoEncodeFunc-8             	 2608023	       458 ns/op	     272 B/op	       2 allocs/op
    BenchmarkPassArrayByPointerIntoEncodeFunc-8    	 2784433	       419 ns/op	     240 B/op	       1 allocs/op
    BenchmarkPassSliceByPointerIntoEncodeFunc-8    	 2813598	       433 ns/op	     240 B/op	       1 allocs/op
    BenchmarkMapGetKeysReflect-8                   	 5064927	       251 ns/op	     101 B/op	       1 allocs/op
    BenchmarkMapGetKeysForeach-8                   	 5790820	       233 ns/op	     109 B/op	       0 allocs/op
    BenchmarkIncrementInt-8                        	1000000000	         0.249 ns/op	       0 B/op	       0 allocs/op
    BenchmarkIncrementIntWithCastFromInterface-8   	95807503	        13.1 ns/op	       8 B/op	       0 allocs/op
    BenchmarkCastIntFromInterface-8                	96132709	        12.7 ns/op	       8 B/op	       0 allocs/op
    BenchmarkCustomHash-8                          	13257229	        87.7 ns/op	      51 B/op	       2 allocs/op
    BenchmarkMD5Hash-8                             	 4416609	       272 ns/op	     211 B/op	       5 allocs/op
    BenchmarkFnvHash-8                             	 7650970	       153 ns/op	     112 B/op	       4 allocs/op
    BenchmarkErrof-8                               	 5872656	       204 ns/op	      64 B/op	       2 allocs/op
    BenchmarkWrap-8                                	 1764667	       670 ns/op	     352 B/op	       4 allocs/op
