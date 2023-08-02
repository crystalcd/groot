package repository_test

import (
	"reflect"
	"testing"

	"github.com/crystal/groot/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func TestBuildUpdate(t *testing.T) {
	// o := Order{
	// 	ID: 1,
	// 	User: User{
	// 		Name: "John",
	// 		Age:  30,
	// 	},

	// 	Items: []string{"item1", "item2"},
	// }
	o1 := Order{
		ID: 1,
		User: User{
			Name: "John",
			Age:  30,
		},
		UserPtr: &User{
			Name: "Test Ptr",
		},
		Items: []string{"item1", "item2"},
		i:     10,
	}
	// expected1 := bson.M{
	// 	"ID": 1,
	// 	"User": bson.M{
	// 		"Name": "John",
	// 		"Age":  30,
	// 	},
	// 	"Items": []string{"item1", "item2"},
	// }
	expected2 := bson.M{
		"ID": 1,
		"User": bson.M{
			"Name": "John",
			"Age":  30,
		},
		"UserPtr": bson.M{
			"Name": "Test Ptr",
		},
		"Items": []string{"item1", "item2"},
	}
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want bson.M
	}{
		// {
		// 	name: "normal",
		// 	args: args{
		// 		i: o,
		// 	},
		// 	want: expected1,
		// },
		{
			name: "ptr",
			args: args{
				i: o1,
			},
			want: expected2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.BuildUpdate(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
