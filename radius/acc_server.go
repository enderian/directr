package radius

import (
	"github.com/enderian/directrd/sessions"
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"layeh.com/radius/rfc2866"
	"log"
)

func startAccServer() {

	if ctx.Conf().Radius.DisabledAccounting {
		return
	}

	handler := func(rw radius.ResponseWriter, r *radius.Request) {
		username := rfc2865.UserName_GetString(r.Packet)
		status := rfc2866.AcctStatusType_Get(r.Packet)
		sessionId := rfc2866.AcctSessionID_GetString(r.Packet)
		identifier := rfc2865.NASIdentifier_GetString(r.Packet)

		switch status {
		case rfc2866.AcctStatusType_Value_Start:
			_ = sessions.Start(sessionId, username, identifier)
		case rfc2866.AcctStatusType_Value_Stop:
			_ = sessions.End(sessionId, username, identifier)
		case rfc2866.AcctStatusType_Value_InterimUpdate:
			_ = sessions.Update(sessionId, username, identifier)
		}

		packet := r.Packet
		packet = r.Packet.Response(radius.CodeAccountingResponse)
		if err := rw.Write(packet); err != nil {
			log.Fatalf("Error while responding to RADIUS accounting: %s", err.Error())
		}
	}

	server := radius.PacketServer{
		Addr:         ":1813",
		Handler:      radius.HandlerFunc(handler),
		SecretSource: radius.StaticSecretSource([]byte(ctx.Conf().Radius.SharedSecret)),
	}

	log.Printf("Starting RADIUS accounting server on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error while starting RADIUS accounting server: %s", err.Error())
	}
}
