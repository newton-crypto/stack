<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Operations;


class PaymentsgetServerInfoResponse
{
	
    public string $contentType;
    
    /**
     * Server information
     * 
     * @var ?\formance\stack\Models\Shared\ServerInfo $serverInfo
     */
	
    public ?\formance\stack\Models\Shared\ServerInfo $serverInfo = null;
    
	
    public int $statusCode;
    
	
    public ?\Psr\Http\Message\ResponseInterface $rawResponse = null;
    
	public function __construct()
	{
		$this->contentType = "";
		$this->serverInfo = null;
		$this->statusCode = 0;
		$this->rawResponse = null;
	}
}