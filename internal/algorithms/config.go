package algorithms

type Algorithm string

const (
	AlgorithmFixedWindow   Algorithm = "fixed-window"
	AlgorithmSlidingWindow Algorithm = "sliding-window"
	AlgorithmTokenBucket   Algorithm = "token-bucket"
)

type Config struct {
	Algorithm  Algorithm
	Limit      int
	WindowSecs int
}
