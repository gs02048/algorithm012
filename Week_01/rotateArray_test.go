package Week_01

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name:"Rotate",
			args:args{
				nums:[]int{1,2,3,4,5,6,7},
				k:3,
			},
			want:[]int{5,6,7,1,2,3,4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate(tt.args.nums, tt.args.k)
			if tt.args.nums[0] == tt.want[0]{
				fmt.Println("success")
			}else{
				fmt.Println("fail")
			}
		})
	}
}
