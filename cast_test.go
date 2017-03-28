package rediskv

import (
	"gopnik"
	"testing"
)

func TestFilecacheCast(t *testing.T) {
	fp := new(RedisKV)
	var v gopnik.KVStore
	v = fp
	_ = v
}
