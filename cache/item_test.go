package cache_test

import (
	"fmt"

	"github.com/m-mizutani/gots/cache"
)

func ExampleItem() {
	item1 := cache.NewItem[string]()
	item1.Set("")

	item2 := cache.NewItem[string]()

	fmt.Println(item1.HasCache(), "=>", item1.Get(), ";")
	fmt.Println(item2.HasCache(), "=>", item2.Get(), ";")
	// Output:
	// true =>  ;
	// false =>  ;
}
