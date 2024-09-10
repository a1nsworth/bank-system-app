package transport

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bank-system-app/internal/models"
	"bank-system-app/internal/services"

	_ "github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type CreateBankRequest struct {
	name        string `validate:"required"`
	bankOffices []*struct {
		address string `validate:"required"`
		status  *models.OfficeStatus
		rental  uint32 `validate:"required"`
	}
	bankAtms []*struct {
		name         string `validate:"required"`
		status       *models.AtmStatus
		amortization uint `validate:"required"`
	}
}

type BankHandler struct {
	service services.BankService
}

func NewBankHandler(service services.BankService, mux *http.ServeMux) BankHandler {
	fmt.Println("new bank handler")
	bankHandler := BankHandler{service: service}
	mux.HandleFunc("/bank/", bankHandler.handleBank)
	return bankHandler
}

func (b BankHandler) getAllBanks(w http.ResponseWriter) {
	bank, err := b.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	renderJSON(w, bank)
}

func (b BankHandler) getBankById(w http.ResponseWriter, r *http.Request, id uint) {
	bank, err := b.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		renderJSON(w, bank)
	}
}

func (b BankHandler) deleteBankById(w http.ResponseWriter, r *http.Request, id uint) {
	err := b.service.DeleteById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		renderJSON(w, struct{}{})
	}
}

func (b BankHandler) handleBank(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/bank/" {
		if r.Method == http.MethodGet {
			b.getAllBanks(w)
		}
	} else {
		args := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(args) != 2 {
			http.Error(w, "invalid request", http.StatusBadRequest)
		}
		id, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
		}
		switch r.Method {
		case http.MethodDelete:
			b.deleteBankById(w, r, uint(id))
		case http.MethodGet:
			b.getBankById(w, r, uint(id))
		}
	}
}
