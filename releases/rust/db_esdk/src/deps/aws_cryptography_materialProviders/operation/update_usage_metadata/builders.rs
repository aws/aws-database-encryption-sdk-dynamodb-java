// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
pub use crate::deps::aws_cryptography_materialProviders::operation::update_usage_metadata::_unit::UnitBuilder;

pub use crate::deps::aws_cryptography_materialProviders::operation::update_usage_metadata::_update_usage_metadata_input::UpdateUsageMetadataInputBuilder;

impl UpdateUsageMetadataInputBuilder {
    /// Sends a request with this input using the given client.
    pub async fn send_with(
        self,
        cryptographic_materials_cache: &crate::deps::aws_cryptography_materialProviders::types::cryptographic_materials_cache::CryptographicMaterialsCacheRef,
    ) -> ::std::result::Result<
        (),
        crate::deps::aws_cryptography_materialProviders::types::error::Error,
    > {
        let mut fluent_builder = cryptographic_materials_cache.update_usage_metadata();
        fluent_builder.inner = self;
        fluent_builder.send().await
    }
}
/// Fluent builder constructing a request to `UpdateUsageMetadata`.
///
#[derive(::std::clone::Clone, ::std::fmt::Debug)]
pub struct UpdateUsageMetadataFluentBuilder {
    cryptographic_materials_cache: crate::deps::aws_cryptography_materialProviders::types::cryptographic_materials_cache::CryptographicMaterialsCacheRef,
    pub(crate) inner: crate::deps::aws_cryptography_materialProviders::operation::update_usage_metadata::builders::UpdateUsageMetadataInputBuilder,
}
impl UpdateUsageMetadataFluentBuilder {
    /// Creates a new `UpdateUsageMetadata`.
    pub(crate) fn new(cryptographic_materials_cache: crate::deps::aws_cryptography_materialProviders::types::cryptographic_materials_cache::CryptographicMaterialsCacheRef) -> Self {
        Self {
            cryptographic_materials_cache,
            inner: ::std::default::Default::default(),
        }
    }
    /// Access the UpdateUsageMetadata as a reference.
    pub fn as_input(&self) -> &crate::deps::aws_cryptography_materialProviders::operation::update_usage_metadata::builders::UpdateUsageMetadataInputBuilder {
        &self.inner
    }
    /// Sends the request and returns the response.
    pub async fn send(
        self,
    ) -> ::std::result::Result<
        (),
        crate::deps::aws_cryptography_materialProviders::types::error::Error,
    > {
        let input = self
            .inner
            .build()
            // Using Opaque since we don't have a validation-specific error yet.
            // Operations' models don't declare their own validation error,
            // and smithy-rs seems to not generate a ValidationError case unless there is.
            // Vanilla smithy-rs uses SdkError::construction_failure, but we aren't using SdkError.
            .map_err(|mut e| {
	     let msg = format!("{:?}", e);
             crate::deps::aws_cryptography_materialProviders::types::error::Error::OpaqueWithText {
                obj: ::dafny_runtime::Object::from_ref(&mut e as &mut dyn ::std::any::Any),
		objMessage: msg
             }})?;
        crate::deps::aws_cryptography_materialProviders::operation::update_usage_metadata::UpdateUsageMetadata::send(&self.cryptographic_materials_cache, input).await
    }

    #[allow(missing_docs)]
pub fn bytes_used(mut self, input: impl ::std::convert::Into<::std::primitive::i32>) -> Self {
    self.inner = self.inner.bytes_used(input.into());
    self
}
#[allow(missing_docs)]
pub fn set_bytes_used(mut self, input: ::std::option::Option<::std::primitive::i32>) -> Self {
    self.inner = self.inner.set_bytes_used(input);
    self
}
#[allow(missing_docs)]
pub fn get_bytes_used(&self) -> &::std::option::Option<::std::primitive::i32> {
    self.inner.get_bytes_used()
}
#[allow(missing_docs)]
pub fn identifier(mut self, input: impl ::std::convert::Into<::aws_smithy_types::Blob>) -> Self {
    self.inner = self.inner.identifier(input.into());
    self
}
#[allow(missing_docs)]
pub fn set_identifier(mut self, input: ::std::option::Option<::aws_smithy_types::Blob>) -> Self {
    self.inner = self.inner.set_identifier(input);
    self
}
#[allow(missing_docs)]
pub fn get_identifier(&self) -> &::std::option::Option<::aws_smithy_types::Blob> {
    self.inner.get_identifier()
}
}