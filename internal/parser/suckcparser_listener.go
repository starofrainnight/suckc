// Code generated from SuckCParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SuckCParser

import "github.com/antlr4-go/antlr/v4"

// SuckCParserListener is a complete listener for a parse tree produced by SuckCParser.
type SuckCParserListener interface {
	antlr.ParseTreeListener

	// EnterTranslationUnit is called when entering the translationUnit production.
	EnterTranslationUnit(c *TranslationUnitContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterIdExpression is called when entering the idExpression production.
	EnterIdExpression(c *IdExpressionContext)

	// EnterUnqualifiedId is called when entering the unqualifiedId production.
	EnterUnqualifiedId(c *UnqualifiedIdContext)

	// EnterQualifiedId is called when entering the qualifiedId production.
	EnterQualifiedId(c *QualifiedIdContext)

	// EnterNestedNameSpecifier is called when entering the nestedNameSpecifier production.
	EnterNestedNameSpecifier(c *NestedNameSpecifierContext)

	// EnterLambdaExpression is called when entering the lambdaExpression production.
	EnterLambdaExpression(c *LambdaExpressionContext)

	// EnterLambdaIntroducer is called when entering the lambdaIntroducer production.
	EnterLambdaIntroducer(c *LambdaIntroducerContext)

	// EnterLambdaCapture is called when entering the lambdaCapture production.
	EnterLambdaCapture(c *LambdaCaptureContext)

	// EnterCaptureDefault is called when entering the captureDefault production.
	EnterCaptureDefault(c *CaptureDefaultContext)

	// EnterCaptureList is called when entering the captureList production.
	EnterCaptureList(c *CaptureListContext)

	// EnterCapture is called when entering the capture production.
	EnterCapture(c *CaptureContext)

	// EnterSimpleCapture is called when entering the simpleCapture production.
	EnterSimpleCapture(c *SimpleCaptureContext)

	// EnterInitcapture is called when entering the initcapture production.
	EnterInitcapture(c *InitcaptureContext)

	// EnterLambdaDeclarator is called when entering the lambdaDeclarator production.
	EnterLambdaDeclarator(c *LambdaDeclaratorContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterTypeIdOfTheTypeId is called when entering the typeIdOfTheTypeId production.
	EnterTypeIdOfTheTypeId(c *TypeIdOfTheTypeIdContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterPseudoDestructorName is called when entering the pseudoDestructorName production.
	EnterPseudoDestructorName(c *PseudoDestructorNameContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterNewExpression_ is called when entering the newExpression_ production.
	EnterNewExpression_(c *NewExpression_Context)

	// EnterNewPlacement is called when entering the newPlacement production.
	EnterNewPlacement(c *NewPlacementContext)

	// EnterNewTypeId is called when entering the newTypeId production.
	EnterNewTypeId(c *NewTypeIdContext)

	// EnterNewDeclarator_ is called when entering the newDeclarator_ production.
	EnterNewDeclarator_(c *NewDeclarator_Context)

	// EnterNoPointerNewDeclarator is called when entering the noPointerNewDeclarator production.
	EnterNoPointerNewDeclarator(c *NoPointerNewDeclaratorContext)

	// EnterNewInitializer_ is called when entering the newInitializer_ production.
	EnterNewInitializer_(c *NewInitializer_Context)

	// EnterDeleteExpression is called when entering the deleteExpression production.
	EnterDeleteExpression(c *DeleteExpressionContext)

	// EnterNoExceptExpression is called when entering the noExceptExpression production.
	EnterNoExceptExpression(c *NoExceptExpressionContext)

	// EnterCastExpression is called when entering the castExpression production.
	EnterCastExpression(c *CastExpressionContext)

	// EnterPointerMemberExpression is called when entering the pointerMemberExpression production.
	EnterPointerMemberExpression(c *PointerMemberExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterShiftExpression is called when entering the shiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

	// EnterShiftOperator is called when entering the shiftOperator production.
	EnterShiftOperator(c *ShiftOperatorContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterExclusiveOrExpression is called when entering the exclusiveOrExpression production.
	EnterExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// EnterInclusiveOrExpression is called when entering the inclusiveOrExpression production.
	EnterInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// EnterLogicalAndExpression is called when entering the logicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterLogicalOrExpression is called when entering the logicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterConditionalExpression is called when entering the conditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConstantExpression is called when entering the constantExpression production.
	EnterConstantExpression(c *ConstantExpressionContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLabeledStatement is called when entering the labeledStatement production.
	EnterLabeledStatement(c *LabeledStatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// EnterStatementSeq is called when entering the statementSeq production.
	EnterStatementSeq(c *StatementSeqContext)

	// EnterSelectionStatement is called when entering the selectionStatement production.
	EnterSelectionStatement(c *SelectionStatementContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterIterationStatement is called when entering the iterationStatement production.
	EnterIterationStatement(c *IterationStatementContext)

	// EnterForInitStatement is called when entering the forInitStatement production.
	EnterForInitStatement(c *ForInitStatementContext)

	// EnterForRangeDeclaration is called when entering the forRangeDeclaration production.
	EnterForRangeDeclaration(c *ForRangeDeclarationContext)

	// EnterForRangeInitializer is called when entering the forRangeInitializer production.
	EnterForRangeInitializer(c *ForRangeInitializerContext)

	// EnterJumpStatement is called when entering the jumpStatement production.
	EnterJumpStatement(c *JumpStatementContext)

	// EnterDeclarationStatement is called when entering the declarationStatement production.
	EnterDeclarationStatement(c *DeclarationStatementContext)

	// EnterDeclarationSeq is called when entering the declarationSeq production.
	EnterDeclarationSeq(c *DeclarationSeqContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterCImportRegion is called when entering the cImportRegion production.
	EnterCImportRegion(c *CImportRegionContext)

	// EnterCCDeclarationSeq is called when entering the cCDeclarationSeq production.
	EnterCCDeclarationSeq(c *CCDeclarationSeqContext)

	// EnterCCExternalDeclaration is called when entering the cCExternalDeclaration production.
	EnterCCExternalDeclaration(c *CCExternalDeclarationContext)

	// EnterCCAttributesOpt is called when entering the cCAttributesOpt production.
	EnterCCAttributesOpt(c *CCAttributesOptContext)

	// EnterCCAttribute is called when entering the cCAttribute production.
	EnterCCAttribute(c *CCAttributeContext)

	// EnterCCAttributeName is called when entering the cCAttributeName production.
	EnterCCAttributeName(c *CCAttributeNameContext)

	// EnterCCParenGroup is called when entering the cCParenGroup production.
	EnterCCParenGroup(c *CCParenGroupContext)

	// EnterCCDeclSpecifiers is called when entering the cCDeclSpecifiers production.
	EnterCCDeclSpecifiers(c *CCDeclSpecifiersContext)

	// EnterCCDeclSpecifierPart is called when entering the cCDeclSpecifierPart production.
	EnterCCDeclSpecifierPart(c *CCDeclSpecifierPartContext)

	// EnterCCQualifier is called when entering the cCQualifier production.
	EnterCCQualifier(c *CCQualifierContext)

	// EnterCCTypedefName is called when entering the cCTypedefName production.
	EnterCCTypedefName(c *CCTypedefNameContext)

	// EnterCCBuiltinTypeSpecifier is called when entering the cCBuiltinTypeSpecifier production.
	EnterCCBuiltinTypeSpecifier(c *CCBuiltinTypeSpecifierContext)

	// EnterCCStructOrUnionSpecifier is called when entering the cCStructOrUnionSpecifier production.
	EnterCCStructOrUnionSpecifier(c *CCStructOrUnionSpecifierContext)

	// EnterCCMemberDeclaration is called when entering the cCMemberDeclaration production.
	EnterCCMemberDeclaration(c *CCMemberDeclarationContext)

	// EnterCCMemberDeclaratorList is called when entering the cCMemberDeclaratorList production.
	EnterCCMemberDeclaratorList(c *CCMemberDeclaratorListContext)

	// EnterCCMemberDeclarator is called when entering the cCMemberDeclarator production.
	EnterCCMemberDeclarator(c *CCMemberDeclaratorContext)

	// EnterCCBitField is called when entering the cCBitField production.
	EnterCCBitField(c *CCBitFieldContext)

	// EnterCCEnumSpecifier is called when entering the cCEnumSpecifier production.
	EnterCCEnumSpecifier(c *CCEnumSpecifierContext)

	// EnterCCEnumeratorList is called when entering the cCEnumeratorList production.
	EnterCCEnumeratorList(c *CCEnumeratorListContext)

	// EnterCCEnumerator is called when entering the cCEnumerator production.
	EnterCCEnumerator(c *CCEnumeratorContext)

	// EnterCCTypeofSpecifier is called when entering the cCTypeofSpecifier production.
	EnterCCTypeofSpecifier(c *CCTypeofSpecifierContext)

	// EnterCCTypeName is called when entering the cCTypeName production.
	EnterCCTypeName(c *CCTypeNameContext)

	// EnterCCIdentifier is called when entering the cCIdentifier production.
	EnterCCIdentifier(c *CCIdentifierContext)

	// EnterCCDeclarator is called when entering the cCDeclarator production.
	EnterCCDeclarator(c *CCDeclaratorContext)

	// EnterCCPlainDeclarator is called when entering the cCPlainDeclarator production.
	EnterCCPlainDeclarator(c *CCPlainDeclaratorContext)

	// EnterCCPointerOperator is called when entering the cCPointerOperator production.
	EnterCCPointerOperator(c *CCPointerOperatorContext)

	// EnterCCPointerQualifier is called when entering the cCPointerQualifier production.
	EnterCCPointerQualifier(c *CCPointerQualifierContext)

	// EnterCCDirectDeclarator is called when entering the cCDirectDeclarator production.
	EnterCCDirectDeclarator(c *CCDirectDeclaratorContext)

	// EnterCCPlainDirectDeclarator is called when entering the cCPlainDirectDeclarator production.
	EnterCCPlainDirectDeclarator(c *CCPlainDirectDeclaratorContext)

	// EnterCCDirectDeclaratorSuffix is called when entering the cCDirectDeclaratorSuffix production.
	EnterCCDirectDeclaratorSuffix(c *CCDirectDeclaratorSuffixContext)

	// EnterCCPlainDirectDeclaratorSuffix is called when entering the cCPlainDirectDeclaratorSuffix production.
	EnterCCPlainDirectDeclaratorSuffix(c *CCPlainDirectDeclaratorSuffixContext)

	// EnterCCAbstractDeclarator is called when entering the cCAbstractDeclarator production.
	EnterCCAbstractDeclarator(c *CCAbstractDeclaratorContext)

	// EnterCCAbstractDirectDeclarator is called when entering the cCAbstractDirectDeclarator production.
	EnterCCAbstractDirectDeclarator(c *CCAbstractDirectDeclaratorContext)

	// EnterCCAbstractDirectDeclaratorSuffix is called when entering the cCAbstractDirectDeclaratorSuffix production.
	EnterCCAbstractDirectDeclaratorSuffix(c *CCAbstractDirectDeclaratorSuffixContext)

	// EnterCCArraySuffix is called when entering the cCArraySuffix production.
	EnterCCArraySuffix(c *CCArraySuffixContext)

	// EnterCCFunctionSuffix is called when entering the cCFunctionSuffix production.
	EnterCCFunctionSuffix(c *CCFunctionSuffixContext)

	// EnterCCParameterTypeList is called when entering the cCParameterTypeList production.
	EnterCCParameterTypeList(c *CCParameterTypeListContext)

	// EnterCCParameterDeclaration is called when entering the cCParameterDeclaration production.
	EnterCCParameterDeclaration(c *CCParameterDeclarationContext)

	// EnterCCInitDeclaratorList is called when entering the cCInitDeclaratorList production.
	EnterCCInitDeclaratorList(c *CCInitDeclaratorListContext)

	// EnterCCInitDeclarator is called when entering the cCInitDeclarator production.
	EnterCCInitDeclarator(c *CCInitDeclaratorContext)

	// EnterCCInitializerClause is called when entering the cCInitializerClause production.
	EnterCCInitializerClause(c *CCInitializerClauseContext)

	// EnterCCBracedInitList is called when entering the cCBracedInitList production.
	EnterCCBracedInitList(c *CCBracedInitListContext)

	// EnterCCCompoundStatement is called when entering the cCCompoundStatement production.
	EnterCCCompoundStatement(c *CCCompoundStatementContext)

	// EnterCCBlockItem is called when entering the cCBlockItem production.
	EnterCCBlockItem(c *CCBlockItemContext)

	// EnterCCLinkageSpecification is called when entering the cCLinkageSpecification production.
	EnterCCLinkageSpecification(c *CCLinkageSpecificationContext)

	// EnterCCStatement is called when entering the cCStatement production.
	EnterCCStatement(c *CCStatementContext)

	// EnterCCExpression is called when entering the cCExpression production.
	EnterCCExpression(c *CCExpressionContext)

	// EnterCCAssignmentExpression is called when entering the cCAssignmentExpression production.
	EnterCCAssignmentExpression(c *CCAssignmentExpressionContext)

	// EnterCCAssignmentOperator is called when entering the cCAssignmentOperator production.
	EnterCCAssignmentOperator(c *CCAssignmentOperatorContext)

	// EnterCCConditionalExpression is called when entering the cCConditionalExpression production.
	EnterCCConditionalExpression(c *CCConditionalExpressionContext)

	// EnterCCLogicalOrExpression is called when entering the cCLogicalOrExpression production.
	EnterCCLogicalOrExpression(c *CCLogicalOrExpressionContext)

	// EnterCCLogicalAndExpression is called when entering the cCLogicalAndExpression production.
	EnterCCLogicalAndExpression(c *CCLogicalAndExpressionContext)

	// EnterCCInclusiveOrExpression is called when entering the cCInclusiveOrExpression production.
	EnterCCInclusiveOrExpression(c *CCInclusiveOrExpressionContext)

	// EnterCCExclusiveOrExpression is called when entering the cCExclusiveOrExpression production.
	EnterCCExclusiveOrExpression(c *CCExclusiveOrExpressionContext)

	// EnterCCAndExpression is called when entering the cCAndExpression production.
	EnterCCAndExpression(c *CCAndExpressionContext)

	// EnterCCEqualityExpression is called when entering the cCEqualityExpression production.
	EnterCCEqualityExpression(c *CCEqualityExpressionContext)

	// EnterCCRelationalExpression is called when entering the cCRelationalExpression production.
	EnterCCRelationalExpression(c *CCRelationalExpressionContext)

	// EnterCCAdditiveExpression is called when entering the cCAdditiveExpression production.
	EnterCCAdditiveExpression(c *CCAdditiveExpressionContext)

	// EnterCCMultiplicativeExpression is called when entering the cCMultiplicativeExpression production.
	EnterCCMultiplicativeExpression(c *CCMultiplicativeExpressionContext)

	// EnterCCUnaryExpression is called when entering the cCUnaryExpression production.
	EnterCCUnaryExpression(c *CCUnaryExpressionContext)

	// EnterCCUnaryOperator is called when entering the cCUnaryOperator production.
	EnterCCUnaryOperator(c *CCUnaryOperatorContext)

	// EnterCCPostfixExpression is called when entering the cCPostfixExpression production.
	EnterCCPostfixExpression(c *CCPostfixExpressionContext)

	// EnterCCPostfixSuffix is called when entering the cCPostfixSuffix production.
	EnterCCPostfixSuffix(c *CCPostfixSuffixContext)

	// EnterCCArgumentExpressionList is called when entering the cCArgumentExpressionList production.
	EnterCCArgumentExpressionList(c *CCArgumentExpressionListContext)

	// EnterCCPrimaryExpression is called when entering the cCPrimaryExpression production.
	EnterCCPrimaryExpression(c *CCPrimaryExpressionContext)

	// EnterCCConstantExpression is called when entering the cCConstantExpression production.
	EnterCCConstantExpression(c *CCConstantExpressionContext)

	// EnterBlockDeclaration is called when entering the blockDeclaration production.
	EnterBlockDeclaration(c *BlockDeclarationContext)

	// EnterAliasDeclaration is called when entering the aliasDeclaration production.
	EnterAliasDeclaration(c *AliasDeclarationContext)

	// EnterFunctionPointerDeclarator is called when entering the functionPointerDeclarator production.
	EnterFunctionPointerDeclarator(c *FunctionPointerDeclaratorContext)

	// EnterSimpleTypedefDeclarator is called when entering the simpleTypedefDeclarator production.
	EnterSimpleTypedefDeclarator(c *SimpleTypedefDeclaratorContext)

	// EnterTypedefDeclaration is called when entering the typedefDeclaration production.
	EnterTypedefDeclaration(c *TypedefDeclarationContext)

	// EnterSimpleDeclaration is called when entering the simpleDeclaration production.
	EnterSimpleDeclaration(c *SimpleDeclarationContext)

	// EnterStaticAssertDeclaration is called when entering the staticAssertDeclaration production.
	EnterStaticAssertDeclaration(c *StaticAssertDeclarationContext)

	// EnterEmptyDeclaration_ is called when entering the emptyDeclaration_ production.
	EnterEmptyDeclaration_(c *EmptyDeclaration_Context)

	// EnterAttributeDeclaration is called when entering the attributeDeclaration production.
	EnterAttributeDeclaration(c *AttributeDeclarationContext)

	// EnterDeclSpecifier is called when entering the declSpecifier production.
	EnterDeclSpecifier(c *DeclSpecifierContext)

	// EnterDeclSpecifierSeq is called when entering the declSpecifierSeq production.
	EnterDeclSpecifierSeq(c *DeclSpecifierSeqContext)

	// EnterStorageClassSpecifier is called when entering the storageClassSpecifier production.
	EnterStorageClassSpecifier(c *StorageClassSpecifierContext)

	// EnterFunctionSpecifier is called when entering the functionSpecifier production.
	EnterFunctionSpecifier(c *FunctionSpecifierContext)

	// EnterTypedefName is called when entering the typedefName production.
	EnterTypedefName(c *TypedefNameContext)

	// EnterTypeSpecifier is called when entering the typeSpecifier production.
	EnterTypeSpecifier(c *TypeSpecifierContext)

	// EnterTrailingTypeSpecifier is called when entering the trailingTypeSpecifier production.
	EnterTrailingTypeSpecifier(c *TrailingTypeSpecifierContext)

	// EnterTypeSpecifierSeq is called when entering the typeSpecifierSeq production.
	EnterTypeSpecifierSeq(c *TypeSpecifierSeqContext)

	// EnterTrailingTypeSpecifierSeq is called when entering the trailingTypeSpecifierSeq production.
	EnterTrailingTypeSpecifierSeq(c *TrailingTypeSpecifierSeqContext)

	// EnterSimpleTypeLengthModifier is called when entering the simpleTypeLengthModifier production.
	EnterSimpleTypeLengthModifier(c *SimpleTypeLengthModifierContext)

	// EnterSimpleTypeSignednessModifier is called when entering the simpleTypeSignednessModifier production.
	EnterSimpleTypeSignednessModifier(c *SimpleTypeSignednessModifierContext)

	// EnterSimpleTypeSpecifier is called when entering the simpleTypeSpecifier production.
	EnterSimpleTypeSpecifier(c *SimpleTypeSpecifierContext)

	// EnterTheTypeName is called when entering the theTypeName production.
	EnterTheTypeName(c *TheTypeNameContext)

	// EnterDecltypeSpecifier is called when entering the decltypeSpecifier production.
	EnterDecltypeSpecifier(c *DecltypeSpecifierContext)

	// EnterElaboratedTypeSpecifier is called when entering the elaboratedTypeSpecifier production.
	EnterElaboratedTypeSpecifier(c *ElaboratedTypeSpecifierContext)

	// EnterEnumName is called when entering the enumName production.
	EnterEnumName(c *EnumNameContext)

	// EnterEnumSpecifier is called when entering the enumSpecifier production.
	EnterEnumSpecifier(c *EnumSpecifierContext)

	// EnterEnumHead is called when entering the enumHead production.
	EnterEnumHead(c *EnumHeadContext)

	// EnterOpaqueEnumDeclaration is called when entering the opaqueEnumDeclaration production.
	EnterOpaqueEnumDeclaration(c *OpaqueEnumDeclarationContext)

	// EnterEnumkey is called when entering the enumkey production.
	EnterEnumkey(c *EnumkeyContext)

	// EnterEnumbase is called when entering the enumbase production.
	EnterEnumbase(c *EnumbaseContext)

	// EnterEnumeratorList is called when entering the enumeratorList production.
	EnterEnumeratorList(c *EnumeratorListContext)

	// EnterEnumeratorDefinition is called when entering the enumeratorDefinition production.
	EnterEnumeratorDefinition(c *EnumeratorDefinitionContext)

	// EnterEnumerator is called when entering the enumerator production.
	EnterEnumerator(c *EnumeratorContext)

	// EnterNamespaceName is called when entering the namespaceName production.
	EnterNamespaceName(c *NamespaceNameContext)

	// EnterOriginalNamespaceName is called when entering the originalNamespaceName production.
	EnterOriginalNamespaceName(c *OriginalNamespaceNameContext)

	// EnterNamespaceDefinition is called when entering the namespaceDefinition production.
	EnterNamespaceDefinition(c *NamespaceDefinitionContext)

	// EnterNamespaceAlias is called when entering the namespaceAlias production.
	EnterNamespaceAlias(c *NamespaceAliasContext)

	// EnterNamespaceAliasDefinition is called when entering the namespaceAliasDefinition production.
	EnterNamespaceAliasDefinition(c *NamespaceAliasDefinitionContext)

	// EnterQualifiednamespaceSpecifier is called when entering the qualifiednamespaceSpecifier production.
	EnterQualifiednamespaceSpecifier(c *QualifiednamespaceSpecifierContext)

	// EnterUsingDeclaration is called when entering the usingDeclaration production.
	EnterUsingDeclaration(c *UsingDeclarationContext)

	// EnterUsingDirective is called when entering the usingDirective production.
	EnterUsingDirective(c *UsingDirectiveContext)

	// EnterAsmDefinition is called when entering the asmDefinition production.
	EnterAsmDefinition(c *AsmDefinitionContext)

	// EnterLinkageSpecification is called when entering the linkageSpecification production.
	EnterLinkageSpecification(c *LinkageSpecificationContext)

	// EnterAttributeSpecifierSeq is called when entering the attributeSpecifierSeq production.
	EnterAttributeSpecifierSeq(c *AttributeSpecifierSeqContext)

	// EnterAttributeSpecifier is called when entering the attributeSpecifier production.
	EnterAttributeSpecifier(c *AttributeSpecifierContext)

	// EnterAlignmentSpecifier is called when entering the alignmentSpecifier production.
	EnterAlignmentSpecifier(c *AlignmentSpecifierContext)

	// EnterAttributeList is called when entering the attributeList production.
	EnterAttributeList(c *AttributeListContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterAttributeNamespace is called when entering the attributeNamespace production.
	EnterAttributeNamespace(c *AttributeNamespaceContext)

	// EnterAttributeArgumentClause is called when entering the attributeArgumentClause production.
	EnterAttributeArgumentClause(c *AttributeArgumentClauseContext)

	// EnterBalancedTokenSeq is called when entering the balancedTokenSeq production.
	EnterBalancedTokenSeq(c *BalancedTokenSeqContext)

	// EnterBalancedtoken is called when entering the balancedtoken production.
	EnterBalancedtoken(c *BalancedtokenContext)

	// EnterInitDeclarator is called when entering the initDeclarator production.
	EnterInitDeclarator(c *InitDeclaratorContext)

	// EnterDeclarator is called when entering the declarator production.
	EnterDeclarator(c *DeclaratorContext)

	// EnterPointerDeclarator is called when entering the pointerDeclarator production.
	EnterPointerDeclarator(c *PointerDeclaratorContext)

	// EnterNoPointerDeclarator is called when entering the noPointerDeclarator production.
	EnterNoPointerDeclarator(c *NoPointerDeclaratorContext)

	// EnterParametersAndQualifiers is called when entering the parametersAndQualifiers production.
	EnterParametersAndQualifiers(c *ParametersAndQualifiersContext)

	// EnterTrailingReturnType is called when entering the trailingReturnType production.
	EnterTrailingReturnType(c *TrailingReturnTypeContext)

	// EnterPointerOperator is called when entering the pointerOperator production.
	EnterPointerOperator(c *PointerOperatorContext)

	// EnterCvQualifierSeq is called when entering the cvQualifierSeq production.
	EnterCvQualifierSeq(c *CvQualifierSeqContext)

	// EnterCvQualifier is called when entering the cvQualifier production.
	EnterCvQualifier(c *CvQualifierContext)

	// EnterDeclaratorid is called when entering the declaratorid production.
	EnterDeclaratorid(c *DeclaratoridContext)

	// EnterTheTypeId is called when entering the theTypeId production.
	EnterTheTypeId(c *TheTypeIdContext)

	// EnterAbstractDeclarator is called when entering the abstractDeclarator production.
	EnterAbstractDeclarator(c *AbstractDeclaratorContext)

	// EnterPointerAbstractDeclarator is called when entering the pointerAbstractDeclarator production.
	EnterPointerAbstractDeclarator(c *PointerAbstractDeclaratorContext)

	// EnterNoPointerAbstractDeclarator is called when entering the noPointerAbstractDeclarator production.
	EnterNoPointerAbstractDeclarator(c *NoPointerAbstractDeclaratorContext)

	// EnterAbstractPackDeclarator is called when entering the abstractPackDeclarator production.
	EnterAbstractPackDeclarator(c *AbstractPackDeclaratorContext)

	// EnterNoPointerAbstractPackDeclarator is called when entering the noPointerAbstractPackDeclarator production.
	EnterNoPointerAbstractPackDeclarator(c *NoPointerAbstractPackDeclaratorContext)

	// EnterParameterDeclarationClause is called when entering the parameterDeclarationClause production.
	EnterParameterDeclarationClause(c *ParameterDeclarationClauseContext)

	// EnterParameterDeclarationList is called when entering the parameterDeclarationList production.
	EnterParameterDeclarationList(c *ParameterDeclarationListContext)

	// EnterParameterDeclaration is called when entering the parameterDeclaration production.
	EnterParameterDeclaration(c *ParameterDeclarationContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterFunctionDeclarator is called when entering the functionDeclarator production.
	EnterFunctionDeclarator(c *FunctionDeclaratorContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterInitializer is called when entering the initializer production.
	EnterInitializer(c *InitializerContext)

	// EnterBraceOrEqualInitializer is called when entering the braceOrEqualInitializer production.
	EnterBraceOrEqualInitializer(c *BraceOrEqualInitializerContext)

	// EnterInitializerClause is called when entering the initializerClause production.
	EnterInitializerClause(c *InitializerClauseContext)

	// EnterInitializerList is called when entering the initializerList production.
	EnterInitializerList(c *InitializerListContext)

	// EnterBracedInitList is called when entering the bracedInitList production.
	EnterBracedInitList(c *BracedInitListContext)

	// EnterClassName is called when entering the className production.
	EnterClassName(c *ClassNameContext)

	// EnterClassSpecifier is called when entering the classSpecifier production.
	EnterClassSpecifier(c *ClassSpecifierContext)

	// EnterClassHead is called when entering the classHead production.
	EnterClassHead(c *ClassHeadContext)

	// EnterClassHeadName is called when entering the classHeadName production.
	EnterClassHeadName(c *ClassHeadNameContext)

	// EnterClassVirtSpecifier is called when entering the classVirtSpecifier production.
	EnterClassVirtSpecifier(c *ClassVirtSpecifierContext)

	// EnterMemberSpecification is called when entering the memberSpecification production.
	EnterMemberSpecification(c *MemberSpecificationContext)

	// EnterMemberdeclaration is called when entering the memberdeclaration production.
	EnterMemberdeclaration(c *MemberdeclarationContext)

	// EnterMemberDeclaratorList is called when entering the memberDeclaratorList production.
	EnterMemberDeclaratorList(c *MemberDeclaratorListContext)

	// EnterMemberDeclarator is called when entering the memberDeclarator production.
	EnterMemberDeclarator(c *MemberDeclaratorContext)

	// EnterVirtualSpecifierSeq is called when entering the virtualSpecifierSeq production.
	EnterVirtualSpecifierSeq(c *VirtualSpecifierSeqContext)

	// EnterVirtualSpecifier is called when entering the virtualSpecifier production.
	EnterVirtualSpecifier(c *VirtualSpecifierContext)

	// EnterPureSpecifier is called when entering the pureSpecifier production.
	EnterPureSpecifier(c *PureSpecifierContext)

	// EnterBaseClause is called when entering the baseClause production.
	EnterBaseClause(c *BaseClauseContext)

	// EnterBaseSpecifierList is called when entering the baseSpecifierList production.
	EnterBaseSpecifierList(c *BaseSpecifierListContext)

	// EnterBaseSpecifier is called when entering the baseSpecifier production.
	EnterBaseSpecifier(c *BaseSpecifierContext)

	// EnterClassOrDeclType is called when entering the classOrDeclType production.
	EnterClassOrDeclType(c *ClassOrDeclTypeContext)

	// EnterBaseTypeSpecifier is called when entering the baseTypeSpecifier production.
	EnterBaseTypeSpecifier(c *BaseTypeSpecifierContext)

	// EnterAccessSpecifier is called when entering the accessSpecifier production.
	EnterAccessSpecifier(c *AccessSpecifierContext)

	// EnterConversionFunctionId is called when entering the conversionFunctionId production.
	EnterConversionFunctionId(c *ConversionFunctionIdContext)

	// EnterConversionTypeId is called when entering the conversionTypeId production.
	EnterConversionTypeId(c *ConversionTypeIdContext)

	// EnterConversionDeclarator is called when entering the conversionDeclarator production.
	EnterConversionDeclarator(c *ConversionDeclaratorContext)

	// EnterConstructorInitializer is called when entering the constructorInitializer production.
	EnterConstructorInitializer(c *ConstructorInitializerContext)

	// EnterMemInitializerList is called when entering the memInitializerList production.
	EnterMemInitializerList(c *MemInitializerListContext)

	// EnterMemInitializer is called when entering the memInitializer production.
	EnterMemInitializer(c *MemInitializerContext)

	// EnterMeminitializerid is called when entering the meminitializerid production.
	EnterMeminitializerid(c *MeminitializeridContext)

	// EnterOperatorFunctionId is called when entering the operatorFunctionId production.
	EnterOperatorFunctionId(c *OperatorFunctionIdContext)

	// EnterLiteralOperatorId is called when entering the literalOperatorId production.
	EnterLiteralOperatorId(c *LiteralOperatorIdContext)

	// EnterTemplateDeclaration is called when entering the templateDeclaration production.
	EnterTemplateDeclaration(c *TemplateDeclarationContext)

	// EnterTemplateparameterList is called when entering the templateparameterList production.
	EnterTemplateparameterList(c *TemplateparameterListContext)

	// EnterTemplateParameter is called when entering the templateParameter production.
	EnterTemplateParameter(c *TemplateParameterContext)

	// EnterTypeParameter is called when entering the typeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterSimpleTemplateId is called when entering the simpleTemplateId production.
	EnterSimpleTemplateId(c *SimpleTemplateIdContext)

	// EnterTemplateId is called when entering the templateId production.
	EnterTemplateId(c *TemplateIdContext)

	// EnterTemplateName is called when entering the templateName production.
	EnterTemplateName(c *TemplateNameContext)

	// EnterTemplateArgumentList is called when entering the templateArgumentList production.
	EnterTemplateArgumentList(c *TemplateArgumentListContext)

	// EnterTemplateArgument is called when entering the templateArgument production.
	EnterTemplateArgument(c *TemplateArgumentContext)

	// EnterTypeNameSpecifier is called when entering the typeNameSpecifier production.
	EnterTypeNameSpecifier(c *TypeNameSpecifierContext)

	// EnterExplicitInstantiation is called when entering the explicitInstantiation production.
	EnterExplicitInstantiation(c *ExplicitInstantiationContext)

	// EnterExplicitSpecialization is called when entering the explicitSpecialization production.
	EnterExplicitSpecialization(c *ExplicitSpecializationContext)

	// EnterTryBlock is called when entering the tryBlock production.
	EnterTryBlock(c *TryBlockContext)

	// EnterFunctionTryBlock is called when entering the functionTryBlock production.
	EnterFunctionTryBlock(c *FunctionTryBlockContext)

	// EnterHandlerSeq is called when entering the handlerSeq production.
	EnterHandlerSeq(c *HandlerSeqContext)

	// EnterHandler is called when entering the handler production.
	EnterHandler(c *HandlerContext)

	// EnterExceptionDeclaration is called when entering the exceptionDeclaration production.
	EnterExceptionDeclaration(c *ExceptionDeclarationContext)

	// EnterThrowExpression is called when entering the throwExpression production.
	EnterThrowExpression(c *ThrowExpressionContext)

	// EnterExceptionSpecification is called when entering the exceptionSpecification production.
	EnterExceptionSpecification(c *ExceptionSpecificationContext)

	// EnterDynamicExceptionSpecification is called when entering the dynamicExceptionSpecification production.
	EnterDynamicExceptionSpecification(c *DynamicExceptionSpecificationContext)

	// EnterTypeIdList is called when entering the typeIdList production.
	EnterTypeIdList(c *TypeIdListContext)

	// EnterNoeExceptSpecification is called when entering the noeExceptSpecification production.
	EnterNoeExceptSpecification(c *NoeExceptSpecificationContext)

	// EnterTheOperator is called when entering the theOperator production.
	EnterTheOperator(c *TheOperatorContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterOctalLiteral is called when entering the octalLiteral production.
	EnterOctalLiteral(c *OctalLiteralContext)

	// EnterHexadecimalLiteral is called when entering the hexadecimalLiteral production.
	EnterHexadecimalLiteral(c *HexadecimalLiteralContext)

	// EnterBinaryLiteral is called when entering the binaryLiteral production.
	EnterBinaryLiteral(c *BinaryLiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterCharacterLiteral is called when entering the characterLiteral production.
	EnterCharacterLiteral(c *CharacterLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterFloatingLiteral is called when entering the floatingLiteral production.
	EnterFloatingLiteral(c *FloatingLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterPointerLiteral is called when entering the pointerLiteral production.
	EnterPointerLiteral(c *PointerLiteralContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitTranslationUnit is called when exiting the translationUnit production.
	ExitTranslationUnit(c *TranslationUnitContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitIdExpression is called when exiting the idExpression production.
	ExitIdExpression(c *IdExpressionContext)

	// ExitUnqualifiedId is called when exiting the unqualifiedId production.
	ExitUnqualifiedId(c *UnqualifiedIdContext)

	// ExitQualifiedId is called when exiting the qualifiedId production.
	ExitQualifiedId(c *QualifiedIdContext)

	// ExitNestedNameSpecifier is called when exiting the nestedNameSpecifier production.
	ExitNestedNameSpecifier(c *NestedNameSpecifierContext)

	// ExitLambdaExpression is called when exiting the lambdaExpression production.
	ExitLambdaExpression(c *LambdaExpressionContext)

	// ExitLambdaIntroducer is called when exiting the lambdaIntroducer production.
	ExitLambdaIntroducer(c *LambdaIntroducerContext)

	// ExitLambdaCapture is called when exiting the lambdaCapture production.
	ExitLambdaCapture(c *LambdaCaptureContext)

	// ExitCaptureDefault is called when exiting the captureDefault production.
	ExitCaptureDefault(c *CaptureDefaultContext)

	// ExitCaptureList is called when exiting the captureList production.
	ExitCaptureList(c *CaptureListContext)

	// ExitCapture is called when exiting the capture production.
	ExitCapture(c *CaptureContext)

	// ExitSimpleCapture is called when exiting the simpleCapture production.
	ExitSimpleCapture(c *SimpleCaptureContext)

	// ExitInitcapture is called when exiting the initcapture production.
	ExitInitcapture(c *InitcaptureContext)

	// ExitLambdaDeclarator is called when exiting the lambdaDeclarator production.
	ExitLambdaDeclarator(c *LambdaDeclaratorContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitTypeIdOfTheTypeId is called when exiting the typeIdOfTheTypeId production.
	ExitTypeIdOfTheTypeId(c *TypeIdOfTheTypeIdContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitPseudoDestructorName is called when exiting the pseudoDestructorName production.
	ExitPseudoDestructorName(c *PseudoDestructorNameContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitNewExpression_ is called when exiting the newExpression_ production.
	ExitNewExpression_(c *NewExpression_Context)

	// ExitNewPlacement is called when exiting the newPlacement production.
	ExitNewPlacement(c *NewPlacementContext)

	// ExitNewTypeId is called when exiting the newTypeId production.
	ExitNewTypeId(c *NewTypeIdContext)

	// ExitNewDeclarator_ is called when exiting the newDeclarator_ production.
	ExitNewDeclarator_(c *NewDeclarator_Context)

	// ExitNoPointerNewDeclarator is called when exiting the noPointerNewDeclarator production.
	ExitNoPointerNewDeclarator(c *NoPointerNewDeclaratorContext)

	// ExitNewInitializer_ is called when exiting the newInitializer_ production.
	ExitNewInitializer_(c *NewInitializer_Context)

	// ExitDeleteExpression is called when exiting the deleteExpression production.
	ExitDeleteExpression(c *DeleteExpressionContext)

	// ExitNoExceptExpression is called when exiting the noExceptExpression production.
	ExitNoExceptExpression(c *NoExceptExpressionContext)

	// ExitCastExpression is called when exiting the castExpression production.
	ExitCastExpression(c *CastExpressionContext)

	// ExitPointerMemberExpression is called when exiting the pointerMemberExpression production.
	ExitPointerMemberExpression(c *PointerMemberExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitShiftExpression is called when exiting the shiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

	// ExitShiftOperator is called when exiting the shiftOperator production.
	ExitShiftOperator(c *ShiftOperatorContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitExclusiveOrExpression is called when exiting the exclusiveOrExpression production.
	ExitExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// ExitInclusiveOrExpression is called when exiting the inclusiveOrExpression production.
	ExitInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// ExitLogicalAndExpression is called when exiting the logicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitLogicalOrExpression is called when exiting the logicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitConditionalExpression is called when exiting the conditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConstantExpression is called when exiting the constantExpression production.
	ExitConstantExpression(c *ConstantExpressionContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLabeledStatement is called when exiting the labeledStatement production.
	ExitLabeledStatement(c *LabeledStatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)

	// ExitStatementSeq is called when exiting the statementSeq production.
	ExitStatementSeq(c *StatementSeqContext)

	// ExitSelectionStatement is called when exiting the selectionStatement production.
	ExitSelectionStatement(c *SelectionStatementContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitIterationStatement is called when exiting the iterationStatement production.
	ExitIterationStatement(c *IterationStatementContext)

	// ExitForInitStatement is called when exiting the forInitStatement production.
	ExitForInitStatement(c *ForInitStatementContext)

	// ExitForRangeDeclaration is called when exiting the forRangeDeclaration production.
	ExitForRangeDeclaration(c *ForRangeDeclarationContext)

	// ExitForRangeInitializer is called when exiting the forRangeInitializer production.
	ExitForRangeInitializer(c *ForRangeInitializerContext)

	// ExitJumpStatement is called when exiting the jumpStatement production.
	ExitJumpStatement(c *JumpStatementContext)

	// ExitDeclarationStatement is called when exiting the declarationStatement production.
	ExitDeclarationStatement(c *DeclarationStatementContext)

	// ExitDeclarationSeq is called when exiting the declarationSeq production.
	ExitDeclarationSeq(c *DeclarationSeqContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitCImportRegion is called when exiting the cImportRegion production.
	ExitCImportRegion(c *CImportRegionContext)

	// ExitCCDeclarationSeq is called when exiting the cCDeclarationSeq production.
	ExitCCDeclarationSeq(c *CCDeclarationSeqContext)

	// ExitCCExternalDeclaration is called when exiting the cCExternalDeclaration production.
	ExitCCExternalDeclaration(c *CCExternalDeclarationContext)

	// ExitCCAttributesOpt is called when exiting the cCAttributesOpt production.
	ExitCCAttributesOpt(c *CCAttributesOptContext)

	// ExitCCAttribute is called when exiting the cCAttribute production.
	ExitCCAttribute(c *CCAttributeContext)

	// ExitCCAttributeName is called when exiting the cCAttributeName production.
	ExitCCAttributeName(c *CCAttributeNameContext)

	// ExitCCParenGroup is called when exiting the cCParenGroup production.
	ExitCCParenGroup(c *CCParenGroupContext)

	// ExitCCDeclSpecifiers is called when exiting the cCDeclSpecifiers production.
	ExitCCDeclSpecifiers(c *CCDeclSpecifiersContext)

	// ExitCCDeclSpecifierPart is called when exiting the cCDeclSpecifierPart production.
	ExitCCDeclSpecifierPart(c *CCDeclSpecifierPartContext)

	// ExitCCQualifier is called when exiting the cCQualifier production.
	ExitCCQualifier(c *CCQualifierContext)

	// ExitCCTypedefName is called when exiting the cCTypedefName production.
	ExitCCTypedefName(c *CCTypedefNameContext)

	// ExitCCBuiltinTypeSpecifier is called when exiting the cCBuiltinTypeSpecifier production.
	ExitCCBuiltinTypeSpecifier(c *CCBuiltinTypeSpecifierContext)

	// ExitCCStructOrUnionSpecifier is called when exiting the cCStructOrUnionSpecifier production.
	ExitCCStructOrUnionSpecifier(c *CCStructOrUnionSpecifierContext)

	// ExitCCMemberDeclaration is called when exiting the cCMemberDeclaration production.
	ExitCCMemberDeclaration(c *CCMemberDeclarationContext)

	// ExitCCMemberDeclaratorList is called when exiting the cCMemberDeclaratorList production.
	ExitCCMemberDeclaratorList(c *CCMemberDeclaratorListContext)

	// ExitCCMemberDeclarator is called when exiting the cCMemberDeclarator production.
	ExitCCMemberDeclarator(c *CCMemberDeclaratorContext)

	// ExitCCBitField is called when exiting the cCBitField production.
	ExitCCBitField(c *CCBitFieldContext)

	// ExitCCEnumSpecifier is called when exiting the cCEnumSpecifier production.
	ExitCCEnumSpecifier(c *CCEnumSpecifierContext)

	// ExitCCEnumeratorList is called when exiting the cCEnumeratorList production.
	ExitCCEnumeratorList(c *CCEnumeratorListContext)

	// ExitCCEnumerator is called when exiting the cCEnumerator production.
	ExitCCEnumerator(c *CCEnumeratorContext)

	// ExitCCTypeofSpecifier is called when exiting the cCTypeofSpecifier production.
	ExitCCTypeofSpecifier(c *CCTypeofSpecifierContext)

	// ExitCCTypeName is called when exiting the cCTypeName production.
	ExitCCTypeName(c *CCTypeNameContext)

	// ExitCCIdentifier is called when exiting the cCIdentifier production.
	ExitCCIdentifier(c *CCIdentifierContext)

	// ExitCCDeclarator is called when exiting the cCDeclarator production.
	ExitCCDeclarator(c *CCDeclaratorContext)

	// ExitCCPlainDeclarator is called when exiting the cCPlainDeclarator production.
	ExitCCPlainDeclarator(c *CCPlainDeclaratorContext)

	// ExitCCPointerOperator is called when exiting the cCPointerOperator production.
	ExitCCPointerOperator(c *CCPointerOperatorContext)

	// ExitCCPointerQualifier is called when exiting the cCPointerQualifier production.
	ExitCCPointerQualifier(c *CCPointerQualifierContext)

	// ExitCCDirectDeclarator is called when exiting the cCDirectDeclarator production.
	ExitCCDirectDeclarator(c *CCDirectDeclaratorContext)

	// ExitCCPlainDirectDeclarator is called when exiting the cCPlainDirectDeclarator production.
	ExitCCPlainDirectDeclarator(c *CCPlainDirectDeclaratorContext)

	// ExitCCDirectDeclaratorSuffix is called when exiting the cCDirectDeclaratorSuffix production.
	ExitCCDirectDeclaratorSuffix(c *CCDirectDeclaratorSuffixContext)

	// ExitCCPlainDirectDeclaratorSuffix is called when exiting the cCPlainDirectDeclaratorSuffix production.
	ExitCCPlainDirectDeclaratorSuffix(c *CCPlainDirectDeclaratorSuffixContext)

	// ExitCCAbstractDeclarator is called when exiting the cCAbstractDeclarator production.
	ExitCCAbstractDeclarator(c *CCAbstractDeclaratorContext)

	// ExitCCAbstractDirectDeclarator is called when exiting the cCAbstractDirectDeclarator production.
	ExitCCAbstractDirectDeclarator(c *CCAbstractDirectDeclaratorContext)

	// ExitCCAbstractDirectDeclaratorSuffix is called when exiting the cCAbstractDirectDeclaratorSuffix production.
	ExitCCAbstractDirectDeclaratorSuffix(c *CCAbstractDirectDeclaratorSuffixContext)

	// ExitCCArraySuffix is called when exiting the cCArraySuffix production.
	ExitCCArraySuffix(c *CCArraySuffixContext)

	// ExitCCFunctionSuffix is called when exiting the cCFunctionSuffix production.
	ExitCCFunctionSuffix(c *CCFunctionSuffixContext)

	// ExitCCParameterTypeList is called when exiting the cCParameterTypeList production.
	ExitCCParameterTypeList(c *CCParameterTypeListContext)

	// ExitCCParameterDeclaration is called when exiting the cCParameterDeclaration production.
	ExitCCParameterDeclaration(c *CCParameterDeclarationContext)

	// ExitCCInitDeclaratorList is called when exiting the cCInitDeclaratorList production.
	ExitCCInitDeclaratorList(c *CCInitDeclaratorListContext)

	// ExitCCInitDeclarator is called when exiting the cCInitDeclarator production.
	ExitCCInitDeclarator(c *CCInitDeclaratorContext)

	// ExitCCInitializerClause is called when exiting the cCInitializerClause production.
	ExitCCInitializerClause(c *CCInitializerClauseContext)

	// ExitCCBracedInitList is called when exiting the cCBracedInitList production.
	ExitCCBracedInitList(c *CCBracedInitListContext)

	// ExitCCCompoundStatement is called when exiting the cCCompoundStatement production.
	ExitCCCompoundStatement(c *CCCompoundStatementContext)

	// ExitCCBlockItem is called when exiting the cCBlockItem production.
	ExitCCBlockItem(c *CCBlockItemContext)

	// ExitCCLinkageSpecification is called when exiting the cCLinkageSpecification production.
	ExitCCLinkageSpecification(c *CCLinkageSpecificationContext)

	// ExitCCStatement is called when exiting the cCStatement production.
	ExitCCStatement(c *CCStatementContext)

	// ExitCCExpression is called when exiting the cCExpression production.
	ExitCCExpression(c *CCExpressionContext)

	// ExitCCAssignmentExpression is called when exiting the cCAssignmentExpression production.
	ExitCCAssignmentExpression(c *CCAssignmentExpressionContext)

	// ExitCCAssignmentOperator is called when exiting the cCAssignmentOperator production.
	ExitCCAssignmentOperator(c *CCAssignmentOperatorContext)

	// ExitCCConditionalExpression is called when exiting the cCConditionalExpression production.
	ExitCCConditionalExpression(c *CCConditionalExpressionContext)

	// ExitCCLogicalOrExpression is called when exiting the cCLogicalOrExpression production.
	ExitCCLogicalOrExpression(c *CCLogicalOrExpressionContext)

	// ExitCCLogicalAndExpression is called when exiting the cCLogicalAndExpression production.
	ExitCCLogicalAndExpression(c *CCLogicalAndExpressionContext)

	// ExitCCInclusiveOrExpression is called when exiting the cCInclusiveOrExpression production.
	ExitCCInclusiveOrExpression(c *CCInclusiveOrExpressionContext)

	// ExitCCExclusiveOrExpression is called when exiting the cCExclusiveOrExpression production.
	ExitCCExclusiveOrExpression(c *CCExclusiveOrExpressionContext)

	// ExitCCAndExpression is called when exiting the cCAndExpression production.
	ExitCCAndExpression(c *CCAndExpressionContext)

	// ExitCCEqualityExpression is called when exiting the cCEqualityExpression production.
	ExitCCEqualityExpression(c *CCEqualityExpressionContext)

	// ExitCCRelationalExpression is called when exiting the cCRelationalExpression production.
	ExitCCRelationalExpression(c *CCRelationalExpressionContext)

	// ExitCCAdditiveExpression is called when exiting the cCAdditiveExpression production.
	ExitCCAdditiveExpression(c *CCAdditiveExpressionContext)

	// ExitCCMultiplicativeExpression is called when exiting the cCMultiplicativeExpression production.
	ExitCCMultiplicativeExpression(c *CCMultiplicativeExpressionContext)

	// ExitCCUnaryExpression is called when exiting the cCUnaryExpression production.
	ExitCCUnaryExpression(c *CCUnaryExpressionContext)

	// ExitCCUnaryOperator is called when exiting the cCUnaryOperator production.
	ExitCCUnaryOperator(c *CCUnaryOperatorContext)

	// ExitCCPostfixExpression is called when exiting the cCPostfixExpression production.
	ExitCCPostfixExpression(c *CCPostfixExpressionContext)

	// ExitCCPostfixSuffix is called when exiting the cCPostfixSuffix production.
	ExitCCPostfixSuffix(c *CCPostfixSuffixContext)

	// ExitCCArgumentExpressionList is called when exiting the cCArgumentExpressionList production.
	ExitCCArgumentExpressionList(c *CCArgumentExpressionListContext)

	// ExitCCPrimaryExpression is called when exiting the cCPrimaryExpression production.
	ExitCCPrimaryExpression(c *CCPrimaryExpressionContext)

	// ExitCCConstantExpression is called when exiting the cCConstantExpression production.
	ExitCCConstantExpression(c *CCConstantExpressionContext)

	// ExitBlockDeclaration is called when exiting the blockDeclaration production.
	ExitBlockDeclaration(c *BlockDeclarationContext)

	// ExitAliasDeclaration is called when exiting the aliasDeclaration production.
	ExitAliasDeclaration(c *AliasDeclarationContext)

	// ExitFunctionPointerDeclarator is called when exiting the functionPointerDeclarator production.
	ExitFunctionPointerDeclarator(c *FunctionPointerDeclaratorContext)

	// ExitSimpleTypedefDeclarator is called when exiting the simpleTypedefDeclarator production.
	ExitSimpleTypedefDeclarator(c *SimpleTypedefDeclaratorContext)

	// ExitTypedefDeclaration is called when exiting the typedefDeclaration production.
	ExitTypedefDeclaration(c *TypedefDeclarationContext)

	// ExitSimpleDeclaration is called when exiting the simpleDeclaration production.
	ExitSimpleDeclaration(c *SimpleDeclarationContext)

	// ExitStaticAssertDeclaration is called when exiting the staticAssertDeclaration production.
	ExitStaticAssertDeclaration(c *StaticAssertDeclarationContext)

	// ExitEmptyDeclaration_ is called when exiting the emptyDeclaration_ production.
	ExitEmptyDeclaration_(c *EmptyDeclaration_Context)

	// ExitAttributeDeclaration is called when exiting the attributeDeclaration production.
	ExitAttributeDeclaration(c *AttributeDeclarationContext)

	// ExitDeclSpecifier is called when exiting the declSpecifier production.
	ExitDeclSpecifier(c *DeclSpecifierContext)

	// ExitDeclSpecifierSeq is called when exiting the declSpecifierSeq production.
	ExitDeclSpecifierSeq(c *DeclSpecifierSeqContext)

	// ExitStorageClassSpecifier is called when exiting the storageClassSpecifier production.
	ExitStorageClassSpecifier(c *StorageClassSpecifierContext)

	// ExitFunctionSpecifier is called when exiting the functionSpecifier production.
	ExitFunctionSpecifier(c *FunctionSpecifierContext)

	// ExitTypedefName is called when exiting the typedefName production.
	ExitTypedefName(c *TypedefNameContext)

	// ExitTypeSpecifier is called when exiting the typeSpecifier production.
	ExitTypeSpecifier(c *TypeSpecifierContext)

	// ExitTrailingTypeSpecifier is called when exiting the trailingTypeSpecifier production.
	ExitTrailingTypeSpecifier(c *TrailingTypeSpecifierContext)

	// ExitTypeSpecifierSeq is called when exiting the typeSpecifierSeq production.
	ExitTypeSpecifierSeq(c *TypeSpecifierSeqContext)

	// ExitTrailingTypeSpecifierSeq is called when exiting the trailingTypeSpecifierSeq production.
	ExitTrailingTypeSpecifierSeq(c *TrailingTypeSpecifierSeqContext)

	// ExitSimpleTypeLengthModifier is called when exiting the simpleTypeLengthModifier production.
	ExitSimpleTypeLengthModifier(c *SimpleTypeLengthModifierContext)

	// ExitSimpleTypeSignednessModifier is called when exiting the simpleTypeSignednessModifier production.
	ExitSimpleTypeSignednessModifier(c *SimpleTypeSignednessModifierContext)

	// ExitSimpleTypeSpecifier is called when exiting the simpleTypeSpecifier production.
	ExitSimpleTypeSpecifier(c *SimpleTypeSpecifierContext)

	// ExitTheTypeName is called when exiting the theTypeName production.
	ExitTheTypeName(c *TheTypeNameContext)

	// ExitDecltypeSpecifier is called when exiting the decltypeSpecifier production.
	ExitDecltypeSpecifier(c *DecltypeSpecifierContext)

	// ExitElaboratedTypeSpecifier is called when exiting the elaboratedTypeSpecifier production.
	ExitElaboratedTypeSpecifier(c *ElaboratedTypeSpecifierContext)

	// ExitEnumName is called when exiting the enumName production.
	ExitEnumName(c *EnumNameContext)

	// ExitEnumSpecifier is called when exiting the enumSpecifier production.
	ExitEnumSpecifier(c *EnumSpecifierContext)

	// ExitEnumHead is called when exiting the enumHead production.
	ExitEnumHead(c *EnumHeadContext)

	// ExitOpaqueEnumDeclaration is called when exiting the opaqueEnumDeclaration production.
	ExitOpaqueEnumDeclaration(c *OpaqueEnumDeclarationContext)

	// ExitEnumkey is called when exiting the enumkey production.
	ExitEnumkey(c *EnumkeyContext)

	// ExitEnumbase is called when exiting the enumbase production.
	ExitEnumbase(c *EnumbaseContext)

	// ExitEnumeratorList is called when exiting the enumeratorList production.
	ExitEnumeratorList(c *EnumeratorListContext)

	// ExitEnumeratorDefinition is called when exiting the enumeratorDefinition production.
	ExitEnumeratorDefinition(c *EnumeratorDefinitionContext)

	// ExitEnumerator is called when exiting the enumerator production.
	ExitEnumerator(c *EnumeratorContext)

	// ExitNamespaceName is called when exiting the namespaceName production.
	ExitNamespaceName(c *NamespaceNameContext)

	// ExitOriginalNamespaceName is called when exiting the originalNamespaceName production.
	ExitOriginalNamespaceName(c *OriginalNamespaceNameContext)

	// ExitNamespaceDefinition is called when exiting the namespaceDefinition production.
	ExitNamespaceDefinition(c *NamespaceDefinitionContext)

	// ExitNamespaceAlias is called when exiting the namespaceAlias production.
	ExitNamespaceAlias(c *NamespaceAliasContext)

	// ExitNamespaceAliasDefinition is called when exiting the namespaceAliasDefinition production.
	ExitNamespaceAliasDefinition(c *NamespaceAliasDefinitionContext)

	// ExitQualifiednamespaceSpecifier is called when exiting the qualifiednamespaceSpecifier production.
	ExitQualifiednamespaceSpecifier(c *QualifiednamespaceSpecifierContext)

	// ExitUsingDeclaration is called when exiting the usingDeclaration production.
	ExitUsingDeclaration(c *UsingDeclarationContext)

	// ExitUsingDirective is called when exiting the usingDirective production.
	ExitUsingDirective(c *UsingDirectiveContext)

	// ExitAsmDefinition is called when exiting the asmDefinition production.
	ExitAsmDefinition(c *AsmDefinitionContext)

	// ExitLinkageSpecification is called when exiting the linkageSpecification production.
	ExitLinkageSpecification(c *LinkageSpecificationContext)

	// ExitAttributeSpecifierSeq is called when exiting the attributeSpecifierSeq production.
	ExitAttributeSpecifierSeq(c *AttributeSpecifierSeqContext)

	// ExitAttributeSpecifier is called when exiting the attributeSpecifier production.
	ExitAttributeSpecifier(c *AttributeSpecifierContext)

	// ExitAlignmentSpecifier is called when exiting the alignmentSpecifier production.
	ExitAlignmentSpecifier(c *AlignmentSpecifierContext)

	// ExitAttributeList is called when exiting the attributeList production.
	ExitAttributeList(c *AttributeListContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitAttributeNamespace is called when exiting the attributeNamespace production.
	ExitAttributeNamespace(c *AttributeNamespaceContext)

	// ExitAttributeArgumentClause is called when exiting the attributeArgumentClause production.
	ExitAttributeArgumentClause(c *AttributeArgumentClauseContext)

	// ExitBalancedTokenSeq is called when exiting the balancedTokenSeq production.
	ExitBalancedTokenSeq(c *BalancedTokenSeqContext)

	// ExitBalancedtoken is called when exiting the balancedtoken production.
	ExitBalancedtoken(c *BalancedtokenContext)

	// ExitInitDeclarator is called when exiting the initDeclarator production.
	ExitInitDeclarator(c *InitDeclaratorContext)

	// ExitDeclarator is called when exiting the declarator production.
	ExitDeclarator(c *DeclaratorContext)

	// ExitPointerDeclarator is called when exiting the pointerDeclarator production.
	ExitPointerDeclarator(c *PointerDeclaratorContext)

	// ExitNoPointerDeclarator is called when exiting the noPointerDeclarator production.
	ExitNoPointerDeclarator(c *NoPointerDeclaratorContext)

	// ExitParametersAndQualifiers is called when exiting the parametersAndQualifiers production.
	ExitParametersAndQualifiers(c *ParametersAndQualifiersContext)

	// ExitTrailingReturnType is called when exiting the trailingReturnType production.
	ExitTrailingReturnType(c *TrailingReturnTypeContext)

	// ExitPointerOperator is called when exiting the pointerOperator production.
	ExitPointerOperator(c *PointerOperatorContext)

	// ExitCvQualifierSeq is called when exiting the cvQualifierSeq production.
	ExitCvQualifierSeq(c *CvQualifierSeqContext)

	// ExitCvQualifier is called when exiting the cvQualifier production.
	ExitCvQualifier(c *CvQualifierContext)

	// ExitDeclaratorid is called when exiting the declaratorid production.
	ExitDeclaratorid(c *DeclaratoridContext)

	// ExitTheTypeId is called when exiting the theTypeId production.
	ExitTheTypeId(c *TheTypeIdContext)

	// ExitAbstractDeclarator is called when exiting the abstractDeclarator production.
	ExitAbstractDeclarator(c *AbstractDeclaratorContext)

	// ExitPointerAbstractDeclarator is called when exiting the pointerAbstractDeclarator production.
	ExitPointerAbstractDeclarator(c *PointerAbstractDeclaratorContext)

	// ExitNoPointerAbstractDeclarator is called when exiting the noPointerAbstractDeclarator production.
	ExitNoPointerAbstractDeclarator(c *NoPointerAbstractDeclaratorContext)

	// ExitAbstractPackDeclarator is called when exiting the abstractPackDeclarator production.
	ExitAbstractPackDeclarator(c *AbstractPackDeclaratorContext)

	// ExitNoPointerAbstractPackDeclarator is called when exiting the noPointerAbstractPackDeclarator production.
	ExitNoPointerAbstractPackDeclarator(c *NoPointerAbstractPackDeclaratorContext)

	// ExitParameterDeclarationClause is called when exiting the parameterDeclarationClause production.
	ExitParameterDeclarationClause(c *ParameterDeclarationClauseContext)

	// ExitParameterDeclarationList is called when exiting the parameterDeclarationList production.
	ExitParameterDeclarationList(c *ParameterDeclarationListContext)

	// ExitParameterDeclaration is called when exiting the parameterDeclaration production.
	ExitParameterDeclaration(c *ParameterDeclarationContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitFunctionDeclarator is called when exiting the functionDeclarator production.
	ExitFunctionDeclarator(c *FunctionDeclaratorContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitInitializer is called when exiting the initializer production.
	ExitInitializer(c *InitializerContext)

	// ExitBraceOrEqualInitializer is called when exiting the braceOrEqualInitializer production.
	ExitBraceOrEqualInitializer(c *BraceOrEqualInitializerContext)

	// ExitInitializerClause is called when exiting the initializerClause production.
	ExitInitializerClause(c *InitializerClauseContext)

	// ExitInitializerList is called when exiting the initializerList production.
	ExitInitializerList(c *InitializerListContext)

	// ExitBracedInitList is called when exiting the bracedInitList production.
	ExitBracedInitList(c *BracedInitListContext)

	// ExitClassName is called when exiting the className production.
	ExitClassName(c *ClassNameContext)

	// ExitClassSpecifier is called when exiting the classSpecifier production.
	ExitClassSpecifier(c *ClassSpecifierContext)

	// ExitClassHead is called when exiting the classHead production.
	ExitClassHead(c *ClassHeadContext)

	// ExitClassHeadName is called when exiting the classHeadName production.
	ExitClassHeadName(c *ClassHeadNameContext)

	// ExitClassVirtSpecifier is called when exiting the classVirtSpecifier production.
	ExitClassVirtSpecifier(c *ClassVirtSpecifierContext)

	// ExitMemberSpecification is called when exiting the memberSpecification production.
	ExitMemberSpecification(c *MemberSpecificationContext)

	// ExitMemberdeclaration is called when exiting the memberdeclaration production.
	ExitMemberdeclaration(c *MemberdeclarationContext)

	// ExitMemberDeclaratorList is called when exiting the memberDeclaratorList production.
	ExitMemberDeclaratorList(c *MemberDeclaratorListContext)

	// ExitMemberDeclarator is called when exiting the memberDeclarator production.
	ExitMemberDeclarator(c *MemberDeclaratorContext)

	// ExitVirtualSpecifierSeq is called when exiting the virtualSpecifierSeq production.
	ExitVirtualSpecifierSeq(c *VirtualSpecifierSeqContext)

	// ExitVirtualSpecifier is called when exiting the virtualSpecifier production.
	ExitVirtualSpecifier(c *VirtualSpecifierContext)

	// ExitPureSpecifier is called when exiting the pureSpecifier production.
	ExitPureSpecifier(c *PureSpecifierContext)

	// ExitBaseClause is called when exiting the baseClause production.
	ExitBaseClause(c *BaseClauseContext)

	// ExitBaseSpecifierList is called when exiting the baseSpecifierList production.
	ExitBaseSpecifierList(c *BaseSpecifierListContext)

	// ExitBaseSpecifier is called when exiting the baseSpecifier production.
	ExitBaseSpecifier(c *BaseSpecifierContext)

	// ExitClassOrDeclType is called when exiting the classOrDeclType production.
	ExitClassOrDeclType(c *ClassOrDeclTypeContext)

	// ExitBaseTypeSpecifier is called when exiting the baseTypeSpecifier production.
	ExitBaseTypeSpecifier(c *BaseTypeSpecifierContext)

	// ExitAccessSpecifier is called when exiting the accessSpecifier production.
	ExitAccessSpecifier(c *AccessSpecifierContext)

	// ExitConversionFunctionId is called when exiting the conversionFunctionId production.
	ExitConversionFunctionId(c *ConversionFunctionIdContext)

	// ExitConversionTypeId is called when exiting the conversionTypeId production.
	ExitConversionTypeId(c *ConversionTypeIdContext)

	// ExitConversionDeclarator is called when exiting the conversionDeclarator production.
	ExitConversionDeclarator(c *ConversionDeclaratorContext)

	// ExitConstructorInitializer is called when exiting the constructorInitializer production.
	ExitConstructorInitializer(c *ConstructorInitializerContext)

	// ExitMemInitializerList is called when exiting the memInitializerList production.
	ExitMemInitializerList(c *MemInitializerListContext)

	// ExitMemInitializer is called when exiting the memInitializer production.
	ExitMemInitializer(c *MemInitializerContext)

	// ExitMeminitializerid is called when exiting the meminitializerid production.
	ExitMeminitializerid(c *MeminitializeridContext)

	// ExitOperatorFunctionId is called when exiting the operatorFunctionId production.
	ExitOperatorFunctionId(c *OperatorFunctionIdContext)

	// ExitLiteralOperatorId is called when exiting the literalOperatorId production.
	ExitLiteralOperatorId(c *LiteralOperatorIdContext)

	// ExitTemplateDeclaration is called when exiting the templateDeclaration production.
	ExitTemplateDeclaration(c *TemplateDeclarationContext)

	// ExitTemplateparameterList is called when exiting the templateparameterList production.
	ExitTemplateparameterList(c *TemplateparameterListContext)

	// ExitTemplateParameter is called when exiting the templateParameter production.
	ExitTemplateParameter(c *TemplateParameterContext)

	// ExitTypeParameter is called when exiting the typeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitSimpleTemplateId is called when exiting the simpleTemplateId production.
	ExitSimpleTemplateId(c *SimpleTemplateIdContext)

	// ExitTemplateId is called when exiting the templateId production.
	ExitTemplateId(c *TemplateIdContext)

	// ExitTemplateName is called when exiting the templateName production.
	ExitTemplateName(c *TemplateNameContext)

	// ExitTemplateArgumentList is called when exiting the templateArgumentList production.
	ExitTemplateArgumentList(c *TemplateArgumentListContext)

	// ExitTemplateArgument is called when exiting the templateArgument production.
	ExitTemplateArgument(c *TemplateArgumentContext)

	// ExitTypeNameSpecifier is called when exiting the typeNameSpecifier production.
	ExitTypeNameSpecifier(c *TypeNameSpecifierContext)

	// ExitExplicitInstantiation is called when exiting the explicitInstantiation production.
	ExitExplicitInstantiation(c *ExplicitInstantiationContext)

	// ExitExplicitSpecialization is called when exiting the explicitSpecialization production.
	ExitExplicitSpecialization(c *ExplicitSpecializationContext)

	// ExitTryBlock is called when exiting the tryBlock production.
	ExitTryBlock(c *TryBlockContext)

	// ExitFunctionTryBlock is called when exiting the functionTryBlock production.
	ExitFunctionTryBlock(c *FunctionTryBlockContext)

	// ExitHandlerSeq is called when exiting the handlerSeq production.
	ExitHandlerSeq(c *HandlerSeqContext)

	// ExitHandler is called when exiting the handler production.
	ExitHandler(c *HandlerContext)

	// ExitExceptionDeclaration is called when exiting the exceptionDeclaration production.
	ExitExceptionDeclaration(c *ExceptionDeclarationContext)

	// ExitThrowExpression is called when exiting the throwExpression production.
	ExitThrowExpression(c *ThrowExpressionContext)

	// ExitExceptionSpecification is called when exiting the exceptionSpecification production.
	ExitExceptionSpecification(c *ExceptionSpecificationContext)

	// ExitDynamicExceptionSpecification is called when exiting the dynamicExceptionSpecification production.
	ExitDynamicExceptionSpecification(c *DynamicExceptionSpecificationContext)

	// ExitTypeIdList is called when exiting the typeIdList production.
	ExitTypeIdList(c *TypeIdListContext)

	// ExitNoeExceptSpecification is called when exiting the noeExceptSpecification production.
	ExitNoeExceptSpecification(c *NoeExceptSpecificationContext)

	// ExitTheOperator is called when exiting the theOperator production.
	ExitTheOperator(c *TheOperatorContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitOctalLiteral is called when exiting the octalLiteral production.
	ExitOctalLiteral(c *OctalLiteralContext)

	// ExitHexadecimalLiteral is called when exiting the hexadecimalLiteral production.
	ExitHexadecimalLiteral(c *HexadecimalLiteralContext)

	// ExitBinaryLiteral is called when exiting the binaryLiteral production.
	ExitBinaryLiteral(c *BinaryLiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitCharacterLiteral is called when exiting the characterLiteral production.
	ExitCharacterLiteral(c *CharacterLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitFloatingLiteral is called when exiting the floatingLiteral production.
	ExitFloatingLiteral(c *FloatingLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitPointerLiteral is called when exiting the pointerLiteral production.
	ExitPointerLiteral(c *PointerLiteralContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
