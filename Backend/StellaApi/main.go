package main

import (
	"mf-stella-api/Services"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/send_response", http.HandlerFunc(Services.SendResponse))
	// mux.Handle("/redirect_node", http.HandlerFunc(Services.RedirectNode))

	mux.Handle("/getmembers", http.HandlerFunc(Services.GetMembers))
	mux.Handle("/getmember", http.HandlerFunc(Services.GetMember))
	mux.Handle("/toggle_live_chat", http.HandlerFunc(Services.ToggleLiveChat))

	mux.Handle("/get_admins", http.HandlerFunc(Services.GetAdmins))

	mux.Handle("/get_chat_history", http.HandlerFunc(Services.GetChatHistory))

	mux.Handle("/get_assignment", http.HandlerFunc(Services.GetAssignments))

	mux.Handle("/get_group", http.HandlerFunc(Services.GetGroups))

	mux.Handle("/get_integrations", http.HandlerFunc(Services.GetIntegrations))
	mux.Handle("/get_integration", http.HandlerFunc(Services.PutIntegrations))
	mux.Handle("/get_integration/", http.HandlerFunc(Services.DeleteIntegrations))

	mux.Handle("/update_channel", http.HandlerFunc(Services.UpdatedChannel))

	mux.Handle("/get_whatsapp_file/", http.HandlerFunc(Services.GetWhatsappFile))
	mux.Handle("/get_whatsapp_template", http.HandlerFunc(Services.GetWhatsappTemplate))

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8001", handler)

}
