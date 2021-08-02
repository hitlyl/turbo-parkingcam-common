package protocol

import (
	"encoding/binary"
	"errors"
	"strconv"
)

type PicData struct {
	Mac uint64
	Id uint8
	Name uint16
	LatestSerial uint16
	HOffset uint16
	VOffset uint16
	Width uint16
	Height uint16
	Len uint32
	Crc uint16
	ReceiveLen uint32
	Data []byte
	CreateNano int64
	Ip string

}

//02 02 0003 0000 0000 0780 0438 000e4d4c d49f
func(p *PicData) ParseInfo(buf []byte) error{
	if len(buf)<18{
		return errors.New("pic info len error")
	}
	p.Id = buf[1]
	p.Name = binary.BigEndian.Uint16(buf[2:4])
	p.HOffset = binary.BigEndian.Uint16(buf[4:6])
	p.VOffset = binary.BigEndian.Uint16(buf[6:8])
	p.Width = binary.BigEndian.Uint16(buf[8:10])
	p.Height = binary.BigEndian.Uint16(buf[10:12])
	p.Len = binary.BigEndian.Uint32(buf[12:16])
	p.Crc = binary.BigEndian.Uint16(buf[16:])
	p.LatestSerial = 0
	return nil
}

//02 02 0003 0000 ffd8f....
func(p *PicData) ParseData(buf []byte) error{
	id:=buf[1]
	if id!=p.Id{
		return errors.New("id error. id="+strconv.Itoa(int(id))+" p.id="+strconv.Itoa(int(p.Id)))
	}
	name:=binary.BigEndian.Uint16(buf[2:4])
	if name !=p.Name{
		return errors.New("name error.")
	}
  	serial:= binary.BigEndian.Uint16(buf[4:6])
  	if p.LatestSerial!=0 && serial!=p.LatestSerial+1{
  		return errors.New("serial error. latestSerial="+strconv.Itoa(int(p.LatestSerial))+" serial="+strconv.Itoa(int(serial)))
	}
	p.LatestSerial = serial
	p.ReceiveLen=p.ReceiveLen+uint32(len(buf))-6
	if p.ReceiveLen>p.Len{
		return errors.New("len big error.")
	}
	p.Data = append(p.Data,buf[6:]...)

	return nil
}

