package diff

import (
	"fmt"

	"github.com/Saisathvik94/codemaxx/internal/ui/colors"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func ShowDiff(original, modified string) {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(original, modified, false)

	for _, d := range diffs {
		switch d.Type {
		case diffmatchpatch.DiffInsert:
			fmt.Println(colors.SuccessStyle.Render("+" + d.Text)) // green color for Added lines
		case diffmatchpatch.DiffDelete:
			fmt.Println(colors.ErrorStyle.Render("-" + d.Text)) // Red color for deletedLines lines
		case diffmatchpatch.DiffEqual:
			fmt.Println(d.Text)
		}
	}
}
