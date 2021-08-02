package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/panjf2000/gnet"
)

type Header struct {
	FixHead uint16
	Len uint16
	Version uint8
	DeviceType uint8
	DeviceId uint64
	ModuleType uint8
	ModuleId uint8
	ProtocolType uint8
	ProtocolSubType uint8
	MessageId uint64
}
type Message struct{
	InternalId int
	MsgHeader *Header
	MsgBody []byte
	Crc uint16
}
var MsgInternalId=0
type CamStartMsg struct {
	Mode     uint8 //0:  停止拍摄,1：单次拍摄,2：连续拍摄
	Cam      uint8 //1 A镜头拍摄,2 B
	PicName  uint16
	AHOffset uint16 //水平起始像素
	AVOffset uint16
	AWidth   uint16
	AHeight  uint16
	BHOffset uint16 //水平起始像素
	BVOffset uint16
	BWidth   uint16
	BHeight  uint16
}
func (c *CamStartMsg)ToBytes()([]byte, error){
	result := make([]byte, 0)
	buffer := bytes.NewBuffer(result)
	if err := binary.Write(buffer, binary.BigEndian, c.Mode); err != nil {
		s := fmt.Sprintf("pack Mode error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, c.Cam); err != nil {
		s := fmt.Sprintf("pack Cam error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, c.PicName); err != nil {
		s := fmt.Sprintf("pack PicName error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, c.AHOffset); err != nil {
		s := fmt.Sprintf("pack AHOffset error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, c.AVOffset); err != nil {
		s := fmt.Sprintf("pack AVOffset error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, c.AWidth); err != nil {
		s := fmt.Sprintf("pack AWidth error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, c.AHeight); err != nil {
		s := fmt.Sprintf("pack AHeight error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, c.BHOffset); err != nil {
		s := fmt.Sprintf("pack BHOffset error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, c.BVOffset); err != nil {
		s := fmt.Sprintf("pack BVOffset error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, c.BWidth); err != nil {
		s := fmt.Sprintf("pack BWidth error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, c.BHeight); err != nil {
		s := fmt.Sprintf("pack BHeight error , %v", err)
		return nil, errors.New(s)
	}
	return buffer.Bytes(), nil
}

var messageId uint64
var HeaderLen = 26

func HeaderFromBytes(buf []byte)(*Header, error){
	if len(buf)!=HeaderLen{
		return nil, errors.New("header from bytes len error.")
	}
	byteBuffer := bytes.NewBuffer(buf)
	var fixHead uint16
	var len uint16
	var version,deviceType uint8
	var deviceId uint64
	var moduleType,moduleId,protocolType,protocolSubType uint8
	var messageId uint64

	_ = binary.Read(byteBuffer, binary.BigEndian, &fixHead)
	_ = binary.Read(byteBuffer, binary.BigEndian, &len)
	_ = binary.Read(byteBuffer, binary.BigEndian, &version)
	_ = binary.Read(byteBuffer, binary.BigEndian, &deviceType)
	_ = binary.Read(byteBuffer, binary.BigEndian, &deviceId)
	_ = binary.Read(byteBuffer, binary.BigEndian, &moduleType)
	_ = binary.Read(byteBuffer, binary.BigEndian, &moduleId)
	_ = binary.Read(byteBuffer, binary.BigEndian, &protocolType)
	_ = binary.Read(byteBuffer, binary.BigEndian, &protocolSubType)
	_ = binary.Read(byteBuffer, binary.BigEndian, &messageId)


	return &Header{
		FixHead: fixHead,
		Len:len,
		Version: version,
		DeviceType: deviceType,
		DeviceId: deviceId,
		ModuleType: moduleType,
		ModuleId: moduleId,
		ProtocolType: protocolType,
		ProtocolSubType: protocolSubType,
		MessageId: messageId,
	},nil
}

func(h *Header)ToBytes() ([]byte,error){
	result := make([]byte, 0)
	buffer := bytes.NewBuffer(result)
	if err := binary.Write(buffer, binary.BigEndian, h.FixHead); err != nil {
		s := fmt.Sprintf("pack fixHead error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, h.Len); err != nil {
		s := fmt.Sprintf("pack Len error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, h.Version); err != nil {
		s := fmt.Sprintf("pack Version error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, h.DeviceType); err != nil {
		s := fmt.Sprintf("pack DeviceType error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, h.DeviceId); err != nil {
		s := fmt.Sprintf("pack DeviceId error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, h.ModuleType); err != nil {
		s := fmt.Sprintf("pack ModuleType error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, h.ModuleId); err != nil {
		s := fmt.Sprintf("pack Len ModuleId , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, h.ProtocolType); err != nil {
		s := fmt.Sprintf("pack ProtocolType error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, h.ProtocolSubType); err != nil {
		s := fmt.Sprintf("pack ProtocolSubType error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, h.MessageId); err != nil {
		s := fmt.Sprintf("pack MessageId error , %v", err)
		return nil, errors.New(s)
	}
	return buffer.Bytes(), nil

}
func(m *Message)ToBytes() ([]byte,error){
	headerBytes,err:=m.MsgHeader.ToBytes()
	if err!=nil{
		return nil, err
	}
	result := make([]byte, 0)
	buffer := bytes.NewBuffer(result)
	if err := binary.Write(buffer, binary.BigEndian, headerBytes); err != nil {
		s := fmt.Sprintf("mesage pack header error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, m.MsgBody); err != nil {
		s := fmt.Sprintf("mesage pack header error , %v", err)
		return nil, errors.New(s)
	}
	if err := binary.Write(buffer, binary.BigEndian, m.Crc); err != nil {
		s := fmt.Sprintf("mesage pack header error , %v", err)
		return nil, errors.New(s)
	}
	return buffer.Bytes(),nil
}

func(m *Message)Encode(c gnet.Conn, buf[] byte)([]byte,error){
	//item := c.Context().(Message)
	//return item.ToBytes()
	return buf, nil
}

func(m *Message)Decode(c gnet.Conn) ([]byte, error) {
	if size, headerBuf := c.ReadN(HeaderLen); size == HeaderLen {
		header,err:=HeaderFromBytes(headerBuf)
		if err!=nil{
			return nil, err
		}
		totalLen:=header.Len+2
		if dataSize, data := c.ReadN(int(totalLen)); dataSize == int(totalLen) {
			msg,err:=MessageFromBytes(data)
			if err!=nil{
				return nil, err
			}
			c.SetContext(msg)
			c.ShiftN(int(totalLen))
			return data, nil
		}
	}
	return nil, errors.New("not enough header data")
}

func MessageFromBytes(buf []byte)(*Message,error){
	len:=len(buf)
	header,err:=HeaderFromBytes(buf[:HeaderLen])
	if err!=nil{
		return nil, err
	}
	MsgInternalId=MsgInternalId+1
	return &Message{
		InternalId: MsgInternalId,
		MsgHeader: header,
		MsgBody: buf[HeaderLen:len-2],
		Crc: binary.BigEndian.Uint16(buf[len-2:len]),
	},nil
}

func NewCamStartPicMsg() ([]byte,error){
	messageId+=1
	header:=&Header{
		FixHead: 0xfefe,
		Version: 0x1,
		DeviceType: 0x05,
		DeviceId: 0x6666666666666666,
		ModuleType: 0x00,
		ModuleId:0x00,
		ProtocolType: 0x02,
		ProtocolSubType: 0x50,
		MessageId: 0x8888888888888888,
	}



	camStartMsg:=&CamStartMsg{
		Mode:     2,
		Cam:      1,
		PicName:  1,
		AHOffset: 0,
		AVOffset: 0,
		AWidth:   1920,
		AHeight:  1080,
		BHOffset: 0,
		BVOffset: 0,
		BWidth: 1920,
		BHeight: 1080,
	}
	bodyBytes,err:=camStartMsg.ToBytes()
	if err!=nil{
		return nil, err
	}
	header.Len = uint16(HeaderLen+len(bodyBytes))
	headerBytes,err:=header.ToBytes()
	if err!=nil{
		return nil, err
	}
	crcBytes:=append(headerBytes, bodyBytes...)

	message:=&Message{
		MsgHeader: header,
		MsgBody: bodyBytes,
		Crc: CheckSum(crcBytes),
	}
	msgBytes,err:=message.ToBytes()
	if err!=nil{
		return nil, err
	}
	return msgBytes, nil
}