# gots: Go Tools Suite [![test](https://github.com/m-mizutani/gots/actions/workflows/test.yml/badge.svg)](https://github.com/m-mizutani/gots/actions/workflows/test.yml) [![gosec](https://github.com/m-mizutani/gots/actions/workflows/gosec.yml/badge.svg)](https://github.com/m-mizutani/gots/actions/workflows/gosec.yml) [![trivy](https://github.com/m-mizutani/gots/actions/workflows/trivy.yml/badge.svg)](https://github.com/m-mizutani/gots/actions/workflows/trivy.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/m-mizutani/gots.svg)](https://pkg.go.dev/github.com/m-mizutani/gots)

Simple and easy, but tedious function set in Go.

## Slice functions

### Map

```go
out := slice.Map([]string{"blue", "orange", "red"}, func(v string) int {
    return len(v)
})
fmt.Println(out)
// Output: [4 6 3]
```

### Filter

```go
out := slice.Filter([]int{1, 2, 3, 4, 5}, func(v int) bool {
    return v%2 == 0
})
fmt.Println(out)
// Output: [2 4]
```

### Reduce

```go
out := slice.Reduce([]int{1, 2, 3, 4, 5}, func(p, c int) int {
    return p + c
}, 10)
fmt.Println(out)
// Output: 25
```

### Uniq

```go
out := slice.Unique([]string{"A", "B", "A", "C", "B"})
fmt.Println(out)
// Output: [A B C]
```

### Count

```go
out := slice.Count([]string{"A", "B", "A", "C", "B"}, "B")
fmt.Println(out)
// Output: 2
```

### CountIf

```go
out := slice.CountIf([]string{"A", "B", "A", "C", "B"}, func(v string) bool {
    return v == "A" || v == "C"
})
fmt.Println(out)
// Output: 3
```

### Contains

```go
fmt.Println(slice.Contains([]string{"A", "B", "A", "C", "B"}, "B"))
// Output: true

fmt.Println(slice.Contains([]string{"A", "B", "A", "C", "B"}, "X"))
// Output: false
```

### Flatten

```go
	fmt.Println(slice.Flatten(
		[]string{"A", "B"},
		[]string{"C", "D", "E"},
		[]string{"F"},
	))
	// Output: [A B C D E F]
```

## Random string

```go
// Returns 6 random characters
fmt.Println(rands.New(6))
// Output: H9ZBvu

// Returns 32 random characters from "最高技術責任者"
fmt.Println(rands.New(32, "最高技術責任者"))
// Output: 技者責責任術任最高責責最者者最任術任術者最任最最高技術術責最高任

// Returns 6 random alphabet (both of lower and upper) characters
fmt.Println(rands.Alphabet(6))
// Output: UJieDx

// Returns 6 random lower case alphabet characters
fmt.Println(rands.LowerCase(6))
// Output: dgffld

// Returns 6 random upper case alphabet characters
fmt.Println(rands.UpperCase(6))
// Output: DUQPEK

// Returns 6 random number characters
fmt.Println(rands.Number(6))
// Output: 784323
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
    ts: ptr.To(time.Now().Unix()),
}
```

## Stack trace

### Get caller

```go
func hello() {
	fmt.Println(stack.GetCaller())
}

func ExampleGetCaller() {
	hello()
	// Output: github.com/m-mizutani/gots/test_test.ExampleGetCaller
}
```


## License

Apache License 2.0