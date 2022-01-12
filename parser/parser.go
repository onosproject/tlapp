package parser

//#include "parser.h"
//TSLanguage *tree_sitter_tlapp();
import "C"
import (
	"context"
	"io/ioutil"
	"unsafe"

	sitter "github.com/smacker/go-tree-sitter"
)

func ParseBytes(ctx context.Context, bytes []byte) (*sitter.Node, error) {
	return sitter.ParseCtx(ctx, bytes, GetLanguage())
}

func ParseFile(ctx context.Context, filename string) (*sitter.Node, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ParseBytes(ctx, bytes)
}

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_tlapp())
	return sitter.NewLanguage(ptr)
}
