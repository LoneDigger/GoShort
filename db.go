package main

import (
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	expiration = 7 * 24 * time.Hour
	cleanup    = 8 * time.Hour
)

var Db = &CacheDb{
	db: cache.New(expiration, cleanup),
}

type CacheDb struct {
	db *cache.Cache
}

func (d *CacheDb) Get(path string) (string, bool) {
	val, b := d.db.Get(path)
	if b {
		return val.(string), true
	}
	return "", false
}

func (d *CacheDb) Set(path, url string) {
	d.db.SetDefault(path, url)
}
