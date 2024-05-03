from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class InitRequest(_message.Message):
    __slots__ = ("policy", "jgf")
    POLICY_FIELD_NUMBER: _ClassVar[int]
    JGF_FIELD_NUMBER: _ClassVar[int]
    policy: str
    jgf: str
    def __init__(self, policy: _Optional[str] = ..., jgf: _Optional[str] = ...) -> None: ...

class InitResponse(_message.Message):
    __slots__ = ("status",)
    class ResultType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        INIT_SUCCESS: _ClassVar[InitResponse.ResultType]
        INIT_ERROR: _ClassVar[InitResponse.ResultType]
        INIT_DENIED: _ClassVar[InitResponse.ResultType]
    INIT_SUCCESS: InitResponse.ResultType
    INIT_ERROR: InitResponse.ResultType
    INIT_DENIED: InitResponse.ResultType
    STATUS_FIELD_NUMBER: _ClassVar[int]
    status: InitResponse.ResultType
    def __init__(self, status: _Optional[_Union[InitResponse.ResultType, str]] = ...) -> None: ...

class MatchRequest(_message.Message):
    __slots__ = ("jobspec", "request", "count")
    JOBSPEC_FIELD_NUMBER: _ClassVar[int]
    REQUEST_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    jobspec: str
    request: str
    count: int
    def __init__(self, jobspec: _Optional[str] = ..., request: _Optional[str] = ..., count: _Optional[int] = ...) -> None: ...

class MatchResponse(_message.Message):
    __slots__ = ("allocation", "jobid", "reserved", "at", "overhead", "status")
    class ResultType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        MATCH_SUCCESS: _ClassVar[MatchResponse.ResultType]
        MATCH_ERROR: _ClassVar[MatchResponse.ResultType]
        MATCH_DENIED: _ClassVar[MatchResponse.ResultType]
    MATCH_SUCCESS: MatchResponse.ResultType
    MATCH_ERROR: MatchResponse.ResultType
    MATCH_DENIED: MatchResponse.ResultType
    ALLOCATION_FIELD_NUMBER: _ClassVar[int]
    JOBID_FIELD_NUMBER: _ClassVar[int]
    RESERVED_FIELD_NUMBER: _ClassVar[int]
    AT_FIELD_NUMBER: _ClassVar[int]
    OVERHEAD_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    allocation: str
    jobid: int
    reserved: bool
    at: int
    overhead: float
    status: MatchResponse.ResultType
    def __init__(self, allocation: _Optional[str] = ..., jobid: _Optional[int] = ..., reserved: bool = ..., at: _Optional[int] = ..., overhead: _Optional[float] = ..., status: _Optional[_Union[MatchResponse.ResultType, str]] = ...) -> None: ...

class CancelRequest(_message.Message):
    __slots__ = ("jobID",)
    JOBID_FIELD_NUMBER: _ClassVar[int]
    jobID: int
    def __init__(self, jobID: _Optional[int] = ...) -> None: ...

class CancelResponse(_message.Message):
    __slots__ = ("jobID", "error", "status")
    class ResultType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        CANCEL_SUCCESS: _ClassVar[CancelResponse.ResultType]
        CANCEL_REQUEST_ERROR: _ClassVar[CancelResponse.ResultType]
        CANCEL_ERROR: _ClassVar[CancelResponse.ResultType]
        CANCEL_DENIED: _ClassVar[CancelResponse.ResultType]
    CANCEL_SUCCESS: CancelResponse.ResultType
    CANCEL_REQUEST_ERROR: CancelResponse.ResultType
    CANCEL_ERROR: CancelResponse.ResultType
    CANCEL_DENIED: CancelResponse.ResultType
    JOBID_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    jobID: int
    error: int
    status: CancelResponse.ResultType
    def __init__(self, jobID: _Optional[int] = ..., error: _Optional[int] = ..., status: _Optional[_Union[CancelResponse.ResultType, str]] = ...) -> None: ...
