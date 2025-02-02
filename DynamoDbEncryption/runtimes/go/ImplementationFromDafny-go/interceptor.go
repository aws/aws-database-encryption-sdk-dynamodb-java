package the_program

import (
	"context"
	"fmt"

	"github.com/aws/aws-database-encryption-sdk-dynamodb/awscryptographydbencryptionsdkdynamodbsmithygeneratedtypes"
	"github.com/aws/aws-database-encryption-sdk-dynamodb/awscryptographydbencryptionsdkdynamodbtransformssmithygenerated"
	"github.com/aws/aws-database-encryption-sdk-dynamodb/awscryptographydbencryptionsdkdynamodbtransformssmithygeneratedtypes"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/smithy-go/middleware"
)

// DbEsdkInterceptor handles encryption/decryption of DynamoDB items
type DbEsdkMiddleware struct {
	client           *awscryptographydbencryptionsdkdynamodbtransformssmithygenerated.Client
	originalRequests map[string]interface{}
}

func NewDbEsdkMiddleware(config awscryptographydbencryptionsdkdynamodbsmithygeneratedtypes.DynamoDbTablesEncryptionConfig) (*DbEsdkMiddleware, error) {
	client, err := awscryptographydbencryptionsdkdynamodbtransformssmithygenerated.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &DbEsdkMiddleware{
		client:           client,
		originalRequests: make(map[string]interface{}),
	}, nil
}

func (m DbEsdkMiddleware) CreateMiddleware() func(options *dynamodb.Options) {
	return func(options *dynamodb.Options) {
		options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
			// Add request interceptor at the beginning of Initialize step
			requestIntercetor := m.createRequestInterceptor()
			if err := stack.Initialize.Add(requestIntercetor, middleware.Before); err != nil {
				return err
			}
			// Add response interceptor at the end of Finalize step
			return stack.Finalize.Add(m.createResponseInterceptor(), middleware.After)
		})
	}
}

func (m DbEsdkMiddleware) createRequestInterceptor() middleware.InitializeMiddleware {
	return middleware.InitializeMiddlewareFunc("RequestInterceptor", func(
		ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler,
	) (
		out middleware.InitializeOutput, metadata middleware.Metadata, err error,
	) {
		m.handleRequestInterception(ctx, in.Parameters)
		return next.HandleInitialize(ctx, in)
	})
}

// handleRequestInterception handles the interception logic before the DynamoDB operation
func (m DbEsdkMiddleware) handleRequestInterception(ctx context.Context, params interface{}) context.Context {
	if v, ok := params.(*dynamodb.PutItemInput); ok {
		ctx = middleware.WithStackValue(ctx, "originalInput", v)
		transformedRequest, err := m.client.PutItemInputTransform(context.TODO(), awscryptographydbencryptionsdkdynamodbtransformssmithygeneratedtypes.PutItemInputTransformInput{
			SdkInput: *v,
		})
		if err != nil {
			fmt.Println(err)
		}
		*v = transformedRequest.TransformedInput
	}
	if v, ok := params.(*dynamodb.BatchExecuteStatementInput); ok {
		BatchExecuteStatementInputTransformOutput, err := m.client.BatchExecuteStatementInputTransform(context.TODO(), awscryptographydbencryptionsdkdynamodbtransformssmithygeneratedtypes.BatchExecuteStatementInputTransformInput{
			SdkInput: *v,
		})
		if err != nil {
			fmt.Println(err)
		}
		*v = BatchExecuteStatementInputTransformOutput.TransformedInput
	}
	return ctx
}

// createResponseInterceptor creates and returns the middleware interceptor for responses
func (m DbEsdkMiddleware) createResponseInterceptor() middleware.FinalizeMiddleware {
	return middleware.FinalizeMiddlewareFunc("ResponseInterceptor", func(
		ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler,
	) (
		out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
	) {
		// First let the request complete
		result, metadata, err := next.HandleFinalize(ctx, in)
		if err != nil {
			return result, metadata, err
		}
		// Then intercept the response
		m.handleResponseInterception(ctx, result.Result)
		return result, metadata, err
	})
}

// handleResponseInterception handles the interception logic after the DynamoDB operation
func (m DbEsdkMiddleware) handleResponseInterception(ctx context.Context, response interface{}) {
	if v, ok := response.(*dynamodb.PutItemOutput); ok {
		fmt.Println(ctx.Value("originalInput").(dynamodb.PutItemInput))
		transformedRequest, err := m.client.PutItemOutputTransform(context.TODO(), awscryptographydbencryptionsdkdynamodbtransformssmithygeneratedtypes.PutItemOutputTransformInput{
			OriginalInput: m.originalRequests["PutItemInput"].(dynamodb.PutItemInput),
			SdkOutput:     *v,
		})
		if err != nil {
			fmt.Println(err)
		}
		*v = transformedRequest.TransformedOutput
	}
	// if getItemOutput, ok := response.(*dynamodb.GetItemOutput); ok {
	// 	fmt.Println("GetItemOutput Response intercepted:")
	// 	if age, ok := getItemOutput.Item["Age"].(*types.AttributeValueMemberN); ok {
	// 		fmt.Println("Age:", age.Value)
	// 	}
	// 	if id, ok := getItemOutput.Item["ID"].(*types.AttributeValueMemberN); ok {
	// 		fmt.Println("ID:", id.Value)
	// 	}
	// 	if name, ok := getItemOutput.Item["Name"].(*types.AttributeValueMemberS); ok {
	// 		fmt.Println("Name:", name.Value)
	// 	}
	// 	if intercepted, ok := getItemOutput.Item["intercepted attribute"].(*types.AttributeValueMemberS); ok {
	// 		fmt.Println("intercepted attribute:", intercepted.Value)
	// 	}
	// 	getItemOutput.Item["intercepted attribute"] = &types.AttributeValueMemberS{Value: "I read your data "}
	// 	// You can modify the response here if needed
	// }
}
