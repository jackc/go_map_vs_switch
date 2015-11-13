# Go Map vs. Switch Benchmarks

This benchmark measures the performance of branching with a switch calling a function vs using a map of functions. For example:

```go
// Given f is []func()
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
* The moderate cache-friendliness level uses the benchmark loop's index to look up the branch to take in procomputed slice. This incurs a memory read that may displace some of the cache. This slice is simply an ascending list of numbers, so the branches are followed in a predictable manner (e.g. `ascInputs[i%len(ascInputs)] % 4` for a case with 4 branches).
* The poor cache-friendliness level uses the benchmark loop's index to look up the branch to take in procomputed slice. This incurs a memory read that may displace some of the cache. This slice is a randomized list of numbers, so branches are followed in an unpredictable manner (e.g. `randInputs[i%len(randInputs)] % 4` for a case with 4 branches).

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
BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc4-8     1000000000           1.99 ns/op
BenchmarkOptimalCacheFriendlinessMapMinimalFunc4-8        1000000000           2.39 ns/op
BenchmarkModerateCacheFriendlinessSwitchMinimalFunc4-8    200000000          8.18 ns/op
BenchmarkModerateCacheFriendlinessMapMinimalFunc4-8       100000000         10.6 ns/op
BenchmarkPoorCacheFriendlinessSwitchMinimalFunc4-8        100000000         14.2 ns/op
BenchmarkPoorCacheFriendlinessMapMinimalFunc4-8           100000000         19.1 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc4-8    500000000          3.48 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc4-8       1000000000           2.69 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc4-8   100000000         11.0 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc4-8      100000000         11.2 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc4-8       100000000         17.7 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc4-8          100000000         19.7 ns/op
BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc8-8     1000000000           2.33 ns/op
BenchmarkOptimalCacheFriendlinessMapMinimalFunc8-8        1000000000           2.39 ns/op
BenchmarkModerateCacheFriendlinessSwitchMinimalFunc8-8    200000000          9.76 ns/op
BenchmarkModerateCacheFriendlinessMapMinimalFunc8-8       100000000         10.5 ns/op
BenchmarkPoorCacheFriendlinessSwitchMinimalFunc8-8        100000000         16.5 ns/op
BenchmarkPoorCacheFriendlinessMapMinimalFunc8-8           100000000         20.4 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc8-8    500000000          3.76 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc8-8       1000000000           2.75 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc8-8   100000000         11.6 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc8-8      100000000         11.0 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc8-8       100000000         20.0 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc8-8          100000000         21.1 ns/op
BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc16-8    1000000000           2.60 ns/op
BenchmarkOptimalCacheFriendlinessMapMinimalFunc16-8       1000000000           2.51 ns/op
BenchmarkModerateCacheFriendlinessSwitchMinimalFunc16-8   200000000          8.98 ns/op
BenchmarkModerateCacheFriendlinessMapMinimalFunc16-8      100000000         10.6 ns/op
BenchmarkPoorCacheFriendlinessSwitchMinimalFunc16-8       100000000         19.4 ns/op
BenchmarkPoorCacheFriendlinessMapMinimalFunc16-8          100000000         21.2 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc16-8   300000000          4.13 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc16-8      1000000000           2.69 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc16-8  100000000         11.8 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc16-8     100000000         11.0 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc16-8      50000000          24.0 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc16-8         100000000         21.9 ns/op
BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc32-8    1000000000           2.76 ns/op
BenchmarkOptimalCacheFriendlinessMapMinimalFunc32-8       1000000000           2.39 ns/op
BenchmarkModerateCacheFriendlinessSwitchMinimalFunc32-8   200000000          9.13 ns/op
BenchmarkModerateCacheFriendlinessMapMinimalFunc32-8      100000000         10.6 ns/op
BenchmarkPoorCacheFriendlinessSwitchMinimalFunc32-8       50000000          22.9 ns/op
BenchmarkPoorCacheFriendlinessMapMinimalFunc32-8          100000000         22.0 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc32-8   300000000          4.24 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc32-8      1000000000           2.75 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc32-8  100000000         11.9 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc32-8     100000000         11.0 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc32-8      50000000          27.6 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc32-8         50000000          22.0 ns/op
BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc64-8    500000000          3.20 ns/op
BenchmarkOptimalCacheFriendlinessMapMinimalFunc64-8       1000000000           2.39 ns/op
BenchmarkModerateCacheFriendlinessSwitchMinimalFunc64-8   200000000          9.36 ns/op
BenchmarkModerateCacheFriendlinessMapMinimalFunc64-8      100000000         10.7 ns/op
BenchmarkPoorCacheFriendlinessSwitchMinimalFunc64-8       50000000          26.7 ns/op
BenchmarkPoorCacheFriendlinessMapMinimalFunc64-8          50000000          22.1 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc64-8   300000000          4.91 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc64-8      500000000          3.21 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc64-8  100000000         12.2 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc64-8     100000000         11.0 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc64-8      50000000          31.9 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc64-8         50000000          23.5 ns/op
BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc128-8   500000000          3.68 ns/op
BenchmarkOptimalCacheFriendlinessMapMinimalFunc128-8      1000000000           2.60 ns/op
BenchmarkModerateCacheFriendlinessSwitchMinimalFunc128-8  100000000         10.2 ns/op
BenchmarkModerateCacheFriendlinessMapMinimalFunc128-8     100000000         10.6 ns/op
BenchmarkPoorCacheFriendlinessSwitchMinimalFunc128-8      50000000          31.0 ns/op
BenchmarkPoorCacheFriendlinessMapMinimalFunc128-8         50000000          22.5 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc128-8  300000000          5.17 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc128-8     500000000          3.25 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc128-8 100000000         12.6 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc128-8    100000000         11.0 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc128-8     50000000          35.6 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc128-8        50000000          23.9 ns/op
BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc256-8   300000000          4.18 ns/op
BenchmarkOptimalCacheFriendlinessMapMinimalFunc256-8      500000000          2.98 ns/op
BenchmarkModerateCacheFriendlinessSwitchMinimalFunc256-8  100000000         10.8 ns/op
BenchmarkModerateCacheFriendlinessMapMinimalFunc256-8     100000000         10.7 ns/op
BenchmarkPoorCacheFriendlinessSwitchMinimalFunc256-8      50000000          34.5 ns/op
BenchmarkPoorCacheFriendlinessMapMinimalFunc256-8         50000000          22.6 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc256-8  300000000          5.87 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc256-8     500000000          4.12 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc256-8 100000000         14.1 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc256-8    100000000         11.8 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc256-8     30000000          44.8 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc256-8        50000000          26.6 ns/op
BenchmarkOptimalCacheFriendlinessSwitchMinimalFunc512-8   300000000          4.57 ns/op
BenchmarkOptimalCacheFriendlinessMapMinimalFunc512-8      500000000          3.30 ns/op
BenchmarkModerateCacheFriendlinessSwitchMinimalFunc512-8  100000000         12.9 ns/op
BenchmarkModerateCacheFriendlinessMapMinimalFunc512-8     100000000         12.3 ns/op
BenchmarkPoorCacheFriendlinessSwitchMinimalFunc512-8      30000000          39.9 ns/op
BenchmarkPoorCacheFriendlinessMapMinimalFunc512-8         100000000         23.5 ns/op
BenchmarkOptimalCacheFriendlinessSwitchNoInlineFunc512-8  100000000         13.2 ns/op
BenchmarkOptimalCacheFriendlinessMapNoInlineFunc512-8     300000000          4.33 ns/op
BenchmarkModerateCacheFriendlinessSwitchNoInlineFunc512-8 100000000         20.1 ns/op
BenchmarkModerateCacheFriendlinessMapNoInlineFunc512-8    100000000         11.6 ns/op
BenchmarkPoorCacheFriendlinessSwitchNoInlineFunc512-8     30000000          51.7 ns/op
BenchmarkPoorCacheFriendlinessMapNoInlineFunc512-8        50000000          27.5 ns/op
```
