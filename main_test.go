package main

import (
	"encoding/json"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
)

func TestJSONMarshal(t *testing.T) {
	o1 := &JSONInt64Proto{Int64A: 1, Int64B: 1111111111111111111}
	o2 := &JSONInt64{Int64A: 1, Int64B: 1111111111111111111}
	o3 := &JSONInt64StringTag{Int64A: 1, Int64B: 1111111111111111111}
	for _, o := range []interface{}{o1, o2, o3} {
		b, err := json.Marshal(o)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("json.Marshal(<%T>=%+v) = %s", o, o, string(b))

		switch o := o.(type) {
		case *JSONInt64Proto:
			b, err = protojson.Marshal(o)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("protojson.Marshal(<%T>=%+v) = %s", o, o, string(b))
		}
	}
}

func TestJSONUnmarshal1(t *testing.T) {
	o1 := &JSONInt64Proto{Int64A: 1, Int64B: 1111111111111111111}
	o2 := &JSONInt64{}
	o3 := &JSONInt64StringTag{}

	b, err := json.Marshal(o1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("json.Marshal(<%T>%+v) = %s", o1, o1, string(b))

	err = json.Unmarshal(b, o2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("json.Unmarshal(%s, <%T>) = %+v", string(b), o2, o2)

	err = json.Unmarshal(b, o3)
	if err == nil {
		t.Fatalf("json.Unmarshal(%s, <%T>) should fail", string(b), o3)
	}
	t.Logf("json.Unmarshal(%s, <%T>) = %s", string(b), o3, err)
}

func TestJSONUnmarshal2(t *testing.T) {
	o1 := &JSONInt64Proto{Int64A: 1, Int64B: 1111111111111111111}
	o2 := &JSONInt64{}
	o3 := &JSONInt64StringTag{}

	b, err := protojson.Marshal(o1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("protojson.Marshal(<%T>%+v) = %s", o1, o1, string(b))

	err = json.Unmarshal(b, o2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("json.Unmarshal(%s, <%T>) = %+v", string(b), o2, o2)

	err = json.Unmarshal(b, o3)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("json.Unmarshal(%s, <%T>) = %+v", string(b), o3, o3)
}

func TestJSONUnmarshal3(t *testing.T) {
	o1 := &JSONInt64{Int64A: 1, Int64B: 1111111111111111111}
	o2 := &JSONInt64Proto{}

	b, err := json.Marshal(o1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("json.Marshal(<%T>%+v) = %s", o1, o1, string(b))

	err = json.Unmarshal(b, o2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("json.Unmarshal(%s, <%T>) = %+v", string(b), o2, o2)

	o2 = &JSONInt64Proto{}
	err = protojson.Unmarshal(b, o2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("protojson.Unmarshal(%s, <%T>) = %+v", string(b), o2, o2)
}

func TestJSONUnmarshal4(t *testing.T) {
	o1 := &JSONInt64StringTag{Int64A: 1, Int64B: 1111111111111111111}
	o2 := &JSONInt64Proto{}

	b, err := json.Marshal(o1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("json.Marshal(<%T>%+v) = %s", o1, o1, string(b))

	err = json.Unmarshal(b, o2)
	if err == nil {
		t.Fatalf("json.Unmarshal(%s, <%T>) should fail", string(b), o2)
	}
	t.Logf("json.Unmarshal(%s, <%T>) = %s", string(b), o2, err)

	o2 = &JSONInt64Proto{}
	err = protojson.Unmarshal(b, o2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("protojson.Unmarshal(%s, <%T>) = %+v", string(b), o2, o2)
}
