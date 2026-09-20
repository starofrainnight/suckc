// Code generated from SuckCParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SuckCParser

import "github.com/antlr4-go/antlr/v4"

type BaseSuckCParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSuckCParserVisitor) VisitTranslationUnit(ctx *TranslationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitIdExpression(ctx *IdExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitUnqualifiedId(ctx *UnqualifiedIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitQualifiedId(ctx *QualifiedIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNestedNameSpecifier(ctx *NestedNameSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitLambdaExpression(ctx *LambdaExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitLambdaIntroducer(ctx *LambdaIntroducerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitLambdaCapture(ctx *LambdaCaptureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitCaptureDefault(ctx *CaptureDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitCaptureList(ctx *CaptureListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitCapture(ctx *CaptureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitSimpleCapture(ctx *SimpleCaptureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitInitcapture(ctx *InitcaptureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitLambdaDeclarator(ctx *LambdaDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTypeIdOfTheTypeId(ctx *TypeIdOfTheTypeIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitPseudoDestructorName(ctx *PseudoDestructorNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitUnaryOperator(ctx *UnaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNewExpression_(ctx *NewExpression_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNewPlacement(ctx *NewPlacementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNewTypeId(ctx *NewTypeIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNewDeclarator_(ctx *NewDeclarator_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNoPointerNewDeclarator(ctx *NoPointerNewDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNewInitializer_(ctx *NewInitializer_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitDeleteExpression(ctx *DeleteExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNoExceptExpression(ctx *NoExceptExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitCastExpression(ctx *CastExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitPointerMemberExpression(ctx *PointerMemberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitShiftExpression(ctx *ShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitShiftOperator(ctx *ShiftOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitConstantExpression(ctx *ConstantExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitLabeledStatement(ctx *LabeledStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitCompoundStatement(ctx *CompoundStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitStatementSeq(ctx *StatementSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitSelectionStatement(ctx *SelectionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitIterationStatement(ctx *IterationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitForInitStatement(ctx *ForInitStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitForRangeDeclaration(ctx *ForRangeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitForRangeInitializer(ctx *ForRangeInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitJumpStatement(ctx *JumpStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitDeclarationStatement(ctx *DeclarationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitDeclarationSeq(ctx *DeclarationSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitBlockDeclaration(ctx *BlockDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAliasDeclaration(ctx *AliasDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitFunctionPointerDeclarator(ctx *FunctionPointerDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitSimpleTypedefDeclarator(ctx *SimpleTypedefDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTypedefDeclaration(ctx *TypedefDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitSimpleDeclaration(ctx *SimpleDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitEmptyDeclaration_(ctx *EmptyDeclaration_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAttributeDeclaration(ctx *AttributeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitDeclSpecifier(ctx *DeclSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitDeclSpecifierSeq(ctx *DeclSpecifierSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitStorageClassSpecifier(ctx *StorageClassSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitFunctionSpecifier(ctx *FunctionSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTypedefName(ctx *TypedefNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTrailingTypeSpecifier(ctx *TrailingTypeSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTypeSpecifierSeq(ctx *TypeSpecifierSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTrailingTypeSpecifierSeq(ctx *TrailingTypeSpecifierSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitSimpleTypeLengthModifier(ctx *SimpleTypeLengthModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitSimpleTypeSignednessModifier(ctx *SimpleTypeSignednessModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitSimpleTypeSpecifier(ctx *SimpleTypeSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTheTypeName(ctx *TheTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitDecltypeSpecifier(ctx *DecltypeSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitElaboratedTypeSpecifier(ctx *ElaboratedTypeSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitEnumName(ctx *EnumNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitEnumSpecifier(ctx *EnumSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitEnumHead(ctx *EnumHeadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitOpaqueEnumDeclaration(ctx *OpaqueEnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitEnumkey(ctx *EnumkeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitEnumbase(ctx *EnumbaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitEnumeratorList(ctx *EnumeratorListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitEnumeratorDefinition(ctx *EnumeratorDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitEnumerator(ctx *EnumeratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNamespaceName(ctx *NamespaceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitOriginalNamespaceName(ctx *OriginalNamespaceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNamespaceDefinition(ctx *NamespaceDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNamespaceAlias(ctx *NamespaceAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNamespaceAliasDefinition(ctx *NamespaceAliasDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitQualifiednamespaceSpecifier(ctx *QualifiednamespaceSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitUsingDeclaration(ctx *UsingDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitUsingDirective(ctx *UsingDirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAsmDefinition(ctx *AsmDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitLinkageSpecification(ctx *LinkageSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAttributeSpecifierSeq(ctx *AttributeSpecifierSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAttributeSpecifier(ctx *AttributeSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAlignmentSpecifier(ctx *AlignmentSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAttributeList(ctx *AttributeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAttribute(ctx *AttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAttributeNamespace(ctx *AttributeNamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAttributeArgumentClause(ctx *AttributeArgumentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitBalancedTokenSeq(ctx *BalancedTokenSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitBalancedtoken(ctx *BalancedtokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitInitDeclarator(ctx *InitDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitDeclarator(ctx *DeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitPointerDeclarator(ctx *PointerDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNoPointerDeclarator(ctx *NoPointerDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitParametersAndQualifiers(ctx *ParametersAndQualifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTrailingReturnType(ctx *TrailingReturnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitPointerOperator(ctx *PointerOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitCvQualifierSeq(ctx *CvQualifierSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitCvQualifier(ctx *CvQualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitRefqualifier(ctx *RefqualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitDeclaratorid(ctx *DeclaratoridContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTheTypeId(ctx *TheTypeIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAbstractDeclarator(ctx *AbstractDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitPointerAbstractDeclarator(ctx *PointerAbstractDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNoPointerAbstractDeclarator(ctx *NoPointerAbstractDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAbstractPackDeclarator(ctx *AbstractPackDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNoPointerAbstractPackDeclarator(ctx *NoPointerAbstractPackDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitParameterDeclarationClause(ctx *ParameterDeclarationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitParameterDeclarationList(ctx *ParameterDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitParameterDeclaration(ctx *ParameterDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitInitializer(ctx *InitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitBraceOrEqualInitializer(ctx *BraceOrEqualInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitInitializerClause(ctx *InitializerClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitInitializerList(ctx *InitializerListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitBracedInitList(ctx *BracedInitListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitClassName(ctx *ClassNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitClassSpecifier(ctx *ClassSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitClassHead(ctx *ClassHeadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitClassHeadName(ctx *ClassHeadNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitClassVirtSpecifier(ctx *ClassVirtSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitClassKey(ctx *ClassKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitMemberSpecification(ctx *MemberSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitMemberdeclaration(ctx *MemberdeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitMemberDeclaratorList(ctx *MemberDeclaratorListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitMemberDeclarator(ctx *MemberDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitVirtualSpecifierSeq(ctx *VirtualSpecifierSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitVirtualSpecifier(ctx *VirtualSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitPureSpecifier(ctx *PureSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitBaseClause(ctx *BaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitBaseSpecifierList(ctx *BaseSpecifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitBaseSpecifier(ctx *BaseSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitClassOrDeclType(ctx *ClassOrDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitBaseTypeSpecifier(ctx *BaseTypeSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitAccessSpecifier(ctx *AccessSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitConversionFunctionId(ctx *ConversionFunctionIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitConversionTypeId(ctx *ConversionTypeIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitConversionDeclarator(ctx *ConversionDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitConstructorInitializer(ctx *ConstructorInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitMemInitializerList(ctx *MemInitializerListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitMemInitializer(ctx *MemInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitMeminitializerid(ctx *MeminitializeridContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitOperatorFunctionId(ctx *OperatorFunctionIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitLiteralOperatorId(ctx *LiteralOperatorIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTemplateDeclaration(ctx *TemplateDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTemplateparameterList(ctx *TemplateparameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTemplateParameter(ctx *TemplateParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitSimpleTemplateId(ctx *SimpleTemplateIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTemplateId(ctx *TemplateIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTemplateName(ctx *TemplateNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTemplateArgumentList(ctx *TemplateArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTemplateArgument(ctx *TemplateArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTypeNameSpecifier(ctx *TypeNameSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitExplicitInstantiation(ctx *ExplicitInstantiationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitExplicitSpecialization(ctx *ExplicitSpecializationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTryBlock(ctx *TryBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitFunctionTryBlock(ctx *FunctionTryBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitHandlerSeq(ctx *HandlerSeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitHandler(ctx *HandlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitExceptionDeclaration(ctx *ExceptionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitThrowExpression(ctx *ThrowExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitExceptionSpecification(ctx *ExceptionSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitDynamicExceptionSpecification(ctx *DynamicExceptionSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTypeIdList(ctx *TypeIdListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitNoeExceptSpecification(ctx *NoeExceptSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitTheOperator(ctx *TheOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitOctalLiteral(ctx *OctalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitHexadecimalLiteral(ctx *HexadecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitBinaryLiteral(ctx *BinaryLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitCharacterLiteral(ctx *CharacterLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitFloatingLiteral(ctx *FloatingLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitPointerLiteral(ctx *PointerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSuckCParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
