// Package model
// Created by RTT.
// Author: teocci@yandex.com on 2021-Aug-30
package model

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"hash/fnv"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	gopg "github.com/go-pg/pg/v10"
	"github.com/jackc/pgx/v4"
)

const jinanBenchConfig = "user=jinan password=jinan#db host=localhost port=5432 dbname=jinan_db sslmode=disable"

var (
	setupOnce sync.Once
	pg        *gopg.DB
)

func HashE(s string) (uint32, error) {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

func Setup() *gopg.DB {
	setupOnce.Do(func() {
		config, err := pgxpool.ParseConfig(jinanBenchConfig)
		if err != nil {
			log.Fatalf("extractConfig failed: %v", err)
		}

		pg, err = openPg(*config.ConnConfig)
		if err != nil {
			log.Fatalf("openPq failed: %v", err)
		}
	})

	return pg
}

func openPg(config pgx.ConnConfig) (*gopg.DB, error) {
	var options gopg.Options

	options.Addr = fmt.Sprintf("%s:%d", config.Host, config.Port)
	_, err := os.Stat(config.Host)
	if err == nil {
		options.Network = "unix"
		if !strings.Contains(config.Host, "/.s.PGSQL.") {
			options.Addr = filepath.Join(config.Host, ".s.PGSQL.5432")
		}
	}

	options.User = config.User
	options.Database = config.Database
	options.Password = config.Password
	options.TLSConfig = config.TLSConfig

	options.MaxRetries = 1
	options.MinRetryBackoff = -1

	options.DialTimeout = 30 * time.Second
	options.ReadTimeout = 10 * time.Second
	options.WriteTimeout = 10 * time.Second

	options.PoolSize = 10
	options.MaxConnAge = 10 * time.Second
	options.PoolTimeout = 30 * time.Second
	options.IdleTimeout = time.Second
	options.IdleCheckFrequency = 100 * time.Millisecond

	return gopg.Connect(&options), nil
}
