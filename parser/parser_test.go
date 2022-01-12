package parser_test

import (
	"context"
	"testing"

	"github.com/onosproject/tlapp/parser"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	module := `
---- MODULE Test ----
EXTENDS foo/bar/Baz
INSTANCE foo/bar/Baz
LOCAL foo == INSTANCE bar/Baz
f(x) == x
====
`
	n, err := sitter.ParseCtx(context.Background(), []byte(module), parser.GetLanguage())
	assert.NoError(t, err)
	assert.Equal(t, "(source_file (module name: (identifier) (extends (module_path (identifier) (identifier) (identifier))) (instance (module_path (identifier) (identifier) (identifier))) (local_definition (module_definition name: (identifier) (def_eq) definition: (instance (module_path (identifier) (identifier))))) (operator_definition name: (identifier) parameter: (identifier) (def_eq) definition: (identifier_ref))))", n.String())
}
