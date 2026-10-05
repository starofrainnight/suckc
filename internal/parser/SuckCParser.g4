/*******************************************************************************
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Camilo Sanchez (Camiloasc1) 2020 Martin Mirchev (Marti2203)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
 * associated documentation files (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute,
 * sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or
 * substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
 * NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 * ****************************************************************************
 */

// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

// Reference to antlr4 grammar CPP14 & C11
parser grammar SuckCParser;

options {
    tokenVocab = SuckCLexer;
}

/*Basic concepts*/

translationUnit: declarationSeq? EOF;

/*Expressions*/

primaryExpression:
    literal+
    | This
    | LeftParen expression RightParen
    | idExpression
    | lambdaExpression
;

idExpression: unqualifiedId | qualifiedId;

unqualifiedId:
    Identifier
    | operatorFunctionId
    | conversionFunctionId
    | literalOperatorId
    | Tilde (className | decltypeSpecifier)
    | templateId
;

qualifiedId: nestedNameSpecifier Template? unqualifiedId;

nestedNameSpecifier: (theTypeName | namespaceName | decltypeSpecifier)? DoubleColon
    | nestedNameSpecifier ( Identifier | Template? simpleTemplateId) DoubleColon
;

lambdaExpression: lambdaIntroducer lambdaDeclarator? compoundStatement;

lambdaIntroducer: LeftBracket lambdaCapture? RightBracket;

lambdaCapture: captureList | captureDefault (Comma captureList)?;

captureDefault: Assign;

captureList: capture (Comma capture)* Ellipsis?;

capture: simpleCapture | initcapture;

simpleCapture: Identifier | This;

initcapture: Identifier initializer;

lambdaDeclarator:
    LeftParen parameterDeclarationClause? RightParen Mutable? exceptionSpecification? attributeSpecifierSeq? trailingReturnType?
;

postfixExpression:
    primaryExpression
    | postfixExpression LeftBracket (expression | bracedInitList) RightBracket
    | postfixExpression LeftParen expressionList? RightParen
    | (simpleTypeSpecifier | typeNameSpecifier) (
        LeftParen expressionList? RightParen
        | bracedInitList
    )
    | postfixExpression (Dot | Arrow) (Template? idExpression | pseudoDestructorName)
    | postfixExpression (PlusPlus | MinusMinus)
    | (Dynamic_cast | Static_cast | Reinterpret_cast | Const_cast) Less theTypeId Greater LeftParen expression RightParen
    | typeIdOfTheTypeId LeftParen (expression | theTypeId) RightParen
;

/*
 add a middle layer to eliminate duplicated function declarations
 */

typeIdOfTheTypeId: Typeid_;

expressionList: initializerList;

pseudoDestructorName:
    nestedNameSpecifier? (theTypeName DoubleColon)? Tilde theTypeName
    | nestedNameSpecifier Template simpleTemplateId DoubleColon Tilde theTypeName
    | Tilde decltypeSpecifier
;

unaryExpression:
    postfixExpression
    | (PlusPlus | MinusMinus | unaryOperator | Sizeof) unaryExpression
    | Sizeof (LeftParen theTypeId RightParen | Ellipsis LeftParen Identifier RightParen)
    | Alignof LeftParen theTypeId RightParen
    | noExceptExpression
    | newExpression_
    | deleteExpression
;

unaryOperator: Or | Star | And | Plus | Tilde | Minus | Not;

newExpression_:
    DoubleColon? New newPlacement? (newTypeId | LeftParen theTypeId RightParen) newInitializer_?
;

newPlacement: LeftParen expressionList RightParen;

newTypeId: typeSpecifierSeq newDeclarator_?;

newDeclarator_: pointerOperator newDeclarator_? | noPointerNewDeclarator;

noPointerNewDeclarator:
    LeftBracket expression RightBracket attributeSpecifierSeq?
    | noPointerNewDeclarator LeftBracket constantExpression RightBracket attributeSpecifierSeq?
;

newInitializer_: LeftParen expressionList? RightParen | bracedInitList;

deleteExpression: DoubleColon? Delete (LeftBracket RightBracket)? castExpression;

noExceptExpression: Noexcept LeftParen expression RightParen;

castExpression: unaryExpression | LeftParen theTypeId RightParen castExpression;

pointerMemberExpression: castExpression ((DotStar | ArrowStar) castExpression)*;

multiplicativeExpression: pointerMemberExpression ((Star | Div | Mod) pointerMemberExpression)*;

additiveExpression: multiplicativeExpression ((Plus | Minus) multiplicativeExpression)*;

shiftExpression: additiveExpression (shiftOperator additiveExpression)*;

shiftOperator: Greater Greater | Less Less;

relationalExpression:
    shiftExpression ((Less | Greater | LessEqual | GreaterEqual) shiftExpression)*
;

equalityExpression: relationalExpression ((Equal | NotEqual) relationalExpression)*;

andExpression: equalityExpression (And equalityExpression)*;

exclusiveOrExpression: andExpression (Caret andExpression)*;

inclusiveOrExpression: exclusiveOrExpression (Or exclusiveOrExpression)*;

logicalAndExpression: inclusiveOrExpression (AndAnd inclusiveOrExpression)*;

logicalOrExpression: logicalAndExpression (OrOr logicalAndExpression)*;

conditionalExpression: logicalOrExpression (Question expression Colon assignmentExpression)?;
assignmentExpression:
    conditionalExpression
    | logicalOrExpression assignmentOperator initializerClause
    | throwExpression
;

assignmentOperator:
    Assign
    | StarAssign
    | DivAssign
    | ModAssign
    | PlusAssign
    | MinusAssign
    | RightShiftAssign
    | LeftShiftAssign
    | AndAssign
    | XorAssign
    | OrAssign
;

expression: assignmentExpression (Comma assignmentExpression)*;

constantExpression: conditionalExpression;

/*Statements*/

statement:
    labeledStatement
    | declarationStatement
    | attributeSpecifierSeq? (
        expressionStatement
        | compoundStatement
        | selectionStatement
        | iterationStatement
        | jumpStatement
        | tryBlock
    )
;

labeledStatement:
    attributeSpecifierSeq? (Identifier | Case constantExpression | Default) Colon statement
;

expressionStatement: expression? Semi;

compoundStatement: LeftBrace statementSeq? RightBrace;

statementSeq: (statement | cImportRegion)+;

selectionStatement:
    If LeftParen condition RightParen statement (Else statement)?
    | Switch LeftParen condition RightParen statement
;

condition:
    expression
    | attributeSpecifierSeq? declSpecifierSeq declarator (
        Assign initializerClause
        | bracedInitList
    )
;

iterationStatement:
    While LeftParen condition RightParen statement
    | Do statement While LeftParen expression RightParen Semi
    | For LeftParen (
        forInitStatement condition? Semi expression?
        | forRangeDeclaration Colon forRangeInitializer
    ) RightParen statement
;

forInitStatement: expressionStatement | simpleDeclaration;

forRangeDeclaration: attributeSpecifierSeq? declSpecifierSeq declarator;

forRangeInitializer: expression | bracedInitList;

jumpStatement: (Break | Continue | Return (expression | bracedInitList)? | Goto Identifier) Semi;

declarationStatement: blockDeclaration;

/*Declarations*/

declarationSeq: declaration+;

declaration:
    blockDeclaration
    | functionDefinition
    // | templateDeclaration
    // | explicitInstantiation
    | explicitSpecialization
    | linkageSpecification
    // We don't have namespace definition in SuckC language
    //| namespaceDefinition
    | emptyDeclaration_
    | attributeDeclaration
    | cImportRegion
;

/*======================================================================
C-scope region contents: C declarations as they appear in a
preprocessed header file, plus just enough statement and expression
grammar to parse a function body, since a region may also wrap inline
definitions.

Two design decisions worth stating, because each fixes a problem that is
invisible until it bites:

1. Attribute, asm-label and typeof spellings (__attribute__, __attribute,
   __declspec, __asm__, __asm, __typeof__, __typeof, typeof, typeof, and
   the Keil __at / __irq family) are all plain Identifier tokens, and each
   is spelled as an identifier followed by a parenthesised group. They are
   matched positionally instead: an attribute is only offered where a
   declarator cannot start, or after a completed declarator, so
   "int f(int a);" never resolves its "f" as an attribute -- the reading
   that treats f(int a) as an attribute leaves no declarator before the
   ';' and dies in ANTLR's viability check. This avoids touching
   SuckCLexer.g4 at all, which matters because adding lexer rules for
   those spellings would change tokenization for the whole language.

2. cCExternalDeclaration separates function definitions from declarations
   STRUCTURALLY, never by rule order. When two alternatives both match,
   ANTLR silently prefers the first with no diagnostic at all, so an
   order-based split turns "int C { int a; };" into a bogus function
   definition named C and nothing complains. Here a definition requires a
   declarator that may end in a function suffix (cCDeclarator), while the
   braced-initializer alternative uses cCPlainDeclarator, which cannot end
   that way. The two readings are therefore never both viable.

Deliberate gaps:

  - '<<' and '>>' are not lexer tokens in SuckCLexer.g4 (only '<<=' and
    '>>=' are), so cCExpression omits the shift operators. Headers spell
    shifts through macros, which cpp has already expanded away.
  - A K&R parameter list is indistinguishable from a prototype with
    unnamed parameters: "f(a, b)" reads as two parameters whose types are
    the identifiers a and b.
  - "int f() { ... };" -- a definition followed by a stray semicolon --
    is rejected. C89 rejects it too; this follows the standard rather
    than the habit.
  - restrict and the vendor storage words (__packed, __sptr, __irq,
    __at, ...) are Identifier tokens, so they pass through the bare
    -Identifier type-specifier path. Token consumption is identical either
    way; recovering the correct type needs the symbol table Stage 2b adds.
======================================================================*/

// A C-scope region: two pragma markers around C code. Valid at file scope
// and inside compoundStatement.
//
// The body is a list of block items rather than a declaration sequence,
// because a region is not limited to declarations. At file scope a
// vendor header contributes declarations; inside a function body the
// natural region wraps statements, which is the shape the language
// documentation shows. cCBlockItem covers both, and '*' keeps an empty
// region valid -- frontend records that as a Region whose StartTok is
// greater than its EndTok.
cImportRegion: CScopeBegin cCBlockItem* CScopeEnd;

cCDeclarationSeq: cCExternalDeclaration+;

cCExternalDeclaration:
      cCAttributesOpt cCDeclSpecifiers cCDeclarator cCDeclarationSeq? cCCompoundStatement
    | cCAttributesOpt cCDeclSpecifiers cCInitDeclaratorList? Semi
;

cCAttributesOpt: cCAttribute*;

cCAttribute: cCAttributeName cCParenGroup;

// __attribute__, __attribute, __declspec, __asm__, __asm, asm, and the
// Keil/IAR suffixes (__at, __irq, __align) all land here. See note 1 in the
// block comment above for why no lexer rule distinguishes them.
cCAttributeName: Identifier | Asm;

// Balanced parentheses over arbitrary tokens. SuckCParser's own
// balancedTokenSeq requires at least one token and nests only brackets,
// so it cannot stand in for an empty "()" nor for the leading '(' of the
// double-parenthesized attribute form.
cCParenGroup: LeftParen (cCParenGroup | ~(LeftParen | RightParen))* RightParen;

// Deliberately not SuckCParser's storageClassSpecifier / typeSpecifier:
// those omit Typedef and Auto, and typeSpecifier drags in decltype,
// templates and Auto-as-a-type, none of which belong to C.
//
// The split into two alternatives is what stops the specifier list from
// swallowing the declarator. "typedef int x;" and "size_t x;" differ only
// in whether the identifier names a new type or one already known, and the
// grammar cannot know which -- C decides it with the symbol table. Left to
// a bare Identifier inside cCDeclSpecifier, the '+' loop keeps consuming and
// ANTLR resolves the ambiguity by taking the greedy branch, silently: no
// error, no warning, and the declarator is simply gone.
//
// So the first alternative cannot end on a bare identifier. When both
// readings of "typedef int x;" are viable -- specifiers "typedef int" plus
// declarator "x", or specifiers "typedef int x" and no declarator -- the
// first alternative is the one that leaves the declarator, and it is tried
// first. The second alternative carries a leading typedef name, which is the
// only reading left for "size_t x;" once "size_t" cannot be part of a
// "typedef unsigned long"-shaped specifier run.
cCDeclSpecifiers
    : cCDeclSpecifierPart+
    | cCDeclSpecifierPart* cCTypedefName cCDeclSpecifierPart*
    ;

cCDeclSpecifierPart
    : Typedef
    | Extern StringLiteral
    | Extern
    | Static
    | Auto
    | Register
    | Inline
    | cCQualifier
    | cCBuiltinTypeSpecifier
    | cCStructOrUnionSpecifier
    | cCEnumSpecifier
    | cCTypeofSpecifier
    ;

cCQualifier: Const | Volatile | Constexpr | Restrict;

// A name already known to be a type. Kept out of cCBuiltinTypeSpecifier on
// purpose -- see cCDeclSpecifiers.
cCTypedefName: Identifier;

cCBuiltinTypeSpecifier:
      Char | Short | Int | Long | Float | Double | Signed | Unsigned | Bool | Void
;

cCStructOrUnionSpecifier:
    (Struct | Union) cCAttributesOpt cCIdentifier? (LeftBrace cCMemberDeclaration* RightBrace)?
;

cCMemberDeclaration: cCAttributesOpt cCDeclSpecifiers cCMemberDeclaratorList? cCBitField? Semi;

cCMemberDeclaratorList: cCMemberDeclarator (Comma cCMemberDeclarator)*;

cCMemberDeclarator: cCDeclarator;

// "int : 0;" -- an unnamed bitfield. Optional rather than an alternative so
// that "int a : 3;", "int : 0;" and an anonymous member (neither present)
// all reduce through one path instead of competing as alternatives.
cCBitField: Colon cCConstantExpression;

cCEnumSpecifier: Enum cCAttributesOpt cCIdentifier? (LeftBrace cCEnumeratorList? RightBrace)?;

cCEnumeratorList: cCEnumerator (Comma cCEnumerator)* Comma?;

cCEnumerator: cCIdentifier (Assign cCConstantExpression)?;

cCTypeofSpecifier: Typeof LeftParen (cCTypeName | cCExpression) RightParen;

cCTypeName: cCDeclSpecifiers (cCDeclarator | cCAbstractDeclarator)?;

// Every SuckC keyword that can legitimately spell an ordinary C identifier
// is accepted here. Inside a C region these words are not SuckC keywords at
// all -- they are C names -- and the region rules are the only place this
// rule is reachable, so widening it cannot affect SuckC code. Without this,
// real headers fail: glibc's stddef.h declares "typedef int wchar_t;" and
// uchar.h declares char16_t/char32_t, all spelled with SuckC keyword tokens.
//
// `true`, `false` and `nullptr` are reached through BooleanLiteral and
// PointerLiteral, not through True_/False_/Nullptr. Those three lexer rules
// exist but can never fire: BooleanLiteral and PointerLiteral are declared
// first and match the same spellings, and a lexer rule referenced only from
// another lexer rule emits no tokens of its own.
//
// Control-flow keywords are deliberately absent. They would create genuine
// ambiguity with cCStatement (a label "foo:" versus a keyword statement) and
// no C header spells a declaration with them.
cCIdentifier: Identifier
             | Asm | Bool | Mutable | Thread_local
             | BooleanLiteral | PointerLiteral
             | Wchar | Char16 | Char32
             | Constexpr | New | Delete
             | Template | Typename_ | Namespace;

// cCPlainDeclarator and cCDeclarator are deliberately parallel hierarchies
// differing only in whether a function suffix is permitted. See note 2 in
// the block comment above before changing that.
cCDeclarator: cCAttributesOpt cCPointerOperator* cCDirectDeclarator;

cCPlainDeclarator: cCAttributesOpt cCPointerOperator* cCPlainDirectDeclarator;

cCPointerOperator: Star cCAttributesOpt cCPointerQualifier*;

cCPointerQualifier: cCQualifier | Identifier;

cCDirectDeclarator:
      cCIdentifier cCDirectDeclaratorSuffix*
    | LeftParen cCDeclarator RightParen cCDirectDeclaratorSuffix*
;

cCPlainDirectDeclarator:
      cCIdentifier cCPlainDirectDeclaratorSuffix*
    | LeftParen cCPlainDeclarator RightParen cCPlainDirectDeclaratorSuffix*
;

// A postfix attribute also covers the Keil and IAR address/interrupt
// suffixes, "int vector __at(0x2000);", which have the same shape.
cCDirectDeclaratorSuffix: cCArraySuffix | cCFunctionSuffix | cCAttribute;

cCPlainDirectDeclaratorSuffix: cCArraySuffix | cCAttribute;

// A declarator with no identifier, as in the parameter of "int f(int)"
// and the type name of "(int *)". It must be able to consume something, or
// it would compete with the optional named declarator above.
cCAbstractDeclarator: cCAttributesOpt (cCPointerOperator+ cCAbstractDirectDeclarator? | cCAbstractDirectDeclarator);

cCAbstractDirectDeclarator:
      cCAbstractDirectDeclaratorSuffix+
    | LeftParen cCAbstractDeclarator RightParen cCAbstractDirectDeclaratorSuffix*
;

cCAbstractDirectDeclaratorSuffix: cCArraySuffix | cCFunctionSuffix | cCAttribute;

cCArraySuffix: LeftBracket cCConstantExpression? RightBracket;

cCFunctionSuffix: LeftParen cCParameterTypeList? RightParen cCAttributesOpt;

cCParameterTypeList:
      cCParameterDeclaration (Comma cCParameterDeclaration)* (Comma Ellipsis)?
    | Ellipsis
;

cCParameterDeclaration: cCAttributesOpt cCDeclSpecifiers (cCDeclarator | cCAbstractDeclarator)?;

cCInitDeclaratorList: cCInitDeclarator (Comma cCInitDeclarator)*;

cCInitDeclarator:
      cCDeclarator Assign cCInitializerClause
    | cCPlainDeclarator cCBracedInitList
    | cCDeclarator
;

cCInitializerClause: cCExpression | cCBracedInitList;

cCBracedInitList: LeftBrace (cCInitializerClause (Comma? cCInitializerClause)*)? RightBrace;

cCCompoundStatement: LeftBrace cCBlockItem* RightBrace;

cCBlockItem: cCDeclarationSeq | cCStatement | cCLinkageSpecification;

cCLinkageSpecification: Extern StringLiteral LeftBrace cCDeclarationSeq? RightBrace;

cCStatement:
      cCCompoundStatement
    | If LeftParen cCExpression RightParen cCStatement (Else cCStatement)?
    | Switch LeftParen cCExpression RightParen cCStatement
    | While LeftParen cCExpression RightParen cCStatement
    | Do cCStatement While LeftParen cCExpression RightParen Semi
    | For LeftParen (cCDeclarationSeq | cCExpression? Semi) cCExpression? Semi cCExpression? RightParen cCStatement
    | Case cCConstantExpression Colon
    | Default Colon
    | Break Semi
    | Continue Semi
    | Return cCExpression? Semi
    | Goto cCIdentifier Semi
    | cCIdentifier Colon cCStatement?
    | cCExpression? Semi
;

cCExpression: cCAssignmentExpression;

cCAssignmentExpression:
      cCUnaryExpression cCAssignmentOperator cCAssignmentExpression
    | cCConditionalExpression
;

cCAssignmentOperator:
      Assign | PlusAssign | MinusAssign | StarAssign | DivAssign | ModAssign
    | AndAssign | OrAssign | XorAssign | LeftShiftAssign | RightShiftAssign
;

cCConditionalExpression:
      cCLogicalOrExpression (Question cCExpression Colon cCConditionalExpression)?
;

cCLogicalOrExpression: cCLogicalAndExpression (OrOr cCLogicalAndExpression)*;

cCLogicalAndExpression: cCInclusiveOrExpression (AndAnd cCInclusiveOrExpression)*;

cCInclusiveOrExpression: cCExclusiveOrExpression (Or cCExclusiveOrExpression)*;

cCExclusiveOrExpression: cCAndExpression (Caret cCAndExpression)*;

cCAndExpression: cCEqualityExpression (And cCEqualityExpression)*;

cCEqualityExpression: cCRelationalExpression ((Equal | NotEqual) cCRelationalExpression)*;

cCRelationalExpression: cCAdditiveExpression ((Less | Greater | LessEqual | GreaterEqual) cCAdditiveExpression)*;

cCAdditiveExpression: cCMultiplicativeExpression ((Plus | Minus) cCMultiplicativeExpression)*;

cCMultiplicativeExpression: cCUnaryExpression ((Star | Div | Mod) cCUnaryExpression)*;

cCUnaryExpression:
      PlusPlus cCUnaryExpression
    | MinusMinus cCUnaryExpression
    | cCUnaryOperator cCUnaryExpression
    | LeftParen cCTypeName RightParen cCUnaryExpression
    | Sizeof cCUnaryExpression
    | cCPostfixExpression
;

cCUnaryOperator: And | Star | Plus | Minus | Tilde | Not;

cCPostfixExpression: cCPrimaryExpression cCPostfixSuffix*;

cCPostfixSuffix:
      LeftBracket cCExpression? RightBracket
    | LeftParen cCArgumentExpressionList? RightParen
    | Dot cCIdentifier
    | Arrow cCIdentifier
    | PlusPlus
    | MinusMinus
;

cCArgumentExpressionList: cCAssignmentExpression (Comma cCAssignmentExpression)*;

// The lexer has no aggregate Literal token, so the literal forms are
// listed individually.
cCPrimaryExpression:
      cCIdentifier
    | DecimalLiteral
    | OctalLiteral
    | HexadecimalLiteral
    | BinaryLiteral
    | FloatingLiteral
    | CharacterLiteral
    | StringLiteral
    | UserDefinedLiteral
    | BooleanLiteral
    | PointerLiteral
    | LeftParen cCExpression RightParen
;

cCConstantExpression: cCConditionalExpression;

blockDeclaration:
    simpleDeclaration
    | asmDefinition
    | namespaceAliasDefinition
    | usingDeclaration
    | usingDirective
    | staticAssertDeclaration
    | typedefDeclaration
    | aliasDeclaration
    | opaqueEnumDeclaration
;

aliasDeclaration: Using Identifier attributeSpecifierSeq? Assign theTypeId Semi;

functionPointerDeclarator:
    typeSpecifier LeftParen Star typedefName RightParen parametersAndQualifiers
;
simpleTypedefDeclarator: typeSpecifier typedefName;

typedefDeclaration: Typedef (functionPointerDeclarator | simpleTypedefDeclarator) Semi;

simpleDeclaration:
    declSpecifierSeq? initDeclarator? Semi
    | attributeSpecifierSeq declSpecifierSeq? initDeclarator Semi
;

staticAssertDeclaration:
    Static_assert LeftParen constantExpression Comma StringLiteral RightParen Semi
;

emptyDeclaration_: Semi;

attributeDeclaration: attributeSpecifierSeq Semi;

declSpecifier: storageClassSpecifier | typeSpecifier | functionSpecifier | Friend | Constexpr;

declSpecifierSeq: declSpecifier+? attributeSpecifierSeq?;

storageClassSpecifier: Register | Static | Thread_local | Extern | Mutable;

functionSpecifier: Inline | Virtual | Explicit;

typedefName: Identifier;

typeSpecifier: trailingTypeSpecifier | classSpecifier | enumSpecifier;

trailingTypeSpecifier:
    simpleTypeSpecifier
    | elaboratedTypeSpecifier
    | typeNameSpecifier
    | cvQualifier
;

typeSpecifierSeq: typeSpecifier+ attributeSpecifierSeq?;

trailingTypeSpecifierSeq: trailingTypeSpecifier+ attributeSpecifierSeq?;

simpleTypeLengthModifier: Short | Long;

simpleTypeSignednessModifier: Unsigned | Signed;

simpleTypeSpecifier:
    nestedNameSpecifier? theTypeName
    | nestedNameSpecifier Template simpleTemplateId
    | Char
    | Char16
    | Char32
    | Wchar
    | Bool
    | Short
    | Int
    | Long
    | Float
    | Signed
    | Unsigned
    | Float
    | Double
    | Void
    | Auto
    | decltypeSpecifier
;

theTypeName: className | enumName | typedefName | simpleTemplateId;

decltypeSpecifier: Decltype LeftParen (expression | Auto) RightParen;

elaboratedTypeSpecifier:
    Struct (
        attributeSpecifierSeq? nestedNameSpecifier? Identifier
        | simpleTemplateId
        | nestedNameSpecifier Template? simpleTemplateId
    )
    | Enum nestedNameSpecifier? Identifier
;

enumName: Identifier;

enumSpecifier: enumHead LeftBrace (enumeratorList Comma?)? RightBrace;

enumHead: enumkey attributeSpecifierSeq? (nestedNameSpecifier? Identifier)? enumbase?;

opaqueEnumDeclaration: enumkey attributeSpecifierSeq? Identifier enumbase? Semi;

enumkey: Enum Struct?;

enumbase: Colon typeSpecifierSeq;

enumeratorList: enumeratorDefinition (Comma enumeratorDefinition)*;

enumeratorDefinition: enumerator (Assign constantExpression)?;

enumerator: Identifier;

namespaceName: originalNamespaceName | namespaceAlias;

originalNamespaceName: Identifier;

namespaceDefinition:
    Inline? Namespace (Identifier | originalNamespaceName)? LeftBrace namespaceBody = declarationSeq? RightBrace
;

namespaceAlias: Identifier;

namespaceAliasDefinition: Namespace Identifier Assign qualifiednamespaceSpecifier Semi;

qualifiednamespaceSpecifier: nestedNameSpecifier? namespaceName;

usingDeclaration: Using (Typename_? nestedNameSpecifier | DoubleColon) unqualifiedId Semi;

usingDirective: attributeSpecifierSeq? Using Namespace nestedNameSpecifier? namespaceName Semi;

asmDefinition: Asm LeftParen StringLiteral RightParen Semi;

linkageSpecification: Extern StringLiteral (LeftBrace declarationSeq? RightBrace | declaration);

attributeSpecifierSeq: attributeSpecifier+;

attributeSpecifier:
    LeftBracket LeftBracket attributeList? RightBracket RightBracket
    | alignmentSpecifier
;

alignmentSpecifier: Alignas LeftParen (theTypeId | constantExpression) Ellipsis? RightParen;

attributeList: attribute (Comma attribute)* Ellipsis?;

attribute: (attributeNamespace DoubleColon)? Identifier attributeArgumentClause?;

attributeNamespace: Identifier;

attributeArgumentClause: LeftParen balancedTokenSeq? RightParen;

balancedTokenSeq: balancedtoken+;

balancedtoken:
    LeftParen balancedTokenSeq RightParen
    | LeftBracket balancedTokenSeq RightBracket
    | LeftBrace balancedTokenSeq RightBrace
    | ~(LeftParen | RightParen | LeftBrace | RightBrace | LeftBracket | RightBracket)+
;

/*Declarators*/

initDeclarator: declarator initializer?;

declarator: pointerDeclarator | noPointerDeclarator parametersAndQualifiers trailingReturnType;

pointerDeclarator: (pointerOperator Const?)* noPointerDeclarator;

noPointerDeclarator:
    declaratorid attributeSpecifierSeq?
    | noPointerDeclarator (
        parametersAndQualifiers
        | LeftBracket constantExpression? RightBracket attributeSpecifierSeq?
    )
    | LeftParen pointerDeclarator RightParen
;

parametersAndQualifiers:
    LeftParen parameterDeclarationClause? RightParen cvQualifierSeq? exceptionSpecification? attributeSpecifierSeq?
;

trailingReturnType: Arrow trailingTypeSpecifierSeq abstractDeclarator?;

pointerOperator: nestedNameSpecifier? Star attributeSpecifierSeq? cvQualifierSeq?;

cvQualifierSeq: cvQualifier+;

cvQualifier: Const | Volatile;

declaratorid: Ellipsis? idExpression;

theTypeId: typeSpecifierSeq abstractDeclarator?;

abstractDeclarator:
    pointerAbstractDeclarator
    | noPointerAbstractDeclarator? parametersAndQualifiers trailingReturnType
    | abstractPackDeclarator
;

pointerAbstractDeclarator: pointerOperator* (noPointerAbstractDeclarator | pointerOperator);

noPointerAbstractDeclarator: (
        parametersAndQualifiers
        | LeftParen pointerAbstractDeclarator RightParen
    ) (
        parametersAndQualifiers
        | LeftBracket constantExpression? RightBracket attributeSpecifierSeq?
    )*
;

abstractPackDeclarator: pointerOperator* noPointerAbstractPackDeclarator;

noPointerAbstractPackDeclarator:
    Ellipsis (
        parametersAndQualifiers
        | LeftBracket constantExpression? RightBracket attributeSpecifierSeq?
    )*
;

parameterDeclarationClause: parameterDeclarationList (Comma? Ellipsis)?;

parameterDeclarationList: parameterDeclaration (Comma parameterDeclaration)*;

parameterDeclaration:
    attributeSpecifierSeq? declSpecifierSeq (declarator | abstractDeclarator?) (
        Assign initializerClause
    )?
;

functionDefinition:
    attributeSpecifierSeq? declSpecifierSeq? functionDeclarator virtualSpecifierSeq? functionBody
;

// Unlike declarator, this always requires a parameter list. Otherwise
// `int C { int a; };` parses as a definition of C returning `int`.
functionDeclarator:
    pointerDeclarator parametersAndQualifiers trailingReturnType?
;

functionBody:
    constructorInitializer? compoundStatement
    | functionTryBlock
    | Assign (Default | Delete) Semi
;

initializer: braceOrEqualInitializer | LeftParen expressionList RightParen;

// int anvar "= 1" or "(...)"
// 
// 1. int anvar = 1
// 2. int anvar(1)
braceOrEqualInitializer: Assign initializerClause | bracedInitList;

initializerClause: assignmentExpression | bracedInitList;

initializerList: initializerClause Ellipsis? (Comma initializerClause Ellipsis?)*;

bracedInitList: LeftBrace (initializerList Comma?)? RightBrace;

/*Classes*/

className: Identifier | simpleTemplateId;

classSpecifier: classHead LeftBrace memberSpecification? RightBrace;

classHead:
    Struct attributeSpecifierSeq? (classHeadName classVirtSpecifier?)? baseClause?
    | Union attributeSpecifierSeq? ( classHeadName classVirtSpecifier?)?
;

classHeadName: nestedNameSpecifier? className;

classVirtSpecifier: Final;

memberSpecification: (memberdeclaration | accessSpecifier Colon)+;

memberdeclaration:
    attributeSpecifierSeq? declSpecifierSeq? memberDeclaratorList? Semi
    | functionDefinition
    | usingDeclaration
    | staticAssertDeclaration
    // | templateDeclaration
    | aliasDeclaration
    | emptyDeclaration_
;

memberDeclaratorList: memberDeclarator (Comma memberDeclarator)*;

memberDeclarator:
    declarator (
        virtualSpecifierSeq
        | { IsPureSpecifierAllowed() }? pureSpecifier
        | { IsPureSpecifierAllowed() }? virtualSpecifierSeq pureSpecifier
        | braceOrEqualInitializer
    )
    | declarator
    | Identifier? attributeSpecifierSeq? Colon constantExpression
;

virtualSpecifierSeq: virtualSpecifier+;

virtualSpecifier: Override | Final;

/*
 pureSpecifier: Assign '0'//Conflicts with the lexer ;
 */

pureSpecifier: Assign integerLiteral;

/*Derived classes*/

baseClause: Colon baseSpecifierList;

baseSpecifierList: baseSpecifier Ellipsis? (Comma baseSpecifier Ellipsis?)*;

baseSpecifier:
    attributeSpecifierSeq? (
        baseTypeSpecifier
        | Virtual accessSpecifier? baseTypeSpecifier
        | accessSpecifier Virtual? baseTypeSpecifier
    )
;

classOrDeclType: nestedNameSpecifier? className | decltypeSpecifier;

baseTypeSpecifier: classOrDeclType;

accessSpecifier: Private | Protected | Public;

/*Special member functions*/

conversionFunctionId: Operator conversionTypeId;

conversionTypeId: typeSpecifierSeq conversionDeclarator?;

conversionDeclarator: pointerOperator conversionDeclarator?;

constructorInitializer: Colon memInitializerList;

memInitializerList: memInitializer Ellipsis? (Comma memInitializer Ellipsis?)*;

memInitializer: meminitializerid (LeftParen expressionList? RightParen | bracedInitList);

meminitializerid: classOrDeclType | Identifier;

/*Overloading*/

operatorFunctionId: Operator theOperator;

literalOperatorId: Operator (StringLiteral Identifier | UserDefinedStringLiteral);

/*Templates*/

templateDeclaration: Template Less templateparameterList Greater declaration;

templateparameterList: templateParameter (Comma templateParameter)*;

templateParameter: typeParameter | parameterDeclaration;

typeParameter: (Template Less templateparameterList Greater)? Typename_ (
        Ellipsis? Identifier?
        | Identifier? Assign theTypeId
    )
;

simpleTemplateId: templateName Less templateArgumentList? Greater;

templateId:
    simpleTemplateId
    | (operatorFunctionId | literalOperatorId) Less templateArgumentList? Greater
;

templateName: Identifier;

templateArgumentList: templateArgument Ellipsis? (Comma templateArgument Ellipsis?)*;

templateArgument: theTypeId | constantExpression | idExpression;

typeNameSpecifier: Typename_ nestedNameSpecifier (Identifier | Template? simpleTemplateId);

explicitInstantiation: Extern? Template declaration;

explicitSpecialization: Template Less Greater declaration;

/*Exception handling*/

tryBlock: Try compoundStatement handlerSeq;

functionTryBlock: Try constructorInitializer? compoundStatement handlerSeq;

handlerSeq: handler+;

handler: Catch LeftParen exceptionDeclaration RightParen compoundStatement;

exceptionDeclaration:
    attributeSpecifierSeq? typeSpecifierSeq (declarator | abstractDeclarator)?
    | Ellipsis
;

throwExpression: Throw assignmentExpression?;

exceptionSpecification: dynamicExceptionSpecification | noeExceptSpecification;

dynamicExceptionSpecification: Throw LeftParen typeIdList? RightParen;

typeIdList: theTypeId Ellipsis? (Comma theTypeId Ellipsis?)*;

noeExceptSpecification: Noexcept LeftParen constantExpression RightParen | Noexcept;

/*Preprocessing directives*/

/*Lexer*/

theOperator:
    New (LeftBracket RightBracket)?
    | Delete (LeftBracket RightBracket)?
    | Plus
    | Minus
    | Star
    | Div
    | Mod
    | Caret
    | And
    | Or
    | Tilde
    | Not
    | Assign
    | Greater
    | Less
    | GreaterEqual
    | PlusAssign
    | MinusAssign
    | StarAssign
    | ModAssign
    | XorAssign
    | AndAssign
    | OrAssign
    | Less Less
    | Greater Greater
    | RightShiftAssign
    | LeftShiftAssign
    | Equal
    | NotEqual
    | LessEqual
    | AndAnd
    | OrOr
    | PlusPlus
    | MinusMinus
    | Comma
    | ArrowStar
    | Arrow
    | LeftParen RightParen
    | LeftBracket RightBracket
;

decimalLiteral     : DecimalLiteral Integersuffix?;
octalLiteral       : OctalLiteral Integersuffix?;
hexadecimalLiteral : HexadecimalLiteral Integersuffix?;
binaryLiteral      : BinaryLiteral Integersuffix?;
integerLiteral     : decimalLiteral | octalLiteral | hexadecimalLiteral | binaryLiteral;
characterLiteral   : CharacterLiteral;
stringLiteral      : StringLiteral;
floatingLiteral    : FloatingLiteral;
booleanLiteral     : BooleanLiteral;
pointerLiteral     : PointerLiteral;

literal:
    integerLiteral
    | characterLiteral
    | floatingLiteral
    | stringLiteral
    | booleanLiteral
    | pointerLiteral
    | UserDefinedLiteral
;