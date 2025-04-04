package bundler_tests

import (
	"testing"

	"github.com/matthewmueller/esbuild_internal/config"
)

var importstar_suite = suite{
	name: "importstar",
}

func TestImportStarUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportImportStarUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './bar'
				let foo = 234
				console.log(foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				import * as ns from './foo'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportImportStarNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './bar'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				import * as ns from './foo'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportImportStarCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './bar'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				import * as ns from './foo'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarAsUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './bar'
				let foo = 234
				console.log(foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				export * as ns from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarAsNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './bar'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				export * as ns from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarAsCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './bar'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				export * as ns from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './bar'
				let foo = 234
				console.log(foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				export * from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './bar'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				export * from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './bar'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				export * from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarCommonJSUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(foo)
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarCommonJSCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarCommonJSNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarAndCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				const ns2 = require('./foo')
				console.log(ns.foo, ns2.foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarNoBundleUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarNoBundleCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarNoBundleNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarMangleNoBundleUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			MinifySyntax:  true,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarMangleNoBundleCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			MinifySyntax:  true,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarMangleNoBundleNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			MinifySyntax:  true,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarOmitAmbiguous(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './common'
				console.log(ns)
			`,
			"/common.js": `
				export * from './foo'
				export * from './bar'
			`,
			"/foo.js": `
				export const x = 1
				export const y = 2
			`,
			"/bar.js": `
				export const y = 3
				export const z = 4
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportExportStarAmbiguousError(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {x, y, z} from './common'
				console.log(x, y, z)
			`,
			"/common.js": `
				export * from './foo'
				export * from './bar'
			`,
			"/foo.js": `
				export const x = 1
				export const y = 2
			`,
			"/bar.js": `
				export const y = 3
				export const z = 4
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: ERROR: Ambiguous import "y" has multiple matching exports
foo.js: NOTE: One matching export is here:
bar.js: NOTE: Another matching export is here:
`,
	})
}

func TestImportExportStarAmbiguousWarning(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './common'
				console.log(ns.x, ns.y, ns.z)
			`,
			"/common.js": `
				export * from './foo'
				export * from './bar'
			`,
			"/foo.js": `
				export const x = 1
				export const y = 2
			`,
			"/bar.js": `
				export const y = 3
				export const z = 4
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "y" will always be undefined because there are multiple matching exports
foo.js: NOTE: One matching export is here:
bar.js: NOTE: Another matching export is here:
`,
	})
}

func TestReExportStarNameCollisionNotAmbiguousImport(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {x, y} from './common'
				console.log(x, y)
			`,
			"/common.js": `
				export * from './a'
				export * from './b'
			`,
			"/a.js": `
				export * from './c'
			`,
			"/b.js": `
				export {x} from './c'
			`,
			"/c.js": `
				export let x = 1, y = 2
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarNameCollisionNotAmbiguousExport(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from './a'
				export * from './b'
			`,
			"/a.js": `
				export * from './c'
			`,
			"/b.js": `
				export {x} from './c'
			`,
			"/c.js": `
				export let x = 1, y = 2
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarNameShadowingNotAmbiguous(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {x} from './a'
				console.log(x)
			`,
			"/a.js": `
				export * from './b'
				export let x = 1
			`,
			"/b.js": `
				export let x = 2
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarNameShadowingNotAmbiguousReExport(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {x} from './a'
				console.log(x)
			`,
			"/a.js": `
				export * from './b'
			`,
			"/b.js": `
				export * from './c'
				export let x = 1
			`,
			"/c.js": `
				export let x = 2
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarOfExportStarAs(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as foo_ns from './foo'
				console.log(foo_ns)
			`,
			"/foo.js": `
				export * as bar_ns from './bar'
			`,
			"/bar.js": `
				export const bar = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportOfExportStar(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {bar} from './foo'
				console.log(bar)
			`,
			"/foo.js": `
				export * from './bar'
			`,
			"/bar.js": `
				// Add some statements to increase the part index (this reproduced a crash)
				statement()
				statement()
				statement()
				statement()
				export const bar = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportOfExportStarOfImport(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {bar} from './foo'
				console.log(bar)
			`,
			"/foo.js": `
				// Add some statements to increase the part index (this reproduced a crash)
				statement()
				statement()
				statement()
				statement()
				export * from './bar'
			`,
			"/bar.js": `
				export {value as bar} from './baz'
			`,
			"/baz.js": `
				export const value = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfIIFE(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				export * from './entry'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatIIFE,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfIIFEWithName(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				export * from './entry'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatIIFE,
			AbsOutputFile: "/out.js",
			GlobalName:    []string{"someName"},
		},
	})
}

func TestExportSelfES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				export * from './entry'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				export * from './entry'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfCommonJSMinified(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				module.exports = {foo: 123}
				console.log(require('./entry'))
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:              config.ModeBundle,
			MinifyIdentifiers: true,
			OutputFormat:      config.FormatCommonJS,
			AbsOutputFile:     "/out.js",
		},
	})
}

func TestImportSelfCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				exports.foo = 123
				import {foo} from './entry'
				console.log(foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfAsNamespaceES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				export * as ns from './entry'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportExportSelfAsNamespaceES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				import * as ns from './entry'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportOtherFileExportSelfAsNamespaceES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from './foo'
			`,
			"/foo.js": `
				export const foo = 123
				export * as ns from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportOtherFileImportExportSelfAsNamespaceES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from './foo'
			`,
			"/foo.js": `
				export const foo = 123
				import * as ns from './foo'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestOtherFileExportSelfAsNamespaceUnusedES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export {foo} from './foo'
			`,
			"/foo.js": `
				export const foo = 123
				export * as ns from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestOtherFileImportExportSelfAsNamespaceUnusedES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export {foo} from './foo'
			`,
			"/foo.js": `
				export const foo = 123
				import * as ns from './foo'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfAsNamespaceCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				export * as ns from './entry'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfAndRequireSelfCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				console.log(require('./entry'))
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfAndImportSelfCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as x from './entry'
				export const foo = 123
				console.log(x)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportOtherAsNamespaceCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as ns from './foo'
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportExportOtherAsNamespaceCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				export {ns}
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestNamespaceImportMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns, ns.foo)
			`,
			"/foo.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "foo" will always be undefined because there is no matching export in "foo.js"
`,
	})
}

func TestExportOtherCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export {bar} from './foo'
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportOtherNestedCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export {y} from './bar'
			`,
			"/bar.js": `
				export {x as y} from './foo'
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestNamespaceImportUnusedMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns.foo)
			`,
			"/foo.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "foo" will always be undefined because there is no matching export in "foo.js"
`,
	})
}

func TestNamespaceImportMissingCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns, ns.foo)
			`,
			"/foo.js": `
				exports.x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestNamespaceImportUnusedMissingCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns.foo)
			`,
			"/foo.js": `
				exports.x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportNamespaceImportMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './foo'
				console.log(ns, ns.foo)
			`,
			"/foo.js": `
				export * as ns from './bar'
			`,
			"/bar.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportNamespaceImportUnusedMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './foo'
				console.log(ns.foo)
			`,
			"/foo.js": `
				export * as ns from './bar'
			`,
			"/bar.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestNamespaceImportReExportMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns, ns.foo)
			`,
			"/foo.js": `
				export {foo} from './bar'
			`,
			"/bar.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `foo.js: ERROR: No matching export in "bar.js" for import "foo"
foo.js: ERROR: No matching export in "bar.js" for import "foo"
`,
	})
}

func TestNamespaceImportReExportUnusedMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns.foo)
			`,
			"/foo.js": `
				export {foo} from './bar'
			`,
			"/bar.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `foo.js: ERROR: No matching export in "bar.js" for import "foo"
foo.js: ERROR: No matching export in "bar.js" for import "foo"
`,
	})
}

func TestNamespaceImportReExportStarMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns, ns.foo)
			`,
			"/foo.js": `
				export * from './bar'
			`,
			"/bar.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "foo" will always be undefined because there is no matching export in "foo.js"
`,
	})
}

func TestNamespaceImportReExportStarUnusedMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns.foo)
			`,
			"/foo.js": `
				export * from './bar'
			`,
			"/bar.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "foo" will always be undefined because there is no matching export in "foo.js"
`,
	})
}

func TestExportStarDefaultExportCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from './foo'
			`,
			"/foo.js": `
				export default 'default' // This should not be picked up
				export let foo = 'foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestIssue176(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as things from './folders'
				console.log(JSON.stringify(things))
			`,
			"/folders/index.js": `
				export * from "./child"
			`,
			"/folders/child/index.js": `
				export { foo } from './foo'
			`,
			"/folders/child/foo.js": `
				export const foo = () => 'hi there'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarExternalIIFE(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatIIFE,
			AbsOutputFile: "/out.js",
			GlobalName:    []string{"mod"},
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"foo": true,
				}},
			},
		},
	})
}

func TestReExportStarExternalES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"foo": true,
				}},
			},
		},
	})
}

func TestReExportStarExternalCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"foo": true,
				}},
			},
		},
	})
}

func TestReExportStarIIFENoBundle(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeConvertFormat,
			OutputFormat:  config.FormatIIFE,
			AbsOutputFile: "/out.js",
			GlobalName:    []string{"mod"},
		},
	})
}

func TestReExportStarES6NoBundle(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeConvertFormat,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarCommonJSNoBundle(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeConvertFormat,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarAsExternalIIFE(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as out from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatIIFE,
			AbsOutputFile: "/out.js",
			GlobalName:    []string{"mod"},
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"foo": true,
				}},
			},
		},
	})
}

func TestReExportStarAsExternalES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as out from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"foo": true,
				}},
			},
		},
	})
}

func TestReExportStarAsExternalCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as out from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"foo": true,
				}},
			},
		},
	})
}

func TestReExportStarAsIIFENoBundle(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as out from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeConvertFormat,
			OutputFormat:  config.FormatIIFE,
			AbsOutputFile: "/out.js",
			GlobalName:    []string{"mod"},
		},
	})
}

func TestReExportStarAsES6NoBundle(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as out from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeConvertFormat,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarAsCommonJSNoBundle(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as out from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeConvertFormat,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportDefaultNamespaceComboIssue446(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/external-default2.js": `
				import def, {default as default2} from 'external'
				console.log(def, default2)
			`,
			"/external-ns.js": `
				import def, * as ns from 'external'
				console.log(def, ns)
			`,
			"/external-ns-default.js": `
				import def, * as ns from 'external'
				console.log(def, ns, ns.default)
			`,
			"/external-ns-def.js": `
				import def, * as ns from 'external'
				console.log(def, ns, ns.def)
			`,
			"/external-default.js": `
				import def, * as ns from 'external'
				console.log(def, ns.default)
			`,
			"/external-def.js": `
				import def, * as ns from 'external'
				console.log(def, ns.def)
			`,
			"/internal-default2.js": `
				import def, {default as default2} from './internal'
				console.log(def, default2)
			`,
			"/internal-ns.js": `
				import def, * as ns from './internal'
				console.log(def, ns)
			`,
			"/internal-ns-default.js": `
				import def, * as ns from './internal'
				console.log(def, ns, ns.default)
			`,
			"/internal-ns-def.js": `
				import def, * as ns from './internal'
				console.log(def, ns, ns.def)
			`,
			"/internal-default.js": `
				import def, * as ns from './internal'
				console.log(def, ns.default)
			`,
			"/internal-def.js": `
				import def, * as ns from './internal'
				console.log(def, ns.def)
			`,
			"/internal.js": `
				export default 123
			`,
		},
		entryPaths: []string{
			"/external-default2.js",
			"/external-ns.js",
			"/external-ns-default.js",
			"/external-ns-def.js",
			"/external-default.js",
			"/external-def.js",
			"/internal-default2.js",
			"/internal-ns.js",
			"/internal-ns-default.js",
			"/internal-ns-def.js",
			"/internal-default.js",
			"/internal-def.js",
		},
		options: config.Options{
			Mode:         config.ModeBundle,
			AbsOutputDir: "/out",
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"external": true,
				}},
			},
		},
		expectedCompileLog: `internal-def.js: WARNING: Import "def" will always be undefined because there is no matching export in "internal.js"
internal-ns-def.js: WARNING: Import "def" will always be undefined because there is no matching export in "internal.js"
`,
	})
}

func TestImportDefaultNamespaceComboNoDefault(t *testing.T) {
	// Note: "entry-dead.js" checks that this warning doesn't happen for dead code
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry-default-ns-prop.js": `import def, * as ns from './foo'; console.log(def, ns, ns.default)`,
			"/entry-default-ns.js":      `import def, * as ns from './foo'; console.log(def, ns)`,
			"/entry-default-prop.js":    `import def, * as ns from './foo'; console.log(def, ns.default)`,
			"/entry-default.js":         `import def from './foo'; console.log(def)`,
			"/entry-prop.js":            `import * as ns from './foo'; console.log(ns.default)`,
			"/entry-dead.js":            `import * as ns from './foo'; 0 && console.log(ns.default)`,
			"/entry-typo.js":            `import * as ns from './foo'; console.log(ns.buton)`,
			"/entry-typo-indirect.js":   `import * as ns from './indirect'; console.log(ns.buton)`,
			"/foo.js":                   `export let button = {}`,
			"/indirect.js":              `export * from './foo'`,
		},
		entryPaths: []string{
			"/entry-default-ns-prop.js",
			"/entry-default-ns.js",
			"/entry-default-prop.js",
			"/entry-default.js",
			"/entry-prop.js",
			"/entry-dead.js",
			"/entry-typo.js",
			"/entry-typo-indirect.js",
		},
		options: config.Options{
			Mode:         config.ModeBundle,
			AbsOutputDir: "/out",
		},
		expectedCompileLog: `entry-default-ns-prop.js: ERROR: No matching export in "foo.js" for import "default"
entry-default-ns-prop.js: WARNING: Import "default" will always be undefined because there is no matching export in "foo.js"
entry-default-ns.js: ERROR: No matching export in "foo.js" for import "default"
entry-default-prop.js: ERROR: No matching export in "foo.js" for import "default"
entry-default-prop.js: WARNING: Import "default" will always be undefined because there is no matching export in "foo.js"
entry-default.js: ERROR: No matching export in "foo.js" for import "default"
entry-prop.js: WARNING: Import "default" will always be undefined because there is no matching export in "foo.js"
entry-typo-indirect.js: WARNING: Import "buton" will always be undefined because there is no matching export in "indirect.js"
foo.js: NOTE: Did you mean to import "button" instead?
entry-typo.js: WARNING: Import "buton" will always be undefined because there is no matching export in "foo.js"
foo.js: NOTE: Did you mean to import "button" instead?
`,
	})
}

func TestImportNamespaceUndefinedPropertyEmptyFile(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry-nope.js": `
				import * as js from './empty.js'
				import * as mjs from './empty.mjs'
				import * as cjs from './empty.cjs'
				console.log(
					js.nope,
					mjs.nope,
					cjs.nope,
				)
			`,

			// Note: For CommonJS-style modules, we automatically assign the exports
			// object to the "default" property if there is no property named "default".
			// This is for compatibility with node. So this test intentionally behaves
			// differently from the test above.
			"/entry-default.js": `
				import * as js from './empty.js'
				import * as mjs from './empty.mjs'
				import * as cjs from './empty.cjs'
				console.log(
					js.default,
					mjs.default,
					cjs.default,
				)
			`,

			"/empty.js":  ``,
			"/empty.mjs": ``,
			"/empty.cjs": ``,
		},
		entryPaths: []string{
			"/entry-nope.js",
			"/entry-default.js",
		},
		options: config.Options{
			Mode:         config.ModeBundle,
			AbsOutputDir: "/out",
		},
		expectedCompileLog: `entry-default.js: WARNING: Import "default" will always be undefined because there is no matching export in "empty.mjs"
entry-nope.js: WARNING: Import "nope" will always be undefined because the file "empty.js" has no exports
entry-nope.js: WARNING: Import "nope" will always be undefined because the file "empty.mjs" has no exports
entry-nope.js: WARNING: Import "nope" will always be undefined because the file "empty.cjs" has no exports
`,
	})
}

func TestImportNamespaceUndefinedPropertySideEffectFreeFile(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry-nope.js": `
				import * as js from './foo/no-side-effects.js'
				import * as mjs from './foo/no-side-effects.mjs'
				import * as cjs from './foo/no-side-effects.cjs'
				console.log(
					js.nope,
					mjs.nope,
					cjs.nope,
				)
			`,

			// Note: For CommonJS-style modules, we automatically assign the exports
			// object to the "default" property if there is no property named "default".
			// This is for compatibility with node. So this test intentionally behaves
			// differently from the test above.
			"/entry-default.js": `
				import * as js from './foo/no-side-effects.js'
				import * as mjs from './foo/no-side-effects.mjs'
				import * as cjs from './foo/no-side-effects.cjs'
				console.log(
					js.default,
					mjs.default,
					cjs.default,
				)
			`,

			"/foo/package.json":        `{ "sideEffects": false }`,
			"/foo/no-side-effects.js":  `console.log('js')`,
			"/foo/no-side-effects.mjs": `console.log('mjs')`,
			"/foo/no-side-effects.cjs": `console.log('cjs')`,
		},
		entryPaths: []string{
			"/entry-nope.js",
			"/entry-default.js",
		},
		options: config.Options{
			Mode:         config.ModeBundle,
			AbsOutputDir: "/out",
		},
		expectedCompileLog: `entry-default.js: WARNING: Import "default" will always be undefined because there is no matching export in "foo/no-side-effects.mjs"
entry-nope.js: WARNING: Import "nope" will always be undefined because the file "foo/no-side-effects.js" has no exports
entry-nope.js: WARNING: Import "nope" will always be undefined because the file "foo/no-side-effects.mjs" has no exports
entry-nope.js: WARNING: Import "nope" will always be undefined because the file "foo/no-side-effects.cjs" has no exports
`,
	})
}

// Failure case due to a bug in https://github.com/evanw/esbuild/pull/2059
func TestReExportStarEntryPointAndInnerFile(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from 'a'
				import * as inner from './inner.js'
				export { inner }
			`,
			"/inner.js": `
				export * from 'b'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:         config.ModeBundle,
			AbsOutputDir: "/out",
			OutputFormat: config.FormatCommonJS,
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{
					Exact: map[string]bool{
						"a": true,
						"b": true,
					},
				},
			},
		},
	})
}
