package tree_sitter_ehir_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_ehir "github.com/meshushkevich/tree-sitter-ehir/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_ehir.Language())
	if language == nil {
		t.Errorf("Error loading EHIR grammar")
	}
}
