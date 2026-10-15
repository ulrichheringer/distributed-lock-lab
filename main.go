package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:6379", "Redis address")
	key := flag.String("key", "lock-lab", "lock key")
	ttl := flag.Duration("ttl", 2*time.Second, "lock expiration")
	hold := flag.Duration("hold", 4*time.Second, "simulated work duration")
	workers := flag.Int("workers", 3, "number of contenders")
	flag.Parse()
	if *workers < 1 || *ttl <= 0 || *hold <= 0 {
		log.Fatal("workers, ttl, and hold must be positive")
	}
	redis := Redis{Addr: *addr}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	for id := 1; id <= *workers; id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			token, err := newToken()
			if err != nil {
				log.Printf("worker=%d: %v", id, err)
				return
			}
			for ctx.Err() == nil {
				acquired, err := redis.Acquire(ctx, *key, token, *ttl)
				if err != nil {
					log.Printf("worker=%d acquire: %v", id, err)
					return
				}
				if acquired {
					fmt.Printf("worker=%d acquired token=%s\n", id, token[:8])
					time.Sleep(*hold)
					released, err := redis.Release(ctx, *key, token)
					fmt.Printf("worker=%d release=%t err=%v\n", id, released, err)
					return
				}
				time.Sleep(100 * time.Millisecond)
			}
			fmt.Printf("worker=%d timed out\n", id)
		}(id)
	}
	wg.Wait()
}

func newToken() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}
	return hex.EncodeToString(b[:]), nil
}
