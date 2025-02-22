package php8_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Demooon86/php-parser/internal/php8"
	"github.com/Demooon86/php-parser/pkg/conf"
	"github.com/Demooon86/php-parser/pkg/version"
)

func BenchmarkPhp8(b *testing.B) {
	src, err := os.ReadFile(filepath.Join("testdata", "test.php"))
	if err != nil {
		b.Fatal("can not read testdata/test.php: " + err.Error())
	}

	for n := 0; n < b.N; n++ {
		config := conf.Config{
			Version: &version.Version{
				Major: 8,
				Minor: 8,
			},
		}
		lexer := php8.NewLexer(src, config)
		php8parser := php8.NewParser(lexer, config)
		php8parser.Parse()
	}
}
