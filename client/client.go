package client

import (
	"crypto/sha1"
	"errors"
	"io"
	"net"
	"time"

	"github.com/icehubin/futu-go/logger"
	"github.com/icehubin/futu-go/pb/common"
	"github.com/icehubin/futu-go/pb/initconnect"

	"github.com/icehubin/futu-go/adapt"
)

func New(host string) (*Client, error) {
	return Create(host, false, "", true)
}

func NewEncrypt(host string, rsa_file string) (*Client, error) {
	return Create(host, true, rsa_file, true)
}

func Create(host string, encrypt bool, rsa_file string, notify bool) (*Client, error) {
	client := &Client{
		host:     host,
		encrypt:  encrypt,
		notify:   notify,
		rsa_file: rsa_file,
	}
	err := client.Init()
	if err != nil {
		return nil, err
	}
	return client, err
}

func (c *Client) Init() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", c.host)
	if err != nil {
		logger.WithFields(logger.Fields{
			"host": c.host,
		}).Error("Host error")
		return errors.New("Host error")
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		logger.WithFields(logger.Fields{
			"host": c.host,
		}).Error("Link error")
		return errors.New("Link error")
	}
	c.conn = conn
	c.status = 1
	PacketEncAlgo := common.PacketEncAlgo_PacketEncAlgo_None
	if c.encrypt {
		PacketEncAlgo = common.PacketEncAlgo_PacketEncAlgo_FTAES_ECB
	}
	c.ProtoWtite(adapt.ProtoID_InitConnect,
		adapt.With("PacketEncAlgo", PacketEncAlgo),
		adapt.With("recvNotify", c.notify),
	)
	_, res := c.protoRead()
	if res.RetType != adapt.RetType_Succeed {
		return errors.New("Init Server Fail")
	}
	c.server = &adapt.Server{
		ConnID:            res.S2C.(*initconnect.S2C).GetConnID(),
		KeepAliveInterval: res.S2C.(*initconnect.S2C).GetKeepAliveInterval(),
		ConnAESKey:        res.S2C.(*initconnect.S2C).GetConnAESKey(),
		AesCBCiv:          res.S2C.(*initconnect.S2C).GetAesCBCiv(),
		LoginUserID:       res.S2C.(*initconnect.S2C).GetLoginUserID(),
		ServerVer:         res.S2C.(*initconnect.S2C).GetServerVer(),
		UserAttribution:   res.S2C.(*initconnect.S2C).GetUserAttribution(),

		Encrypt: c.encrypt,
	}
	return nil
}

func (c *Client) KeepAlive() {
	c.Async(adapt.ProtoID_KeepAlive)
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) Async(protoID uint32, dopts ...adapt.Option) error {
	_, err := c.ProtoWtite(protoID, dopts...)
	if err != nil {
		return err
	}
	return nil
}

/*
写一次
*/
func (c *Client) ProtoWtite(protoID uint32, dopts ...adapt.Option) (uint32, error) {
	if createProto, ok := adapt.DataAdaptMap[protoID]; ok {
		//need to rewrite fix PacketID!
		da := createProto.(func(dopts ...adapt.Option) adapt.AdaptInterface)(dopts...)
		serialNO := c.genSerialNO()
		if nil != c.server {
			da.SetPacketID(adapt.PacketID{
				ConnID:   c.server.ConnID,
				SerialNo: serialNO,
			})
		}

		if bodyByte, ok := da.PackBody(); ok {
			bodyByte = c.bodyEncrypt(protoID, bodyByte)
			headerByte := c.genHeader(protoID, serialNO, bodyByte)
			logger.WithFields(logger.Fields{
				"protoID":  protoID,
				"serialNO": serialNO,
				"request":  da.GetC2S(),
			}).Debug("ProtoWtite")
			_, err := c.conn.Write(append(headerByte, bodyByte...))
			return serialNO, err
		}
		logger.WithFields(logger.Fields{
			"protoID": protoID,
			"request": da.GetC2S(),
			"errMsg":  "PackBody error",
		}).Error("ProtoWtite")
		return 0, errors.New("PackBody error")
	}
	logger.WithFields(logger.Fields{
		"protoID": protoID,
		"errMsg":  "unKnow ProtoID",
	}).Error("ProtoWtite")
	return 0, errors.New("unKnow ProtoID")
}

/*
读一个返回包
*/
func (c *Client) protoRead() (*adapt.Header, *adapt.Response) {
	header, bodyByte := c.readAPack()
	if nil == header || nil == bodyByte {
		packErr := adapt.PackErr()
		return nil, &packErr
	}
	//todo error fix

	protoID := header.GetProtoID()
	if createProto, ok := adapt.DataAdaptMap[protoID]; ok {
		da := createProto.(func(dopts ...adapt.Option) adapt.AdaptInterface)()
		bodyByte = c.bodyDecrypt(protoID, bodyByte)
		response := da.UnPackBody(bodyByte)
		logger.WithFields(logger.Fields{
			"ProtoID": protoID,
			"SeriaNo": header.GetSeriaNo(),
			"RetType": response.RetType,
			"RetMsg":  response.RetMsg,
			"ErrCode": response.ErrCode,
			"S2C":     response.S2C,
		}).Debug("protoRead")
		return header, &response
	}
	logger.WithFields(logger.Fields{
		"ProtoID": protoID,
		"SeriaNo": header.GetSeriaNo(),
	}).Warn("protoRead")
	packDefault := adapt.DefaultErr()
	return header, &packDefault
}

func (c *Client) Read() *ResPack {
	header, res := c.protoRead()
	return &ResPack{
		Header:   header,
		Response: res,
	}
}

func (c *Client) Sync(protoID uint32, dopts ...adapt.Option) adapt.Response {
	_, err := c.ProtoWtite(protoID, dopts...)
	if err != nil {
		return adapt.DefaultErr()
	}
	_, res := c.protoRead()
	return *res
}

//------------

func (c *Client) readAPack() (*adapt.Header, []byte) {
	h := make([]byte, adapt.HEADER_LENGTH)
	lenth, err := c.readBytes(h)
	if err == nil && lenth == adapt.HEADER_LENGTH {
		header := &adapt.Header{}
		header.UnPack(h)
		bodyLen := header.GetBodyLen()
		bodyByte := make([]byte, 0)

		for bLen := 0; bLen < int(bodyLen); {
			b := make([]byte, int(bodyLen)-bLen)
			lenth, err = c.readBytes(b)
			bodyByte = append(bodyByte, b...)
			bLen += lenth
		}
		return header, bodyByte
	}
	return nil, nil
}

func (c *Client) readBytes(b []byte) (int, error) {
	lenth, err := c.conn.Read(b)

	if err != nil && err == io.EOF {
		logger.WithFields(logger.Fields{
			"host":   c.host,
			"server": c.server,
		}).Fatal("Connection lost")
		c.status = 110
		panic("Please Reconnect")
	}
	return lenth, err
}

func (c *Client) genHeader(protoID uint32, serialNO uint32, body []byte) []byte {
	header := adapt.Header{}
	header.SetProtoID(protoID)
	header.SetProtoFmtType(uint8(0))
	header.SetProtoVer(0)
	header.SetSeriaNo(serialNO)
	header.SetBodyLen(uint32(len(body)))

	hs := sha1.New()
	io.WriteString(hs, string(body))
	s := hs.Sum(nil)
	header.SetBodySHA1(s)
	hp := header.Pack()
	return hp
}

func (c *Client) genSerialNO() uint32 {
	//todo lock
	if c.serialNO <= 10000 {
		c.serialNO = 10000
	}
	c.serialNO++
	if c.serialNO >= 4294967295 {
		c.serialNO = uint32(time.Now().Unix()%10000 + 10000)
	}
	return c.serialNO
}

func (c *Client) GetEncrypt() bool {
	return c.encrypt
}

func (c *Client) OK() bool {
	return c.status == 1
}

type ResPack struct {
	Header   *adapt.Header
	Response *adapt.Response
}

type Client struct {
	host     string
	encrypt  bool
	notify   bool
	rsa_file string
	conn     *net.TCPConn
	server   *adapt.Server
	status   uint8
	serialNO uint32
}
