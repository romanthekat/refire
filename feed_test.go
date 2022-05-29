package main

import (
	"reflect"
	"testing"
)

func Test_filterFeedByKeywords(t *testing.T) {
	type args struct {
		entries  []Entry
		keywords []string
	}
	tests := []struct {
		name string
		args args
		want []Entry
	}{
		{
			name: "several basic cases for simple keywords detection",
			args: args{
				entries: []Entry{
					{
						Id:    "1",
						Title: "title keyword1 only",
					},
					{
						Id:    "2",
						Title: "title keyword2",
					},
					{
						Id:    "2.1",
						Title: "a keyword - partially similar",
					},
					{
						Id:    "3",
						Title: "title keyword1 and keyword2 at the same time",
					},
					{
						Id:    "4",
						Title: "title keyword3 only",
					},
				},
				keywords: []string{"keyword1", "keyword2"},
			},
			want: []Entry{
				{
					Id:    "1",
					Title: "title keyword1 only",
				},
				{
					Id:    "2",
					Title: "title keyword2",
				},
				{
					Id:    "3",
					Title: "title keyword1 and keyword2 at the same time",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterFeedByKeywords(tt.args.entries, tt.args.keywords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterFeedByKeywords() = %v, want %v", got, tt.want)
			}
		})
	}
}
