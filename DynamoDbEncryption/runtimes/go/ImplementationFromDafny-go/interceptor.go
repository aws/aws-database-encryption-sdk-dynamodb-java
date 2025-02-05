package the_program

import (
	"context"
	"fmt"

	"github.com/aws/aws-database-encryption-sdk-dynamodb/awscryptographydbencryptionsdkdynamodbsmithygeneratedtypes"
	"github.com/aws/aws-database-encryption-sdk-dynamodb/awscryptographydbencryptionsdkdynamodbtransformssmithygenerated"
	"github.com/aws/aws-database-encryption-sdk-dynamodb/awscryptographydbencryptionsdkdynamodbtransformssmithygeneratedtypes"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
		ctx = m.handleRequestInterception(ctx, in.Parameters)
		return next.HandleInitialize(ctx, in)
	})
}

// handleRequestInterception handles the interception logic before the DynamoDB operation
func (m DbEsdkMiddleware) handleRequestInterception(ctx context.Context, params interface{}) context.Context {
	switch v := params.(type) {
	case *dynamodb.PutItemInput:
		ctx = middleware.WithStackValue(ctx, "originalInput", *DeepCopyPutItemInput(v))
		transformedRequest, err := m.client.PutItemInputTransform(context.TODO(), awscryptographydbencryptionsdkdynamodbtransformssmithygeneratedtypes.PutItemInputTransformInput{
			SdkInput: *v,
		})
		if err != nil {
			fmt.Println(err)
		}
		*v = transformedRequest.TransformedInput
	case *dynamodb.GetItemInput:
		ctx = middleware.WithStackValue(ctx, "originalInput", *DeepCopyGetItemInput(v))
		transformedRequest, err := m.client.GetItemInputTransform(context.TODO(), awscryptographydbencryptionsdkdynamodbtransformssmithygeneratedtypes.GetItemInputTransformInput{
			SdkInput: *v,
		})
		if err != nil {
			fmt.Println(err)
		}
		*v = transformedRequest.TransformedInput
		// case *dynamodb.BatchExecuteStatementInput:
		// 	m.originalRequests["BatchExecuteStatementInput"] = *DeepCopyBatchExecuteStatementInput(v)
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
	switch v := response.(type) {
	case *dynamodb.PutItemOutput:
		transformedRequest, err := m.client.PutItemOutputTransform(context.TODO(), awscryptographydbencryptionsdkdynamodbtransformssmithygeneratedtypes.PutItemOutputTransformInput{
			OriginalInput: middleware.GetStackValue(ctx, "originalInput").(dynamodb.PutItemInput),
			SdkOutput:     *v,
		})
		if err != nil {
			fmt.Println(err)
		}
		*v = transformedRequest.TransformedOutput
	case *dynamodb.GetItemOutput:
		transformedRequest, err := m.client.GetItemOutputTransform(context.TODO(), awscryptographydbencryptionsdkdynamodbtransformssmithygeneratedtypes.GetItemOutputTransformInput{
			OriginalInput: middleware.GetStackValue(ctx, "originalInput").(dynamodb.GetItemInput),
			SdkOutput:     *v,
		})
		if err != nil {
			fmt.Println(err)
		}
		*v = transformedRequest.TransformedOutput
	}

}

// DeepCopyPutItemInput performs a deep copy of a PutItemInput struct.
func DeepCopyPutItemInput(input *dynamodb.PutItemInput) *dynamodb.PutItemInput {
	if input == nil {
		return nil
	}

	copyItem := make(map[string]types.AttributeValue, len(input.Item))
	for k, v := range input.Item {
		copyItem[k] = deepCopyAttributeValue(v)
	}

	copyExpected := make(map[string]types.ExpectedAttributeValue, len(input.Expected))
	for k, v := range input.Expected {
		copyExpected[k] = v // ExpectedAttributeValue is a struct, so direct assignment is fine
	}

	copyExprNames := make(map[string]string, len(input.ExpressionAttributeNames))
	for k, v := range input.ExpressionAttributeNames {
		copyExprNames[k] = v
	}

	copyExprValues := make(map[string]types.AttributeValue, len(input.ExpressionAttributeValues))
	for k, v := range input.ExpressionAttributeValues {
		copyExprValues[k] = deepCopyAttributeValue(v)
	}

	// Copying string pointers
	var tableName *string
	if input.TableName != nil {
		t := *input.TableName
		tableName = &t
	}

	var conditionExpression *string
	if input.ConditionExpression != nil {
		ce := *input.ConditionExpression
		conditionExpression = &ce
	}

	return &dynamodb.PutItemInput{
		Item:                                copyItem,
		TableName:                           tableName,
		ConditionExpression:                 conditionExpression,
		ConditionalOperator:                 input.ConditionalOperator,
		Expected:                            copyExpected,
		ExpressionAttributeNames:            copyExprNames,
		ExpressionAttributeValues:           copyExprValues,
		ReturnConsumedCapacity:              input.ReturnConsumedCapacity,
		ReturnItemCollectionMetrics:         input.ReturnItemCollectionMetrics,
		ReturnValues:                        input.ReturnValues,
		ReturnValuesOnConditionCheckFailure: input.ReturnValuesOnConditionCheckFailure,
	}
}

// DeepCopyGetItemInput performs a deep copy of a GetItemInput struct.
func DeepCopyGetItemInput(input *dynamodb.GetItemInput) *dynamodb.GetItemInput {
	if input == nil {
		return nil
	}

	// Deep copy Key map
	copyKey := make(map[string]types.AttributeValue, len(input.Key))
	for k, v := range input.Key {
		copyKey[k] = deepCopyAttributeValue(v)
	}

	// Deep copy ExpressionAttributeNames map
	copyExprNames := make(map[string]string, len(input.ExpressionAttributeNames))
	for k, v := range input.ExpressionAttributeNames {
		copyExprNames[k] = v
	}

	// Deep copy AttributesToGet slice
	copyAttributesToGet := make([]string, len(input.AttributesToGet))
	copy(copyAttributesToGet, input.AttributesToGet)

	// Copy string and bool pointers
	var tableName *string
	if input.TableName != nil {
		t := *input.TableName
		tableName = &t
	}

	var projectionExpression *string
	if input.ProjectionExpression != nil {
		pe := *input.ProjectionExpression
		projectionExpression = &pe
	}

	var consistentRead *bool
	if input.ConsistentRead != nil {
		cr := *input.ConsistentRead
		consistentRead = &cr
	}

	return &dynamodb.GetItemInput{
		Key:                      copyKey,
		TableName:                tableName,
		AttributesToGet:          copyAttributesToGet,
		ConsistentRead:           consistentRead,
		ExpressionAttributeNames: copyExprNames,
		ProjectionExpression:     projectionExpression,
		ReturnConsumedCapacity:   input.ReturnConsumedCapacity,
	}
}

// deepCopyAttributeValue performs a deep copy of AttributeValue.
func deepCopyAttributeValue(attr types.AttributeValue) types.AttributeValue {
	switch v := attr.(type) {
	case *types.AttributeValueMemberS:
		return &types.AttributeValueMemberS{Value: v.Value}
	case *types.AttributeValueMemberN:
		return &types.AttributeValueMemberN{Value: v.Value}
	case *types.AttributeValueMemberB:
		b := make([]byte, len(v.Value))
		copy(b, v.Value)
		return &types.AttributeValueMemberB{Value: b}
	case *types.AttributeValueMemberBOOL:
		return &types.AttributeValueMemberBOOL{Value: v.Value}
	case *types.AttributeValueMemberNULL:
		return &types.AttributeValueMemberNULL{Value: v.Value}
	case *types.AttributeValueMemberM:
		newMap := make(map[string]types.AttributeValue, len(v.Value))
		for key, value := range v.Value {
			newMap[key] = deepCopyAttributeValue(value)
		}
		return &types.AttributeValueMemberM{Value: newMap}
	case *types.AttributeValueMemberL:
		newList := make([]types.AttributeValue, len(v.Value))
		for i, value := range v.Value {
			newList[i] = deepCopyAttributeValue(value)
		}
		return &types.AttributeValueMemberL{Value: newList}
	case *types.AttributeValueMemberSS:
		newSS := make([]string, len(v.Value))
		copy(newSS, v.Value)
		return &types.AttributeValueMemberSS{Value: newSS}
	case *types.AttributeValueMemberNS:
		newNS := make([]string, len(v.Value))
		copy(newNS, v.Value)
		return &types.AttributeValueMemberNS{Value: newNS}
	case *types.AttributeValueMemberBS:
		newBS := make([][]byte, len(v.Value))
		for i, b := range v.Value {
			newBS[i] = make([]byte, len(b))
			copy(newBS[i], b)
		}
		return &types.AttributeValueMemberBS{Value: newBS}
	default:
		return attr // Unknown type, return as is (might be a bug)
	}
}
