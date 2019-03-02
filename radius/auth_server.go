package radius

import (
	"github.com/enderian/directrd/users"
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"log"
)

func startAuthServer() {

	if ctx.Conf().Radius.DisabledAuth {
		return
	}

	handler := func(rw radius.ResponseWriter, r *radius.Request) {
		username := rfc2865.UserName_GetString(r.Packet)
		password := rfc2865.UserPassword_GetString(r.Packet)
		identifier := rfc2865.NASIdentifier_GetString(r.Packet)

		packet := r.Packet
		if err := users.Login(username, password, identifier); err == nil {
			packet = r.Packet.Response(radius.CodeAccessAccept)
		} else {
			packet = r.Packet.Response(radius.CodeAccessReject)
			_ = rfc2865.ReplyMessage_AddString(packet, err.Error())
		}
		if err := rw.Write(packet); err != nil {
			log.Printf("Error while writing RADIUS response: %s", err.Error())
		}
	}

	server := radius.PacketServer{
		Addr:         ":1812",
		Handler:      radius.HandlerFunc(handler),
		SecretSource: radius.StaticSecretSource([]byte(ctx.Conf().Radius.SharedSecret)),
	}

	log.Printf("Starting RADIUS authentication server on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error while starting RADIUS auth server: %s", err.Error())
	}
}
