// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
#[allow(missing_docs)] // documentation missing in model
#[non_exhaustive]
#[derive(::std::clone::Clone, ::std::cmp::PartialEq, ::std::fmt::Debug)]
pub struct UpdateItemOutputTransformOutput {
    #[allow(missing_docs)] // documentation missing in model
    pub transformed_output:
        ::std::option::Option<aws_sdk_dynamodb::operation::update_item::UpdateItemOutput>,
}
impl UpdateItemOutputTransformOutput {
    #[allow(missing_docs)] // documentation missing in model
    pub fn transformed_output(
        &self,
    ) -> &::std::option::Option<aws_sdk_dynamodb::operation::update_item::UpdateItemOutput> {
        &self.transformed_output
    }
}
impl UpdateItemOutputTransformOutput {
    /// Creates a new builder-style object to manufacture [`UpdateItemOutputTransformOutput`](crate::types::UpdateItemOutputTransformOutput).
    pub fn builder() -> crate::types::builders::UpdateItemOutputTransformOutputBuilder {
        crate::types::builders::UpdateItemOutputTransformOutputBuilder::default()
    }
}

/// A builder for [`UpdateItemOutputTransformOutput`](crate::types::UpdateItemOutputTransformOutput).
#[non_exhaustive]
#[derive(
    ::std::clone::Clone, ::std::cmp::PartialEq, ::std::default::Default, ::std::fmt::Debug,
)]
pub struct UpdateItemOutputTransformOutputBuilder {
    pub(crate) transformed_output:
        ::std::option::Option<aws_sdk_dynamodb::operation::update_item::UpdateItemOutput>,
}
impl UpdateItemOutputTransformOutputBuilder {
    #[allow(missing_docs)] // documentation missing in model
    pub fn transformed_output(
        mut self,
        input: impl ::std::convert::Into<aws_sdk_dynamodb::operation::update_item::UpdateItemOutput>,
    ) -> Self {
        self.transformed_output = ::std::option::Option::Some(input.into());
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn set_transformed_output(
        mut self,
        input: ::std::option::Option<aws_sdk_dynamodb::operation::update_item::UpdateItemOutput>,
    ) -> Self {
        self.transformed_output = input;
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn get_transformed_output(
        &self,
    ) -> &::std::option::Option<aws_sdk_dynamodb::operation::update_item::UpdateItemOutput> {
        &self.transformed_output
    }
    /// Consumes the builder and constructs a [`UpdateItemOutputTransformOutput`](crate::types::UpdateItemOutputTransformOutput).
    pub fn build(
        self,
    ) -> ::std::result::Result<
        crate::types::UpdateItemOutputTransformOutput,
        ::aws_smithy_types::error::operation::BuildError,
    > {
        ::std::result::Result::Ok(crate::types::UpdateItemOutputTransformOutput {
            transformed_output: self.transformed_output,
        })
    }
}