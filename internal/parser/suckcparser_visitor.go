// Code generated from SuckCParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SuckCParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SuckCParser.
type SuckCParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SuckCParser#translationUnit.
	VisitTranslationUnit(ctx *TranslationUnitContext) interface{}

	// Visit a parse tree produced by SuckCParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#idExpression.
	VisitIdExpression(ctx *IdExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#unqualifiedId.
	VisitUnqualifiedId(ctx *UnqualifiedIdContext) interface{}

	// Visit a parse tree produced by SuckCParser#qualifiedId.
	VisitQualifiedId(ctx *QualifiedIdContext) interface{}

	// Visit a parse tree produced by SuckCParser#nestedNameSpecifier.
	VisitNestedNameSpecifier(ctx *NestedNameSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#lambdaExpression.
	VisitLambdaExpression(ctx *LambdaExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#lambdaIntroducer.
	VisitLambdaIntroducer(ctx *LambdaIntroducerContext) interface{}

	// Visit a parse tree produced by SuckCParser#lambdaCapture.
	VisitLambdaCapture(ctx *LambdaCaptureContext) interface{}

	// Visit a parse tree produced by SuckCParser#captureDefault.
	VisitCaptureDefault(ctx *CaptureDefaultContext) interface{}

	// Visit a parse tree produced by SuckCParser#captureList.
	VisitCaptureList(ctx *CaptureListContext) interface{}

	// Visit a parse tree produced by SuckCParser#capture.
	VisitCapture(ctx *CaptureContext) interface{}

	// Visit a parse tree produced by SuckCParser#simpleCapture.
	VisitSimpleCapture(ctx *SimpleCaptureContext) interface{}

	// Visit a parse tree produced by SuckCParser#initcapture.
	VisitInitcapture(ctx *InitcaptureContext) interface{}

	// Visit a parse tree produced by SuckCParser#lambdaDeclarator.
	VisitLambdaDeclarator(ctx *LambdaDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#typeIdOfTheTypeId.
	VisitTypeIdOfTheTypeId(ctx *TypeIdOfTheTypeIdContext) interface{}

	// Visit a parse tree produced by SuckCParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by SuckCParser#pseudoDestructorName.
	VisitPseudoDestructorName(ctx *PseudoDestructorNameContext) interface{}

	// Visit a parse tree produced by SuckCParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#unaryOperator.
	VisitUnaryOperator(ctx *UnaryOperatorContext) interface{}

	// Visit a parse tree produced by SuckCParser#newExpression_.
	VisitNewExpression_(ctx *NewExpression_Context) interface{}

	// Visit a parse tree produced by SuckCParser#newPlacement.
	VisitNewPlacement(ctx *NewPlacementContext) interface{}

	// Visit a parse tree produced by SuckCParser#newTypeId.
	VisitNewTypeId(ctx *NewTypeIdContext) interface{}

	// Visit a parse tree produced by SuckCParser#newDeclarator_.
	VisitNewDeclarator_(ctx *NewDeclarator_Context) interface{}

	// Visit a parse tree produced by SuckCParser#noPointerNewDeclarator.
	VisitNoPointerNewDeclarator(ctx *NoPointerNewDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#newInitializer_.
	VisitNewInitializer_(ctx *NewInitializer_Context) interface{}

	// Visit a parse tree produced by SuckCParser#deleteExpression.
	VisitDeleteExpression(ctx *DeleteExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#noExceptExpression.
	VisitNoExceptExpression(ctx *NoExceptExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#castExpression.
	VisitCastExpression(ctx *CastExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#pointerMemberExpression.
	VisitPointerMemberExpression(ctx *PointerMemberExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#shiftExpression.
	VisitShiftExpression(ctx *ShiftExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#shiftOperator.
	VisitShiftOperator(ctx *ShiftOperatorContext) interface{}

	// Visit a parse tree produced by SuckCParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#exclusiveOrExpression.
	VisitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#inclusiveOrExpression.
	VisitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#conditionalExpression.
	VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by SuckCParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#constantExpression.
	VisitConstantExpression(ctx *ConstantExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by SuckCParser#labeledStatement.
	VisitLabeledStatement(ctx *LabeledStatementContext) interface{}

	// Visit a parse tree produced by SuckCParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by SuckCParser#compoundStatement.
	VisitCompoundStatement(ctx *CompoundStatementContext) interface{}

	// Visit a parse tree produced by SuckCParser#statementSeq.
	VisitStatementSeq(ctx *StatementSeqContext) interface{}

	// Visit a parse tree produced by SuckCParser#selectionStatement.
	VisitSelectionStatement(ctx *SelectionStatementContext) interface{}

	// Visit a parse tree produced by SuckCParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by SuckCParser#iterationStatement.
	VisitIterationStatement(ctx *IterationStatementContext) interface{}

	// Visit a parse tree produced by SuckCParser#forInitStatement.
	VisitForInitStatement(ctx *ForInitStatementContext) interface{}

	// Visit a parse tree produced by SuckCParser#forRangeDeclaration.
	VisitForRangeDeclaration(ctx *ForRangeDeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#forRangeInitializer.
	VisitForRangeInitializer(ctx *ForRangeInitializerContext) interface{}

	// Visit a parse tree produced by SuckCParser#jumpStatement.
	VisitJumpStatement(ctx *JumpStatementContext) interface{}

	// Visit a parse tree produced by SuckCParser#declarationStatement.
	VisitDeclarationStatement(ctx *DeclarationStatementContext) interface{}

	// Visit a parse tree produced by SuckCParser#declarationSeq.
	VisitDeclarationSeq(ctx *DeclarationSeqContext) interface{}

	// Visit a parse tree produced by SuckCParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#blockDeclaration.
	VisitBlockDeclaration(ctx *BlockDeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#aliasDeclaration.
	VisitAliasDeclaration(ctx *AliasDeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#functionPointerDeclarator.
	VisitFunctionPointerDeclarator(ctx *FunctionPointerDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#simpleTypedefDeclarator.
	VisitSimpleTypedefDeclarator(ctx *SimpleTypedefDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#typedefDeclaration.
	VisitTypedefDeclaration(ctx *TypedefDeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#simpleDeclaration.
	VisitSimpleDeclaration(ctx *SimpleDeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#staticAssertDeclaration.
	VisitStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#emptyDeclaration_.
	VisitEmptyDeclaration_(ctx *EmptyDeclaration_Context) interface{}

	// Visit a parse tree produced by SuckCParser#attributeDeclaration.
	VisitAttributeDeclaration(ctx *AttributeDeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#declSpecifier.
	VisitDeclSpecifier(ctx *DeclSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#declSpecifierSeq.
	VisitDeclSpecifierSeq(ctx *DeclSpecifierSeqContext) interface{}

	// Visit a parse tree produced by SuckCParser#storageClassSpecifier.
	VisitStorageClassSpecifier(ctx *StorageClassSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#functionSpecifier.
	VisitFunctionSpecifier(ctx *FunctionSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#typedefName.
	VisitTypedefName(ctx *TypedefNameContext) interface{}

	// Visit a parse tree produced by SuckCParser#typeSpecifier.
	VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#trailingTypeSpecifier.
	VisitTrailingTypeSpecifier(ctx *TrailingTypeSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#typeSpecifierSeq.
	VisitTypeSpecifierSeq(ctx *TypeSpecifierSeqContext) interface{}

	// Visit a parse tree produced by SuckCParser#trailingTypeSpecifierSeq.
	VisitTrailingTypeSpecifierSeq(ctx *TrailingTypeSpecifierSeqContext) interface{}

	// Visit a parse tree produced by SuckCParser#simpleTypeLengthModifier.
	VisitSimpleTypeLengthModifier(ctx *SimpleTypeLengthModifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#simpleTypeSignednessModifier.
	VisitSimpleTypeSignednessModifier(ctx *SimpleTypeSignednessModifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#simpleTypeSpecifier.
	VisitSimpleTypeSpecifier(ctx *SimpleTypeSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#theTypeName.
	VisitTheTypeName(ctx *TheTypeNameContext) interface{}

	// Visit a parse tree produced by SuckCParser#decltypeSpecifier.
	VisitDecltypeSpecifier(ctx *DecltypeSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#elaboratedTypeSpecifier.
	VisitElaboratedTypeSpecifier(ctx *ElaboratedTypeSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#enumName.
	VisitEnumName(ctx *EnumNameContext) interface{}

	// Visit a parse tree produced by SuckCParser#enumSpecifier.
	VisitEnumSpecifier(ctx *EnumSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#enumHead.
	VisitEnumHead(ctx *EnumHeadContext) interface{}

	// Visit a parse tree produced by SuckCParser#opaqueEnumDeclaration.
	VisitOpaqueEnumDeclaration(ctx *OpaqueEnumDeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#enumkey.
	VisitEnumkey(ctx *EnumkeyContext) interface{}

	// Visit a parse tree produced by SuckCParser#enumbase.
	VisitEnumbase(ctx *EnumbaseContext) interface{}

	// Visit a parse tree produced by SuckCParser#enumeratorList.
	VisitEnumeratorList(ctx *EnumeratorListContext) interface{}

	// Visit a parse tree produced by SuckCParser#enumeratorDefinition.
	VisitEnumeratorDefinition(ctx *EnumeratorDefinitionContext) interface{}

	// Visit a parse tree produced by SuckCParser#enumerator.
	VisitEnumerator(ctx *EnumeratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#namespaceName.
	VisitNamespaceName(ctx *NamespaceNameContext) interface{}

	// Visit a parse tree produced by SuckCParser#originalNamespaceName.
	VisitOriginalNamespaceName(ctx *OriginalNamespaceNameContext) interface{}

	// Visit a parse tree produced by SuckCParser#namespaceDefinition.
	VisitNamespaceDefinition(ctx *NamespaceDefinitionContext) interface{}

	// Visit a parse tree produced by SuckCParser#namespaceAlias.
	VisitNamespaceAlias(ctx *NamespaceAliasContext) interface{}

	// Visit a parse tree produced by SuckCParser#namespaceAliasDefinition.
	VisitNamespaceAliasDefinition(ctx *NamespaceAliasDefinitionContext) interface{}

	// Visit a parse tree produced by SuckCParser#qualifiednamespaceSpecifier.
	VisitQualifiednamespaceSpecifier(ctx *QualifiednamespaceSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#usingDeclaration.
	VisitUsingDeclaration(ctx *UsingDeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#usingDirective.
	VisitUsingDirective(ctx *UsingDirectiveContext) interface{}

	// Visit a parse tree produced by SuckCParser#asmDefinition.
	VisitAsmDefinition(ctx *AsmDefinitionContext) interface{}

	// Visit a parse tree produced by SuckCParser#linkageSpecification.
	VisitLinkageSpecification(ctx *LinkageSpecificationContext) interface{}

	// Visit a parse tree produced by SuckCParser#attributeSpecifierSeq.
	VisitAttributeSpecifierSeq(ctx *AttributeSpecifierSeqContext) interface{}

	// Visit a parse tree produced by SuckCParser#attributeSpecifier.
	VisitAttributeSpecifier(ctx *AttributeSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#alignmentSpecifier.
	VisitAlignmentSpecifier(ctx *AlignmentSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#attributeList.
	VisitAttributeList(ctx *AttributeListContext) interface{}

	// Visit a parse tree produced by SuckCParser#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by SuckCParser#attributeNamespace.
	VisitAttributeNamespace(ctx *AttributeNamespaceContext) interface{}

	// Visit a parse tree produced by SuckCParser#attributeArgumentClause.
	VisitAttributeArgumentClause(ctx *AttributeArgumentClauseContext) interface{}

	// Visit a parse tree produced by SuckCParser#balancedTokenSeq.
	VisitBalancedTokenSeq(ctx *BalancedTokenSeqContext) interface{}

	// Visit a parse tree produced by SuckCParser#balancedtoken.
	VisitBalancedtoken(ctx *BalancedtokenContext) interface{}

	// Visit a parse tree produced by SuckCParser#initDeclarator.
	VisitInitDeclarator(ctx *InitDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#declarator.
	VisitDeclarator(ctx *DeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#pointerDeclarator.
	VisitPointerDeclarator(ctx *PointerDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#noPointerDeclarator.
	VisitNoPointerDeclarator(ctx *NoPointerDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#parametersAndQualifiers.
	VisitParametersAndQualifiers(ctx *ParametersAndQualifiersContext) interface{}

	// Visit a parse tree produced by SuckCParser#trailingReturnType.
	VisitTrailingReturnType(ctx *TrailingReturnTypeContext) interface{}

	// Visit a parse tree produced by SuckCParser#pointerOperator.
	VisitPointerOperator(ctx *PointerOperatorContext) interface{}

	// Visit a parse tree produced by SuckCParser#cvQualifierSeq.
	VisitCvQualifierSeq(ctx *CvQualifierSeqContext) interface{}

	// Visit a parse tree produced by SuckCParser#cvQualifier.
	VisitCvQualifier(ctx *CvQualifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#declaratorid.
	VisitDeclaratorid(ctx *DeclaratoridContext) interface{}

	// Visit a parse tree produced by SuckCParser#theTypeId.
	VisitTheTypeId(ctx *TheTypeIdContext) interface{}

	// Visit a parse tree produced by SuckCParser#abstractDeclarator.
	VisitAbstractDeclarator(ctx *AbstractDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#pointerAbstractDeclarator.
	VisitPointerAbstractDeclarator(ctx *PointerAbstractDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#noPointerAbstractDeclarator.
	VisitNoPointerAbstractDeclarator(ctx *NoPointerAbstractDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#abstractPackDeclarator.
	VisitAbstractPackDeclarator(ctx *AbstractPackDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#noPointerAbstractPackDeclarator.
	VisitNoPointerAbstractPackDeclarator(ctx *NoPointerAbstractPackDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#parameterDeclarationClause.
	VisitParameterDeclarationClause(ctx *ParameterDeclarationClauseContext) interface{}

	// Visit a parse tree produced by SuckCParser#parameterDeclarationList.
	VisitParameterDeclarationList(ctx *ParameterDeclarationListContext) interface{}

	// Visit a parse tree produced by SuckCParser#parameterDeclaration.
	VisitParameterDeclaration(ctx *ParameterDeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by SuckCParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by SuckCParser#initializer.
	VisitInitializer(ctx *InitializerContext) interface{}

	// Visit a parse tree produced by SuckCParser#braceOrEqualInitializer.
	VisitBraceOrEqualInitializer(ctx *BraceOrEqualInitializerContext) interface{}

	// Visit a parse tree produced by SuckCParser#initializerClause.
	VisitInitializerClause(ctx *InitializerClauseContext) interface{}

	// Visit a parse tree produced by SuckCParser#initializerList.
	VisitInitializerList(ctx *InitializerListContext) interface{}

	// Visit a parse tree produced by SuckCParser#bracedInitList.
	VisitBracedInitList(ctx *BracedInitListContext) interface{}

	// Visit a parse tree produced by SuckCParser#className.
	VisitClassName(ctx *ClassNameContext) interface{}

	// Visit a parse tree produced by SuckCParser#classSpecifier.
	VisitClassSpecifier(ctx *ClassSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#classHead.
	VisitClassHead(ctx *ClassHeadContext) interface{}

	// Visit a parse tree produced by SuckCParser#classHeadName.
	VisitClassHeadName(ctx *ClassHeadNameContext) interface{}

	// Visit a parse tree produced by SuckCParser#classVirtSpecifier.
	VisitClassVirtSpecifier(ctx *ClassVirtSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#classKey.
	VisitClassKey(ctx *ClassKeyContext) interface{}

	// Visit a parse tree produced by SuckCParser#memberSpecification.
	VisitMemberSpecification(ctx *MemberSpecificationContext) interface{}

	// Visit a parse tree produced by SuckCParser#memberdeclaration.
	VisitMemberdeclaration(ctx *MemberdeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#memberDeclaratorList.
	VisitMemberDeclaratorList(ctx *MemberDeclaratorListContext) interface{}

	// Visit a parse tree produced by SuckCParser#memberDeclarator.
	VisitMemberDeclarator(ctx *MemberDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#virtualSpecifierSeq.
	VisitVirtualSpecifierSeq(ctx *VirtualSpecifierSeqContext) interface{}

	// Visit a parse tree produced by SuckCParser#virtualSpecifier.
	VisitVirtualSpecifier(ctx *VirtualSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#pureSpecifier.
	VisitPureSpecifier(ctx *PureSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#baseClause.
	VisitBaseClause(ctx *BaseClauseContext) interface{}

	// Visit a parse tree produced by SuckCParser#baseSpecifierList.
	VisitBaseSpecifierList(ctx *BaseSpecifierListContext) interface{}

	// Visit a parse tree produced by SuckCParser#baseSpecifier.
	VisitBaseSpecifier(ctx *BaseSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#classOrDeclType.
	VisitClassOrDeclType(ctx *ClassOrDeclTypeContext) interface{}

	// Visit a parse tree produced by SuckCParser#baseTypeSpecifier.
	VisitBaseTypeSpecifier(ctx *BaseTypeSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#accessSpecifier.
	VisitAccessSpecifier(ctx *AccessSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#conversionFunctionId.
	VisitConversionFunctionId(ctx *ConversionFunctionIdContext) interface{}

	// Visit a parse tree produced by SuckCParser#conversionTypeId.
	VisitConversionTypeId(ctx *ConversionTypeIdContext) interface{}

	// Visit a parse tree produced by SuckCParser#conversionDeclarator.
	VisitConversionDeclarator(ctx *ConversionDeclaratorContext) interface{}

	// Visit a parse tree produced by SuckCParser#constructorInitializer.
	VisitConstructorInitializer(ctx *ConstructorInitializerContext) interface{}

	// Visit a parse tree produced by SuckCParser#memInitializerList.
	VisitMemInitializerList(ctx *MemInitializerListContext) interface{}

	// Visit a parse tree produced by SuckCParser#memInitializer.
	VisitMemInitializer(ctx *MemInitializerContext) interface{}

	// Visit a parse tree produced by SuckCParser#meminitializerid.
	VisitMeminitializerid(ctx *MeminitializeridContext) interface{}

	// Visit a parse tree produced by SuckCParser#operatorFunctionId.
	VisitOperatorFunctionId(ctx *OperatorFunctionIdContext) interface{}

	// Visit a parse tree produced by SuckCParser#literalOperatorId.
	VisitLiteralOperatorId(ctx *LiteralOperatorIdContext) interface{}

	// Visit a parse tree produced by SuckCParser#templateDeclaration.
	VisitTemplateDeclaration(ctx *TemplateDeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#templateparameterList.
	VisitTemplateparameterList(ctx *TemplateparameterListContext) interface{}

	// Visit a parse tree produced by SuckCParser#templateParameter.
	VisitTemplateParameter(ctx *TemplateParameterContext) interface{}

	// Visit a parse tree produced by SuckCParser#typeParameter.
	VisitTypeParameter(ctx *TypeParameterContext) interface{}

	// Visit a parse tree produced by SuckCParser#simpleTemplateId.
	VisitSimpleTemplateId(ctx *SimpleTemplateIdContext) interface{}

	// Visit a parse tree produced by SuckCParser#templateId.
	VisitTemplateId(ctx *TemplateIdContext) interface{}

	// Visit a parse tree produced by SuckCParser#templateName.
	VisitTemplateName(ctx *TemplateNameContext) interface{}

	// Visit a parse tree produced by SuckCParser#templateArgumentList.
	VisitTemplateArgumentList(ctx *TemplateArgumentListContext) interface{}

	// Visit a parse tree produced by SuckCParser#templateArgument.
	VisitTemplateArgument(ctx *TemplateArgumentContext) interface{}

	// Visit a parse tree produced by SuckCParser#typeNameSpecifier.
	VisitTypeNameSpecifier(ctx *TypeNameSpecifierContext) interface{}

	// Visit a parse tree produced by SuckCParser#explicitInstantiation.
	VisitExplicitInstantiation(ctx *ExplicitInstantiationContext) interface{}

	// Visit a parse tree produced by SuckCParser#explicitSpecialization.
	VisitExplicitSpecialization(ctx *ExplicitSpecializationContext) interface{}

	// Visit a parse tree produced by SuckCParser#tryBlock.
	VisitTryBlock(ctx *TryBlockContext) interface{}

	// Visit a parse tree produced by SuckCParser#functionTryBlock.
	VisitFunctionTryBlock(ctx *FunctionTryBlockContext) interface{}

	// Visit a parse tree produced by SuckCParser#handlerSeq.
	VisitHandlerSeq(ctx *HandlerSeqContext) interface{}

	// Visit a parse tree produced by SuckCParser#handler.
	VisitHandler(ctx *HandlerContext) interface{}

	// Visit a parse tree produced by SuckCParser#exceptionDeclaration.
	VisitExceptionDeclaration(ctx *ExceptionDeclarationContext) interface{}

	// Visit a parse tree produced by SuckCParser#throwExpression.
	VisitThrowExpression(ctx *ThrowExpressionContext) interface{}

	// Visit a parse tree produced by SuckCParser#exceptionSpecification.
	VisitExceptionSpecification(ctx *ExceptionSpecificationContext) interface{}

	// Visit a parse tree produced by SuckCParser#dynamicExceptionSpecification.
	VisitDynamicExceptionSpecification(ctx *DynamicExceptionSpecificationContext) interface{}

	// Visit a parse tree produced by SuckCParser#typeIdList.
	VisitTypeIdList(ctx *TypeIdListContext) interface{}

	// Visit a parse tree produced by SuckCParser#noeExceptSpecification.
	VisitNoeExceptSpecification(ctx *NoeExceptSpecificationContext) interface{}

	// Visit a parse tree produced by SuckCParser#theOperator.
	VisitTheOperator(ctx *TheOperatorContext) interface{}

	// Visit a parse tree produced by SuckCParser#decimalLiteral.
	VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{}

	// Visit a parse tree produced by SuckCParser#octalLiteral.
	VisitOctalLiteral(ctx *OctalLiteralContext) interface{}

	// Visit a parse tree produced by SuckCParser#hexadecimalLiteral.
	VisitHexadecimalLiteral(ctx *HexadecimalLiteralContext) interface{}

	// Visit a parse tree produced by SuckCParser#binaryLiteral.
	VisitBinaryLiteral(ctx *BinaryLiteralContext) interface{}

	// Visit a parse tree produced by SuckCParser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by SuckCParser#characterLiteral.
	VisitCharacterLiteral(ctx *CharacterLiteralContext) interface{}

	// Visit a parse tree produced by SuckCParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by SuckCParser#floatingLiteral.
	VisitFloatingLiteral(ctx *FloatingLiteralContext) interface{}

	// Visit a parse tree produced by SuckCParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by SuckCParser#pointerLiteral.
	VisitPointerLiteral(ctx *PointerLiteralContext) interface{}

	// Visit a parse tree produced by SuckCParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
