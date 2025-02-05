package dbesdkmiddleware

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// deepCopyPutItemInput performs a deep copy of a PutItemInput struct.
func deepCopyPutItemInput(input *dynamodb.PutItemInput) *dynamodb.PutItemInput {
	if input == nil {
		return nil
	}
	copyItem := make(map[string]types.AttributeValue, len(input.Item))
	for k, v := range input.Item {
		copyItem[k] = deepCopyAttributeValue(v)
	}
	copyExpected := make(map[string]types.ExpectedAttributeValue, len(input.Expected))
	for k, v := range input.Expected {
		copyExpected[k] = v
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

// deepCopyGetItemInput performs a deep copy of a GetItemInput struct.
func deepCopyGetItemInput(input *dynamodb.GetItemInput) *dynamodb.GetItemInput {
	if input == nil {
		return nil
	}
	copyKey := make(map[string]types.AttributeValue, len(input.Key))
	for k, v := range input.Key {
		copyKey[k] = deepCopyAttributeValue(v)
	}
	copyExprNames := make(map[string]string, len(input.ExpressionAttributeNames))
	for k, v := range input.ExpressionAttributeNames {
		copyExprNames[k] = v
	}
	copyAttributesToGet := make([]string, len(input.AttributesToGet))
	copy(copyAttributesToGet, input.AttributesToGet)
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
		panic("Unknown AttributeValue type.")
	}
}
