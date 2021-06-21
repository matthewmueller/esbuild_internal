package js_parser

import (
	"fmt"

	"github.com/matthewmueller/esbuild_internal/helpers"
	"github.com/matthewmueller/esbuild_internal/js_ast"
	"github.com/matthewmueller/esbuild_internal/js_lexer"
	"github.com/matthewmueller/esbuild_internal/logger"
)

type jsonParser struct {
	log                            logger.Log
	source                         logger.Source
	tracker                        logger.LineColumnTracker
	lexer                          js_lexer.Lexer
	options                        JSONOptions
	suppressWarningsAboutWeirdCode bool
}

func (p *jsonParser) parseMaybeTrailingComma(closeToken js_lexer.T) bool {
	commaRange := p.lexer.Range()
	p.lexer.Expect(js_lexer.TComma)

	if p.lexer.Token == closeToken {
		if !p.options.AllowTrailingCommas {
			p.log.AddRangeError(&p.tracker, commaRange, "JSON does not support trailing commas")
		}
		return false
	}

	return true
}

func (p *jsonParser) parseExpr() js_ast.Expr {
	loc := p.lexer.Loc()

	switch p.lexer.Token {
	case js_lexer.TFalse:
		p.lexer.Next()
		return js_ast.Expr{Loc: loc, Data: &js_ast.EBoolean{Value: false}}

	case js_lexer.TTrue:
		p.lexer.Next()
		return js_ast.Expr{Loc: loc, Data: &js_ast.EBoolean{Value: true}}

	case js_lexer.TNull:
		p.lexer.Next()
		return js_ast.Expr{Loc: loc, Data: js_ast.ENullShared}

	case js_lexer.TStringLiteral:
		value := p.lexer.StringLiteral()
		p.lexer.Next()
		return js_ast.Expr{Loc: loc, Data: &js_ast.EString{Value: value}}

	case js_lexer.TNumericLiteral:
		value := p.lexer.Number
		p.lexer.Next()
		return js_ast.Expr{Loc: loc, Data: &js_ast.ENumber{Value: value}}

	case js_lexer.TMinus:
		p.lexer.Next()
		value := p.lexer.Number
		p.lexer.Expect(js_lexer.TNumericLiteral)
		return js_ast.Expr{Loc: loc, Data: &js_ast.ENumber{Value: -value}}

	case js_lexer.TOpenBracket:
		p.lexer.Next()
		isSingleLine := !p.lexer.HasNewlineBefore
		items := []js_ast.Expr{}

		for p.lexer.Token != js_lexer.TCloseBracket {
			if len(items) > 0 {
				if p.lexer.HasNewlineBefore {
					isSingleLine = false
				}
				if !p.parseMaybeTrailingComma(js_lexer.TCloseBracket) {
					break
				}
				if p.lexer.HasNewlineBefore {
					isSingleLine = false
				}
			}

			item := p.parseExpr()
			items = append(items, item)
		}

		if p.lexer.HasNewlineBefore {
			isSingleLine = false
		}
		p.lexer.Expect(js_lexer.TCloseBracket)
		return js_ast.Expr{Loc: loc, Data: &js_ast.EArray{
			Items:        items,
			IsSingleLine: isSingleLine,
		}}

	case js_lexer.TOpenBrace:
		p.lexer.Next()
		isSingleLine := !p.lexer.HasNewlineBefore
		properties := []js_ast.Property{}
		duplicates := make(map[string]logger.Range)

		for p.lexer.Token != js_lexer.TCloseBrace {
			if len(properties) > 0 {
				if p.lexer.HasNewlineBefore {
					isSingleLine = false
				}
				if !p.parseMaybeTrailingComma(js_lexer.TCloseBrace) {
					break
				}
				if p.lexer.HasNewlineBefore {
					isSingleLine = false
				}
			}

			keyString := p.lexer.StringLiteral()
			keyRange := p.lexer.Range()
			key := js_ast.Expr{Loc: keyRange.Loc, Data: &js_ast.EString{Value: keyString}}
			p.lexer.Expect(js_lexer.TStringLiteral)

			// Warn about duplicate keys
			if !p.suppressWarningsAboutWeirdCode {
				keyText := js_lexer.UTF16ToString(keyString)
				if prevRange, ok := duplicates[keyText]; ok {
					p.log.AddRangeWarningWithNotes(&p.tracker, keyRange, fmt.Sprintf("Duplicate key %q in object literal", keyText),
						[]logger.MsgData{logger.RangeData(&p.tracker, prevRange, fmt.Sprintf("The original %q is here", keyText))})
				} else {
					duplicates[keyText] = keyRange
				}
			}

			p.lexer.Expect(js_lexer.TColon)
			value := p.parseExpr()

			property := js_ast.Property{
				Kind:       js_ast.PropertyNormal,
				Key:        key,
				ValueOrNil: value,
			}
			properties = append(properties, property)
		}

		if p.lexer.HasNewlineBefore {
			isSingleLine = false
		}
		p.lexer.Expect(js_lexer.TCloseBrace)
		return js_ast.Expr{Loc: loc, Data: &js_ast.EObject{
			Properties:   properties,
			IsSingleLine: isSingleLine,
		}}

	default:
		p.lexer.Unexpected()
		return js_ast.Expr{}
	}
}

type JSONOptions struct {
	AllowComments       bool
	AllowTrailingCommas bool
}

func ParseJSON(log logger.Log, source logger.Source, options JSONOptions) (result js_ast.Expr, ok bool) {
	ok = true
	defer func() {
		r := recover()
		if _, isLexerPanic := r.(js_lexer.LexerPanic); isLexerPanic {
			ok = false
		} else if r != nil {
			panic(r)
		}
	}()

	p := &jsonParser{
		log:                            log,
		source:                         source,
		tracker:                        logger.MakeLineColumnTracker(&source),
		options:                        options,
		lexer:                          js_lexer.NewLexerJSON(log, source, options.AllowComments),
		suppressWarningsAboutWeirdCode: helpers.IsInsideNodeModules(source.KeyPath.Text),
	}

	result = p.parseExpr()
	p.lexer.Expect(js_lexer.TEndOfFile)
	return
}