package usecases

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseJwt(t *testing.T) {
	testDate := "2020-01-01T03:00:00+09:00"
	_, err := parseJstWithRFC3339(testDate)
	if err != nil {
		t.Fatalf("%v", "parse date error."+err.Error())
	}
}

func TestParseRightfulness(t *testing.T) {
	testDate := "2020-01-01T03:00:00+09:00"
	parsed, err := parseJstWithRFC3339(testDate)
	if err != nil {
		t.Fatalf("%v", "parse date error."+err.Error())
	}

	formatDate, _ := time.Parse("2006-01-02T15:04:05Z07:00", testDate)
	assert.Equal(t, formatDate.Format("2006-01-02T15:04:05Z07:00"), parsed.Format("2006-01-02T15:04:05Z07:00"))

	rfcDate, _ := time.Parse(time.RFC3339, testDate)
	assert.Equal(t, rfcDate.Format(time.RFC3339), parsed.Format(time.RFC3339))
}
