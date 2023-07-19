## SONIC vs JSON 

#### SONIC
| Name                            | Operation | ns/op        | B/op        | allocs/op    |
|---------------------------------|-----------|--------------|-------------|--------------|
| BenchmarkSonic/input_size_10-4  | 184616    | 10799 ns/op  | 3787 B/op   | 11 allocs/op |
| BenchmarkSonic/input_size_50-4  | 39327     | 31660 ns/op  | 15188 B/op  | 12 allocs/op |
| BenchmarkSonic/input_size_100-4 | 20794     | 55688 ns/op  | 34839 B/op  | 13 allocs/op |
| BenchmarkSonic/input_size_400-4 | 5713      | 213816 ns/op | 130859 B/op | 14 allocs/op |

#### JSON
| Name                           | Operation | ns/op         | B/op       | allocs/op      |
|--------------------------------|-----------|---------------|------------|----------------|
| BenchmarkJSON/input_size_10-4  | 21874     | 57925 ns/op   | 2689 B/op  | 57 allocs/op   |
| BenchmarkJSON/input_size_50-4  | 4152      | 273070 ns/op  | 12053 B/op | 265 allocs/op  |
| BenchmarkJSON/input_size_100-4 | 2942      | 390126 ns/op  | 24157 B/op | 519 allocs/op  |
| BenchmarkJSON/input_size_400-4 | 930       | 1406397 ns/op | 96853 B/op | 2025 allocs/op |