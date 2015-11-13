require "rake/clean"
require "fileutils"

CLEAN.include("bench_test.go", "funcs.go")

file "bench_test.go" => "bench_test.go.erb" do
  sh "erb -T - bench_test.go.erb | gofmt > bench_test.go"
end

file "funcs.go" => "funcs.go.erb" do
  sh "erb -T - funcs.go.erb | gofmt > funcs.go"
end

desc "Run Go benchamrks"
task :benchmark => ["bench_test.go", "funcs.go"] do
  sh "go test -test.bench=."
end

task :default => :benchmark
