package radius

import (
	"github.com/enderian.directrd/sessions"
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"layeh.com/radius/rfc2866"
	"log"
)

func startAccServer() {

	handler := func(w radius.ResponseWriter, r *radius.Request) {
		username := rfc2865.UserName_GetString(r.Packet)
		status := rfc2866.AcctStatusType_Get(r.Packet)
		_ = rfc2866.AcctSessionID_GetString(r.Packet)
		identifier := rfc2865.NASIdentifier_GetString(r.Packet)

		switch status {
		case rfc2866.AcctStatusType_Value_Start:
			_ = sessions.Start(username, identifier)
		case rfc2866.AcctStatusType_Value_Stop:
			_ = sessions.End(username, identifier)
		case rfc2866.AcctStatusType_Value_InterimUpdate:
			_ = sessions.Update(username, identifier)
		}
	}

	server := radius.PacketServer{
		Addr:         conf.Radius.AccountingAddress,
		Handler:      radius.HandlerFunc(handler),
		SecretSource: radius.StaticSecretSource([]byte(conf.Radius.SharedSecret)),
	}
	if server.Addr == "" {
		server.Addr = ":1813"
	}

	log.Printf("Starting RADIUS accounting server on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error while starting RADIUS accounting server: %s", err.Error())
	}
}
