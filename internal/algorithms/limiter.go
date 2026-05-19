package algorithms
 
import "context"
 
// Limiter is the contract every rate limiting algorithm must satisfy.
// Any algorithm that implements these 3 methods can be used interchangeably.
type Limiter interface {
    // Allow checks if clientID is allowed to make a request.
	    // Returns: allowed bool, remaining int, err error
    Allow(ctx context.Context, clientID string) (bool, int, error)
}
 
// Config holds configuration for a rate limiter instance.
type Config struct {
    Algorithm  string // "fixed_window", "sliding_window", "token_bucket"
    Limit      int    // max requests per window
    WindowSecs int    // window duration in seconds
}