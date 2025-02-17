/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Expose, Transform } from "class-transformer";

export enum MigrationInfoState {
  ToDo = "to do",
  Done = "done",
}

export class MigrationInfo extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "date" })
  @Transform(({ value }) => new Date(value), { toClassOnly: true })
  date?: Date;

  @SpeakeasyMetadata()
  @Expose({ name: "name" })
  name?: string;

  @SpeakeasyMetadata()
  @Expose({ name: "state" })
  state?: MigrationInfoState;

  @SpeakeasyMetadata()
  @Expose({ name: "version" })
  version?: number;
}
