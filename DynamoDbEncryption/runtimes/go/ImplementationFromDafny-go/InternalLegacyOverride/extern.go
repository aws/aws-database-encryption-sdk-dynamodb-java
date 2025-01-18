package InternalLegacyOverride

import (
	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/smithy-dafny-standard-library/Wrappers"
	"github.com/aws/aws-database-encryption-sdk-dynamodb/awscryptographydbencryptionsdkdynamodbitemencryptorsmithygeneratedtypes"
)

func (CompanionStruct_InternalLegacyOverride_) Build(config interface{}) Wrappers.Result {
	// Create error for unsupported legacy configuration
	err := &awscryptographydbencryptionsdkdynamodbitemencryptorsmithygeneratedtypes.DynamoDbItemEncryptorException{
		Message: "Legacy configuration unsupported.",
	}
	// Create and return a failure result
	return Wrappers.Companion_Result_.Create_Failure_(err)
}
