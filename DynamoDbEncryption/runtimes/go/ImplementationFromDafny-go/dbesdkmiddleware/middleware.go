package dbesdkmiddleware

import (
	"context"
	"fmt"

	"github.com/aws/aws-database-encryption-sdk-dynamodb/awscryptographydbencryptionsdkdynamodbsmithygeneratedtypes"
	"github.com/aws/aws-database-encryption-sdk-dynamodb/awscryptographydbencryptionsdkdynamodbtransformssmithygenerated"
	"github.com/aws/aws-database-encryption-sdk-dynamodb/awscryptographydbencryptionsdkdynamodbtransformssmithygeneratedtypes"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/smithy-go/middleware"
)

type DBEsdkMiddleware struct {
	client *awscryptographydbencryptionsdkdynamodbtransformssmithygenerated.Client
}

func NewDBEsdkMiddleware(config awscryptographydbencryptionsdkdynamodbsmithygeneratedtypes.DynamoDbTablesEncryptionConfig) (*DBEsdkMiddleware, error) {
	client, err := awscryptographydbencryptionsdkdynamodbtransformssmithygenerated.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &DBEsdkMiddleware{
		client: client,
	}, nil
}

func (m DBEsdkMiddleware) CreateMiddleware() func(options *dynamodb.Options) {
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

func (m DBEsdkMiddleware) createRequestInterceptor() middleware.InitializeMiddleware {
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
func (m DBEsdkMiddleware) handleRequestInterception(ctx context.Context, params interface{}) context.Context {
	switch v := params.(type) {
	case *dynamodb.PutItemInput:
		ctx = middleware.WithStackValue(ctx, "originalInput", *deepCopyPutItemInput(v))
		transformedRequest, err := m.client.PutItemInputTransform(context.TODO(), awscryptographydbencryptionsdkdynamodbtransformssmithygeneratedtypes.PutItemInputTransformInput{
			SdkInput: *v,
		})
		if err != nil {
			fmt.Println(err)
		}
		*v = transformedRequest.TransformedInput
	case *dynamodb.GetItemInput:
		ctx = middleware.WithStackValue(ctx, "originalInput", *deepCopyGetItemInput(v))
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
func (m DBEsdkMiddleware) createResponseInterceptor() middleware.FinalizeMiddleware {
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
func (m DBEsdkMiddleware) handleResponseInterception(ctx context.Context, response interface{}) {
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
