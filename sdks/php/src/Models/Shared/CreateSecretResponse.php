<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


/**
 * CreateSecretResponse - Created secret
 * 
 * @package formance\stack\Models\Shared
 * @access public
 */
class CreateSecretResponse
{
	#[\JMS\Serializer\Annotation\SerializedName('data')]
    #[\JMS\Serializer\Annotation\Type('formance\stack\Models\Shared\Secret')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?Secret $data = null;
    
	public function __construct()
	{
		$this->data = null;
	}
}