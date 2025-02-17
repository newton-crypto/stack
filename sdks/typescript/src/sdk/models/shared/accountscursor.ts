/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { PaymentsAccount } from "./paymentsaccount";
import { Expose, Type } from "class-transformer";

export class AccountsCursorCursor extends SpeakeasyBase {
  @SpeakeasyMetadata({ elemType: PaymentsAccount })
  @Expose({ name: "data" })
  @Type(() => PaymentsAccount)
  data: PaymentsAccount[];

  @SpeakeasyMetadata()
  @Expose({ name: "hasMore" })
  hasMore: boolean;

  @SpeakeasyMetadata()
  @Expose({ name: "next" })
  next?: string;

  @SpeakeasyMetadata()
  @Expose({ name: "pageSize" })
  pageSize: number;

  @SpeakeasyMetadata()
  @Expose({ name: "previous" })
  previous?: string;
}

/**
 * OK
 */
export class AccountsCursor extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "cursor" })
  @Type(() => AccountsCursorCursor)
  cursor: AccountsCursorCursor;
}
