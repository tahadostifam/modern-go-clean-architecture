package utils_test

import (
	"testing"

	"github.com/tahadostifam/modern-go-clean-architecture/utils"
)

func TestEncodeToString(t *testing.T) {
	max := 6

	result := utils.EncodeToString(max)

	// checking len of generated code
	if len(result) != max {
		t.Errorf("Expected length of result to be %d, but got %d", max, len(result))
	}

	// checking members of generated code
	for _, char := range result {
		found := false
		for _, validChar := range utils.Table {
			if char == rune(validChar) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Unexpected character in result: %c", char)
		}
	}
}
