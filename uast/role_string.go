// Code generated by "stringer -type=Role"; DO NOT EDIT

package uast

import "fmt"

const _Role_name = "SimpleIdentifierQualifiedIdentifierExpressionStatementFilePackageDeclarationImportDeclarationImportPathImportAliasFunctionDeclarationTypeDeclarationVisibleFromInstanceVisibleFromTypeVisibleFromSubtypeVisibleFromPackageVisibleFromSubpackageVisibleFromModuleVisibleFromFriendVisibleFromWorldIfIfConditionIfBodyIfElseSwitchSwitchCaseSwitchCaseConditionSwitchCaseBodySwitchDefaultForForInitForExpressionForUpdateForBodyForEachWhileWhileConditionWhileBodyDoWhileDoWhileConditionDoWhileBodyBreakContinueBlockBlockScopeReturnTryTryBodyTryCatchTryFinallyThrowAssertMethodInvocationMethodInvocationObjectMethodInvocationNameMethodInvocationArgumentNoopLiteralNullLiteralStringLiteralNumberLiteralTypeLiteralTypePrimitiveTypeAssignmentAssignmentVariableAssignmentValueThisCommentDocumentationWhitespace"

var _Role_index = [...]uint16{0, 16, 35, 45, 54, 58, 76, 93, 103, 114, 133, 148, 167, 182, 200, 218, 239, 256, 273, 289, 291, 302, 308, 314, 320, 330, 349, 363, 376, 379, 386, 399, 408, 415, 422, 427, 441, 450, 457, 473, 484, 489, 497, 502, 512, 518, 521, 528, 536, 546, 551, 557, 573, 595, 615, 639, 643, 650, 661, 674, 687, 698, 702, 715, 725, 743, 758, 762, 769, 782, 792}

func (i Role) String() string {
	if i < 0 || i >= Role(len(_Role_index)-1) {
		return fmt.Sprintf("Role(%d)", i)
	}
	return _Role_name[_Role_index[i]:_Role_index[i+1]]
}
