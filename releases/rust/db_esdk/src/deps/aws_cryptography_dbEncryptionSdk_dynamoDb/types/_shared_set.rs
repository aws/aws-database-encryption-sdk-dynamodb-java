// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
#[allow(missing_docs)] // documentation missing in model
#[non_exhaustive]
#[derive(::std::clone::Clone, ::std::cmp::PartialEq, ::std::fmt::Debug)]
pub struct SharedSet {
    #[allow(missing_docs)] // documentation missing in model
    pub other: ::std::option::Option<::std::string::String>,
}
impl SharedSet {
    #[allow(missing_docs)] // documentation missing in model
    pub fn other(&self) -> &::std::option::Option<::std::string::String> {
        &self.other
    }
}
impl SharedSet {
    /// Creates a new builder-style object to manufacture [`SharedSet`](crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::SharedSet).
    pub fn builder(
    ) -> crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::builders::SharedSetBuilder
    {
        crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::builders::SharedSetBuilder::default()
    }
}

/// A builder for [`SharedSet`](crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::SharedSet).
#[non_exhaustive]
#[derive(
    ::std::clone::Clone, ::std::cmp::PartialEq, ::std::default::Default, ::std::fmt::Debug,
)]
pub struct SharedSetBuilder {
    pub(crate) other: ::std::option::Option<::std::string::String>,
}
impl SharedSetBuilder {
    #[allow(missing_docs)] // documentation missing in model
    pub fn other(mut self, input: impl ::std::convert::Into<::std::string::String>) -> Self {
        self.other = ::std::option::Option::Some(input.into());
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn set_other(mut self, input: ::std::option::Option<::std::string::String>) -> Self {
        self.other = input;
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn get_other(&self) -> &::std::option::Option<::std::string::String> {
        &self.other
    }
    /// Consumes the builder and constructs a [`SharedSet`](crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::SharedSet).
    pub fn build(
        self,
    ) -> ::std::result::Result<
        crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::SharedSet,
        ::aws_smithy_types::error::operation::BuildError,
    > {
        ::std::result::Result::Ok(
            crate::deps::aws_cryptography_dbEncryptionSdk_dynamoDb::types::SharedSet {
                other: self.other,
            },
        )
    }
}