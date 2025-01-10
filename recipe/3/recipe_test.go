package recipe_test

import (
	"testing"

	"github.com/qba73/recipe"
)

func BenchmarkCookTurkey(b *testing.B) {
	recipe.CookTurkey()
}
