package terminals

import (
	"log"
	"net"

	"github.com/enderian/directrd/types"
)

func Event(event *types.Event, source net.Addr) {
	if event.Terminal == "" {
		return
	}

	var terminal = &types.Terminal{}
	if err := ctx.DB().Where("hostname = ?", event.Terminal).Find(terminal).Error; err != nil {
		log.Printf("terminal attempted to connect but does not exist: %v: %v", event.Terminal, err)
		return
	}

	sourceAddr := source.(*net.UDPAddr)
	if sourceAddr.IP.String() != terminal.Addr {
		log.Printf("terminal attempted to connect but has wrong address: %v: %v", event.Terminal, sourceAddr.IP.String())
		return
	}

	switch event.Type {
	case types.Event_Greetings:

		terminal.OperatingSystem = event.Data["os"]
		udpAddr := source.(*net.UDPAddr)
		terminal.Addr = udpAddr.IP.String()

		ctx.DB().Save(terminal)

	case types.Event_KeepAlive:
		if terminal.Status == types.StatusOffline {
			terminal.Status = types.StatusOnline
		}
	case types.Event_Goodbye:
		terminal.Addr = ""
		terminal.Status = types.StatusOffline
	}

	terminal.SaveRedis()
}
