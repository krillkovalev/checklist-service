// To do найти панику при записи в кеш
package models

import (
	"context"
	"fmt"
	"time"
	"github.com/redis/go-redis/v9"
	"reflect"
)

const (
	SetTimeExp = 300 * time.Minute
	KeyFormat = "task:id:%d"
)

func ToRedisSet(ctx context.Context, rdb *redis.Client, key string, task *Task) error {
	val := reflect.ValueOf(task).Elem()

	setter := func(p redis.Pipeliner) error {
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			tag := field.Tag.Get("redis")

			if err := p.HSet(ctx, key, tag, val.Field(i).Interface()).Err(); err != nil {
				return err
			}
		}

		if err := p.Expire(ctx, key, SetTimeExp).Err(); err != nil {
			return err
		}
		return nil
	}

	if _, err := rdb.Pipelined(ctx, setter); err != nil {
		return err
	}

	return nil

}

func DeleteFromCache(ctx context.Context, rdb *redis.Client, id int) error{
	key := fmt.Sprintf(KeyFormat, id)
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

