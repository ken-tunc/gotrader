package gotrader

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestOpen(t *testing.T) {
	type args struct {
		dialector gorm.Dialector
		config    *gorm.Config
	}
	tests := []struct {
		name    string
		args    args
		check   func(db *gorm.DB) error
		wantErr bool
	}{
		{
			name: "Open in memory sqlite database",
			args: args{
				dialector: sqlite.Open(":memory:"),
				config:    &gorm.Config{},
			},
			check: func(db *gorm.DB) error {
				sqlDb, err := db.DB()
				if err != nil {
					return err
				}
				return sqlDb.Ping()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OpenDb(tt.args.dialector, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err = tt.check(got); err != nil {
				t.Errorf("Open() check error = %v", err)
			}
		})
	}
}
