<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


class ConfigUser
{
	#[\JMS\Serializer\Annotation\SerializedName('endpoint')]
    #[\JMS\Serializer\Annotation\Type('string')]
    public string $endpoint;
    
    /**
     * $eventTypes
     * 
     * @var array<string> $eventTypes
     */
	#[\JMS\Serializer\Annotation\SerializedName('eventTypes')]
    #[\JMS\Serializer\Annotation\Type('array<string>')]
    public array $eventTypes;
    
	#[\JMS\Serializer\Annotation\SerializedName('secret')]
    #[\JMS\Serializer\Annotation\Type('string')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?string $secret = null;
    
	public function __construct()
	{
		$this->endpoint = "";
		$this->eventTypes = [];
		$this->secret = null;
	}
}
