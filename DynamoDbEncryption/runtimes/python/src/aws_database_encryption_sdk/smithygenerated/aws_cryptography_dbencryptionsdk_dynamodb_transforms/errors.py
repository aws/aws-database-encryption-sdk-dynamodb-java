# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0
# Do not modify this file. This file is machine generated, and any changes to it will be overwritten.

import _dafny
from aws_cryptography_internal_dynamodb.smithygenerated.com_amazonaws_dynamodb.shim import (
    _sdk_error_to_dafny_error as com_amazonaws_dynamodb_sdk_error_to_dafny_error,
)
import aws_database_encryption_sdk.internaldafny.generated
import aws_database_encryption_sdk.internaldafny.generated.AwsCryptographyDbEncryptionSdkDynamoDbTransformsTypes
from aws_database_encryption_sdk.smithygenerated.aws_cryptography_dbencryptionsdk_dynamodb.errors import (
    _smithy_error_to_dafny_error as aws_cryptography_dbencryptionsdk_dynamodb_smithy_error_to_dafny_error,
)
from aws_database_encryption_sdk.smithygenerated.aws_cryptography_dbencryptionsdk_dynamodb_itemencryptor.errors import (
    _smithy_error_to_dafny_error as aws_cryptography_dbencryptionsdk_dynamodb_itemencryptor_smithy_error_to_dafny_error,
)
import aws_database_encryption_sdk.smithygenerated.aws_cryptography_dbencryptionsdk_dynamodb_transforms.errors
from aws_database_encryption_sdk.smithygenerated.aws_cryptography_dbencryptionsdk_structuredencryption.errors import (
    _smithy_error_to_dafny_error as aws_cryptography_dbencryptionsdk_structuredencryption_smithy_error_to_dafny_error,
)
from typing import Any, Dict, Generic, List, Literal, TypeVar


class ServiceError(Exception):
    """Base error for all errors in the service."""

    pass


T = TypeVar("T")


class ApiError(ServiceError, Generic[T]):
    """Base error for all api errors in the service."""

    code: T

    def __init__(self, message: str):
        super().__init__(message)
        self.message = message


class UnknownApiError(ApiError[Literal["Unknown"]]):
    """Error representing any unknown api errors."""

    code: Literal["Unknown"] = "Unknown"


class DynamoDbEncryptionTransformsException(
    ApiError[Literal["DynamoDbEncryptionTransformsException"]]
):
    code: Literal["DynamoDbEncryptionTransformsException"] = (
        "DynamoDbEncryptionTransformsException"
    )
    message: str

    def __init__(
        self,
        *,
        message: str,
    ):
        super().__init__(message)

    def as_dict(self) -> Dict[str, Any]:
        """Converts the DynamoDbEncryptionTransformsException to a
        dictionary."""
        return {
            "message": self.message,
            "code": self.code,
        }

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> "DynamoDbEncryptionTransformsException":
        """Creates a DynamoDbEncryptionTransformsException from a
        dictionary."""
        kwargs: Dict[str, Any] = {
            "message": d["message"],
        }

        return DynamoDbEncryptionTransformsException(**kwargs)

    def __repr__(self) -> str:
        result = "DynamoDbEncryptionTransformsException("
        if self.message is not None:
            result += f"message={repr(self.message)}"

        return result + ")"

    def __eq__(self, other: Any) -> bool:
        if not isinstance(other, DynamoDbEncryptionTransformsException):
            return False
        attributes: list[str] = [
            "message",
            "message",
        ]
        return all(getattr(self, a) == getattr(other, a) for a in attributes)


class DynamoDbEncryptionTransformsException(
    ApiError[Literal["DynamoDbEncryptionTransformsException"]]
):
    code: Literal["DynamoDbEncryptionTransformsException"] = (
        "DynamoDbEncryptionTransformsException"
    )
    message: str


class ComAmazonawsDynamodb(ApiError[Literal["ComAmazonawsDynamodb"]]):
    ComAmazonawsDynamodb: Any


class DynamoDbEncryption(ApiError[Literal["DynamoDbEncryption"]]):
    DynamoDbEncryption: Any


class DynamoDbItemEncryptor(ApiError[Literal["DynamoDbItemEncryptor"]]):
    DynamoDbItemEncryptor: Any


class StructuredEncryption(ApiError[Literal["StructuredEncryption"]]):
    StructuredEncryption: Any


class CollectionOfErrors(ApiError[Literal["CollectionOfErrors"]]):
    code: Literal["CollectionOfErrors"] = "CollectionOfErrors"
    message: str
    list: List[ServiceError]

    def __init__(self, *, message: str, list):
        super().__init__(message)
        self.list = list

    def as_dict(self) -> Dict[str, Any]:
        """Converts the CollectionOfErrors to a dictionary.

        The dictionary uses the modeled shape names rather than the
        parameter names as keys to be mostly compatible with boto3.
        """
        return {
            "message": self.message,
            "code": self.code,
            "list": self.list,
        }

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> "CollectionOfErrors":
        """Creates a CollectionOfErrors from a dictionary.

        The dictionary is expected to use the modeled shape names rather
        than the parameter names as keys to be mostly compatible with
        boto3.
        """
        kwargs: Dict[str, Any] = {"message": d["message"], "list": d["list"]}

        return CollectionOfErrors(**kwargs)

    def __repr__(self) -> str:
        result = "CollectionOfErrors("
        result += f"message={self.message},"
        if self.message is not None:
            result += f"message={repr(self.message)}"
        result += f"list={self.list}"
        result += ")"
        return result

    def __eq__(self, other: Any) -> bool:
        if not isinstance(other, CollectionOfErrors):
            return False
        if not (self.list == other.list):
            return False
        attributes: list[str] = ["message", "message"]
        return all(getattr(self, a) == getattr(other, a) for a in attributes)


class OpaqueError(ApiError[Literal["OpaqueError"]]):
    code: Literal["OpaqueError"] = "OpaqueError"
    obj: Any  # As an OpaqueError, type of obj is unknown

    def __init__(self, *, obj):
        super().__init__("")
        self.obj = obj

    def as_dict(self) -> Dict[str, Any]:
        """Converts the OpaqueError to a dictionary.

        The dictionary uses the modeled shape names rather than the
        parameter names as keys to be mostly compatible with boto3.
        """
        return {
            "message": self.message,
            "code": self.code,
            "obj": self.obj,
        }

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> "OpaqueError":
        """Creates a OpaqueError from a dictionary.

        The dictionary is expected to use the modeled shape names rather
        than the parameter names as keys to be mostly compatible with
        boto3.
        """
        kwargs: Dict[str, Any] = {"message": d["message"], "obj": d["obj"]}

        return OpaqueError(**kwargs)

    def __repr__(self) -> str:
        result = "OpaqueError("
        result += f"message={self.message},"
        if self.message is not None:
            result += f"message={repr(self.message)}"
        result += f"obj={self.obj}"
        result += ")"
        return result

    def __eq__(self, other: Any) -> bool:
        if not isinstance(other, OpaqueError):
            return False
        if not (self.obj == other.obj):
            return False
        attributes: list[str] = ["message", "message"]
        return all(getattr(self, a) == getattr(other, a) for a in attributes)


class OpaqueWithTextError(ApiError[Literal["OpaqueWithTextError"]]):
    code: Literal["OpaqueWithTextError"] = "OpaqueWithTextError"
    obj: Any  # As an OpaqueWithTextError, type of obj is unknown
    obj_message: str  # obj_message is a message representing the details of obj

    def __init__(self, *, obj, obj_message):
        super().__init__("")
        self.obj = obj
        self.obj_message = obj_message

    def as_dict(self) -> Dict[str, Any]:
        """Converts the OpaqueWithTextError to a dictionary.

        The dictionary uses the modeled shape names rather than the
        parameter names as keys to be mostly compatible with boto3.
        """
        return {
            "message": self.message,
            "code": self.code,
            "obj": self.obj,
            "obj_message": self.obj_message,
        }

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> "OpaqueWithTextError":
        """Creates a OpaqueWithTextError from a dictionary.

        The dictionary is expected to use the modeled shape names rather
        than the parameter names as keys to be mostly compatible with
        boto3.
        """
        kwargs: Dict[str, Any] = {
            "message": d["message"],
            "obj": d["obj"],
            "obj_message": d["obj_message"],
        }

        return OpaqueWithTextError(**kwargs)

    def __repr__(self) -> str:
        result = "OpaqueWithTextError("
        result += f"message={self.message},"
        if self.message is not None:
            result += f"message={repr(self.message)}"
        result += f"obj={self.obj}"
        result += f"obj_message={self.obj_message}"
        result += ")"
        return result

    def __eq__(self, other: Any) -> bool:
        if not isinstance(other, OpaqueWithTextError):
            return False
        if not (self.obj == other.obj):
            return False
        attributes: list[str] = ["message", "message"]
        return all(getattr(self, a) == getattr(other, a) for a in attributes)


def _smithy_error_to_dafny_error(e: ServiceError):
    """Converts the provided native Smithy-modeled error into the corresponding
    Dafny error."""
    if isinstance(
        e,
        aws_database_encryption_sdk.smithygenerated.aws_cryptography_dbencryptionsdk_dynamodb_transforms.errors.DynamoDbEncryptionTransformsException,
    ):
        return aws_database_encryption_sdk.internaldafny.generated.AwsCryptographyDbEncryptionSdkDynamoDbTransformsTypes.Error_DynamoDbEncryptionTransformsException(
            message=_dafny.Seq(e.message)
        )

    if isinstance(e, ComAmazonawsDynamodb):
        return aws_database_encryption_sdk.internaldafny.generated.AwsCryptographyDbEncryptionSdkDynamoDbTransformsTypes.Error_ComAmazonawsDynamodb(
            com_amazonaws_dynamodb_sdk_error_to_dafny_error(e.message)
        )

    if isinstance(e, DynamoDbEncryption):
        return aws_database_encryption_sdk.internaldafny.generated.AwsCryptographyDbEncryptionSdkDynamoDbTransformsTypes.Error_AwsCryptographyDbEncryptionSdkDynamoDb(
            aws_cryptography_dbencryptionsdk_dynamodb_smithy_error_to_dafny_error(
                e.message
            )
        )

    if isinstance(e, DynamoDbItemEncryptor):
        return aws_database_encryption_sdk.internaldafny.generated.AwsCryptographyDbEncryptionSdkDynamoDbTransformsTypes.Error_AwsCryptographyDbEncryptionSdkDynamoDbItemEncryptor(
            aws_cryptography_dbencryptionsdk_dynamodb_itemencryptor_smithy_error_to_dafny_error(
                e.message
            )
        )

    if isinstance(e, StructuredEncryption):
        return aws_database_encryption_sdk.internaldafny.generated.AwsCryptographyDbEncryptionSdkDynamoDbTransformsTypes.Error_AwsCryptographyDbEncryptionSdkStructuredEncryption(
            aws_cryptography_dbencryptionsdk_structuredencryption_smithy_error_to_dafny_error(
                e.message
            )
        )

    if isinstance(e, CollectionOfErrors):
        return aws_database_encryption_sdk.internaldafny.generated.AwsCryptographyDbEncryptionSdkDynamoDbTransformsTypes.Error_CollectionOfErrors(
            message=_dafny.Seq(e.message),
            list=_dafny.Seq(
                _smithy_error_to_dafny_error(native_err) for native_err in e.list
            ),
        )

    if isinstance(e, OpaqueError):
        return aws_database_encryption_sdk.internaldafny.generated.AwsCryptographyDbEncryptionSdkDynamoDbTransformsTypes.Error_Opaque(
            obj=e.obj
        )

    if isinstance(e, OpaqueWithTextError):
        return aws_database_encryption_sdk.internaldafny.generated.AwsCryptographyDbEncryptionSdkDynamoDbTransformsTypes.Error_OpaqueWithText(
            obj=e.obj, objMessage=e.obj_message
        )

    else:
        return aws_database_encryption_sdk.internaldafny.generated.AwsCryptographyDbEncryptionSdkDynamoDbTransformsTypes.Error_Opaque(
            obj=e
        )