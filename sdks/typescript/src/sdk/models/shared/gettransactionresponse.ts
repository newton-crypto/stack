/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { ExpandedTransaction } from "./expandedtransaction";
import { Expose, Type } from "class-transformer";

/**
 * OK
 */
export class GetTransactionResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "data" })
  @Type(() => ExpandedTransaction)
  data: ExpandedTransaction;
}
