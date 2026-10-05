// Code generated from SuckCParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SuckCParser

import "github.com/antlr4-go/antlr/v4"

// BaseSuckCParserListener is a complete listener for a parse tree produced by SuckCParser.
type BaseSuckCParserListener struct{}

var _ SuckCParserListener = &BaseSuckCParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSuckCParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSuckCParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSuckCParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSuckCParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTranslationUnit is called when production translationUnit is entered.
func (s *BaseSuckCParserListener) EnterTranslationUnit(ctx *TranslationUnitContext) {}

// ExitTranslationUnit is called when production translationUnit is exited.
func (s *BaseSuckCParserListener) ExitTranslationUnit(ctx *TranslationUnitContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseSuckCParserListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseSuckCParserListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterIdExpression is called when production idExpression is entered.
func (s *BaseSuckCParserListener) EnterIdExpression(ctx *IdExpressionContext) {}

// ExitIdExpression is called when production idExpression is exited.
func (s *BaseSuckCParserListener) ExitIdExpression(ctx *IdExpressionContext) {}

// EnterUnqualifiedId is called when production unqualifiedId is entered.
func (s *BaseSuckCParserListener) EnterUnqualifiedId(ctx *UnqualifiedIdContext) {}

// ExitUnqualifiedId is called when production unqualifiedId is exited.
func (s *BaseSuckCParserListener) ExitUnqualifiedId(ctx *UnqualifiedIdContext) {}

// EnterQualifiedId is called when production qualifiedId is entered.
func (s *BaseSuckCParserListener) EnterQualifiedId(ctx *QualifiedIdContext) {}

// ExitQualifiedId is called when production qualifiedId is exited.
func (s *BaseSuckCParserListener) ExitQualifiedId(ctx *QualifiedIdContext) {}

// EnterNestedNameSpecifier is called when production nestedNameSpecifier is entered.
func (s *BaseSuckCParserListener) EnterNestedNameSpecifier(ctx *NestedNameSpecifierContext) {}

// ExitNestedNameSpecifier is called when production nestedNameSpecifier is exited.
func (s *BaseSuckCParserListener) ExitNestedNameSpecifier(ctx *NestedNameSpecifierContext) {}

// EnterLambdaExpression is called when production lambdaExpression is entered.
func (s *BaseSuckCParserListener) EnterLambdaExpression(ctx *LambdaExpressionContext) {}

// ExitLambdaExpression is called when production lambdaExpression is exited.
func (s *BaseSuckCParserListener) ExitLambdaExpression(ctx *LambdaExpressionContext) {}

// EnterLambdaIntroducer is called when production lambdaIntroducer is entered.
func (s *BaseSuckCParserListener) EnterLambdaIntroducer(ctx *LambdaIntroducerContext) {}

// ExitLambdaIntroducer is called when production lambdaIntroducer is exited.
func (s *BaseSuckCParserListener) ExitLambdaIntroducer(ctx *LambdaIntroducerContext) {}

// EnterLambdaCapture is called when production lambdaCapture is entered.
func (s *BaseSuckCParserListener) EnterLambdaCapture(ctx *LambdaCaptureContext) {}

// ExitLambdaCapture is called when production lambdaCapture is exited.
func (s *BaseSuckCParserListener) ExitLambdaCapture(ctx *LambdaCaptureContext) {}

// EnterCaptureDefault is called when production captureDefault is entered.
func (s *BaseSuckCParserListener) EnterCaptureDefault(ctx *CaptureDefaultContext) {}

// ExitCaptureDefault is called when production captureDefault is exited.
func (s *BaseSuckCParserListener) ExitCaptureDefault(ctx *CaptureDefaultContext) {}

// EnterCaptureList is called when production captureList is entered.
func (s *BaseSuckCParserListener) EnterCaptureList(ctx *CaptureListContext) {}

// ExitCaptureList is called when production captureList is exited.
func (s *BaseSuckCParserListener) ExitCaptureList(ctx *CaptureListContext) {}

// EnterCapture is called when production capture is entered.
func (s *BaseSuckCParserListener) EnterCapture(ctx *CaptureContext) {}

// ExitCapture is called when production capture is exited.
func (s *BaseSuckCParserListener) ExitCapture(ctx *CaptureContext) {}

// EnterSimpleCapture is called when production simpleCapture is entered.
func (s *BaseSuckCParserListener) EnterSimpleCapture(ctx *SimpleCaptureContext) {}

// ExitSimpleCapture is called when production simpleCapture is exited.
func (s *BaseSuckCParserListener) ExitSimpleCapture(ctx *SimpleCaptureContext) {}

// EnterInitcapture is called when production initcapture is entered.
func (s *BaseSuckCParserListener) EnterInitcapture(ctx *InitcaptureContext) {}

// ExitInitcapture is called when production initcapture is exited.
func (s *BaseSuckCParserListener) ExitInitcapture(ctx *InitcaptureContext) {}

// EnterLambdaDeclarator is called when production lambdaDeclarator is entered.
func (s *BaseSuckCParserListener) EnterLambdaDeclarator(ctx *LambdaDeclaratorContext) {}

// ExitLambdaDeclarator is called when production lambdaDeclarator is exited.
func (s *BaseSuckCParserListener) ExitLambdaDeclarator(ctx *LambdaDeclaratorContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseSuckCParserListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseSuckCParserListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterTypeIdOfTheTypeId is called when production typeIdOfTheTypeId is entered.
func (s *BaseSuckCParserListener) EnterTypeIdOfTheTypeId(ctx *TypeIdOfTheTypeIdContext) {}

// ExitTypeIdOfTheTypeId is called when production typeIdOfTheTypeId is exited.
func (s *BaseSuckCParserListener) ExitTypeIdOfTheTypeId(ctx *TypeIdOfTheTypeIdContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseSuckCParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseSuckCParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterPseudoDestructorName is called when production pseudoDestructorName is entered.
func (s *BaseSuckCParserListener) EnterPseudoDestructorName(ctx *PseudoDestructorNameContext) {}

// ExitPseudoDestructorName is called when production pseudoDestructorName is exited.
func (s *BaseSuckCParserListener) ExitPseudoDestructorName(ctx *PseudoDestructorNameContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseSuckCParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseSuckCParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseSuckCParserListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseSuckCParserListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterNewExpression_ is called when production newExpression_ is entered.
func (s *BaseSuckCParserListener) EnterNewExpression_(ctx *NewExpression_Context) {}

// ExitNewExpression_ is called when production newExpression_ is exited.
func (s *BaseSuckCParserListener) ExitNewExpression_(ctx *NewExpression_Context) {}

// EnterNewPlacement is called when production newPlacement is entered.
func (s *BaseSuckCParserListener) EnterNewPlacement(ctx *NewPlacementContext) {}

// ExitNewPlacement is called when production newPlacement is exited.
func (s *BaseSuckCParserListener) ExitNewPlacement(ctx *NewPlacementContext) {}

// EnterNewTypeId is called when production newTypeId is entered.
func (s *BaseSuckCParserListener) EnterNewTypeId(ctx *NewTypeIdContext) {}

// ExitNewTypeId is called when production newTypeId is exited.
func (s *BaseSuckCParserListener) ExitNewTypeId(ctx *NewTypeIdContext) {}

// EnterNewDeclarator_ is called when production newDeclarator_ is entered.
func (s *BaseSuckCParserListener) EnterNewDeclarator_(ctx *NewDeclarator_Context) {}

// ExitNewDeclarator_ is called when production newDeclarator_ is exited.
func (s *BaseSuckCParserListener) ExitNewDeclarator_(ctx *NewDeclarator_Context) {}

// EnterNoPointerNewDeclarator is called when production noPointerNewDeclarator is entered.
func (s *BaseSuckCParserListener) EnterNoPointerNewDeclarator(ctx *NoPointerNewDeclaratorContext) {}

// ExitNoPointerNewDeclarator is called when production noPointerNewDeclarator is exited.
func (s *BaseSuckCParserListener) ExitNoPointerNewDeclarator(ctx *NoPointerNewDeclaratorContext) {}

// EnterNewInitializer_ is called when production newInitializer_ is entered.
func (s *BaseSuckCParserListener) EnterNewInitializer_(ctx *NewInitializer_Context) {}

// ExitNewInitializer_ is called when production newInitializer_ is exited.
func (s *BaseSuckCParserListener) ExitNewInitializer_(ctx *NewInitializer_Context) {}

// EnterDeleteExpression is called when production deleteExpression is entered.
func (s *BaseSuckCParserListener) EnterDeleteExpression(ctx *DeleteExpressionContext) {}

// ExitDeleteExpression is called when production deleteExpression is exited.
func (s *BaseSuckCParserListener) ExitDeleteExpression(ctx *DeleteExpressionContext) {}

// EnterNoExceptExpression is called when production noExceptExpression is entered.
func (s *BaseSuckCParserListener) EnterNoExceptExpression(ctx *NoExceptExpressionContext) {}

// ExitNoExceptExpression is called when production noExceptExpression is exited.
func (s *BaseSuckCParserListener) ExitNoExceptExpression(ctx *NoExceptExpressionContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *BaseSuckCParserListener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *BaseSuckCParserListener) ExitCastExpression(ctx *CastExpressionContext) {}

// EnterPointerMemberExpression is called when production pointerMemberExpression is entered.
func (s *BaseSuckCParserListener) EnterPointerMemberExpression(ctx *PointerMemberExpressionContext) {}

// ExitPointerMemberExpression is called when production pointerMemberExpression is exited.
func (s *BaseSuckCParserListener) ExitPointerMemberExpression(ctx *PointerMemberExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseSuckCParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseSuckCParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseSuckCParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseSuckCParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *BaseSuckCParserListener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *BaseSuckCParserListener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

// EnterShiftOperator is called when production shiftOperator is entered.
func (s *BaseSuckCParserListener) EnterShiftOperator(ctx *ShiftOperatorContext) {}

// ExitShiftOperator is called when production shiftOperator is exited.
func (s *BaseSuckCParserListener) ExitShiftOperator(ctx *ShiftOperatorContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseSuckCParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseSuckCParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseSuckCParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseSuckCParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseSuckCParserListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseSuckCParserListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *BaseSuckCParserListener) EnterExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *BaseSuckCParserListener) ExitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *BaseSuckCParserListener) EnterInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *BaseSuckCParserListener) ExitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BaseSuckCParserListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BaseSuckCParserListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BaseSuckCParserListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BaseSuckCParserListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseSuckCParserListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseSuckCParserListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BaseSuckCParserListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BaseSuckCParserListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseSuckCParserListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseSuckCParserListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSuckCParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSuckCParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *BaseSuckCParserListener) EnterConstantExpression(ctx *ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *BaseSuckCParserListener) ExitConstantExpression(ctx *ConstantExpressionContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSuckCParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSuckCParserListener) ExitStatement(ctx *StatementContext) {}

// EnterLabeledStatement is called when production labeledStatement is entered.
func (s *BaseSuckCParserListener) EnterLabeledStatement(ctx *LabeledStatementContext) {}

// ExitLabeledStatement is called when production labeledStatement is exited.
func (s *BaseSuckCParserListener) ExitLabeledStatement(ctx *LabeledStatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseSuckCParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseSuckCParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BaseSuckCParserListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BaseSuckCParserListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterStatementSeq is called when production statementSeq is entered.
func (s *BaseSuckCParserListener) EnterStatementSeq(ctx *StatementSeqContext) {}

// ExitStatementSeq is called when production statementSeq is exited.
func (s *BaseSuckCParserListener) ExitStatementSeq(ctx *StatementSeqContext) {}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *BaseSuckCParserListener) EnterSelectionStatement(ctx *SelectionStatementContext) {}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *BaseSuckCParserListener) ExitSelectionStatement(ctx *SelectionStatementContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseSuckCParserListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseSuckCParserListener) ExitCondition(ctx *ConditionContext) {}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *BaseSuckCParserListener) EnterIterationStatement(ctx *IterationStatementContext) {}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *BaseSuckCParserListener) ExitIterationStatement(ctx *IterationStatementContext) {}

// EnterForInitStatement is called when production forInitStatement is entered.
func (s *BaseSuckCParserListener) EnterForInitStatement(ctx *ForInitStatementContext) {}

// ExitForInitStatement is called when production forInitStatement is exited.
func (s *BaseSuckCParserListener) ExitForInitStatement(ctx *ForInitStatementContext) {}

// EnterForRangeDeclaration is called when production forRangeDeclaration is entered.
func (s *BaseSuckCParserListener) EnterForRangeDeclaration(ctx *ForRangeDeclarationContext) {}

// ExitForRangeDeclaration is called when production forRangeDeclaration is exited.
func (s *BaseSuckCParserListener) ExitForRangeDeclaration(ctx *ForRangeDeclarationContext) {}

// EnterForRangeInitializer is called when production forRangeInitializer is entered.
func (s *BaseSuckCParserListener) EnterForRangeInitializer(ctx *ForRangeInitializerContext) {}

// ExitForRangeInitializer is called when production forRangeInitializer is exited.
func (s *BaseSuckCParserListener) ExitForRangeInitializer(ctx *ForRangeInitializerContext) {}

// EnterJumpStatement is called when production jumpStatement is entered.
func (s *BaseSuckCParserListener) EnterJumpStatement(ctx *JumpStatementContext) {}

// ExitJumpStatement is called when production jumpStatement is exited.
func (s *BaseSuckCParserListener) ExitJumpStatement(ctx *JumpStatementContext) {}

// EnterDeclarationStatement is called when production declarationStatement is entered.
func (s *BaseSuckCParserListener) EnterDeclarationStatement(ctx *DeclarationStatementContext) {}

// ExitDeclarationStatement is called when production declarationStatement is exited.
func (s *BaseSuckCParserListener) ExitDeclarationStatement(ctx *DeclarationStatementContext) {}

// EnterDeclarationSeq is called when production declarationSeq is entered.
func (s *BaseSuckCParserListener) EnterDeclarationSeq(ctx *DeclarationSeqContext) {}

// ExitDeclarationSeq is called when production declarationSeq is exited.
func (s *BaseSuckCParserListener) ExitDeclarationSeq(ctx *DeclarationSeqContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseSuckCParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseSuckCParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterCImportRegion is called when production cImportRegion is entered.
func (s *BaseSuckCParserListener) EnterCImportRegion(ctx *CImportRegionContext) {}

// ExitCImportRegion is called when production cImportRegion is exited.
func (s *BaseSuckCParserListener) ExitCImportRegion(ctx *CImportRegionContext) {}

// EnterCCDeclarationSeq is called when production cCDeclarationSeq is entered.
func (s *BaseSuckCParserListener) EnterCCDeclarationSeq(ctx *CCDeclarationSeqContext) {}

// ExitCCDeclarationSeq is called when production cCDeclarationSeq is exited.
func (s *BaseSuckCParserListener) ExitCCDeclarationSeq(ctx *CCDeclarationSeqContext) {}

// EnterCCExternalDeclaration is called when production cCExternalDeclaration is entered.
func (s *BaseSuckCParserListener) EnterCCExternalDeclaration(ctx *CCExternalDeclarationContext) {}

// ExitCCExternalDeclaration is called when production cCExternalDeclaration is exited.
func (s *BaseSuckCParserListener) ExitCCExternalDeclaration(ctx *CCExternalDeclarationContext) {}

// EnterCCAttributesOpt is called when production cCAttributesOpt is entered.
func (s *BaseSuckCParserListener) EnterCCAttributesOpt(ctx *CCAttributesOptContext) {}

// ExitCCAttributesOpt is called when production cCAttributesOpt is exited.
func (s *BaseSuckCParserListener) ExitCCAttributesOpt(ctx *CCAttributesOptContext) {}

// EnterCCAttribute is called when production cCAttribute is entered.
func (s *BaseSuckCParserListener) EnterCCAttribute(ctx *CCAttributeContext) {}

// ExitCCAttribute is called when production cCAttribute is exited.
func (s *BaseSuckCParserListener) ExitCCAttribute(ctx *CCAttributeContext) {}

// EnterCCAttributeName is called when production cCAttributeName is entered.
func (s *BaseSuckCParserListener) EnterCCAttributeName(ctx *CCAttributeNameContext) {}

// ExitCCAttributeName is called when production cCAttributeName is exited.
func (s *BaseSuckCParserListener) ExitCCAttributeName(ctx *CCAttributeNameContext) {}

// EnterCCParenGroup is called when production cCParenGroup is entered.
func (s *BaseSuckCParserListener) EnterCCParenGroup(ctx *CCParenGroupContext) {}

// ExitCCParenGroup is called when production cCParenGroup is exited.
func (s *BaseSuckCParserListener) ExitCCParenGroup(ctx *CCParenGroupContext) {}

// EnterCCDeclSpecifiers is called when production cCDeclSpecifiers is entered.
func (s *BaseSuckCParserListener) EnterCCDeclSpecifiers(ctx *CCDeclSpecifiersContext) {}

// ExitCCDeclSpecifiers is called when production cCDeclSpecifiers is exited.
func (s *BaseSuckCParserListener) ExitCCDeclSpecifiers(ctx *CCDeclSpecifiersContext) {}

// EnterCCDeclSpecifierPart is called when production cCDeclSpecifierPart is entered.
func (s *BaseSuckCParserListener) EnterCCDeclSpecifierPart(ctx *CCDeclSpecifierPartContext) {}

// ExitCCDeclSpecifierPart is called when production cCDeclSpecifierPart is exited.
func (s *BaseSuckCParserListener) ExitCCDeclSpecifierPart(ctx *CCDeclSpecifierPartContext) {}

// EnterCCQualifier is called when production cCQualifier is entered.
func (s *BaseSuckCParserListener) EnterCCQualifier(ctx *CCQualifierContext) {}

// ExitCCQualifier is called when production cCQualifier is exited.
func (s *BaseSuckCParserListener) ExitCCQualifier(ctx *CCQualifierContext) {}

// EnterCCTypedefName is called when production cCTypedefName is entered.
func (s *BaseSuckCParserListener) EnterCCTypedefName(ctx *CCTypedefNameContext) {}

// ExitCCTypedefName is called when production cCTypedefName is exited.
func (s *BaseSuckCParserListener) ExitCCTypedefName(ctx *CCTypedefNameContext) {}

// EnterCCBuiltinTypeSpecifier is called when production cCBuiltinTypeSpecifier is entered.
func (s *BaseSuckCParserListener) EnterCCBuiltinTypeSpecifier(ctx *CCBuiltinTypeSpecifierContext) {}

// ExitCCBuiltinTypeSpecifier is called when production cCBuiltinTypeSpecifier is exited.
func (s *BaseSuckCParserListener) ExitCCBuiltinTypeSpecifier(ctx *CCBuiltinTypeSpecifierContext) {}

// EnterCCStructOrUnionSpecifier is called when production cCStructOrUnionSpecifier is entered.
func (s *BaseSuckCParserListener) EnterCCStructOrUnionSpecifier(ctx *CCStructOrUnionSpecifierContext) {
}

// ExitCCStructOrUnionSpecifier is called when production cCStructOrUnionSpecifier is exited.
func (s *BaseSuckCParserListener) ExitCCStructOrUnionSpecifier(ctx *CCStructOrUnionSpecifierContext) {
}

// EnterCCMemberDeclaration is called when production cCMemberDeclaration is entered.
func (s *BaseSuckCParserListener) EnterCCMemberDeclaration(ctx *CCMemberDeclarationContext) {}

// ExitCCMemberDeclaration is called when production cCMemberDeclaration is exited.
func (s *BaseSuckCParserListener) ExitCCMemberDeclaration(ctx *CCMemberDeclarationContext) {}

// EnterCCMemberDeclaratorList is called when production cCMemberDeclaratorList is entered.
func (s *BaseSuckCParserListener) EnterCCMemberDeclaratorList(ctx *CCMemberDeclaratorListContext) {}

// ExitCCMemberDeclaratorList is called when production cCMemberDeclaratorList is exited.
func (s *BaseSuckCParserListener) ExitCCMemberDeclaratorList(ctx *CCMemberDeclaratorListContext) {}

// EnterCCMemberDeclarator is called when production cCMemberDeclarator is entered.
func (s *BaseSuckCParserListener) EnterCCMemberDeclarator(ctx *CCMemberDeclaratorContext) {}

// ExitCCMemberDeclarator is called when production cCMemberDeclarator is exited.
func (s *BaseSuckCParserListener) ExitCCMemberDeclarator(ctx *CCMemberDeclaratorContext) {}

// EnterCCBitField is called when production cCBitField is entered.
func (s *BaseSuckCParserListener) EnterCCBitField(ctx *CCBitFieldContext) {}

// ExitCCBitField is called when production cCBitField is exited.
func (s *BaseSuckCParserListener) ExitCCBitField(ctx *CCBitFieldContext) {}

// EnterCCEnumSpecifier is called when production cCEnumSpecifier is entered.
func (s *BaseSuckCParserListener) EnterCCEnumSpecifier(ctx *CCEnumSpecifierContext) {}

// ExitCCEnumSpecifier is called when production cCEnumSpecifier is exited.
func (s *BaseSuckCParserListener) ExitCCEnumSpecifier(ctx *CCEnumSpecifierContext) {}

// EnterCCEnumeratorList is called when production cCEnumeratorList is entered.
func (s *BaseSuckCParserListener) EnterCCEnumeratorList(ctx *CCEnumeratorListContext) {}

// ExitCCEnumeratorList is called when production cCEnumeratorList is exited.
func (s *BaseSuckCParserListener) ExitCCEnumeratorList(ctx *CCEnumeratorListContext) {}

// EnterCCEnumerator is called when production cCEnumerator is entered.
func (s *BaseSuckCParserListener) EnterCCEnumerator(ctx *CCEnumeratorContext) {}

// ExitCCEnumerator is called when production cCEnumerator is exited.
func (s *BaseSuckCParserListener) ExitCCEnumerator(ctx *CCEnumeratorContext) {}

// EnterCCTypeofSpecifier is called when production cCTypeofSpecifier is entered.
func (s *BaseSuckCParserListener) EnterCCTypeofSpecifier(ctx *CCTypeofSpecifierContext) {}

// ExitCCTypeofSpecifier is called when production cCTypeofSpecifier is exited.
func (s *BaseSuckCParserListener) ExitCCTypeofSpecifier(ctx *CCTypeofSpecifierContext) {}

// EnterCCTypeName is called when production cCTypeName is entered.
func (s *BaseSuckCParserListener) EnterCCTypeName(ctx *CCTypeNameContext) {}

// ExitCCTypeName is called when production cCTypeName is exited.
func (s *BaseSuckCParserListener) ExitCCTypeName(ctx *CCTypeNameContext) {}

// EnterCCIdentifier is called when production cCIdentifier is entered.
func (s *BaseSuckCParserListener) EnterCCIdentifier(ctx *CCIdentifierContext) {}

// ExitCCIdentifier is called when production cCIdentifier is exited.
func (s *BaseSuckCParserListener) ExitCCIdentifier(ctx *CCIdentifierContext) {}

// EnterCCDeclarator is called when production cCDeclarator is entered.
func (s *BaseSuckCParserListener) EnterCCDeclarator(ctx *CCDeclaratorContext) {}

// ExitCCDeclarator is called when production cCDeclarator is exited.
func (s *BaseSuckCParserListener) ExitCCDeclarator(ctx *CCDeclaratorContext) {}

// EnterCCPlainDeclarator is called when production cCPlainDeclarator is entered.
func (s *BaseSuckCParserListener) EnterCCPlainDeclarator(ctx *CCPlainDeclaratorContext) {}

// ExitCCPlainDeclarator is called when production cCPlainDeclarator is exited.
func (s *BaseSuckCParserListener) ExitCCPlainDeclarator(ctx *CCPlainDeclaratorContext) {}

// EnterCCPointerOperator is called when production cCPointerOperator is entered.
func (s *BaseSuckCParserListener) EnterCCPointerOperator(ctx *CCPointerOperatorContext) {}

// ExitCCPointerOperator is called when production cCPointerOperator is exited.
func (s *BaseSuckCParserListener) ExitCCPointerOperator(ctx *CCPointerOperatorContext) {}

// EnterCCPointerQualifier is called when production cCPointerQualifier is entered.
func (s *BaseSuckCParserListener) EnterCCPointerQualifier(ctx *CCPointerQualifierContext) {}

// ExitCCPointerQualifier is called when production cCPointerQualifier is exited.
func (s *BaseSuckCParserListener) ExitCCPointerQualifier(ctx *CCPointerQualifierContext) {}

// EnterCCDirectDeclarator is called when production cCDirectDeclarator is entered.
func (s *BaseSuckCParserListener) EnterCCDirectDeclarator(ctx *CCDirectDeclaratorContext) {}

// ExitCCDirectDeclarator is called when production cCDirectDeclarator is exited.
func (s *BaseSuckCParserListener) ExitCCDirectDeclarator(ctx *CCDirectDeclaratorContext) {}

// EnterCCPlainDirectDeclarator is called when production cCPlainDirectDeclarator is entered.
func (s *BaseSuckCParserListener) EnterCCPlainDirectDeclarator(ctx *CCPlainDirectDeclaratorContext) {}

// ExitCCPlainDirectDeclarator is called when production cCPlainDirectDeclarator is exited.
func (s *BaseSuckCParserListener) ExitCCPlainDirectDeclarator(ctx *CCPlainDirectDeclaratorContext) {}

// EnterCCDirectDeclaratorSuffix is called when production cCDirectDeclaratorSuffix is entered.
func (s *BaseSuckCParserListener) EnterCCDirectDeclaratorSuffix(ctx *CCDirectDeclaratorSuffixContext) {
}

// ExitCCDirectDeclaratorSuffix is called when production cCDirectDeclaratorSuffix is exited.
func (s *BaseSuckCParserListener) ExitCCDirectDeclaratorSuffix(ctx *CCDirectDeclaratorSuffixContext) {
}

// EnterCCPlainDirectDeclaratorSuffix is called when production cCPlainDirectDeclaratorSuffix is entered.
func (s *BaseSuckCParserListener) EnterCCPlainDirectDeclaratorSuffix(ctx *CCPlainDirectDeclaratorSuffixContext) {
}

// ExitCCPlainDirectDeclaratorSuffix is called when production cCPlainDirectDeclaratorSuffix is exited.
func (s *BaseSuckCParserListener) ExitCCPlainDirectDeclaratorSuffix(ctx *CCPlainDirectDeclaratorSuffixContext) {
}

// EnterCCAbstractDeclarator is called when production cCAbstractDeclarator is entered.
func (s *BaseSuckCParserListener) EnterCCAbstractDeclarator(ctx *CCAbstractDeclaratorContext) {}

// ExitCCAbstractDeclarator is called when production cCAbstractDeclarator is exited.
func (s *BaseSuckCParserListener) ExitCCAbstractDeclarator(ctx *CCAbstractDeclaratorContext) {}

// EnterCCAbstractDirectDeclarator is called when production cCAbstractDirectDeclarator is entered.
func (s *BaseSuckCParserListener) EnterCCAbstractDirectDeclarator(ctx *CCAbstractDirectDeclaratorContext) {
}

// ExitCCAbstractDirectDeclarator is called when production cCAbstractDirectDeclarator is exited.
func (s *BaseSuckCParserListener) ExitCCAbstractDirectDeclarator(ctx *CCAbstractDirectDeclaratorContext) {
}

// EnterCCAbstractDirectDeclaratorSuffix is called when production cCAbstractDirectDeclaratorSuffix is entered.
func (s *BaseSuckCParserListener) EnterCCAbstractDirectDeclaratorSuffix(ctx *CCAbstractDirectDeclaratorSuffixContext) {
}

// ExitCCAbstractDirectDeclaratorSuffix is called when production cCAbstractDirectDeclaratorSuffix is exited.
func (s *BaseSuckCParserListener) ExitCCAbstractDirectDeclaratorSuffix(ctx *CCAbstractDirectDeclaratorSuffixContext) {
}

// EnterCCArraySuffix is called when production cCArraySuffix is entered.
func (s *BaseSuckCParserListener) EnterCCArraySuffix(ctx *CCArraySuffixContext) {}

// ExitCCArraySuffix is called when production cCArraySuffix is exited.
func (s *BaseSuckCParserListener) ExitCCArraySuffix(ctx *CCArraySuffixContext) {}

// EnterCCFunctionSuffix is called when production cCFunctionSuffix is entered.
func (s *BaseSuckCParserListener) EnterCCFunctionSuffix(ctx *CCFunctionSuffixContext) {}

// ExitCCFunctionSuffix is called when production cCFunctionSuffix is exited.
func (s *BaseSuckCParserListener) ExitCCFunctionSuffix(ctx *CCFunctionSuffixContext) {}

// EnterCCParameterTypeList is called when production cCParameterTypeList is entered.
func (s *BaseSuckCParserListener) EnterCCParameterTypeList(ctx *CCParameterTypeListContext) {}

// ExitCCParameterTypeList is called when production cCParameterTypeList is exited.
func (s *BaseSuckCParserListener) ExitCCParameterTypeList(ctx *CCParameterTypeListContext) {}

// EnterCCParameterDeclaration is called when production cCParameterDeclaration is entered.
func (s *BaseSuckCParserListener) EnterCCParameterDeclaration(ctx *CCParameterDeclarationContext) {}

// ExitCCParameterDeclaration is called when production cCParameterDeclaration is exited.
func (s *BaseSuckCParserListener) ExitCCParameterDeclaration(ctx *CCParameterDeclarationContext) {}

// EnterCCInitDeclaratorList is called when production cCInitDeclaratorList is entered.
func (s *BaseSuckCParserListener) EnterCCInitDeclaratorList(ctx *CCInitDeclaratorListContext) {}

// ExitCCInitDeclaratorList is called when production cCInitDeclaratorList is exited.
func (s *BaseSuckCParserListener) ExitCCInitDeclaratorList(ctx *CCInitDeclaratorListContext) {}

// EnterCCInitDeclarator is called when production cCInitDeclarator is entered.
func (s *BaseSuckCParserListener) EnterCCInitDeclarator(ctx *CCInitDeclaratorContext) {}

// ExitCCInitDeclarator is called when production cCInitDeclarator is exited.
func (s *BaseSuckCParserListener) ExitCCInitDeclarator(ctx *CCInitDeclaratorContext) {}

// EnterCCInitializerClause is called when production cCInitializerClause is entered.
func (s *BaseSuckCParserListener) EnterCCInitializerClause(ctx *CCInitializerClauseContext) {}

// ExitCCInitializerClause is called when production cCInitializerClause is exited.
func (s *BaseSuckCParserListener) ExitCCInitializerClause(ctx *CCInitializerClauseContext) {}

// EnterCCBracedInitList is called when production cCBracedInitList is entered.
func (s *BaseSuckCParserListener) EnterCCBracedInitList(ctx *CCBracedInitListContext) {}

// ExitCCBracedInitList is called when production cCBracedInitList is exited.
func (s *BaseSuckCParserListener) ExitCCBracedInitList(ctx *CCBracedInitListContext) {}

// EnterCCCompoundStatement is called when production cCCompoundStatement is entered.
func (s *BaseSuckCParserListener) EnterCCCompoundStatement(ctx *CCCompoundStatementContext) {}

// ExitCCCompoundStatement is called when production cCCompoundStatement is exited.
func (s *BaseSuckCParserListener) ExitCCCompoundStatement(ctx *CCCompoundStatementContext) {}

// EnterCCBlockItem is called when production cCBlockItem is entered.
func (s *BaseSuckCParserListener) EnterCCBlockItem(ctx *CCBlockItemContext) {}

// ExitCCBlockItem is called when production cCBlockItem is exited.
func (s *BaseSuckCParserListener) ExitCCBlockItem(ctx *CCBlockItemContext) {}

// EnterCCLinkageSpecification is called when production cCLinkageSpecification is entered.
func (s *BaseSuckCParserListener) EnterCCLinkageSpecification(ctx *CCLinkageSpecificationContext) {}

// ExitCCLinkageSpecification is called when production cCLinkageSpecification is exited.
func (s *BaseSuckCParserListener) ExitCCLinkageSpecification(ctx *CCLinkageSpecificationContext) {}

// EnterCCStatement is called when production cCStatement is entered.
func (s *BaseSuckCParserListener) EnterCCStatement(ctx *CCStatementContext) {}

// ExitCCStatement is called when production cCStatement is exited.
func (s *BaseSuckCParserListener) ExitCCStatement(ctx *CCStatementContext) {}

// EnterCCExpression is called when production cCExpression is entered.
func (s *BaseSuckCParserListener) EnterCCExpression(ctx *CCExpressionContext) {}

// ExitCCExpression is called when production cCExpression is exited.
func (s *BaseSuckCParserListener) ExitCCExpression(ctx *CCExpressionContext) {}

// EnterCCAssignmentExpression is called when production cCAssignmentExpression is entered.
func (s *BaseSuckCParserListener) EnterCCAssignmentExpression(ctx *CCAssignmentExpressionContext) {}

// ExitCCAssignmentExpression is called when production cCAssignmentExpression is exited.
func (s *BaseSuckCParserListener) ExitCCAssignmentExpression(ctx *CCAssignmentExpressionContext) {}

// EnterCCAssignmentOperator is called when production cCAssignmentOperator is entered.
func (s *BaseSuckCParserListener) EnterCCAssignmentOperator(ctx *CCAssignmentOperatorContext) {}

// ExitCCAssignmentOperator is called when production cCAssignmentOperator is exited.
func (s *BaseSuckCParserListener) ExitCCAssignmentOperator(ctx *CCAssignmentOperatorContext) {}

// EnterCCConditionalExpression is called when production cCConditionalExpression is entered.
func (s *BaseSuckCParserListener) EnterCCConditionalExpression(ctx *CCConditionalExpressionContext) {}

// ExitCCConditionalExpression is called when production cCConditionalExpression is exited.
func (s *BaseSuckCParserListener) ExitCCConditionalExpression(ctx *CCConditionalExpressionContext) {}

// EnterCCLogicalOrExpression is called when production cCLogicalOrExpression is entered.
func (s *BaseSuckCParserListener) EnterCCLogicalOrExpression(ctx *CCLogicalOrExpressionContext) {}

// ExitCCLogicalOrExpression is called when production cCLogicalOrExpression is exited.
func (s *BaseSuckCParserListener) ExitCCLogicalOrExpression(ctx *CCLogicalOrExpressionContext) {}

// EnterCCLogicalAndExpression is called when production cCLogicalAndExpression is entered.
func (s *BaseSuckCParserListener) EnterCCLogicalAndExpression(ctx *CCLogicalAndExpressionContext) {}

// ExitCCLogicalAndExpression is called when production cCLogicalAndExpression is exited.
func (s *BaseSuckCParserListener) ExitCCLogicalAndExpression(ctx *CCLogicalAndExpressionContext) {}

// EnterCCInclusiveOrExpression is called when production cCInclusiveOrExpression is entered.
func (s *BaseSuckCParserListener) EnterCCInclusiveOrExpression(ctx *CCInclusiveOrExpressionContext) {}

// ExitCCInclusiveOrExpression is called when production cCInclusiveOrExpression is exited.
func (s *BaseSuckCParserListener) ExitCCInclusiveOrExpression(ctx *CCInclusiveOrExpressionContext) {}

// EnterCCExclusiveOrExpression is called when production cCExclusiveOrExpression is entered.
func (s *BaseSuckCParserListener) EnterCCExclusiveOrExpression(ctx *CCExclusiveOrExpressionContext) {}

// ExitCCExclusiveOrExpression is called when production cCExclusiveOrExpression is exited.
func (s *BaseSuckCParserListener) ExitCCExclusiveOrExpression(ctx *CCExclusiveOrExpressionContext) {}

// EnterCCAndExpression is called when production cCAndExpression is entered.
func (s *BaseSuckCParserListener) EnterCCAndExpression(ctx *CCAndExpressionContext) {}

// ExitCCAndExpression is called when production cCAndExpression is exited.
func (s *BaseSuckCParserListener) ExitCCAndExpression(ctx *CCAndExpressionContext) {}

// EnterCCEqualityExpression is called when production cCEqualityExpression is entered.
func (s *BaseSuckCParserListener) EnterCCEqualityExpression(ctx *CCEqualityExpressionContext) {}

// ExitCCEqualityExpression is called when production cCEqualityExpression is exited.
func (s *BaseSuckCParserListener) ExitCCEqualityExpression(ctx *CCEqualityExpressionContext) {}

// EnterCCRelationalExpression is called when production cCRelationalExpression is entered.
func (s *BaseSuckCParserListener) EnterCCRelationalExpression(ctx *CCRelationalExpressionContext) {}

// ExitCCRelationalExpression is called when production cCRelationalExpression is exited.
func (s *BaseSuckCParserListener) ExitCCRelationalExpression(ctx *CCRelationalExpressionContext) {}

// EnterCCAdditiveExpression is called when production cCAdditiveExpression is entered.
func (s *BaseSuckCParserListener) EnterCCAdditiveExpression(ctx *CCAdditiveExpressionContext) {}

// ExitCCAdditiveExpression is called when production cCAdditiveExpression is exited.
func (s *BaseSuckCParserListener) ExitCCAdditiveExpression(ctx *CCAdditiveExpressionContext) {}

// EnterCCMultiplicativeExpression is called when production cCMultiplicativeExpression is entered.
func (s *BaseSuckCParserListener) EnterCCMultiplicativeExpression(ctx *CCMultiplicativeExpressionContext) {
}

// ExitCCMultiplicativeExpression is called when production cCMultiplicativeExpression is exited.
func (s *BaseSuckCParserListener) ExitCCMultiplicativeExpression(ctx *CCMultiplicativeExpressionContext) {
}

// EnterCCUnaryExpression is called when production cCUnaryExpression is entered.
func (s *BaseSuckCParserListener) EnterCCUnaryExpression(ctx *CCUnaryExpressionContext) {}

// ExitCCUnaryExpression is called when production cCUnaryExpression is exited.
func (s *BaseSuckCParserListener) ExitCCUnaryExpression(ctx *CCUnaryExpressionContext) {}

// EnterCCUnaryOperator is called when production cCUnaryOperator is entered.
func (s *BaseSuckCParserListener) EnterCCUnaryOperator(ctx *CCUnaryOperatorContext) {}

// ExitCCUnaryOperator is called when production cCUnaryOperator is exited.
func (s *BaseSuckCParserListener) ExitCCUnaryOperator(ctx *CCUnaryOperatorContext) {}

// EnterCCPostfixExpression is called when production cCPostfixExpression is entered.
func (s *BaseSuckCParserListener) EnterCCPostfixExpression(ctx *CCPostfixExpressionContext) {}

// ExitCCPostfixExpression is called when production cCPostfixExpression is exited.
func (s *BaseSuckCParserListener) ExitCCPostfixExpression(ctx *CCPostfixExpressionContext) {}

// EnterCCPostfixSuffix is called when production cCPostfixSuffix is entered.
func (s *BaseSuckCParserListener) EnterCCPostfixSuffix(ctx *CCPostfixSuffixContext) {}

// ExitCCPostfixSuffix is called when production cCPostfixSuffix is exited.
func (s *BaseSuckCParserListener) ExitCCPostfixSuffix(ctx *CCPostfixSuffixContext) {}

// EnterCCArgumentExpressionList is called when production cCArgumentExpressionList is entered.
func (s *BaseSuckCParserListener) EnterCCArgumentExpressionList(ctx *CCArgumentExpressionListContext) {
}

// ExitCCArgumentExpressionList is called when production cCArgumentExpressionList is exited.
func (s *BaseSuckCParserListener) ExitCCArgumentExpressionList(ctx *CCArgumentExpressionListContext) {
}

// EnterCCPrimaryExpression is called when production cCPrimaryExpression is entered.
func (s *BaseSuckCParserListener) EnterCCPrimaryExpression(ctx *CCPrimaryExpressionContext) {}

// ExitCCPrimaryExpression is called when production cCPrimaryExpression is exited.
func (s *BaseSuckCParserListener) ExitCCPrimaryExpression(ctx *CCPrimaryExpressionContext) {}

// EnterCCConstantExpression is called when production cCConstantExpression is entered.
func (s *BaseSuckCParserListener) EnterCCConstantExpression(ctx *CCConstantExpressionContext) {}

// ExitCCConstantExpression is called when production cCConstantExpression is exited.
func (s *BaseSuckCParserListener) ExitCCConstantExpression(ctx *CCConstantExpressionContext) {}

// EnterBlockDeclaration is called when production blockDeclaration is entered.
func (s *BaseSuckCParserListener) EnterBlockDeclaration(ctx *BlockDeclarationContext) {}

// ExitBlockDeclaration is called when production blockDeclaration is exited.
func (s *BaseSuckCParserListener) ExitBlockDeclaration(ctx *BlockDeclarationContext) {}

// EnterAliasDeclaration is called when production aliasDeclaration is entered.
func (s *BaseSuckCParserListener) EnterAliasDeclaration(ctx *AliasDeclarationContext) {}

// ExitAliasDeclaration is called when production aliasDeclaration is exited.
func (s *BaseSuckCParserListener) ExitAliasDeclaration(ctx *AliasDeclarationContext) {}

// EnterFunctionPointerDeclarator is called when production functionPointerDeclarator is entered.
func (s *BaseSuckCParserListener) EnterFunctionPointerDeclarator(ctx *FunctionPointerDeclaratorContext) {
}

// ExitFunctionPointerDeclarator is called when production functionPointerDeclarator is exited.
func (s *BaseSuckCParserListener) ExitFunctionPointerDeclarator(ctx *FunctionPointerDeclaratorContext) {
}

// EnterSimpleTypedefDeclarator is called when production simpleTypedefDeclarator is entered.
func (s *BaseSuckCParserListener) EnterSimpleTypedefDeclarator(ctx *SimpleTypedefDeclaratorContext) {}

// ExitSimpleTypedefDeclarator is called when production simpleTypedefDeclarator is exited.
func (s *BaseSuckCParserListener) ExitSimpleTypedefDeclarator(ctx *SimpleTypedefDeclaratorContext) {}

// EnterTypedefDeclaration is called when production typedefDeclaration is entered.
func (s *BaseSuckCParserListener) EnterTypedefDeclaration(ctx *TypedefDeclarationContext) {}

// ExitTypedefDeclaration is called when production typedefDeclaration is exited.
func (s *BaseSuckCParserListener) ExitTypedefDeclaration(ctx *TypedefDeclarationContext) {}

// EnterSimpleDeclaration is called when production simpleDeclaration is entered.
func (s *BaseSuckCParserListener) EnterSimpleDeclaration(ctx *SimpleDeclarationContext) {}

// ExitSimpleDeclaration is called when production simpleDeclaration is exited.
func (s *BaseSuckCParserListener) ExitSimpleDeclaration(ctx *SimpleDeclarationContext) {}

// EnterStaticAssertDeclaration is called when production staticAssertDeclaration is entered.
func (s *BaseSuckCParserListener) EnterStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) {}

// ExitStaticAssertDeclaration is called when production staticAssertDeclaration is exited.
func (s *BaseSuckCParserListener) ExitStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) {}

// EnterEmptyDeclaration_ is called when production emptyDeclaration_ is entered.
func (s *BaseSuckCParserListener) EnterEmptyDeclaration_(ctx *EmptyDeclaration_Context) {}

// ExitEmptyDeclaration_ is called when production emptyDeclaration_ is exited.
func (s *BaseSuckCParserListener) ExitEmptyDeclaration_(ctx *EmptyDeclaration_Context) {}

// EnterAttributeDeclaration is called when production attributeDeclaration is entered.
func (s *BaseSuckCParserListener) EnterAttributeDeclaration(ctx *AttributeDeclarationContext) {}

// ExitAttributeDeclaration is called when production attributeDeclaration is exited.
func (s *BaseSuckCParserListener) ExitAttributeDeclaration(ctx *AttributeDeclarationContext) {}

// EnterDeclSpecifier is called when production declSpecifier is entered.
func (s *BaseSuckCParserListener) EnterDeclSpecifier(ctx *DeclSpecifierContext) {}

// ExitDeclSpecifier is called when production declSpecifier is exited.
func (s *BaseSuckCParserListener) ExitDeclSpecifier(ctx *DeclSpecifierContext) {}

// EnterDeclSpecifierSeq is called when production declSpecifierSeq is entered.
func (s *BaseSuckCParserListener) EnterDeclSpecifierSeq(ctx *DeclSpecifierSeqContext) {}

// ExitDeclSpecifierSeq is called when production declSpecifierSeq is exited.
func (s *BaseSuckCParserListener) ExitDeclSpecifierSeq(ctx *DeclSpecifierSeqContext) {}

// EnterStorageClassSpecifier is called when production storageClassSpecifier is entered.
func (s *BaseSuckCParserListener) EnterStorageClassSpecifier(ctx *StorageClassSpecifierContext) {}

// ExitStorageClassSpecifier is called when production storageClassSpecifier is exited.
func (s *BaseSuckCParserListener) ExitStorageClassSpecifier(ctx *StorageClassSpecifierContext) {}

// EnterFunctionSpecifier is called when production functionSpecifier is entered.
func (s *BaseSuckCParserListener) EnterFunctionSpecifier(ctx *FunctionSpecifierContext) {}

// ExitFunctionSpecifier is called when production functionSpecifier is exited.
func (s *BaseSuckCParserListener) ExitFunctionSpecifier(ctx *FunctionSpecifierContext) {}

// EnterTypedefName is called when production typedefName is entered.
func (s *BaseSuckCParserListener) EnterTypedefName(ctx *TypedefNameContext) {}

// ExitTypedefName is called when production typedefName is exited.
func (s *BaseSuckCParserListener) ExitTypedefName(ctx *TypedefNameContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BaseSuckCParserListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BaseSuckCParserListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterTrailingTypeSpecifier is called when production trailingTypeSpecifier is entered.
func (s *BaseSuckCParserListener) EnterTrailingTypeSpecifier(ctx *TrailingTypeSpecifierContext) {}

// ExitTrailingTypeSpecifier is called when production trailingTypeSpecifier is exited.
func (s *BaseSuckCParserListener) ExitTrailingTypeSpecifier(ctx *TrailingTypeSpecifierContext) {}

// EnterTypeSpecifierSeq is called when production typeSpecifierSeq is entered.
func (s *BaseSuckCParserListener) EnterTypeSpecifierSeq(ctx *TypeSpecifierSeqContext) {}

// ExitTypeSpecifierSeq is called when production typeSpecifierSeq is exited.
func (s *BaseSuckCParserListener) ExitTypeSpecifierSeq(ctx *TypeSpecifierSeqContext) {}

// EnterTrailingTypeSpecifierSeq is called when production trailingTypeSpecifierSeq is entered.
func (s *BaseSuckCParserListener) EnterTrailingTypeSpecifierSeq(ctx *TrailingTypeSpecifierSeqContext) {
}

// ExitTrailingTypeSpecifierSeq is called when production trailingTypeSpecifierSeq is exited.
func (s *BaseSuckCParserListener) ExitTrailingTypeSpecifierSeq(ctx *TrailingTypeSpecifierSeqContext) {
}

// EnterSimpleTypeLengthModifier is called when production simpleTypeLengthModifier is entered.
func (s *BaseSuckCParserListener) EnterSimpleTypeLengthModifier(ctx *SimpleTypeLengthModifierContext) {
}

// ExitSimpleTypeLengthModifier is called when production simpleTypeLengthModifier is exited.
func (s *BaseSuckCParserListener) ExitSimpleTypeLengthModifier(ctx *SimpleTypeLengthModifierContext) {
}

// EnterSimpleTypeSignednessModifier is called when production simpleTypeSignednessModifier is entered.
func (s *BaseSuckCParserListener) EnterSimpleTypeSignednessModifier(ctx *SimpleTypeSignednessModifierContext) {
}

// ExitSimpleTypeSignednessModifier is called when production simpleTypeSignednessModifier is exited.
func (s *BaseSuckCParserListener) ExitSimpleTypeSignednessModifier(ctx *SimpleTypeSignednessModifierContext) {
}

// EnterSimpleTypeSpecifier is called when production simpleTypeSpecifier is entered.
func (s *BaseSuckCParserListener) EnterSimpleTypeSpecifier(ctx *SimpleTypeSpecifierContext) {}

// ExitSimpleTypeSpecifier is called when production simpleTypeSpecifier is exited.
func (s *BaseSuckCParserListener) ExitSimpleTypeSpecifier(ctx *SimpleTypeSpecifierContext) {}

// EnterTheTypeName is called when production theTypeName is entered.
func (s *BaseSuckCParserListener) EnterTheTypeName(ctx *TheTypeNameContext) {}

// ExitTheTypeName is called when production theTypeName is exited.
func (s *BaseSuckCParserListener) ExitTheTypeName(ctx *TheTypeNameContext) {}

// EnterDecltypeSpecifier is called when production decltypeSpecifier is entered.
func (s *BaseSuckCParserListener) EnterDecltypeSpecifier(ctx *DecltypeSpecifierContext) {}

// ExitDecltypeSpecifier is called when production decltypeSpecifier is exited.
func (s *BaseSuckCParserListener) ExitDecltypeSpecifier(ctx *DecltypeSpecifierContext) {}

// EnterElaboratedTypeSpecifier is called when production elaboratedTypeSpecifier is entered.
func (s *BaseSuckCParserListener) EnterElaboratedTypeSpecifier(ctx *ElaboratedTypeSpecifierContext) {}

// ExitElaboratedTypeSpecifier is called when production elaboratedTypeSpecifier is exited.
func (s *BaseSuckCParserListener) ExitElaboratedTypeSpecifier(ctx *ElaboratedTypeSpecifierContext) {}

// EnterEnumName is called when production enumName is entered.
func (s *BaseSuckCParserListener) EnterEnumName(ctx *EnumNameContext) {}

// ExitEnumName is called when production enumName is exited.
func (s *BaseSuckCParserListener) ExitEnumName(ctx *EnumNameContext) {}

// EnterEnumSpecifier is called when production enumSpecifier is entered.
func (s *BaseSuckCParserListener) EnterEnumSpecifier(ctx *EnumSpecifierContext) {}

// ExitEnumSpecifier is called when production enumSpecifier is exited.
func (s *BaseSuckCParserListener) ExitEnumSpecifier(ctx *EnumSpecifierContext) {}

// EnterEnumHead is called when production enumHead is entered.
func (s *BaseSuckCParserListener) EnterEnumHead(ctx *EnumHeadContext) {}

// ExitEnumHead is called when production enumHead is exited.
func (s *BaseSuckCParserListener) ExitEnumHead(ctx *EnumHeadContext) {}

// EnterOpaqueEnumDeclaration is called when production opaqueEnumDeclaration is entered.
func (s *BaseSuckCParserListener) EnterOpaqueEnumDeclaration(ctx *OpaqueEnumDeclarationContext) {}

// ExitOpaqueEnumDeclaration is called when production opaqueEnumDeclaration is exited.
func (s *BaseSuckCParserListener) ExitOpaqueEnumDeclaration(ctx *OpaqueEnumDeclarationContext) {}

// EnterEnumkey is called when production enumkey is entered.
func (s *BaseSuckCParserListener) EnterEnumkey(ctx *EnumkeyContext) {}

// ExitEnumkey is called when production enumkey is exited.
func (s *BaseSuckCParserListener) ExitEnumkey(ctx *EnumkeyContext) {}

// EnterEnumbase is called when production enumbase is entered.
func (s *BaseSuckCParserListener) EnterEnumbase(ctx *EnumbaseContext) {}

// ExitEnumbase is called when production enumbase is exited.
func (s *BaseSuckCParserListener) ExitEnumbase(ctx *EnumbaseContext) {}

// EnterEnumeratorList is called when production enumeratorList is entered.
func (s *BaseSuckCParserListener) EnterEnumeratorList(ctx *EnumeratorListContext) {}

// ExitEnumeratorList is called when production enumeratorList is exited.
func (s *BaseSuckCParserListener) ExitEnumeratorList(ctx *EnumeratorListContext) {}

// EnterEnumeratorDefinition is called when production enumeratorDefinition is entered.
func (s *BaseSuckCParserListener) EnterEnumeratorDefinition(ctx *EnumeratorDefinitionContext) {}

// ExitEnumeratorDefinition is called when production enumeratorDefinition is exited.
func (s *BaseSuckCParserListener) ExitEnumeratorDefinition(ctx *EnumeratorDefinitionContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *BaseSuckCParserListener) EnterEnumerator(ctx *EnumeratorContext) {}

// ExitEnumerator is called when production enumerator is exited.
func (s *BaseSuckCParserListener) ExitEnumerator(ctx *EnumeratorContext) {}

// EnterNamespaceName is called when production namespaceName is entered.
func (s *BaseSuckCParserListener) EnterNamespaceName(ctx *NamespaceNameContext) {}

// ExitNamespaceName is called when production namespaceName is exited.
func (s *BaseSuckCParserListener) ExitNamespaceName(ctx *NamespaceNameContext) {}

// EnterOriginalNamespaceName is called when production originalNamespaceName is entered.
func (s *BaseSuckCParserListener) EnterOriginalNamespaceName(ctx *OriginalNamespaceNameContext) {}

// ExitOriginalNamespaceName is called when production originalNamespaceName is exited.
func (s *BaseSuckCParserListener) ExitOriginalNamespaceName(ctx *OriginalNamespaceNameContext) {}

// EnterNamespaceDefinition is called when production namespaceDefinition is entered.
func (s *BaseSuckCParserListener) EnterNamespaceDefinition(ctx *NamespaceDefinitionContext) {}

// ExitNamespaceDefinition is called when production namespaceDefinition is exited.
func (s *BaseSuckCParserListener) ExitNamespaceDefinition(ctx *NamespaceDefinitionContext) {}

// EnterNamespaceAlias is called when production namespaceAlias is entered.
func (s *BaseSuckCParserListener) EnterNamespaceAlias(ctx *NamespaceAliasContext) {}

// ExitNamespaceAlias is called when production namespaceAlias is exited.
func (s *BaseSuckCParserListener) ExitNamespaceAlias(ctx *NamespaceAliasContext) {}

// EnterNamespaceAliasDefinition is called when production namespaceAliasDefinition is entered.
func (s *BaseSuckCParserListener) EnterNamespaceAliasDefinition(ctx *NamespaceAliasDefinitionContext) {
}

// ExitNamespaceAliasDefinition is called when production namespaceAliasDefinition is exited.
func (s *BaseSuckCParserListener) ExitNamespaceAliasDefinition(ctx *NamespaceAliasDefinitionContext) {
}

// EnterQualifiednamespaceSpecifier is called when production qualifiednamespaceSpecifier is entered.
func (s *BaseSuckCParserListener) EnterQualifiednamespaceSpecifier(ctx *QualifiednamespaceSpecifierContext) {
}

// ExitQualifiednamespaceSpecifier is called when production qualifiednamespaceSpecifier is exited.
func (s *BaseSuckCParserListener) ExitQualifiednamespaceSpecifier(ctx *QualifiednamespaceSpecifierContext) {
}

// EnterUsingDeclaration is called when production usingDeclaration is entered.
func (s *BaseSuckCParserListener) EnterUsingDeclaration(ctx *UsingDeclarationContext) {}

// ExitUsingDeclaration is called when production usingDeclaration is exited.
func (s *BaseSuckCParserListener) ExitUsingDeclaration(ctx *UsingDeclarationContext) {}

// EnterUsingDirective is called when production usingDirective is entered.
func (s *BaseSuckCParserListener) EnterUsingDirective(ctx *UsingDirectiveContext) {}

// ExitUsingDirective is called when production usingDirective is exited.
func (s *BaseSuckCParserListener) ExitUsingDirective(ctx *UsingDirectiveContext) {}

// EnterAsmDefinition is called when production asmDefinition is entered.
func (s *BaseSuckCParserListener) EnterAsmDefinition(ctx *AsmDefinitionContext) {}

// ExitAsmDefinition is called when production asmDefinition is exited.
func (s *BaseSuckCParserListener) ExitAsmDefinition(ctx *AsmDefinitionContext) {}

// EnterLinkageSpecification is called when production linkageSpecification is entered.
func (s *BaseSuckCParserListener) EnterLinkageSpecification(ctx *LinkageSpecificationContext) {}

// ExitLinkageSpecification is called when production linkageSpecification is exited.
func (s *BaseSuckCParserListener) ExitLinkageSpecification(ctx *LinkageSpecificationContext) {}

// EnterAttributeSpecifierSeq is called when production attributeSpecifierSeq is entered.
func (s *BaseSuckCParserListener) EnterAttributeSpecifierSeq(ctx *AttributeSpecifierSeqContext) {}

// ExitAttributeSpecifierSeq is called when production attributeSpecifierSeq is exited.
func (s *BaseSuckCParserListener) ExitAttributeSpecifierSeq(ctx *AttributeSpecifierSeqContext) {}

// EnterAttributeSpecifier is called when production attributeSpecifier is entered.
func (s *BaseSuckCParserListener) EnterAttributeSpecifier(ctx *AttributeSpecifierContext) {}

// ExitAttributeSpecifier is called when production attributeSpecifier is exited.
func (s *BaseSuckCParserListener) ExitAttributeSpecifier(ctx *AttributeSpecifierContext) {}

// EnterAlignmentSpecifier is called when production alignmentSpecifier is entered.
func (s *BaseSuckCParserListener) EnterAlignmentSpecifier(ctx *AlignmentSpecifierContext) {}

// ExitAlignmentSpecifier is called when production alignmentSpecifier is exited.
func (s *BaseSuckCParserListener) ExitAlignmentSpecifier(ctx *AlignmentSpecifierContext) {}

// EnterAttributeList is called when production attributeList is entered.
func (s *BaseSuckCParserListener) EnterAttributeList(ctx *AttributeListContext) {}

// ExitAttributeList is called when production attributeList is exited.
func (s *BaseSuckCParserListener) ExitAttributeList(ctx *AttributeListContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseSuckCParserListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseSuckCParserListener) ExitAttribute(ctx *AttributeContext) {}

// EnterAttributeNamespace is called when production attributeNamespace is entered.
func (s *BaseSuckCParserListener) EnterAttributeNamespace(ctx *AttributeNamespaceContext) {}

// ExitAttributeNamespace is called when production attributeNamespace is exited.
func (s *BaseSuckCParserListener) ExitAttributeNamespace(ctx *AttributeNamespaceContext) {}

// EnterAttributeArgumentClause is called when production attributeArgumentClause is entered.
func (s *BaseSuckCParserListener) EnterAttributeArgumentClause(ctx *AttributeArgumentClauseContext) {}

// ExitAttributeArgumentClause is called when production attributeArgumentClause is exited.
func (s *BaseSuckCParserListener) ExitAttributeArgumentClause(ctx *AttributeArgumentClauseContext) {}

// EnterBalancedTokenSeq is called when production balancedTokenSeq is entered.
func (s *BaseSuckCParserListener) EnterBalancedTokenSeq(ctx *BalancedTokenSeqContext) {}

// ExitBalancedTokenSeq is called when production balancedTokenSeq is exited.
func (s *BaseSuckCParserListener) ExitBalancedTokenSeq(ctx *BalancedTokenSeqContext) {}

// EnterBalancedtoken is called when production balancedtoken is entered.
func (s *BaseSuckCParserListener) EnterBalancedtoken(ctx *BalancedtokenContext) {}

// ExitBalancedtoken is called when production balancedtoken is exited.
func (s *BaseSuckCParserListener) ExitBalancedtoken(ctx *BalancedtokenContext) {}

// EnterInitDeclarator is called when production initDeclarator is entered.
func (s *BaseSuckCParserListener) EnterInitDeclarator(ctx *InitDeclaratorContext) {}

// ExitInitDeclarator is called when production initDeclarator is exited.
func (s *BaseSuckCParserListener) ExitInitDeclarator(ctx *InitDeclaratorContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *BaseSuckCParserListener) EnterDeclarator(ctx *DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *BaseSuckCParserListener) ExitDeclarator(ctx *DeclaratorContext) {}

// EnterPointerDeclarator is called when production pointerDeclarator is entered.
func (s *BaseSuckCParserListener) EnterPointerDeclarator(ctx *PointerDeclaratorContext) {}

// ExitPointerDeclarator is called when production pointerDeclarator is exited.
func (s *BaseSuckCParserListener) ExitPointerDeclarator(ctx *PointerDeclaratorContext) {}

// EnterNoPointerDeclarator is called when production noPointerDeclarator is entered.
func (s *BaseSuckCParserListener) EnterNoPointerDeclarator(ctx *NoPointerDeclaratorContext) {}

// ExitNoPointerDeclarator is called when production noPointerDeclarator is exited.
func (s *BaseSuckCParserListener) ExitNoPointerDeclarator(ctx *NoPointerDeclaratorContext) {}

// EnterParametersAndQualifiers is called when production parametersAndQualifiers is entered.
func (s *BaseSuckCParserListener) EnterParametersAndQualifiers(ctx *ParametersAndQualifiersContext) {}

// ExitParametersAndQualifiers is called when production parametersAndQualifiers is exited.
func (s *BaseSuckCParserListener) ExitParametersAndQualifiers(ctx *ParametersAndQualifiersContext) {}

// EnterTrailingReturnType is called when production trailingReturnType is entered.
func (s *BaseSuckCParserListener) EnterTrailingReturnType(ctx *TrailingReturnTypeContext) {}

// ExitTrailingReturnType is called when production trailingReturnType is exited.
func (s *BaseSuckCParserListener) ExitTrailingReturnType(ctx *TrailingReturnTypeContext) {}

// EnterPointerOperator is called when production pointerOperator is entered.
func (s *BaseSuckCParserListener) EnterPointerOperator(ctx *PointerOperatorContext) {}

// ExitPointerOperator is called when production pointerOperator is exited.
func (s *BaseSuckCParserListener) ExitPointerOperator(ctx *PointerOperatorContext) {}

// EnterCvQualifierSeq is called when production cvQualifierSeq is entered.
func (s *BaseSuckCParserListener) EnterCvQualifierSeq(ctx *CvQualifierSeqContext) {}

// ExitCvQualifierSeq is called when production cvQualifierSeq is exited.
func (s *BaseSuckCParserListener) ExitCvQualifierSeq(ctx *CvQualifierSeqContext) {}

// EnterCvQualifier is called when production cvQualifier is entered.
func (s *BaseSuckCParserListener) EnterCvQualifier(ctx *CvQualifierContext) {}

// ExitCvQualifier is called when production cvQualifier is exited.
func (s *BaseSuckCParserListener) ExitCvQualifier(ctx *CvQualifierContext) {}

// EnterDeclaratorid is called when production declaratorid is entered.
func (s *BaseSuckCParserListener) EnterDeclaratorid(ctx *DeclaratoridContext) {}

// ExitDeclaratorid is called when production declaratorid is exited.
func (s *BaseSuckCParserListener) ExitDeclaratorid(ctx *DeclaratoridContext) {}

// EnterTheTypeId is called when production theTypeId is entered.
func (s *BaseSuckCParserListener) EnterTheTypeId(ctx *TheTypeIdContext) {}

// ExitTheTypeId is called when production theTypeId is exited.
func (s *BaseSuckCParserListener) ExitTheTypeId(ctx *TheTypeIdContext) {}

// EnterAbstractDeclarator is called when production abstractDeclarator is entered.
func (s *BaseSuckCParserListener) EnterAbstractDeclarator(ctx *AbstractDeclaratorContext) {}

// ExitAbstractDeclarator is called when production abstractDeclarator is exited.
func (s *BaseSuckCParserListener) ExitAbstractDeclarator(ctx *AbstractDeclaratorContext) {}

// EnterPointerAbstractDeclarator is called when production pointerAbstractDeclarator is entered.
func (s *BaseSuckCParserListener) EnterPointerAbstractDeclarator(ctx *PointerAbstractDeclaratorContext) {
}

// ExitPointerAbstractDeclarator is called when production pointerAbstractDeclarator is exited.
func (s *BaseSuckCParserListener) ExitPointerAbstractDeclarator(ctx *PointerAbstractDeclaratorContext) {
}

// EnterNoPointerAbstractDeclarator is called when production noPointerAbstractDeclarator is entered.
func (s *BaseSuckCParserListener) EnterNoPointerAbstractDeclarator(ctx *NoPointerAbstractDeclaratorContext) {
}

// ExitNoPointerAbstractDeclarator is called when production noPointerAbstractDeclarator is exited.
func (s *BaseSuckCParserListener) ExitNoPointerAbstractDeclarator(ctx *NoPointerAbstractDeclaratorContext) {
}

// EnterAbstractPackDeclarator is called when production abstractPackDeclarator is entered.
func (s *BaseSuckCParserListener) EnterAbstractPackDeclarator(ctx *AbstractPackDeclaratorContext) {}

// ExitAbstractPackDeclarator is called when production abstractPackDeclarator is exited.
func (s *BaseSuckCParserListener) ExitAbstractPackDeclarator(ctx *AbstractPackDeclaratorContext) {}

// EnterNoPointerAbstractPackDeclarator is called when production noPointerAbstractPackDeclarator is entered.
func (s *BaseSuckCParserListener) EnterNoPointerAbstractPackDeclarator(ctx *NoPointerAbstractPackDeclaratorContext) {
}

// ExitNoPointerAbstractPackDeclarator is called when production noPointerAbstractPackDeclarator is exited.
func (s *BaseSuckCParserListener) ExitNoPointerAbstractPackDeclarator(ctx *NoPointerAbstractPackDeclaratorContext) {
}

// EnterParameterDeclarationClause is called when production parameterDeclarationClause is entered.
func (s *BaseSuckCParserListener) EnterParameterDeclarationClause(ctx *ParameterDeclarationClauseContext) {
}

// ExitParameterDeclarationClause is called when production parameterDeclarationClause is exited.
func (s *BaseSuckCParserListener) ExitParameterDeclarationClause(ctx *ParameterDeclarationClauseContext) {
}

// EnterParameterDeclarationList is called when production parameterDeclarationList is entered.
func (s *BaseSuckCParserListener) EnterParameterDeclarationList(ctx *ParameterDeclarationListContext) {
}

// ExitParameterDeclarationList is called when production parameterDeclarationList is exited.
func (s *BaseSuckCParserListener) ExitParameterDeclarationList(ctx *ParameterDeclarationListContext) {
}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *BaseSuckCParserListener) EnterParameterDeclaration(ctx *ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *BaseSuckCParserListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseSuckCParserListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseSuckCParserListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterFunctionDeclarator is called when production functionDeclarator is entered.
func (s *BaseSuckCParserListener) EnterFunctionDeclarator(ctx *FunctionDeclaratorContext) {}

// ExitFunctionDeclarator is called when production functionDeclarator is exited.
func (s *BaseSuckCParserListener) ExitFunctionDeclarator(ctx *FunctionDeclaratorContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseSuckCParserListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseSuckCParserListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *BaseSuckCParserListener) EnterInitializer(ctx *InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *BaseSuckCParserListener) ExitInitializer(ctx *InitializerContext) {}

// EnterBraceOrEqualInitializer is called when production braceOrEqualInitializer is entered.
func (s *BaseSuckCParserListener) EnterBraceOrEqualInitializer(ctx *BraceOrEqualInitializerContext) {}

// ExitBraceOrEqualInitializer is called when production braceOrEqualInitializer is exited.
func (s *BaseSuckCParserListener) ExitBraceOrEqualInitializer(ctx *BraceOrEqualInitializerContext) {}

// EnterInitializerClause is called when production initializerClause is entered.
func (s *BaseSuckCParserListener) EnterInitializerClause(ctx *InitializerClauseContext) {}

// ExitInitializerClause is called when production initializerClause is exited.
func (s *BaseSuckCParserListener) ExitInitializerClause(ctx *InitializerClauseContext) {}

// EnterInitializerList is called when production initializerList is entered.
func (s *BaseSuckCParserListener) EnterInitializerList(ctx *InitializerListContext) {}

// ExitInitializerList is called when production initializerList is exited.
func (s *BaseSuckCParserListener) ExitInitializerList(ctx *InitializerListContext) {}

// EnterBracedInitList is called when production bracedInitList is entered.
func (s *BaseSuckCParserListener) EnterBracedInitList(ctx *BracedInitListContext) {}

// ExitBracedInitList is called when production bracedInitList is exited.
func (s *BaseSuckCParserListener) ExitBracedInitList(ctx *BracedInitListContext) {}

// EnterClassName is called when production className is entered.
func (s *BaseSuckCParserListener) EnterClassName(ctx *ClassNameContext) {}

// ExitClassName is called when production className is exited.
func (s *BaseSuckCParserListener) ExitClassName(ctx *ClassNameContext) {}

// EnterClassSpecifier is called when production classSpecifier is entered.
func (s *BaseSuckCParserListener) EnterClassSpecifier(ctx *ClassSpecifierContext) {}

// ExitClassSpecifier is called when production classSpecifier is exited.
func (s *BaseSuckCParserListener) ExitClassSpecifier(ctx *ClassSpecifierContext) {}

// EnterClassHead is called when production classHead is entered.
func (s *BaseSuckCParserListener) EnterClassHead(ctx *ClassHeadContext) {}

// ExitClassHead is called when production classHead is exited.
func (s *BaseSuckCParserListener) ExitClassHead(ctx *ClassHeadContext) {}

// EnterClassHeadName is called when production classHeadName is entered.
func (s *BaseSuckCParserListener) EnterClassHeadName(ctx *ClassHeadNameContext) {}

// ExitClassHeadName is called when production classHeadName is exited.
func (s *BaseSuckCParserListener) ExitClassHeadName(ctx *ClassHeadNameContext) {}

// EnterClassVirtSpecifier is called when production classVirtSpecifier is entered.
func (s *BaseSuckCParserListener) EnterClassVirtSpecifier(ctx *ClassVirtSpecifierContext) {}

// ExitClassVirtSpecifier is called when production classVirtSpecifier is exited.
func (s *BaseSuckCParserListener) ExitClassVirtSpecifier(ctx *ClassVirtSpecifierContext) {}

// EnterMemberSpecification is called when production memberSpecification is entered.
func (s *BaseSuckCParserListener) EnterMemberSpecification(ctx *MemberSpecificationContext) {}

// ExitMemberSpecification is called when production memberSpecification is exited.
func (s *BaseSuckCParserListener) ExitMemberSpecification(ctx *MemberSpecificationContext) {}

// EnterMemberdeclaration is called when production memberdeclaration is entered.
func (s *BaseSuckCParserListener) EnterMemberdeclaration(ctx *MemberdeclarationContext) {}

// ExitMemberdeclaration is called when production memberdeclaration is exited.
func (s *BaseSuckCParserListener) ExitMemberdeclaration(ctx *MemberdeclarationContext) {}

// EnterMemberDeclaratorList is called when production memberDeclaratorList is entered.
func (s *BaseSuckCParserListener) EnterMemberDeclaratorList(ctx *MemberDeclaratorListContext) {}

// ExitMemberDeclaratorList is called when production memberDeclaratorList is exited.
func (s *BaseSuckCParserListener) ExitMemberDeclaratorList(ctx *MemberDeclaratorListContext) {}

// EnterMemberDeclarator is called when production memberDeclarator is entered.
func (s *BaseSuckCParserListener) EnterMemberDeclarator(ctx *MemberDeclaratorContext) {}

// ExitMemberDeclarator is called when production memberDeclarator is exited.
func (s *BaseSuckCParserListener) ExitMemberDeclarator(ctx *MemberDeclaratorContext) {}

// EnterVirtualSpecifierSeq is called when production virtualSpecifierSeq is entered.
func (s *BaseSuckCParserListener) EnterVirtualSpecifierSeq(ctx *VirtualSpecifierSeqContext) {}

// ExitVirtualSpecifierSeq is called when production virtualSpecifierSeq is exited.
func (s *BaseSuckCParserListener) ExitVirtualSpecifierSeq(ctx *VirtualSpecifierSeqContext) {}

// EnterVirtualSpecifier is called when production virtualSpecifier is entered.
func (s *BaseSuckCParserListener) EnterVirtualSpecifier(ctx *VirtualSpecifierContext) {}

// ExitVirtualSpecifier is called when production virtualSpecifier is exited.
func (s *BaseSuckCParserListener) ExitVirtualSpecifier(ctx *VirtualSpecifierContext) {}

// EnterPureSpecifier is called when production pureSpecifier is entered.
func (s *BaseSuckCParserListener) EnterPureSpecifier(ctx *PureSpecifierContext) {}

// ExitPureSpecifier is called when production pureSpecifier is exited.
func (s *BaseSuckCParserListener) ExitPureSpecifier(ctx *PureSpecifierContext) {}

// EnterBaseClause is called when production baseClause is entered.
func (s *BaseSuckCParserListener) EnterBaseClause(ctx *BaseClauseContext) {}

// ExitBaseClause is called when production baseClause is exited.
func (s *BaseSuckCParserListener) ExitBaseClause(ctx *BaseClauseContext) {}

// EnterBaseSpecifierList is called when production baseSpecifierList is entered.
func (s *BaseSuckCParserListener) EnterBaseSpecifierList(ctx *BaseSpecifierListContext) {}

// ExitBaseSpecifierList is called when production baseSpecifierList is exited.
func (s *BaseSuckCParserListener) ExitBaseSpecifierList(ctx *BaseSpecifierListContext) {}

// EnterBaseSpecifier is called when production baseSpecifier is entered.
func (s *BaseSuckCParserListener) EnterBaseSpecifier(ctx *BaseSpecifierContext) {}

// ExitBaseSpecifier is called when production baseSpecifier is exited.
func (s *BaseSuckCParserListener) ExitBaseSpecifier(ctx *BaseSpecifierContext) {}

// EnterClassOrDeclType is called when production classOrDeclType is entered.
func (s *BaseSuckCParserListener) EnterClassOrDeclType(ctx *ClassOrDeclTypeContext) {}

// ExitClassOrDeclType is called when production classOrDeclType is exited.
func (s *BaseSuckCParserListener) ExitClassOrDeclType(ctx *ClassOrDeclTypeContext) {}

// EnterBaseTypeSpecifier is called when production baseTypeSpecifier is entered.
func (s *BaseSuckCParserListener) EnterBaseTypeSpecifier(ctx *BaseTypeSpecifierContext) {}

// ExitBaseTypeSpecifier is called when production baseTypeSpecifier is exited.
func (s *BaseSuckCParserListener) ExitBaseTypeSpecifier(ctx *BaseTypeSpecifierContext) {}

// EnterAccessSpecifier is called when production accessSpecifier is entered.
func (s *BaseSuckCParserListener) EnterAccessSpecifier(ctx *AccessSpecifierContext) {}

// ExitAccessSpecifier is called when production accessSpecifier is exited.
func (s *BaseSuckCParserListener) ExitAccessSpecifier(ctx *AccessSpecifierContext) {}

// EnterConversionFunctionId is called when production conversionFunctionId is entered.
func (s *BaseSuckCParserListener) EnterConversionFunctionId(ctx *ConversionFunctionIdContext) {}

// ExitConversionFunctionId is called when production conversionFunctionId is exited.
func (s *BaseSuckCParserListener) ExitConversionFunctionId(ctx *ConversionFunctionIdContext) {}

// EnterConversionTypeId is called when production conversionTypeId is entered.
func (s *BaseSuckCParserListener) EnterConversionTypeId(ctx *ConversionTypeIdContext) {}

// ExitConversionTypeId is called when production conversionTypeId is exited.
func (s *BaseSuckCParserListener) ExitConversionTypeId(ctx *ConversionTypeIdContext) {}

// EnterConversionDeclarator is called when production conversionDeclarator is entered.
func (s *BaseSuckCParserListener) EnterConversionDeclarator(ctx *ConversionDeclaratorContext) {}

// ExitConversionDeclarator is called when production conversionDeclarator is exited.
func (s *BaseSuckCParserListener) ExitConversionDeclarator(ctx *ConversionDeclaratorContext) {}

// EnterConstructorInitializer is called when production constructorInitializer is entered.
func (s *BaseSuckCParserListener) EnterConstructorInitializer(ctx *ConstructorInitializerContext) {}

// ExitConstructorInitializer is called when production constructorInitializer is exited.
func (s *BaseSuckCParserListener) ExitConstructorInitializer(ctx *ConstructorInitializerContext) {}

// EnterMemInitializerList is called when production memInitializerList is entered.
func (s *BaseSuckCParserListener) EnterMemInitializerList(ctx *MemInitializerListContext) {}

// ExitMemInitializerList is called when production memInitializerList is exited.
func (s *BaseSuckCParserListener) ExitMemInitializerList(ctx *MemInitializerListContext) {}

// EnterMemInitializer is called when production memInitializer is entered.
func (s *BaseSuckCParserListener) EnterMemInitializer(ctx *MemInitializerContext) {}

// ExitMemInitializer is called when production memInitializer is exited.
func (s *BaseSuckCParserListener) ExitMemInitializer(ctx *MemInitializerContext) {}

// EnterMeminitializerid is called when production meminitializerid is entered.
func (s *BaseSuckCParserListener) EnterMeminitializerid(ctx *MeminitializeridContext) {}

// ExitMeminitializerid is called when production meminitializerid is exited.
func (s *BaseSuckCParserListener) ExitMeminitializerid(ctx *MeminitializeridContext) {}

// EnterOperatorFunctionId is called when production operatorFunctionId is entered.
func (s *BaseSuckCParserListener) EnterOperatorFunctionId(ctx *OperatorFunctionIdContext) {}

// ExitOperatorFunctionId is called when production operatorFunctionId is exited.
func (s *BaseSuckCParserListener) ExitOperatorFunctionId(ctx *OperatorFunctionIdContext) {}

// EnterLiteralOperatorId is called when production literalOperatorId is entered.
func (s *BaseSuckCParserListener) EnterLiteralOperatorId(ctx *LiteralOperatorIdContext) {}

// ExitLiteralOperatorId is called when production literalOperatorId is exited.
func (s *BaseSuckCParserListener) ExitLiteralOperatorId(ctx *LiteralOperatorIdContext) {}

// EnterTemplateDeclaration is called when production templateDeclaration is entered.
func (s *BaseSuckCParserListener) EnterTemplateDeclaration(ctx *TemplateDeclarationContext) {}

// ExitTemplateDeclaration is called when production templateDeclaration is exited.
func (s *BaseSuckCParserListener) ExitTemplateDeclaration(ctx *TemplateDeclarationContext) {}

// EnterTemplateparameterList is called when production templateparameterList is entered.
func (s *BaseSuckCParserListener) EnterTemplateparameterList(ctx *TemplateparameterListContext) {}

// ExitTemplateparameterList is called when production templateparameterList is exited.
func (s *BaseSuckCParserListener) ExitTemplateparameterList(ctx *TemplateparameterListContext) {}

// EnterTemplateParameter is called when production templateParameter is entered.
func (s *BaseSuckCParserListener) EnterTemplateParameter(ctx *TemplateParameterContext) {}

// ExitTemplateParameter is called when production templateParameter is exited.
func (s *BaseSuckCParserListener) ExitTemplateParameter(ctx *TemplateParameterContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseSuckCParserListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseSuckCParserListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterSimpleTemplateId is called when production simpleTemplateId is entered.
func (s *BaseSuckCParserListener) EnterSimpleTemplateId(ctx *SimpleTemplateIdContext) {}

// ExitSimpleTemplateId is called when production simpleTemplateId is exited.
func (s *BaseSuckCParserListener) ExitSimpleTemplateId(ctx *SimpleTemplateIdContext) {}

// EnterTemplateId is called when production templateId is entered.
func (s *BaseSuckCParserListener) EnterTemplateId(ctx *TemplateIdContext) {}

// ExitTemplateId is called when production templateId is exited.
func (s *BaseSuckCParserListener) ExitTemplateId(ctx *TemplateIdContext) {}

// EnterTemplateName is called when production templateName is entered.
func (s *BaseSuckCParserListener) EnterTemplateName(ctx *TemplateNameContext) {}

// ExitTemplateName is called when production templateName is exited.
func (s *BaseSuckCParserListener) ExitTemplateName(ctx *TemplateNameContext) {}

// EnterTemplateArgumentList is called when production templateArgumentList is entered.
func (s *BaseSuckCParserListener) EnterTemplateArgumentList(ctx *TemplateArgumentListContext) {}

// ExitTemplateArgumentList is called when production templateArgumentList is exited.
func (s *BaseSuckCParserListener) ExitTemplateArgumentList(ctx *TemplateArgumentListContext) {}

// EnterTemplateArgument is called when production templateArgument is entered.
func (s *BaseSuckCParserListener) EnterTemplateArgument(ctx *TemplateArgumentContext) {}

// ExitTemplateArgument is called when production templateArgument is exited.
func (s *BaseSuckCParserListener) ExitTemplateArgument(ctx *TemplateArgumentContext) {}

// EnterTypeNameSpecifier is called when production typeNameSpecifier is entered.
func (s *BaseSuckCParserListener) EnterTypeNameSpecifier(ctx *TypeNameSpecifierContext) {}

// ExitTypeNameSpecifier is called when production typeNameSpecifier is exited.
func (s *BaseSuckCParserListener) ExitTypeNameSpecifier(ctx *TypeNameSpecifierContext) {}

// EnterExplicitInstantiation is called when production explicitInstantiation is entered.
func (s *BaseSuckCParserListener) EnterExplicitInstantiation(ctx *ExplicitInstantiationContext) {}

// ExitExplicitInstantiation is called when production explicitInstantiation is exited.
func (s *BaseSuckCParserListener) ExitExplicitInstantiation(ctx *ExplicitInstantiationContext) {}

// EnterExplicitSpecialization is called when production explicitSpecialization is entered.
func (s *BaseSuckCParserListener) EnterExplicitSpecialization(ctx *ExplicitSpecializationContext) {}

// ExitExplicitSpecialization is called when production explicitSpecialization is exited.
func (s *BaseSuckCParserListener) ExitExplicitSpecialization(ctx *ExplicitSpecializationContext) {}

// EnterTryBlock is called when production tryBlock is entered.
func (s *BaseSuckCParserListener) EnterTryBlock(ctx *TryBlockContext) {}

// ExitTryBlock is called when production tryBlock is exited.
func (s *BaseSuckCParserListener) ExitTryBlock(ctx *TryBlockContext) {}

// EnterFunctionTryBlock is called when production functionTryBlock is entered.
func (s *BaseSuckCParserListener) EnterFunctionTryBlock(ctx *FunctionTryBlockContext) {}

// ExitFunctionTryBlock is called when production functionTryBlock is exited.
func (s *BaseSuckCParserListener) ExitFunctionTryBlock(ctx *FunctionTryBlockContext) {}

// EnterHandlerSeq is called when production handlerSeq is entered.
func (s *BaseSuckCParserListener) EnterHandlerSeq(ctx *HandlerSeqContext) {}

// ExitHandlerSeq is called when production handlerSeq is exited.
func (s *BaseSuckCParserListener) ExitHandlerSeq(ctx *HandlerSeqContext) {}

// EnterHandler is called when production handler is entered.
func (s *BaseSuckCParserListener) EnterHandler(ctx *HandlerContext) {}

// ExitHandler is called when production handler is exited.
func (s *BaseSuckCParserListener) ExitHandler(ctx *HandlerContext) {}

// EnterExceptionDeclaration is called when production exceptionDeclaration is entered.
func (s *BaseSuckCParserListener) EnterExceptionDeclaration(ctx *ExceptionDeclarationContext) {}

// ExitExceptionDeclaration is called when production exceptionDeclaration is exited.
func (s *BaseSuckCParserListener) ExitExceptionDeclaration(ctx *ExceptionDeclarationContext) {}

// EnterThrowExpression is called when production throwExpression is entered.
func (s *BaseSuckCParserListener) EnterThrowExpression(ctx *ThrowExpressionContext) {}

// ExitThrowExpression is called when production throwExpression is exited.
func (s *BaseSuckCParserListener) ExitThrowExpression(ctx *ThrowExpressionContext) {}

// EnterExceptionSpecification is called when production exceptionSpecification is entered.
func (s *BaseSuckCParserListener) EnterExceptionSpecification(ctx *ExceptionSpecificationContext) {}

// ExitExceptionSpecification is called when production exceptionSpecification is exited.
func (s *BaseSuckCParserListener) ExitExceptionSpecification(ctx *ExceptionSpecificationContext) {}

// EnterDynamicExceptionSpecification is called when production dynamicExceptionSpecification is entered.
func (s *BaseSuckCParserListener) EnterDynamicExceptionSpecification(ctx *DynamicExceptionSpecificationContext) {
}

// ExitDynamicExceptionSpecification is called when production dynamicExceptionSpecification is exited.
func (s *BaseSuckCParserListener) ExitDynamicExceptionSpecification(ctx *DynamicExceptionSpecificationContext) {
}

// EnterTypeIdList is called when production typeIdList is entered.
func (s *BaseSuckCParserListener) EnterTypeIdList(ctx *TypeIdListContext) {}

// ExitTypeIdList is called when production typeIdList is exited.
func (s *BaseSuckCParserListener) ExitTypeIdList(ctx *TypeIdListContext) {}

// EnterNoeExceptSpecification is called when production noeExceptSpecification is entered.
func (s *BaseSuckCParserListener) EnterNoeExceptSpecification(ctx *NoeExceptSpecificationContext) {}

// ExitNoeExceptSpecification is called when production noeExceptSpecification is exited.
func (s *BaseSuckCParserListener) ExitNoeExceptSpecification(ctx *NoeExceptSpecificationContext) {}

// EnterTheOperator is called when production theOperator is entered.
func (s *BaseSuckCParserListener) EnterTheOperator(ctx *TheOperatorContext) {}

// ExitTheOperator is called when production theOperator is exited.
func (s *BaseSuckCParserListener) ExitTheOperator(ctx *TheOperatorContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseSuckCParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseSuckCParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterOctalLiteral is called when production octalLiteral is entered.
func (s *BaseSuckCParserListener) EnterOctalLiteral(ctx *OctalLiteralContext) {}

// ExitOctalLiteral is called when production octalLiteral is exited.
func (s *BaseSuckCParserListener) ExitOctalLiteral(ctx *OctalLiteralContext) {}

// EnterHexadecimalLiteral is called when production hexadecimalLiteral is entered.
func (s *BaseSuckCParserListener) EnterHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// ExitHexadecimalLiteral is called when production hexadecimalLiteral is exited.
func (s *BaseSuckCParserListener) ExitHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// EnterBinaryLiteral is called when production binaryLiteral is entered.
func (s *BaseSuckCParserListener) EnterBinaryLiteral(ctx *BinaryLiteralContext) {}

// ExitBinaryLiteral is called when production binaryLiteral is exited.
func (s *BaseSuckCParserListener) ExitBinaryLiteral(ctx *BinaryLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseSuckCParserListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseSuckCParserListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterCharacterLiteral is called when production characterLiteral is entered.
func (s *BaseSuckCParserListener) EnterCharacterLiteral(ctx *CharacterLiteralContext) {}

// ExitCharacterLiteral is called when production characterLiteral is exited.
func (s *BaseSuckCParserListener) ExitCharacterLiteral(ctx *CharacterLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseSuckCParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseSuckCParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterFloatingLiteral is called when production floatingLiteral is entered.
func (s *BaseSuckCParserListener) EnterFloatingLiteral(ctx *FloatingLiteralContext) {}

// ExitFloatingLiteral is called when production floatingLiteral is exited.
func (s *BaseSuckCParserListener) ExitFloatingLiteral(ctx *FloatingLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseSuckCParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseSuckCParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterPointerLiteral is called when production pointerLiteral is entered.
func (s *BaseSuckCParserListener) EnterPointerLiteral(ctx *PointerLiteralContext) {}

// ExitPointerLiteral is called when production pointerLiteral is exited.
func (s *BaseSuckCParserListener) ExitPointerLiteral(ctx *PointerLiteralContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseSuckCParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseSuckCParserListener) ExitLiteral(ctx *LiteralContext) {}
