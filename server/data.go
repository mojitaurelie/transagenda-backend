package server

import (
	"github.com/transagenda-back/database"
	"net/http"
)

func Appointments(w http.ResponseWriter, r *http.Request) {
	userId, err := userIdFromContext(r.Context())
	if err != nil {
		internalServerError(w, r)
		return
	}
	appointments, err := database.AppointmentsByUserId(userId)
	if err != nil {
		internalServerError(w, r)
		return
	}
	ok(appointments, w, r)
}

func Prescriptions(w http.ResponseWriter, r *http.Request) {
	userId, err := userIdFromContext(r.Context())
	if err != nil {
		internalServerError(w, r)
		return
	}
	prescriptions, err := database.PrescriptionsByUserId(userId)
	if err != nil {
		internalServerError(w, r)
		return
	}
	ok(prescriptions, w, r)
}

func Contacts(w http.ResponseWriter, r *http.Request) {
	userId, err := userIdFromContext(r.Context())
	if err != nil {
		internalServerError(w, r)
		return
	}
	contacts, err := database.ContactsByUserId(userId)
	if err != nil {
		internalServerError(w, r)
		return
	}
	ok(contacts, w, r)
}
