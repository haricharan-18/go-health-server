package config_test
 
import (
    "os"
    "testing"
    "github.com/Zartex-the-art/sei-ratelimiter/internal/config"
)
 
func TestConfig_DefaultValues(t *testing.T) {
    os.Unsetenv("REDIS_URL")
    os.Unsetenv("NODE_ID")
    os.Unsetenv("PORT")
 
    cfg := config.Load()
 
    tests := []struct {
        name string
        got  string
        want string
    }{
        {"RedisURL default", cfg.RedisURL, "localhost:6379"},
        {"NodeID default",   cfg.NodeID,   "node-1"},
        {"Port default",     cfg.Port,     "8080"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.got != tt.want {
                t.Errorf("got %q, want %q", tt.got, tt.want)
            }
        })
    }
}
 
func TestConfig_ReadsFromEnvironment(t *testing.T) {
    os.Setenv("REDIS_URL", "test-redis:6379")
    os.Setenv("NODE_ID",   "node-test")
    os.Setenv("PORT",      "9999")
    defer func() {
        os.Unsetenv("REDIS_URL")
        os.Unsetenv("NODE_ID")
        os.Unsetenv("PORT")
    }()
 
    cfg := config.Load()
 
    if cfg.RedisURL != "test-redis:6379" {
        t.Errorf("RedisURL: got %q, want %q", cfg.RedisURL, "test-redis:6379")
    }
    if cfg.NodeID != "node-test" {
        t.Errorf("NodeID: got %q, want %q", cfg.NodeID, "node-test")
    }
    if cfg.Port != "9999" {
        t.Errorf("Port: got %q, want %q", cfg.Port, "9999")
    }
}
