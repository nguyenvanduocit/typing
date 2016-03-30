package typing

import (
	"strings"
	"fmt"
	"time"
)

func Text(text string) {
	sentences := strings.SplitAfter(text, ".")
	for _, sentence := range sentences {
		parts := strings.SplitAfter(sentence, ",")
		for _, part := range parts {
			for _, char := range part {
				fmt.Printf("%c", char)
				time.Sleep(50 * time.Millisecond)
			}
			time.Sleep(100 * time.Millisecond)
		}

		time.Sleep(200 * time.Millisecond)
	}
}
