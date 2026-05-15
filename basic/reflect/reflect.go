package reflect

import (
	"container/list"
	"fmt"
	"reflect"
)

// GetReflect explore unexported function len, and use it to count args.
// output 2
func GetReflect() {
	l := list.New()
	l.PushFront("foo")
	l.PushFront("bar")

	// Get a reflect.Value fv for the unexported field len.
	fv := reflect.ValueOf(l).Elem().FieldByName("len")
	fmt.Println(fv.Int()) // 2
}
