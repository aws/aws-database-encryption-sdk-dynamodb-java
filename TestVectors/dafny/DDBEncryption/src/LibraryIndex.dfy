// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

include "../Model/AwsCryptographyDynamoDbEncryptionTypesWrapped.dfy"

module WrappedDynamoDbEncryption refines WrappedAbstractAwsCryptographyDynamoDbEncryptionService
{

  import WrappedService = DynamoDbEncryption

  function method WrappedDefaultDynamoDbEncryptionConfig(): DynamoDbEncryptionConfig
  {
    DynamoDbEncryptionConfig
  }

}
