package client

import "github.com/icehubin/futu-go/adapt"

func (c *Client) bodyEncrypt(protoID uint32, body []byte) []byte {
	if !c.encrypt {
		return body
	}
	if protoID == adapt.ProtoID_InitConnect {
		return c.rsaEncrypt(body)
	} else {
		return c.aesEncrypt(body)
	}
}

func (c *Client) bodyDecrypt(protoID uint32, body []byte) []byte {
	if !c.encrypt {
		return body
	}
	if protoID == adapt.ProtoID_InitConnect {
		return c.rsaDecrypt(body)
	} else {
		return c.aesDecrypt(body)
	}
}

//---------- Todo encode and decode

func (c *Client) rsaEncrypt(body []byte) []byte {
	return body
}

func (c *Client) rsaDecrypt(body []byte) []byte {
	return body
}

func (c *Client) aesEncrypt(body []byte) []byte {
	return body
}

func (c *Client) aesDecrypt(body []byte) []byte {
	return body
}
