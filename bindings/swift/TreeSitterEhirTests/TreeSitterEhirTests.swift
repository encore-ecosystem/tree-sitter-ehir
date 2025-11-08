import XCTest
import SwiftTreeSitter
import TreeSitterEhir

final class TreeSitterEhirTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_ehir())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading EHIR grammar")
    }
}
