package config

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	conn := "username:password@tcp(mysql:3306)/todo_db?charset=utf8mb4&parseTime=true"
	port := "5000"
	t.Setenv("MYSQL_CONN", conn)
	t.Setenv("PORT", port)

	tests := []struct {
		name string
		want Config
	}{
		// TODO: Add test cases.
		{
			"normal case",
			Config{MySQLConn: conn, Port: port},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
