module github.com/aws/aws-database-encryption-sdk-dynamodb/testvectors/test

go 1.23.2

replace (
	github.com/aws/aws-cryptographic-material-providers-library/releases/go/dynamodb => ../../../../submodules/MaterialProviders/ComAmazonawsDynamodb/runtimes/go/ImplementationFromDafny-go/
	github.com/aws/aws-cryptographic-material-providers-library/releases/go/kms => ../../../../submodules/MaterialProviders/ComAmazonawsKms/runtimes/go/ImplementationFromDafny-go/
	github.com/aws/aws-cryptographic-material-providers-library/releases/go/mpl => ../../../../submodules/MaterialProviders/AwsCryptographicMaterialProviders/runtimes/go/ImplementationFromDafny-go/
	github.com/aws/aws-cryptographic-material-providers-library/releases/go/primitives => ../../../../submodules/MaterialProviders/AwsCryptographyPrimitives/runtimes/go/ImplementationFromDafny-go/
	github.com/aws/aws-cryptographic-material-providers-library/releases/go/smithy-dafny-standard-library => ../../../../submodules/MaterialProviders/StandardLibrary/runtimes/go/ImplementationFromDafny-go/
	github.com/aws/aws-cryptographic-material-providers-library/testvectors => ../../../../submodules/MaterialProviders/TestVectorsAwsCryptographicMaterialProviders/runtimes/go/ImplementationFromDafny-go/
	github.com/aws/aws-database-encryption-sdk-dynamodb => ../../../../DynamoDbEncryption/runtimes/go/ImplementationFromDafny-go/
	github.com/aws/aws-database-encryption-sdk-dynamodb/testvectors => ../ImplementationFromDafny-go/
)
