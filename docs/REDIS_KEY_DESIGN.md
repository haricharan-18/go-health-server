# Redis Key Design

## Fixed Window
Key: fw:{user}:{window}
Example: fw:alice:window1
Type: String (counter)
Commands: INCR, EXPIRE

## Sliding Window
Key: sw:{user}
Type: Sorted Set
Score: timestamp
Commands: ZADD, ZREMRANGEBYSCORE, ZCARD

## Token Bucket
Key: tb:{user}
Type: String (counter)
Commands: SET, DECR, INCR

## Notes
- Fixed Window uses expiry to reset counters
- Sliding Window removes old requests dynamically
- Token Bucket allows burst traffic with refill