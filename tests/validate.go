package tests

import (
	"regexp"
	"strings"
	"testing"
	"time"
)

// does allow whitespace in end of rows
func validateGetMessages(body string, t *testing.T) {
	rows := strings.Split(body, "\n")
	for _, row := range rows {
		parts := strings.Split(row, " ")
		for i, part := range parts {
			if part == "" {
				continue
			}

			switch i {
			case 0:
				ValidateTimestamp(part, t)
			case 1:
				ValidateTopic(part, t)
			case 2:
				ValidateTopicName(part, t)
			default:
				ValidateIsWhitespace(part, t)
			}
		}
	}
}
func ValidateTimestamp(s string, t *testing.T) {
	_, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		t.Error("Invalid date format in response: ", s)
	}
}

func ValidateTopic(s string, t *testing.T) {
	if s != "Topic" {
		t.Error("Invalid Topic format in response: ", s)
	}
}

func ValidateTopicName(s string, t *testing.T) {
	if s != "my.i" && s != "my.o" {
		t.Error("Invalid Topic name in response: ", s)
	}
}

func ValidateIsWhitespace(s string, t *testing.T) {
	_, err := regexp.Match(`\d`, []byte(s))
	if err != nil {
		t.Error("Invalid row format, extra ouput found in the end of row")
	}
}
