/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { StageSendDestinationAccount } from "./stagesenddestinationaccount";
import { StageSendDestinationPayment } from "./stagesenddestinationpayment";
import { StageSendDestinationWallet } from "./stagesenddestinationwallet";
import { Expose, Type } from "class-transformer";

export class StageSendDestination extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "account" })
  @Type(() => StageSendDestinationAccount)
  account?: StageSendDestinationAccount;

  @SpeakeasyMetadata()
  @Expose({ name: "payment" })
  @Type(() => StageSendDestinationPayment)
  payment?: StageSendDestinationPayment;

  @SpeakeasyMetadata()
  @Expose({ name: "wallet" })
  @Type(() => StageSendDestinationWallet)
  wallet?: StageSendDestinationWallet;
}
