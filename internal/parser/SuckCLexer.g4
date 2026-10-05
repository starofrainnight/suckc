// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar SuckCLexer;

CharacterLiteral: ('u' | 'U' | 'L')? '\'' Cchar+ '\'';

FloatingLiteral:
    FractionalConstant ExponentPart? FloatingSuffix?
    | DigitSequence ExponentPart FloatingSuffix?
;

StringLiteral: EncodingPrefix? (RawString | '"' Schar* '"');

BooleanLiteral: False_ | True_;

PointerLiteral: Nullptr;

// C-scope region markers. MUST stay before MultiLineMacro and Directive:
// all three match the same '#' prefix, and for equal-length matches ANTLR
// picks the rule declared first. Lexer rule names must start uppercase.
//
// The trailing `[ \t]* '\r'?` is required, not cosmetic. Directive matches
// '#' ~[\n]*, which swallows any trailing blanks and the CR of a CRLF line,
// so a marker written as "#pragma c_scope_begin " or with CRLF endings would
// lose the longest-match tie and silently degrade to a hidden Directive --
// the region would never open. Real trailing content still loses to
// Directive ("#pragma c_scope_begin x" is matched further by Directive) and
// stays an ordinary directive.
CScopeBegin: '#pragma' [ \t]+ 'c_scope_begin' [ \t]* '\r'?;

CScopeEnd: '#pragma' [ \t]+ 'c_scope_end' [ \t]* '\r'?;

MultiLineMacro: '#' (~[\n]*? '\\' '\r'? '\n')+ ~ [\n]+ -> channel (HIDDEN);

Directive: '#' ~ [\n]* -> channel (HIDDEN);

/*Keywords*/

Alignas: 'alignas';

Alignof: 'alignof';

Asm: 'asm';

Auto: 'auto';

Bool: 'bool';

Break: 'break';

Case: 'case';

Catch: 'catch';

Char: 'char';

Char16: 'char16_t';

Char32: 'char32_t';

Const: 'const';

Constexpr: 'constexpr';

Const_cast: 'const_cast';

Continue: 'continue';

Decltype: 'decltype';

Default: 'default';

Delete: 'delete';

Do: 'do';

Double: 'double';

Dynamic_cast: 'dynamic_cast';

Else: 'else';

Enum: 'enum';

Explicit: 'explicit';

Export: 'export';

Extern: 'extern';

//DO NOT RENAME - PYTHON NEEDS True and False
False_: 'false';

Final: 'final';

Float: 'float';

For: 'for';

Friend: 'friend';

Goto: 'goto';

If: 'if';

Inline: 'inline';

Int: 'int';

Long: 'long';

Mutable: 'mutable';

Namespace: 'namespace';

New: 'new';

Noexcept: 'noexcept';

Nullptr: 'nullptr';

Operator: 'operator';

Override: 'override';

Private: 'private';

Protected: 'protected';

Public: 'public';

Register: 'register';

Reinterpret_cast: 'reinterpret_cast';

Return: 'return';

Short: 'short';

Signed: 'signed';

Sizeof: 'sizeof';

Static: 'static';

Static_assert: 'static_assert';

Static_cast: 'static_cast';

Struct: 'struct';

Switch: 'switch';

Template: 'template';

This: 'this';

Thread_local: 'thread_local';

Throw: 'throw';

//DO NOT RENAME - PYTHON NEEDS True and False
True_: 'true';

Try: 'try';

Typedef: 'typedef';

Typeid_: 'typeid';

Typename_: 'typename';

Union: 'union';

Unsigned: 'unsigned';

Using: 'using';

Virtual: 'virtual';

Void: 'void';

Volatile: 'volatile';

Wchar: 'wchar_t';

While: 'while';
/*Operators*/

LeftParen: '(';

RightParen: ')';

LeftBracket: '[';

RightBracket: ']';

LeftBrace: '{';

RightBrace: '}';

Plus: '+';

Minus: '-';

Star: '*';

Div: '/';

Mod: '%';

Caret: '^';

And: '&';

Or: '|';

Tilde: '~';

Not: '!' | 'not';

Assign: '=';

Less: '<';

Greater: '>';

PlusAssign: '+=';

MinusAssign: '-=';

StarAssign: '*=';

DivAssign: '/=';

ModAssign: '%=';

XorAssign: '^=';

AndAssign: '&=';

OrAssign: '|=';

LeftShiftAssign: '<<=';

RightShiftAssign: '>>=';

Equal: '==';

NotEqual: '!=';

LessEqual: '<=';

GreaterEqual: '>=';

AndAnd: '&&' | 'and';

OrOr: '||' | 'or';

PlusPlus: '++';

MinusMinus: '--';

Comma: ',';

ArrowStar: '->*';

Arrow: '->';

Question: '?';

Colon: ':';

DoubleColon: '::';

Semi: ';';

Dot: '.';

DotStar: '.*';

Ellipsis: '...';

fragment Hexquad: HEXADECIMALDIGIT HEXADECIMALDIGIT HEXADECIMALDIGIT HEXADECIMALDIGIT;

fragment Universalcharactername: '\\u' Hexquad | '\\U' Hexquad Hexquad;

// 'restrict' is a C type qualifier that SuckC does not have. It is a real
// token rather than a bare Identifier because the C declarator grammar puts
// qualifiers after the '*' -- "int * restrict p" declares p as
// "int * restrict" -- so without a distinct token the subgrammar has to read
// 'restrict' as the declarator's name and the actual name 'p' becomes
// orphaned. Must precede Identifier: the spellings tie with it on length and
// ANTLR then takes the rule declared first.
Restrict: 'restrict' | '__restrict' | '__restrict__';

// 'typeof' / '__typeof__' / '__typeof' need a real token for the same reason
// Restrict does, and worse: spelled as a bare Identifier, the region
// subgrammar's typeof form ("Identifier '(' type-name ')'") matches ANY
// declaration-shaped call. "void f(void);" parses as a typeof specifier
// followed by nothing, and f disappears. Must precede Identifier, same
// tie-break reason as Restrict.
Typeof: 'typeof' | '__typeof__' | '__typeof';

Identifier:
    /*
	 Identifiernondigit | Identifier Identifiernondigit | Identifier DIGIT
	 */ Identifiernondigit (Identifiernondigit | DIGIT)*
;

fragment Identifiernondigit: NONDIGIT | Universalcharactername;

fragment NONDIGIT: [a-zA-Z_];

fragment DIGIT: [0-9];

DecimalLiteral: NONZERODIGIT ('\''? DIGIT)*;

OctalLiteral: '0' ('\''? OCTALDIGIT)*;

HexadecimalLiteral: ('0x' | '0X') HEXADECIMALDIGIT ( '\''? HEXADECIMALDIGIT)*;

BinaryLiteral: ('0b' | '0B') BINARYDIGIT ('\''? BINARYDIGIT)*;

// MUST stay after every plain numeric literal rule. ANTLR resolves an
// equal-length match by declaration order, and 0x0, 0b1 and 0u are each
// matched by both this rule (as a C++11 user-defined literal: digits plus a
// suffix) and by the plain literal rule. Declared first, it wins every tie,
// so every hexadecimal literal and every suffixed integer literal was
// tokenised as UserDefinedLiteral. Longest-match still holds, so 1_km and
// 1.5_km keep tokenising as user-defined literals.
UserDefinedLiteral:
    UserDefinedIntegerLiteral
    | UserDefinedFloatingLiteral
    | UserDefinedStringLiteral
    | UserDefinedCharacterLiteral
;

fragment NONZERODIGIT: [1-9];

fragment OCTALDIGIT: [0-7];

fragment HEXADECIMALDIGIT: [0-9a-fA-F];

fragment BINARYDIGIT: [01];

Integersuffix:
    Unsignedsuffix Longsuffix?
    | Unsignedsuffix Longlongsuffix?
    | Longsuffix Unsignedsuffix?
    | Longlongsuffix Unsignedsuffix?
;

fragment Unsignedsuffix: [uU];

fragment Longsuffix: [lL];

fragment Longlongsuffix: 'll' | 'LL';

fragment Cchar: ~ ['\\\r\n] | Escapesequence | Universalcharactername;

fragment Escapesequence: Simpleescapesequence | Octalescapesequence | Hexadecimalescapesequence;

fragment Simpleescapesequence:
    '\\\''
    | '\\"'
    | '\\?'
    | '\\\\'
    | '\\a'
    | '\\b'
    | '\\f'
    | '\\n'
    | '\\r'
    | '\\' ('\r' '\n'? | '\n')
    | '\\t'
    | '\\v'
;

fragment Octalescapesequence:
    '\\' OCTALDIGIT
    | '\\' OCTALDIGIT OCTALDIGIT
    | '\\' OCTALDIGIT OCTALDIGIT OCTALDIGIT
;

fragment Hexadecimalescapesequence: '\\x' HEXADECIMALDIGIT+;

fragment FractionalConstant: DigitSequence? '.' DigitSequence | DigitSequence '.';

fragment ExponentPart: 'e' SIGN? DigitSequence | 'E' SIGN? DigitSequence;

fragment SIGN: [+-];

fragment DigitSequence: DIGIT ('\''? DIGIT)*;

fragment FloatingSuffix: [flFL];

fragment EncodingPrefix: 'u8' | 'u' | 'U' | 'L';

fragment Schar: ~ ["\\\r\n] | Escapesequence | Universalcharactername;

fragment RawString: 'R"' ( '\\' ["()] | ~[\r\n (])*? '(' ~[)]*? ')' ( '\\' ["()] | ~[\r\n "])*? '"';

UserDefinedIntegerLiteral:
    DecimalLiteral Udsuffix
    | OctalLiteral Udsuffix
    | HexadecimalLiteral Udsuffix
    | BinaryLiteral Udsuffix
;

UserDefinedFloatingLiteral:
    FractionalConstant ExponentPart? Udsuffix
    | DigitSequence ExponentPart Udsuffix
;

UserDefinedStringLiteral: StringLiteral Udsuffix;

UserDefinedCharacterLiteral: CharacterLiteral Udsuffix;

fragment Udsuffix: Identifier;

Whitespace: [ \t]+ -> channel(HIDDEN);

Newline: ('\r' '\n'? | '\n') -> channel(HIDDEN);

BlockComment: '/*' .*? '*/' -> channel(HIDDEN);

LineComment: '//' ~[\r\n]* -> channel(HIDDEN);