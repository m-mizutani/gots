# gots: Go Tools Suite

## Slice functions

### Map

```go
out := gots.Map([]string{"blue", "orange", "red"}, func(v string) int {
    return len(v)
})
fmt.Println(out)
// Output: [4 6 3]
```

### Filter

```go
out := gots.Filter([]int{1, 2, 3, 4, 5}, func(v int) bool {
    return v%2 == 0
})
fmt.Println(out)
// Output: [2 4]
```

### Reduce

```go
out := gots.Reduce([]int{1, 2, 3, 4, 5}, func(p, c int) int {
    return p + c
}, 10)
fmt.Println(out)
// Output: 25
```

### Uniq

```go
out := gots.Unique([]string{"A", "B", "A", "C", "B"})
fmt.Println(out)
// Output: [A B C]
```

### Count

```go
out := gots.Count([]string{"A", "B", "A", "C", "B"}, "B")
fmt.Println(out)
// Output: 2
```

### CountIf

```go
out := gots.CountIf([]string{"A", "B", "A", "C", "B"}, func(v string) bool {
    return v == "A" || v == "C"
})
fmt.Println(out)
// Output: 3
```

## Random string

```go
fmt.Println(gots.RandomString(6))
// Output: H9ZBvu

fmt.Println(gots.RandomString(32, "最高技術責任者"))
// Output: 技者責責任術任最高責責最者者最任術任術者最任最最高技術術責最高任
```

## Value to Pointer

```go
type input struct {
    ts *int64
}

// Compile error
// i := input {
//     ts: &(time.Now().Unix()),
// }


i := input {
    ts: gots.Ptr(time.Now().Unix())
}
```

## License

Apache License 2.0