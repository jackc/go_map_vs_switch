# Go Map vs. Switch Benchmarks

This benchmark measures the performance of branching with a switch calling a function vs using a map of functions. For example:

```go
// Given f is map[int]func()
f[n](i)
```

vs.

```go
    switch n {
    case 0:
      a()
    case 1:
      b()
    case 2:
      c()
    case 3:
      d()
    }
```

## Other Benchmark Dimensions

### Number of Branches

Cases with 4, 8, 16, 32, 64, 128, 256, and 512 branches are included in this benchmark.

### Function Inlining

The `switch` statement may benefit from inlining simple functions. This benchmark tests the difference between functions that can be inlined and those that cannot. The Go compiler does not inline functions that can `panic`. The non-inlinable functions include an unreachable `panic` to block inlining.

### Branch Predictability

Some tests sequentially follow branch 0, 1, 2, etc. Others visit branches in a random order.

### Branch Destination Computation or Lookup

* Compute destination benchmarks use the benchmark loop's index as the branch discriminator. This loops through branches in a predictable manner (e.g. `i % 4` for a case with 4 branches).
* Lookup destination benchmarks use the loop's index to look up the branch to take in pre-computed slice. (e.g. `ascInputs[i%len(ascInputs)] % 4` for a case with 4 branches).

## Running the Benchmarks

```
go get github.com/jackc/go_map_vs_switch
cd $GOPATH/src/github.com/jackc/go_map_vs_switch
go test -test.bench=.
```

These benchmarks contain a great deal of repetitive code. The Ruby tools `rake` and `erb` are used to automate the generation of these benchmarks. You do not need Ruby to run the benchmarks. However, if you wish to make changes you will need a Ruby install. Simply change the `*.erb` files and run `rake`.

## Results

The following results were produced from a Intel i7-4790K running `go1.5.1 linux/amd64` on Ubuntu 14.04.

```
BenchmarkPredictableComputedSwitchInlineFunc4-8     2000000000           1.99 ns/op
BenchmarkPredictableComputedMapInlineFunc4-8        1000000000           2.38 ns/op
BenchmarkPredictableLookupSwitchInlineFunc4-8       200000000          8.16 ns/op
BenchmarkPredictableLookupMapInlineFunc4-8          100000000         10.6 ns/op
BenchmarkUnpredictableLookupSwitchInlineFunc4-8     100000000         14.2 ns/op
BenchmarkUnpredictableLookupMapInlineFunc4-8        100000000         19.0 ns/op
BenchmarkPredictableComputedSwitchNoInlineFunc4-8   500000000          3.46 ns/op
BenchmarkPredictableComputedMapNoInlineFunc4-8      1000000000           2.69 ns/op
BenchmarkPredictableLookupSwitchNoInlineFunc4-8     100000000         11.0 ns/op
BenchmarkPredictableLookupMapNoInlineFunc4-8        100000000         11.1 ns/op
BenchmarkUnpredictableLookupSwitchNoInlineFunc4-8   100000000         17.6 ns/op
BenchmarkUnpredictableLookupMapNoInlineFunc4-8      100000000         19.7 ns/op
BenchmarkPredictableComputedSwitchInlineFunc8-8     1000000000           2.33 ns/op
BenchmarkPredictableComputedMapInlineFunc8-8        1000000000           2.39 ns/op
BenchmarkPredictableLookupSwitchInlineFunc8-8       200000000          9.70 ns/op
BenchmarkPredictableLookupMapInlineFunc8-8          100000000         10.5 ns/op
BenchmarkUnpredictableLookupSwitchInlineFunc8-8     100000000         16.5 ns/op
BenchmarkUnpredictableLookupMapInlineFunc8-8        100000000         20.4 ns/op
BenchmarkPredictableComputedSwitchNoInlineFunc8-8   500000000          3.76 ns/op
BenchmarkPredictableComputedMapNoInlineFunc8-8      1000000000           2.75 ns/op
BenchmarkPredictableLookupSwitchNoInlineFunc8-8     100000000         11.6 ns/op
BenchmarkPredictableLookupMapNoInlineFunc8-8        100000000         10.9 ns/op
BenchmarkUnpredictableLookupSwitchNoInlineFunc8-8   100000000         20.0 ns/op
BenchmarkUnpredictableLookupMapNoInlineFunc8-8      100000000         21.1 ns/op
BenchmarkPredictableComputedSwitchInlineFunc16-8    1000000000           2.60 ns/op
BenchmarkPredictableComputedMapInlineFunc16-8       1000000000           2.52 ns/op
BenchmarkPredictableLookupSwitchInlineFunc16-8      200000000          8.96 ns/op
BenchmarkPredictableLookupMapInlineFunc16-8         100000000         10.6 ns/op
BenchmarkUnpredictableLookupSwitchInlineFunc16-8    100000000         19.3 ns/op
BenchmarkUnpredictableLookupMapInlineFunc16-8       100000000         21.1 ns/op
BenchmarkPredictableComputedSwitchNoInlineFunc16-8  300000000          4.14 ns/op
BenchmarkPredictableComputedMapNoInlineFunc16-8     1000000000           2.69 ns/op
BenchmarkPredictableLookupSwitchNoInlineFunc16-8    100000000         11.7 ns/op
BenchmarkPredictableLookupMapNoInlineFunc16-8       100000000         10.9 ns/op
BenchmarkUnpredictableLookupSwitchNoInlineFunc16-8  50000000          23.9 ns/op
BenchmarkUnpredictableLookupMapNoInlineFunc16-8     100000000         21.9 ns/op
BenchmarkPredictableComputedSwitchInlineFunc32-8    1000000000           2.75 ns/op
BenchmarkPredictableComputedMapInlineFunc32-8       1000000000           2.39 ns/op
BenchmarkPredictableLookupSwitchInlineFunc32-8      200000000          9.10 ns/op
BenchmarkPredictableLookupMapInlineFunc32-8         100000000         10.6 ns/op
BenchmarkUnpredictableLookupSwitchInlineFunc32-8    50000000          22.8 ns/op
BenchmarkUnpredictableLookupMapInlineFunc32-8       100000000         21.7 ns/op
BenchmarkPredictableComputedSwitchNoInlineFunc32-8  300000000          4.24 ns/op
BenchmarkPredictableComputedMapNoInlineFunc32-8     1000000000           2.75 ns/op
BenchmarkPredictableLookupSwitchNoInlineFunc32-8    100000000         11.9 ns/op
BenchmarkPredictableLookupMapNoInlineFunc32-8       100000000         11.0 ns/op
BenchmarkUnpredictableLookupSwitchNoInlineFunc32-8  50000000          27.5 ns/op
BenchmarkUnpredictableLookupMapNoInlineFunc32-8     50000000          22.0 ns/op
BenchmarkPredictableComputedSwitchInlineFunc64-8    500000000          3.16 ns/op
BenchmarkPredictableComputedMapInlineFunc64-8       1000000000           2.38 ns/op
BenchmarkPredictableLookupSwitchInlineFunc64-8      200000000          9.37 ns/op
BenchmarkPredictableLookupMapInlineFunc64-8         100000000         10.6 ns/op
BenchmarkUnpredictableLookupSwitchInlineFunc64-8    50000000          26.8 ns/op
BenchmarkUnpredictableLookupMapInlineFunc64-8       50000000          22.1 ns/op
BenchmarkPredictableComputedSwitchNoInlineFunc64-8  300000000          4.91 ns/op
BenchmarkPredictableComputedMapNoInlineFunc64-8     500000000          3.20 ns/op
BenchmarkPredictableLookupSwitchNoInlineFunc64-8    100000000         12.2 ns/op
BenchmarkPredictableLookupMapNoInlineFunc64-8       100000000         11.0 ns/op
BenchmarkUnpredictableLookupSwitchNoInlineFunc64-8  50000000          31.9 ns/op
BenchmarkUnpredictableLookupMapNoInlineFunc64-8     50000000          23.6 ns/op
BenchmarkPredictableComputedSwitchInlineFunc128-8   500000000          3.68 ns/op
BenchmarkPredictableComputedMapInlineFunc128-8      1000000000           2.60 ns/op
BenchmarkPredictableLookupSwitchInlineFunc128-8     100000000         10.2 ns/op
BenchmarkPredictableLookupMapInlineFunc128-8        100000000         10.6 ns/op
BenchmarkUnpredictableLookupSwitchInlineFunc128-8   50000000          31.0 ns/op
BenchmarkUnpredictableLookupMapInlineFunc128-8      50000000          22.4 ns/op
BenchmarkPredictableComputedSwitchNoInlineFunc128-8 300000000          5.17 ns/op
BenchmarkPredictableComputedMapNoInlineFunc128-8    500000000          3.24 ns/op
BenchmarkPredictableLookupSwitchNoInlineFunc128-8   100000000         12.5 ns/op
BenchmarkPredictableLookupMapNoInlineFunc128-8      100000000         11.0 ns/op
BenchmarkUnpredictableLookupSwitchNoInlineFunc128-8 50000000          35.5 ns/op
BenchmarkUnpredictableLookupMapNoInlineFunc128-8    50000000          23.9 ns/op
BenchmarkPredictableComputedSwitchInlineFunc256-8   300000000          4.13 ns/op
BenchmarkPredictableComputedMapInlineFunc256-8      1000000000           2.98 ns/op
BenchmarkPredictableLookupSwitchInlineFunc256-8     100000000         10.8 ns/op
BenchmarkPredictableLookupMapInlineFunc256-8        100000000         10.7 ns/op
BenchmarkUnpredictableLookupSwitchInlineFunc256-8   50000000          34.4 ns/op
BenchmarkUnpredictableLookupMapInlineFunc256-8      50000000          22.6 ns/op
BenchmarkPredictableComputedSwitchNoInlineFunc256-8 300000000          5.76 ns/op
BenchmarkPredictableComputedMapNoInlineFunc256-8    500000000          3.81 ns/op
BenchmarkPredictableLookupSwitchNoInlineFunc256-8   100000000         13.2 ns/op
BenchmarkPredictableLookupMapNoInlineFunc256-8      100000000         11.1 ns/op
BenchmarkUnpredictableLookupSwitchNoInlineFunc256-8 30000000          41.0 ns/op
BenchmarkUnpredictableLookupMapNoInlineFunc256-8    50000000          25.5 ns/op
BenchmarkPredictableComputedSwitchInlineFunc512-8   300000000          4.39 ns/op
BenchmarkPredictableComputedMapInlineFunc512-8      500000000          2.86 ns/op
BenchmarkPredictableLookupSwitchInlineFunc512-8     100000000         11.4 ns/op
BenchmarkPredictableLookupMapInlineFunc512-8        100000000         11.1 ns/op
BenchmarkUnpredictableLookupSwitchInlineFunc512-8   30000000          39.3 ns/op
BenchmarkUnpredictableLookupMapInlineFunc512-8      50000000          22.7 ns/op
BenchmarkPredictableComputedSwitchNoInlineFunc512-8 100000000         12.5 ns/op
BenchmarkPredictableComputedMapNoInlineFunc512-8    300000000          3.91 ns/op
BenchmarkPredictableLookupSwitchNoInlineFunc512-8   100000000         18.2 ns/op
BenchmarkPredictableLookupMapNoInlineFunc512-8      100000000         11.6 ns/op
BenchmarkUnpredictableLookupSwitchNoInlineFunc512-8 30000000          46.9 ns/op
BenchmarkUnpredictableLookupMapNoInlineFunc512-8    50000000          27.2 ns/op

```
