package pipefilter

import (
	"reflect"
	"testing"
)

func TestStringSplit(t *testing.T) {
	sf := NewSplitFilter(",")
	resp, err := sf.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	parts, ok := resp.([]string)
	if !ok {
		t.Fatalf("Repsonse type is %T, but the expected type is string", parts)
	}
	if !reflect.DeepEqual(parts, []string{"1", "2", "3"}) {
		t.Errorf("Expected value is {\"1\",\"2\",\"3\"}, but actual is %v", parts)
	}
}

func TestWrongInput(t *testing.T) {
	sf := NewSplitFilter(",")
	_, err := sf.Process(123)
	if err == nil {
		t.Fatal("An error is expected.")
	}
}
