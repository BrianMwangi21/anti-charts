package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	analysis "github.com/BrianMwangi21/anti-charts.git/pkg/analysis"
	"github.com/charmbracelet/log"
)

func ValidateInput(input []string) (*analysis.AnalysisRequest, error) {
	log.Debug("Validating input...")

	symbol := strings.ToUpper(input[0])
	if len(symbol) == 0 {
		return nil, errors.New("Symbol entry is invalid")
	}

	duration, err := strconv.Atoi(input[1])
	if err != nil {
		return nil, errors.New("Duration entry is invalid")
	}

	interval := input[2]
	pattern := `^(1m|3m|5m|15m|30m|1h|2h|4h|6h|8h|12h|1d|3d|1w|1M)$`
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, errors.New("Error compiling regex")
	}

	fmt.Println("interval got", interval)
	if !r.MatchString(interval) {
		return nil, errors.New("Interval entry is invalid")
	}

	return &analysis.AnalysisRequest{
		Symbol:   symbol,
		Duration: duration,
		Interval: interval,
	}, nil
}
