/**
 * @file Parser for EHIR
 * @author Maksim Shushkevich <m.e.shushkevich@yandex.com>
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "ehir",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
