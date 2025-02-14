package tools

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	dr, err := ParseDuration("7d")
	fmt.Println(dr, err)
}

func TestTimeDuration(t *testing.T) {
	d := strings.TrimSpace("7d")
	fmt.Println(d)
	dr, err := time.ParseDuration("7d")
	fmt.Println(dr, err)
}
