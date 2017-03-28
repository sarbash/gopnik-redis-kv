package rediskv

import (
	"fmt"

	"plugins"

	"github.com/go-redis/redis"
	"github.com/op/go-logging"
	json "github.com/orofarne/strict-json"
)

var log = logging.MustGetLogger("global")

type RedisKV struct {
	config *redis.Options
	client *redis.Client
}

type RedisKVFactory struct {
}

func (self *RedisKVFactory) Name() string {
	return "RedisKVPlugin"
}

func (self *RedisKVFactory) New(cfg json.RawMessage) (interface{}, error) {
	var res = new(RedisKV)
	if err := res.Configure(cfg); err != nil {
		return nil, err
	}
	return res, nil
}

func init() {
	plugins.DefaultPluginStore.AddPlugin(new(RedisKVFactory))
}

func (self *RedisKV) Configure(cfg json.RawMessage) error {
	// Unmarshal config
	self.config = new(redis.Options)
	if err := json.Unmarshal(cfg, self.config); err != nil {
		return err
	}

	// Connect
	self.client = redis.NewClient(self.config)
	if _, err := self.client.Ping().Result(); err != nil {
		return fmt.Errorf("Failed to connect to Redis: %v", err)
	}

	return nil
}

func (self *RedisKV) Get(key string) (data []byte, err error) {
	log.Debug("Request data by key '%v' from Redis...", key)
	if data, err = self.client.Get(key).Bytes(); err != nil {
		log.Debug("Key '%v' not found", key)
	}
	return
}

func (self *RedisKV) Set(key string, value []byte) (err error) {
	log.Debug("Save data by key '%v' to Redis", key)
	if err = self.client.Set(key, value, 0).Err(); err != nil {
		return fmt.Errorf("Failed to save data by key '%v' to Redis: %v", key, err)
	}
	return
}

func (self *RedisKV) Delete(key string) (err error) {
	log.Debug("Delete data by key '%v' from Redis", key)
	if err = self.client.Del(key).Err(); err != nil {
		return fmt.Errorf("Failed to delete data by key '%v' from Redis: %v", key, err)
	}
	return
}
