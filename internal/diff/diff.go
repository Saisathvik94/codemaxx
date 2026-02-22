package diff

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func ShowDiff(original, modified string) {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(original, modified, false)

	for _, d := range diffs {
		switch d.Type {
		case diffmatchpatch.DiffInsert:
			fmt.Printf("\033[32m+%s\033[0m", d.Text) // green color for Added lines
		case diffmatchpatch.DiffDelete:
			fmt.Printf("\033[31m-%s\033[0m", d.Text) // Red color for deletedLines lines
		case diffmatchpatch.DiffEqual:
			fmt.Println(d.Text)
		}
	}
}
