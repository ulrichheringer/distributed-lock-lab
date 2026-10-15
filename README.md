# distributed-lock-lab

Run concurrent lock contenders against Redis and observe TTL expiration and owner-safe release.

Install with go install github.com/ulrichheringer/distributed-lock-lab@latest. Start Redis with docker compose up -d, then run distributed-lock-lab --ttl 2s --hold 4s --workers 3; contenders retry until they acquire the lock or time out.
