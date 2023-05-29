package cache

import (
	"fmt"
	"strings"

	"github.com/imJayanth/go-modules/config"
	"github.com/imJayanth/go-modules/errors"

	"github.com/gomodule/redigo/redis"
)

func RedisGet(appConfig *config.AppConfig, key string) (string, errors.RestAPIError) {
	key = strings.ToLower(key)
	conn := appConfig.RedisConfig.Pool.Get()
	defer conn.Close()

	objStr, err := redis.String(conn.Do("Get", key))
	if err != nil {
		if err.Error() == "redigo: nil returned" {
			return "", errors.NewNotFoundError("Record not found in Redis cache")
		} else {
			return "", errors.NewNotFoundError(err.Error())
		}
	}
	if strings.TrimSpace(objStr) == "" {
		return objStr, errors.NewNotFoundError("Record empty in Redis cache")
	}
	return objStr, errors.NO_ERROR()
}

func RedisSet(appConfig *config.AppConfig, key string, timeoutInSec int, value string) errors.RestAPIError {
	key = strings.ToLower(key)
	conn := appConfig.RedisConfig.Pool.Get()
	defer conn.Close()

	var err error
	if timeoutInSec > 0 {
		_, err = conn.Do("SETEX", key, timeoutInSec, value)
	} else {
		_, err = conn.Do("SET", key, value)
	}
	if err != nil {
		fmt.Println(err)
		return errors.NewNotFoundError(err.Error())
	}
	return errors.NO_ERROR()
}

func RedisGetKeys(appConfig *config.AppConfig, prefix string) ([]string, errors.RestAPIError) {
	prefix = strings.ToLower(prefix)
	conn := appConfig.RedisConfig.Pool.Get()
	defer conn.Close()

	p := "*" + prefix + "*"
	keys, err := redis.Strings(conn.Do("Keys", p))
	if err != nil {
		return keys, errors.NewNotFoundError(err.Error())
	}
	return keys, errors.NO_ERROR()
}

func RedisDeleteKeys(appConfig *config.AppConfig, prefix string) errors.RestAPIError {
	prefix = strings.ToLower(prefix)
	conn := appConfig.RedisConfig.Pool.Get()
	defer conn.Close()
	p := "*"
	if prefix != "" {
		p = "*" + prefix + "*"
	}

	keys, err := redis.Strings(conn.Do("KEYS", p))
	if err != nil {
		return errors.NewNotFoundError(err.Error())
	}

	for _, k := range keys {
		redis.String(conn.Do("DEL", k))
	}
	return errors.NO_ERROR()
}
