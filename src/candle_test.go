package gotrader

import (
	"reflect"
	"testing"
	"time"
)

func TestTruncateDateTime(t *testing.T) {
	type args struct {
		dateTime time.Time
		duration CandleDuration
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "truncate by minute",
			args: args{
				dateTime: time.Date(2021, time.January, 1, 12, 34, 56, 78, time.UTC),
				duration: MINUTE,
			},
			want: time.Date(2021, time.January, 1, 12, 34, 0, 0, time.UTC),
		},
		{
			name: "truncate by hour",
			args: args{
				dateTime: time.Date(2021, time.January, 1, 12, 34, 56, 78, time.UTC),
				duration: HOUR,
			},
			want: time.Date(2021, time.January, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			name: "truncate by day",
			args: args{
				dateTime: time.Date(2021, time.January, 1, 12, 34, 56, 78, time.UTC),
				duration: DAY,
			},
			want: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TruncateDateTime(tt.args.dateTime, tt.args.duration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruncateDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
