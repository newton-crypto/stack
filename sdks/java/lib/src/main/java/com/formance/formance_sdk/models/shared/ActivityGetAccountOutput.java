/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

public class ActivityGetAccountOutput {
    @JsonProperty("data")
    public AccountWithVolumesAndBalances data;

    public ActivityGetAccountOutput withData(AccountWithVolumesAndBalances data) {
        this.data = data;
        return this;
    }
    
    public ActivityGetAccountOutput(@JsonProperty("data") AccountWithVolumesAndBalances data) {
        this.data = data;
  }
}