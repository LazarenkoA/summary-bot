package storage

import (
	"fmt"
	"github.com/pkg/errors"
	"log/slog"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	pool   *redis.Pool
	logger *slog.Logger
}

func NewRedis(stringConnect string) (*Redis, error) {
	r := &Redis{
		pool:   initPool(stringConnect),
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)).With("name", "redis"),
	}
	return r, r.pool.TestOnBorrow(r.pool.Get(), time.Now())
}

func initPool(stringConnect string) *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: time.Second * 10,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(stringConnect)
			if err != nil {
				return nil, err
			} else {
				return c, err
			}
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return pool
}

func (R *Redis) KeyExists(key string) bool {
	exists, err := redis.Bool(R.pool.Get().Do("EXISTS", key))
	if err != nil {
		R.logger.Error(fmt.Sprintf("Redis. Ошибка при выполнении KeyExists\n%v\n", err))
	}

	return exists
}

func (R *Redis) Keys() []string {
	keys, err := redis.Strings(R.pool.Get().Do("KEYS", "*"))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		R.logger.Error(fmt.Sprintf("Redis. Ошибка при выполнении KEYS. %v\n", err))
	}

	return keys
}

func (R *Redis) Count(key string) int {
	count, err := redis.Int(R.pool.Get().Do("SCARD", key))
	if err != nil {
		R.logger.Error("Redis. Ошибка при выполнении Count")
	}
	return count
}

func (R *Redis) Delete(key string) error {
	_, err := R.pool.Get().Do("DEL", key)
	if err != nil {
		R.logger.Error("Redis. Ошибка при выполнении Delete")
	}
	return err
}

// Установка значения
// ttl - через сколько будет очищено значение (минимальное значение 1 секунда)
func (R *Redis) Set(key, value string, ttl time.Duration) error {
	param := []interface{}{key, value}
	if ttl >= time.Second {
		param = append(param, "EX", ttl.Seconds())
	}

	_, err := R.pool.Get().Do("SET", param...)
	if err != nil {
		R.logger.Error(fmt.Sprintf("Redis. Ошибка при выполнении Set. %v\n", err))
	}
	return err
}

func (R *Redis) Get(key string) (string, error) {
	v, err := redis.String(R.pool.Get().Do("GET", key))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		R.logger.Error(fmt.Sprintf("Redis. Ошибка при выполнении Get. %v\n", err))
	}
	return v, err
}

func (R *Redis) DeleteItems(key, value string) error {
	_, err := R.pool.Get().Do("SREM", key, value)
	if err != nil && !errors.Is(err, redis.ErrNil) {
		R.logger.Error("Redis. Ошибка при выполнении DeleteItems")
	}
	return err
}

func (R *Redis) Items(key string) ([]string, error) {
	items, err := redis.Strings(R.pool.Get().Do("SMEMBERS", key))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		R.logger.Error("Redis. Ошибка при выполнении Items")
		return items, nil
	}
	return items, err
}

func (R *Redis) LPOP(key string) string {
	item, err := redis.String(R.pool.Get().Do("LPOP", key))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		R.logger.Error("Redis. Ошибка при выполнении LPOP")
	}
	return item
}

func (R *Redis) RPUSH(key, value string) error {
	_, err := R.pool.Get().Do("RPUSH", key, value)
	if err != nil && err != redis.ErrNil {
		R.logger.Error("Redis. Ошибка при выполнении RPUSH")
	}
	return err
}

// Добавляет в неупорядоченную коллекцию значение
func (R *Redis) AppendItems(key, value string) error {
	_, err := R.pool.Get().Do("SADD", key, value)
	if err != nil && !errors.Is(err, redis.ErrNil) {
		R.logger.Error("Redis. Ошибка при выполнении AppendItems")
		return nil
	}

	return err
}

func (R *Redis) SetMap(key string, value map[string]string) {
	for k, v := range value {
		_, err := R.pool.Get().Do("HSET", key, k, v)
		if err != nil {
			R.logger.Error("Redis. Ошибка при выполнении SetMap")
			break
		}
	}
}

func (R *Redis) AppendTimeData(key string, t time.Time, data []byte) error {
	_, err := R.pool.Get().Do("ZADD", key, t.Unix(), data)
	if err != nil {
		return errors.Wrap(err, "redis ZADD error")
	}

	return nil
}

func (R *Redis) GetMessageData(key string, tstart, tfinish time.Time) ([][]byte, error) {
	return R.ZRangeByScore(key, tstart.Unix(), tfinish.Unix())
}

func (R *Redis) DeleteMessageDataByTime(key string, t time.Time) error {
	_, err := R.pool.Get().Do("ZREMRANGEBYSCORE", key, t.Unix(), t.Unix())
	if err != nil {
		return errors.Wrap(err, "redis ZREMRANGEBYSCORE error")
	}

	return nil
}

func (R *Redis) DeleteMessageData(key string, t time.Time, data []byte) error {
	_, err := R.pool.Get().Do("ZREM", key, t.Unix(), data)
	if err != nil {
		return errors.Wrap(err, "redis ZREM error")
	}

	return nil
}

func (R *Redis) GetMessageDataForClear(key string, tfinish time.Time) ([][]byte, error) {
	return R.ZRangeByScore(key, 0, tfinish.Unix())
}

func (R *Redis) ZRangeByScore(key string, start, finish int64) ([][]byte, error) {
	data, err := R.pool.Get().Do("ZRANGEBYSCORE", key, start, finish)
	if err != nil {
		return nil, errors.Wrap(err, "redis ZRANGEBYSCORE error")
	}

	result := make([][]byte, len(data.([]interface{})))
	for i, item := range data.([]interface{}) {
		result[i], _ = item.([]byte)
	}

	return result, nil
}

func (R *Redis) StringMap(key string) (map[string]string, error) {
	value, err := redis.StringMap(R.pool.Get().Do("HGETALL", key))
	if err != nil && !errors.Is(err, redis.ErrNil) {
		R.logger.Error(fmt.Sprintf("Redis. Ошибка при выполнении StringMap: %v", err))
	}
	return value, err
}

// Начало транзакции
func (R *Redis) Begin() {
	R.pool.Get().Do("MULTI")
}

// Фиксация транзакции
func (R *Redis) Commit() {
	R.pool.Get().Do("EXEC")
}

// Откат транзакции
func (R *Redis) Rollback() {
	R.pool.Get().Do("DISCARD")
}
