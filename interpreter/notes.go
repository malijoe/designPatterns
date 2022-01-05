package interpreter

/*
    Overview:
        Textual input needs to processed
            e.g., turned into linked structures
            AST = Abstract Syntax Tree
        Examples:
            Programming language compilers, interpreters, IDEs
            HTML, XML, and similar
            Numeric expressions (3+4/5)
            Regular expressions
        turning strings into linked structures is a complicated process

        Interpreter: a component that processes structured text data. Does so by turning it into separate lexical tokens (lexing) and then interpreting sequences of said tokens (parsing).

        Summary:
            Barring simple cases, an interpreter acts in two stages
            Lexing turns text into a set of tokens
            Parsing tokens into meaningful constructs
            Parsed data can then be traversed using the Visitor pattern
 */
