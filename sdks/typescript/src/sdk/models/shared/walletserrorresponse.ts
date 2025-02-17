/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Expose } from "class-transformer";

export enum WalletsErrorResponseErrorCode {
  Validation = "VALIDATION",
}

/**
 * Error
 */
export class WalletsErrorResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "errorCode" })
  errorCode: WalletsErrorResponseErrorCode;

  @SpeakeasyMetadata()
  @Expose({ name: "errorMessage" })
  errorMessage: string;
}
