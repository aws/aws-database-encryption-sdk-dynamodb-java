// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
/// Wraps up an arbitrary Rust Error value as a Dafny Error
pub fn to_opaque_error(value: String) ->
    ::std::rc::Rc<crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error>
{
    let error_msg = value.clone();
    let error_msg = ::dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&error_msg);
    let error_obj: ::dafny_runtime::Object<dyn::std::any::Any> = ::dafny_runtime::Object(Some(
        ::std::rc::Rc::new(::std::cell::UnsafeCell::new(value)),
    ));
    ::std::rc::Rc::new(
        crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::OpaqueWithText {
            obj: error_obj,
	    objMessage: error_msg
        },
    )
}

/// Wraps up an arbitrary Rust Error value as a Dafny Result<T, Error>.Failure
pub fn to_opaque_error_result<T: ::dafny_runtime::DafnyType>(value: String) ->
    ::std::rc::Rc<
        crate::_Wrappers_Compile::Result<
            T,
            ::std::rc::Rc<crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error>
        >
    >
{
    ::std::rc::Rc::new(crate::_Wrappers_Compile::Result::Failure {
        error: to_opaque_error(value),
    })
}
pub fn to_dafny(
    value: crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error,
) -> ::std::rc::Rc<crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error> {
    ::std::rc::Rc::new(match value {
        crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::DynamoDbEncryptionException { message } =>
    crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::DynamoDbEncryptionException {
        message: ::dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&message),
    },
crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::AwsCryptographicPrimitivesError { error } =>
    crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::AwsCryptographyPrimitives {
        AwsCryptographyPrimitives: crate::deps::aws_cryptography_primitives::conversions::error::to_dafny(error),
    },
crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::DynamoDB_20120810Error { error } =>
    crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::ComAmazonawsDynamodb {
        ComAmazonawsDynamodb: crate::deps::com_amazonaws_dynamodb::conversions::error::to_dafny(error),
    },
crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::AwsCryptographicMaterialProvidersError { error } =>
    crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::AwsCryptographyMaterialProviders {
        AwsCryptographyMaterialProviders: crate::deps::aws_cryptography_materialProviders::conversions::error::to_dafny(error),
    },
crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::StructuredEncryptionError { error } =>
    crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::AwsCryptographyDbEncryptionSdkStructuredEncryption {
        AwsCryptographyDbEncryptionSdkStructuredEncryption: crate::deps::aws_cryptography_dbEncryptionSdk_structuredEncryption::conversions::error::to_dafny(error),
    },
crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::KeyStoreError { error } =>
    crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::AwsCryptographyKeyStore {
        AwsCryptographyKeyStore: crate::deps::aws_cryptography_keyStore::conversions::error::to_dafny(error),
    },
        crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::CollectionOfErrors { list, message } =>
            crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::CollectionOfErrors {
                message: ::dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&message),
                list: ::dafny_runtime::dafny_runtime_conversions::vec_to_dafny_sequence(&list, |e| to_dafny(e.clone()))
            },
        crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::ValidationError(inner) =>
            crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::Opaque {
                obj: {
                    let rc = ::std::rc::Rc::new(inner) as ::std::rc::Rc<dyn ::std::any::Any>;
                    // safety: `rc` is new, ensuring it has refcount 1 and is uniquely owned.
                    // we should use `dafny_runtime_conversions::rc_struct_to_dafny_class` once it
                    // accepts unsized types (https://github.com/dafny-lang/dafny/pull/5769)
                    unsafe { ::dafny_runtime::Object::from_rc(rc) }
                },
            },
            crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::Opaque { obj } =>
            crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::Opaque {
                obj: ::dafny_runtime::Object(obj.0)
            },
            crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::OpaqueWithText { obj, objMessage } =>
            crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::OpaqueWithText {
                obj: ::dafny_runtime::Object(obj.0),
                objMessage: ::dafny_runtime::dafny_runtime_conversions::unicode_chars_false::string_to_dafny_string(&objMessage),
            },
    })
}

#[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error,
    >,
) -> crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error {
    match ::std::borrow::Borrow::borrow(&dafny_value) {
        crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::DynamoDbEncryptionException { message } =>
    crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::DynamoDbEncryptionException {
        message: ::dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(&message),
    },
crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::AwsCryptographyPrimitives { AwsCryptographyPrimitives } =>
    crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::AwsCryptographicPrimitivesError {
        error: crate::deps::aws_cryptography_primitives::conversions::error::from_dafny(AwsCryptographyPrimitives.clone()),
    },
crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::ComAmazonawsDynamodb { ComAmazonawsDynamodb } =>
    crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::DynamoDB_20120810Error {
        error: crate::deps::com_amazonaws_dynamodb::conversions::error::from_dafny(ComAmazonawsDynamodb.clone()),
    },
crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::AwsCryptographyMaterialProviders { AwsCryptographyMaterialProviders } =>
    crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::AwsCryptographicMaterialProvidersError {
        error: crate::deps::aws_cryptography_materialProviders::conversions::error::from_dafny(AwsCryptographyMaterialProviders.clone()),
    },
crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::AwsCryptographyDbEncryptionSdkStructuredEncryption { AwsCryptographyDbEncryptionSdkStructuredEncryption } =>
    crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::StructuredEncryptionError {
        error: crate::deps::aws_cryptography_dbEncryptionSdk_structuredEncryption::conversions::error::from_dafny(AwsCryptographyDbEncryptionSdkStructuredEncryption.clone()),
    },
crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::AwsCryptographyKeyStore { AwsCryptographyKeyStore } =>
    crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::KeyStoreError {
        error: crate::deps::aws_cryptography_keyStore::conversions::error::from_dafny(AwsCryptographyKeyStore.clone()),
    },
        crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::CollectionOfErrors { list, message } =>
            crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::CollectionOfErrors {
                message: ::dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(&message),
                list: ::dafny_runtime::dafny_runtime_conversions::dafny_sequence_to_vec(&list, |e| from_dafny(e.clone()))
            },
        crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::Opaque { obj } =>
            crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::Opaque {
                obj: obj.clone()
            },
            crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::Opaque { obj } =>
            {
                use ::std::any::Any;
                if ::dafny_runtime::is_object!(obj, crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::ValidationError) {
                    let typed = ::dafny_runtime::cast_object!(obj.clone(), crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::ValidationError);
                    crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::ValidationError(
                        // safety: dafny_class_to_struct will increment ValidationError's Rc
                        unsafe {
                            ::dafny_runtime::dafny_runtime_conversions::object::dafny_class_to_struct(typed)
                        }
                    )
                } else {
                    crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::Opaque {
                        obj: obj.clone()
                    }
                }
            },
            crate::r#software::amazon::cryptography::dbencryptionsdk::dynamodb::internaldafny::types::Error::OpaqueWithText { obj, objMessage } =>
            {
                use ::std::any::Any;
                if ::dafny_runtime::is_object!(obj, crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::ValidationError) {
                    let typed = ::dafny_runtime::cast_object!(obj.clone(), crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::ValidationError);
                    crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::ValidationError(
                        // safety: dafny_class_to_struct will increment ValidationError's Rc
                        unsafe {
                            ::dafny_runtime::dafny_runtime_conversions::object::dafny_class_to_struct(typed)
                        }
                    )
                } else {
                    crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::error::Error::OpaqueWithText {
                        obj: obj.clone(),
                        objMessage: ::dafny_runtime::dafny_runtime_conversions::unicode_chars_false::dafny_string_to_string(&objMessage),
                    }
                }
            },
    }
}