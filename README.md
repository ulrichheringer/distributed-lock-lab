# distributed-lock-lab

A small Redis-backed simulation of workers competing for a lock. Set the lock TTL shorter than the work time to see another worker acquire it, and the original owner fail to release a lock it no longer holds.

Install with go install github.com/ulrichheringer/distributed-lock-lab@latest.
Start Redis with docker compose up -d, then run distributed-lock-lab --ttl 2s --hold 4s --workers 3.
