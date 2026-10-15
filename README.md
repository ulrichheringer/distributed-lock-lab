# distributed-lock-lab — Redis Lock Simulator

[![Go 1.22](https://img.shields.io/badge/Go-1.22-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Redis](https://img.shields.io/badge/Store-Redis-DC382D?logo=redis&logoColor=white)](https://redis.io/)
[![Docker Compose](https://img.shields.io/badge/Docker%20Compose-Required-2496ED?logo=docker&logoColor=white)](compose.yaml)
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A small Redis-backed simulation of workers competing for a lock. Set the lock TTL shorter than the work time to see another worker acquire it, and the original owner fail to release a lock it no longer holds.

Install with go install github.com/ulrichheringer/distributed-lock-lab@latest.
Start Redis with docker compose up -d, then run distributed-lock-lab --ttl 2s --hold 4s --workers 3.
