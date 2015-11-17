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

### CPU Cache-Friendliness

The cache friendliness any algorithm can dramatically alter its performance characteristics. This benchmark includes 3 levels of cache-friendliness.

* The optimal cache-friendliness level directly uses the benchmark loop's index as the branch discriminator. This loops through branches in a predictable manner (e.g. `i % 4` for a case with 4 branches).
* The moderate cache-friendliness level uses the benchmark loop's index to look up the branch to take in pre-computed slice. This incurs a memory read that may displace some of the cache. This slice is simply an ascending list of numbers, so the branches are followed in a predictable manner (e.g. `ascInputs[i%len(ascInputs)] % 4` for a case with 4 branches).
* The poor cache-friendliness level uses the benchmark loop's index to look up the branch to take in pre-computed slice. This incurs a memory read that may displace some of the cache. This slice is a randomized list of numbers, so branches are followed in an unpredictable manner (e.g. `randInputs[i%len(randInputs)] % 4` for a case with 4 branches).

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
BenchmarkOptimalCacheFriendlinessSwitchInlineFunc4-8      1000000000           1.98 ns/op
BenchmarkOptimalCacheFriendlinessMapInlineFunc4-8         1000000000           2.43 ns/op
BenchmarkModerateCacheFriendlinessSwitchInlineFunc4-8     200000000          8.16 ns/op
BenchmarkModerateCacheFriendlinessMapInlineFunc4-8        100000000         10.6 ns/op
BenchmarkPoorCacheFriendlinessSwitchInlineFunc4-8         100000000         14.2 ns/op
BenchmarkPoorCacheFriendlinessMapInlineFunc4-8            100000000         19.0 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc4-8    500000000          3.48 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc4-8       1000000000           2.68 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc4-8   100000000         10.9 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc4-8      100000000         11.1 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc4-8       100000000         17.6 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc4-8          100000000         19.6 ns/op
BenchmarkOptimalCacheFriendlinessSwitchInlineFunc8-8      1000000000           2.33 ns/op
BenchmarkOptimalCacheFriendlinessMapInlineFunc8-8         1000000000           2.38 ns/op
BenchmarkModerateCacheFriendlinessSwitchInlineFunc8-8     200000000          9.68 ns/op
BenchmarkModerateCacheFriendlinessMapInlineFunc8-8        100000000         10.5 ns/op
BenchmarkPoorCacheFriendlinessSwitchInlineFunc8-8         100000000         16.5 ns/op
BenchmarkPoorCacheFriendlinessMapInlineFunc8-8            100000000         20.3 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc8-8    500000000          3.76 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc8-8       1000000000           2.74 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc8-8   100000000         11.6 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc8-8      100000000         10.9 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc8-8       100000000         19.9 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc8-8          100000000         21.1 ns/op
BenchmarkOptimalCacheFriendlinessSwitchInlineFunc16-8     1000000000           2.60 ns/op
BenchmarkOptimalCacheFriendlinessMapInlineFunc16-8        1000000000           2.51 ns/op
BenchmarkModerateCacheFriendlinessSwitchInlineFunc16-8    200000000          8.94 ns/op
BenchmarkModerateCacheFriendlinessMapInlineFunc16-8       100000000         10.6 ns/op
BenchmarkPoorCacheFriendlinessSwitchInlineFunc16-8        100000000         19.3 ns/op
BenchmarkPoorCacheFriendlinessMapInlineFunc16-8           50000000          21.1 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc16-8   300000000          4.13 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc16-8      1000000000           2.68 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc16-8  100000000         11.8 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc16-8     100000000         10.9 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc16-8      50000000          23.8 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc16-8         100000000         21.8 ns/op
BenchmarkOptimalCacheFriendlinessSwitchInlineFunc32-8     1000000000           2.75 ns/op
BenchmarkOptimalCacheFriendlinessMapInlineFunc32-8        1000000000           2.38 ns/op
BenchmarkModerateCacheFriendlinessSwitchInlineFunc32-8    200000000          9.09 ns/op
BenchmarkModerateCacheFriendlinessMapInlineFunc32-8       100000000         10.6 ns/op
BenchmarkPoorCacheFriendlinessSwitchInlineFunc32-8        50000000          22.8 ns/op
BenchmarkPoorCacheFriendlinessMapInlineFunc32-8           100000000         21.7 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc32-8   300000000          4.23 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc32-8      1000000000           2.74 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc32-8  100000000         11.8 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc32-8     100000000         10.9 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc32-8      50000000          27.5 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc32-8         100000000         22.0 ns/op
BenchmarkOptimalCacheFriendlinessSwitchInlineFunc64-8     500000000          3.16 ns/op
BenchmarkOptimalCacheFriendlinessMapInlineFunc64-8        1000000000           2.38 ns/op
BenchmarkModerateCacheFriendlinessSwitchInlineFunc64-8    200000000          9.33 ns/op
BenchmarkModerateCacheFriendlinessMapInlineFunc64-8       100000000         10.6 ns/op
BenchmarkPoorCacheFriendlinessSwitchInlineFunc64-8        50000000          26.6 ns/op
BenchmarkPoorCacheFriendlinessMapInlineFunc64-8           100000000         22.1 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc64-8   300000000          4.90 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc64-8      500000000          3.19 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc64-8  100000000         12.2 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc64-8     100000000         10.9 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc64-8      50000000          31.8 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc64-8         50000000          23.5 ns/op
BenchmarkOptimalCacheFriendlinessSwitchInlineFunc128-8    500000000          3.66 ns/op
BenchmarkOptimalCacheFriendlinessMapInlineFunc128-8       1000000000           2.58 ns/op
BenchmarkModerateCacheFriendlinessSwitchInlineFunc128-8   100000000         10.2 ns/op
BenchmarkModerateCacheFriendlinessMapInlineFunc128-8      100000000         10.6 ns/op
BenchmarkPoorCacheFriendlinessSwitchInlineFunc128-8       50000000          31.5 ns/op
BenchmarkPoorCacheFriendlinessMapInlineFunc128-8          50000000          22.4 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc128-8  300000000          5.15 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc128-8     500000000          3.24 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc128-8 100000000         12.5 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc128-8    100000000         11.0 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc128-8     50000000          35.4 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc128-8        50000000          23.8 ns/op
BenchmarkOptimalCacheFriendlinessSwitchInlineFunc256-8    300000000          4.12 ns/op
BenchmarkOptimalCacheFriendlinessMapInlineFunc256-8       500000000          2.97 ns/op
BenchmarkModerateCacheFriendlinessSwitchInlineFunc256-8   100000000         10.7 ns/op
BenchmarkModerateCacheFriendlinessMapInlineFunc256-8      100000000         10.6 ns/op
BenchmarkPoorCacheFriendlinessSwitchInlineFunc256-8       50000000          34.4 ns/op
BenchmarkPoorCacheFriendlinessMapInlineFunc256-8          100000000         22.5 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc256-8  300000000          5.77 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc256-8     500000000          3.77 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc256-8 100000000         13.2 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc256-8    100000000         11.1 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc256-8     30000000          40.8 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc256-8        50000000          25.4 ns/op
BenchmarkOptimalCacheFriendlinessSwitchInlineFunc512-8    300000000          4.40 ns/op
BenchmarkOptimalCacheFriendlinessMapInlineFunc512-8       1000000000           2.84 ns/op
BenchmarkModerateCacheFriendlinessSwitchInlineFunc512-8   100000000         11.5 ns/op
BenchmarkModerateCacheFriendlinessMapInlineFunc512-8      100000000         11.0 ns/op
BenchmarkPoorCacheFriendlinessSwitchInlineFunc512-8       30000000          39.2 ns/op
BenchmarkPoorCacheFriendlinessMapInlineFunc512-8          50000000          22.7 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc512-8  100000000         12.5 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc512-8     300000000          4.01 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc512-8 100000000         18.2 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc512-8    100000000         11.6 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc512-8     30000000          46.7 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc512-8        50000000          27.1 ns/op
```
