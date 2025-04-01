package runtime_test

import (
	"testing"

	"github.com/matthewmueller/esbuild_internal/compat"
	"github.com/matthewmueller/esbuild_internal/config"
	"github.com/matthewmueller/esbuild_internal/js_parser"
	"github.com/matthewmueller/esbuild_internal/logger"
	"github.com/matthewmueller/esbuild_internal/runtime"
)

func TestUnsupportedFeatures(t *testing.T) {
	for key, feature := range compat.StringToJSFeature {
		t.Run(key, func(t *testing.T) {
			source := runtime.Source(feature)
			log := logger.NewDeferLog(logger.DeferLogAll, nil)

			js_parser.Parse(log, source, js_parser.OptionsFromConfig(&config.Options{
				UnsupportedJSFeatures: feature,
				TreeShaking:           true,
			}))

			if log.HasErrors() {
				msgs := "Internal error: failed to parse runtime:\n"
				for _, msg := range log.Done() {
					msgs += msg.String(logger.OutputOptions{IncludeSource: true}, logger.TerminalInfo{})
				}
				t.Fatal(msgs[:len(msgs)-1])
			}
		})
	}
}
