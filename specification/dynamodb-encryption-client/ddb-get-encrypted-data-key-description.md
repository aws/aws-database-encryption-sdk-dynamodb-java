[//]: # "Copyright Amazon.com Inc. or its affiliates. All Rights Reserved."
[//]: # "SPDX-License-Identifier: CC-BY-SA-4.0"

# DynamoDb Get Encrypted Data Key Description

## Overview

Currently, Keyring's Encrypted Data Key serialization includes metadata like key provider id, key provider info, branch Key id (only Hierarchy Keyring), and the version of the branch Key (only Hierarchy Keyring). Encrypted Data Key Description operation provides a way for the customers to get these meta data. This operation will provide them with an insight of the keys the data is encrypted with. 

## Operation

### Input

This operation MUST take in either of the following:
- A binary [header](https://github.com/aws/aws-database-encryption-sdk-dynamodb/blob/main/specification/structured-encryption/header.md)
- A [encrypted DynamoDB item](https://github.com/aws/aws-database-encryption-sdk-dynamodb/blob/ff9f08a355a20c81540e4ca652e09aaeffe90c4b/specification/dynamodb-encryption-client/encrypt-item.md#encrypted-dynamodb-item)

### Output

This operation MUST return a list of following:

- [keyProviderId](https://github.com/aws/aws-database-encryption-sdk-dynamodb/blob/main/specification/structured-encryption/header.md#key-provider-id)
- [keyProviderInfo](https://github.com/aws/aws-database-encryption-sdk-dynamodb/blob/main/specification/structured-encryption/header.md#key-provider-information) (only for AWS Cryptographic Materials Provider Keyring)
- [branchKeyId](https://github.com/aws/aws-database-encryption-sdk-dynamodb/blob/main/specification/structured-encryption/header.md#key-provider-information) (only for hierarchy keyring)
- [branchKeyVersion](https://github.com/aws/aws-database-encryption-sdk-dynamodb/blob/main/specification/structured-encryption/header.md#key-provider-information) (only for hierarchy keyring)

### Behavior

This operation should behave in following ways:

- The operation MUST NEVER DECRYPT the Data Keys.
- If the input is a encrypted DynamoDB item, it MUST attempt to extract "aws_dbe_head" attribute from the DynamoDB item to get binary header. 
- This operation MUST deserialize the header bytes according to the header format.
- This operation MUST extract the Format Flavor from the deserialize header. Format Flavor is used to identify the algorithm suite. 
- This operation MUST extract the dataKeys from the deserialize header. 
- For every Data Key in Data Keys, the operation MUST attempt to extract a description of the Data Key.
    - If the Data Key does not belong to AWS Cryptographic Materials Provider Keyring, the operation will only return keyProviderId. 
    - This description SHOULD be as detailed as possible, parsing but NOT decrypting the cipher-text field of the Data Key for embedded fields such as the Hierarchical Keyring’s branchKeyVersion.

