package sdmscim

import (
	"reflect"
	"testing"
)

type testFixture func()

func executeTests(t *testing.T, testGroupType reflect.Type, beforeEach testFixture, afterEach testFixture) {
	testGroup := reflect.New(testGroupType).Elem().Interface()
	for i := 0; i < testGroupType.NumMethod(); i++ {
		m := testGroupType.Method(i)
		t.Run(m.Name, func(t *testing.T) {
			if beforeEach != nil {
				beforeEach()
			}

			in := []reflect.Value{reflect.ValueOf(testGroup), reflect.ValueOf(t)}
			m.Func.Call(in)

			if afterEach != nil {
				afterEach()
			}
		})
	}
}
