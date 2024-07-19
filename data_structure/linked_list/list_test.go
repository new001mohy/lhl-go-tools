package linked_list

import (
	"testing"
)

func TestList_AddLast(t *testing.T) {
	type fields struct {
		head *Node
		size uint64
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		args args
	}{
		{
			args: args{
				value: 1,
			},
		},
		{
			args: args{
				value: "hello world",
			},
		},
		{
			args: args{
				value: nil,
			},
		},
	}
	l := &List{}
	for _, tt := range tests {
		l.AddLast(tt.args.value)
		t.Logf("list size: %d\n", l.size)
	}
}

func TestList_Find(t *testing.T) {
	l := &List{}
	if index, ok := l.Find(1); !ok {
		t.Logf("index: %d\n", index)
	}
	l.AddLast(2)
	l.AddLast(1)
	l.AddLast(map[int]string{1: "one", 2: "two"})
	if index, ok := l.Find(1); ok {
		t.Logf("index: %d\n", index)
	}
	if index, ok := l.Find(map[int]string{1: "one", 2: "two"}); ok {
		t.Logf("index: %d\n", index)
	}
}

func TestList_Get(t *testing.T) {
	l := &List{}
	l.AddLast(1)
	l.AddLast(2)

	if got := l.Get(0); got != nil {
		t.Log(got)
	} else if got = l.Get(2); got != nil {
		t.Errorf("Get() = %v, want %v", got, nil)

	}
}

func TestList_IsEmpty(t *testing.T) {
	type fields struct {
		head *Node
		size uint64
	}
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "IsEmpty",
			want: true,
		},
		{
			name: "IsNotEmpty",
			want: false,
		},
	}

	l := &List{}
	if got := l.IsEmpty(); got != tests[0].want {
		t.Errorf("IsEmpty() = %v, want %v", got, tests[0].want)
	}
	l.AddLast(1)
	if got := l.IsEmpty(); got != tests[1].want {
		t.Errorf("IsEmpty() = %v, want %v", got, tests[1].want)
	}

}

func TestList_RemoveAt(t *testing.T) {
	type fields struct {
		head *Node
		size uint64
	}
	type args struct {
		index uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			if err := l.RemoveAt(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("RemoveAt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestList_Set(t *testing.T) {
	type fields struct {
		head *Node
		size uint64
	}
	type args struct {
		index uint64
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			if err := l.Set(tt.args.index, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestList_Size(t *testing.T) {
	type fields struct {
		head *Node
		size uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			if got := l.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewList(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"newList"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(); got == nil {
				t.Errorf("NewList() = %v, want %v", got, nil)
			}
		})
	}
}
