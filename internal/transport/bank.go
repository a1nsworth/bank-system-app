package transport

import (
	"fmt"
	"net/http"

	"bank-system-app/internal/services"
)

type BankHandler struct {
	service services.BankService
}

func NewBankHandler(service services.BankService, mux *http.ServeMux) BankHandler {
	fmt.Println("new bank handler")
	bankHandler := BankHandler{service: service}
	mux.HandleFunc("/bank/", bankHandler.handleBank)
	return bankHandler
}

func (b BankHandler) getAllBanks(
	w http.ResponseWriter, r *http.Request,
) {
	bank, err := b.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	renderJSON(w, bank)
}

func getBankById(w http.ResponseWriter, r *http.Request, id uint) {
}

func deleteBankById(w http.ResponseWriter, r *http.Request, id uint) {
}

func (b BankHandler) handleBank(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/bank/" {
		if r.Method == http.MethodGet {
			b.getAllBanks(w, r)
		}
	} else {
		fmt.Println("ELSE")
	}
}
