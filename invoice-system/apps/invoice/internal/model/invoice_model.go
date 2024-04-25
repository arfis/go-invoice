package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Company struct {
	ID               *uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name             string     `json:"name" gorm:"default:'Curry labs s.r.o'"`
	CompanyNumber    string     `json:"companyNumber" gorm:"default:12312312312;unique"`
	TaxNumber        string     `json:"taxNumber" gorm:"default:SK12312312312"`
	CompanyTaxNumber string     `json:"companyTaxNumber gorm:"default:12312312312""`
	Street           string     `json:"street" gorm:"default:Damborskeho 6"`
	City             string     `json:"city" gorm:"default:Bratislava"`
	Country          string     `json:"country" gorm:"default:Slovakia"`
	Zipcode          string     `json:"zipcode" gorm:"default:12312312312"`
}

type InvoiceItem struct {
	ID        *uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	InvoiceID *uuid.UUID `json:"invoiceId"`
	Code      string     `json:"code" gorm:"default:'A123'"`
	Name      string     `json:"name" gorm:"default:'Vykopove prace'"`
	Price     uint       `json:"price" gorm:"default:100"`
	Invoice   Invoice    `gorm:"foreignKey:InvoiceID"`
}

type Invoice struct {
	ID                 *uuid.UUID    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	OwnerId            string        `json:"ownerId" gorm:"column:owner_id"`
	CreatedDate        *time.Time    `json:"createdDate" gorm:"default:CURRENT_TIMESTAMP"`
	ValidUntil         *time.Time    `json:"validUntil" gorm:"default:CURRENT_TIMESTAMP"`
	DaysToPay          int           `json:"daysToPay" gorm:"default:30"`
	EmailTo            *string       `json:"emailTo" gorm:"default:'default@email.com'"`
	ReceivingCompanyID *uuid.UUID    `json:"receivingCompanyId"`
	ReceivingCompany   Company       `json:"receivingCompany" gorm:"foreignKey:ReceivingCompanyID"`
	IssuingCompanyID   *uuid.UUID    `json:"issuingCompanyId"`
	IssuingCompany     Company       `json:"issuingCompany" gorm:"foreignKey:IssuingCompanyID"`
	InvoiceItems       []InvoiceItem `json:"invoiceItems" gorm:"foreignKey:InvoiceID"` // Correct setup for 1-to-many
}

func (invoice Invoice) String() string {
	return fmt.Sprintf("\nCURRENT INVOICE code: %s price: %d owner: %s \n", invoice.ReceivingCompany.Name, invoice.DaysToPay, invoice.OwnerId)
}

type InvoiceResponse struct {
	Code    string `json:"code"`
	Price   uint   `json:"price"`
	OwnerId string `json:"ownerId"`
}

func (invoice InvoiceResponse) String() string {
	return fmt.Sprintf("\nCURRENT INVOICE code: %s price: %d owner: %s \n", invoice.Code, invoice.Price, invoice.OwnerId)
}

type CreateInvoiceInput struct {
	Code    string `json:"code"`
	Price   uint   `json:"price"`
	OwnerId string `json:"ownerId"`
}
