package tcpserver

import (
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"github.com/sirupsen/logrus"
	"net"
	"sync"
	"time"
	"github.com/hitlyl/turbo-parkingcam-common/protocol"
)



type PicProcess interface {
	AddDevice(deviceIp string)
	DelDevice(deviceIp string)
	AddMessage(deviceIp string, msg *protocol.Message)
}
type CamTcpServer struct {
	*gnet.EventServer
	picProcessor PicProcess
	codec gnet.ICodec
	workerPool *goroutine.Pool
	addr       string
	multicore  bool
	connectedSockets sync.Map
	logger *logrus.Entry
}

func (cs *CamTcpServer) Dispose(){
	cs.connectedSockets.Range(func(k,v interface{})bool{  //多的删掉
		v.(gnet.Conn).Close()
		return true
	})
	gnet.Stop(nil, cs.addr)
}

func (cs *CamTcpServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	cs.logger.Debugf("CamTcpServer is listening on %s (multi-cores: %t, loops: %d)",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)

	return
}
func (cs *CamTcpServer) getIp(remoteAddr net.Addr) string{

	switch addr := remoteAddr.(type) {
	case *net.UDPAddr:
		return addr.IP.String()

	case *net.TCPAddr:
		return addr.IP.String()

	}
	return ""
}
func(cs *CamTcpServer)OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	cs.logger.Debugf("OnOpened %s",c.RemoteAddr())

	cs.connectedSockets.Store(cs.getIp(c.RemoteAddr()), c)

	cs.picProcessor.AddDevice(cs.getIp(c.RemoteAddr()))
	startMsg,err:=protocol.NewCamStartPicMsg()
	if err!=nil{
		action = gnet.Close
		cs.logger.Error(err)
	}
	c.AsyncWrite(startMsg)

	return
}
func(cs *CamTcpServer)OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	cs.logger.Debugf("OnClosed %s",c.RemoteAddr())
	cs.connectedSockets.Delete(cs.getIp(c.RemoteAddr()))
	cs.picProcessor.DelDevice(cs.getIp(c.RemoteAddr()))
	return
}

func (cs *CamTcpServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {

	msg,err:=protocol.MessageFromBytes(frame)
	if err!=nil{
		cs.logger.Error(err)
	}
	cs.picProcessor.AddMessage(cs.getIp(c.RemoteAddr()), msg)
	//cs.logger.Debugf("id=%d, data=%s",msg.InternalId, hex.EncodeToString(frame))
	return
}

func (cs *CamTcpServer) Tick() (delay time.Duration, action gnet.Action) {

	return
}

func NewCamTcpServer(addr string, picProcessor PicProcess, logger *logrus.Entry) *CamTcpServer {

	codec := &protocol.Message{}
	server:=&CamTcpServer{
		addr:addr, picProcessor:picProcessor, multicore: true, codec: codec,workerPool: goroutine.Default(), logger: logger,
	}
	err:=gnet.Serve(server,addr,gnet.WithMulticore(true), gnet.WithTCPKeepAlive(time.Minute*5),gnet.WithCodec(codec))
	if(err!=nil){
		panic(err)
	}

	return server
}