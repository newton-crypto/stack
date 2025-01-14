<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


enum PaymentType: string
{
    case PAY_IN = 'PAY-IN';
    case PAYOUT = 'PAYOUT';
    case TRANSFER = 'TRANSFER';
    case OTHER = 'OTHER';
}
