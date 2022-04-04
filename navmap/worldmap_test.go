package navmap

import (
	"testing"
)

func TestWorldMap_Idx(t *testing.T) {
	type args struct {
		x int
		y int
	}
	wMap := MakeMap(10, 10)
	tests := []struct {
		name string
		wm   *WorldMap
		args args
		want int
	}{
		// TODO: Add test cases.
		{"a", &wMap, args{1, 0}, 1},
		{"b", &wMap, args{0, 1}, 10},
		{"c", &wMap, args{1, 1}, 11},
		{"d", &wMap, args{1, 2}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wm.Idx(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("WorldMap.Idx() = %v, want %v", got, tt.want)
			}
		})
	}
}
