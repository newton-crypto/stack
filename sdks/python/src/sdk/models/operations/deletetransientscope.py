"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
import requests as requests_http
from typing import Optional


@dataclasses.dataclass
class DeleteTransientScopeRequest:
    
    scope_id: str = dataclasses.field(metadata={'path_param': { 'field_name': 'scopeId', 'style': 'simple', 'explode': False }})
    r"""Scope ID"""
    transient_scope_id: str = dataclasses.field(metadata={'path_param': { 'field_name': 'transientScopeId', 'style': 'simple', 'explode': False }})
    r"""Transient scope ID"""
    

@dataclasses.dataclass
class DeleteTransientScopeResponse:
    
    content_type: str = dataclasses.field()
    status_code: int = dataclasses.field()
    raw_response: Optional[requests_http.Response] = dataclasses.field(default=None)
    