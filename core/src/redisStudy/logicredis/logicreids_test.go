package logicredis

import (
	"github.com/go-redis/redis"
	"reflect"
	"testing"
)

func TestGetRedisClient(t *testing.T) {
	tests := []struct {
		name string
		want *redis.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRedisClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRedisClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRedisClient(t *testing.T) {
	type args struct {
		url      string
		db       int
		poolSize int
	}
	tests := []struct {
		name    string
		args    args
		want    *redis.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRedisClient(tt.args.url, tt.args.db, tt.args.poolSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRedisClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedisClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
