package model

import (
	"github.com/google/uuid"
	"time"
)

type Company struct {
	ID               *uuid.UUID `json:"id,omitempty"`
	Name             string     `json:"name"`
	CompanyNumber    string     `json:"companyNumber"`
	TaxNumber        string     `json:"taxNumber"`
	CompanyTaxNumber string     `json:"companyTaxNumber"`
	Street           string     `json:"street"`
	City             string     `json:"city"`
	Country          string     `json:"country"`
	Zipcode          string     `json:"zipcode"`
}

type InvoiceItem struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	InvoiceID uint       `json:"invoiceId,omitempty"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Price     uint       `json:"price"`
}

type Invoice struct {
	ID                 *uuid.UUID    `json:"id,omitempty"`
	OwnerId            string        `json:"ownerId" gorm:"column:owner_id"`
	CreatedDate        *time.Time    `json:"createdDate"`
	ValidUntil         *time.Time    `json:"validUntil"`
	DaysToPay          int           `json:"daysToPay"`
	EmailTo            string        `json:"emailTo"`
	ReceivingCompanyID *uuid.UUID    `json:"receivingCompanyId,omitempty"`
	ReceivingCompany   Company       `json:"receivingCompany" gorm:"foreignKey:ReceivingCompanyID"`
	IssuingCompanyID   *uuid.UUID    `json:"issuingCompanyId,omitempty"`
	IssuingCompany     Company       `json:"issuingCompany" gorm:"foreignKey:IssuingCompanyID"`
	InvoiceItems       []InvoiceItem `json:"invoiceItems" gorm:"foreignKey:InvoiceID"`
}
