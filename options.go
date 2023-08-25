package generator

// UnicodeTranslateFunc ...
type UnicodeTranslateFunc func(string) string

// Options for Document
type Options struct {
	AutoPrint bool `json:"auto_print,omitempty"`

	CurrencySymbol    string `default:"RSD " json:"currency_symbol,omitempty"`
	CurrencyPrecision int    `default:"2" json:"currency_precision,omitempty"`
	CurrencyDecimal   string `default:"." json:"currency_decimal,omitempty"`
	CurrencyThousand  string `default:" " json:"currency_thousand,omitempty"`

	TextTypeInvoice      string `default:"FAKTURA" json:"text_type_invoice,omitempty"`
	TextTypeQuotation    string `default:"QUOTATION" json:"text_type_quotation,omitempty"`
	TextTypeDeliveryNote string `default:"DELIVERY NOTE" json:"text_type_delivery_note,omitempty"`

	TextRefTitle         string `default:"Referenca" json:"text_ref_title,omitempty"`
	TextVersionTitle     string `default:"Verzija" json:"text_version_title,omitempty"`
	TextDateTitle        string `default:"Datum" json:"text_date_title,omitempty"`
	TextPaymentTermTitle string `default:"Rok za placanje" json:"text_payment_term_title,omitempty"`

	TextItemsNameTitle     string `default:"Naziv" json:"text_items_name_title,omitempty"`
	TextItemsUnitCostTitle string `default:"Cena" json:"text_items_unit_cost_title,omitempty"`
	TextItemsQuantityTitle string `default:"Kol." json:"text_items_quantity_title,omitempty"`
	TextItemsTotalHTTitle  string `default:"Ukupno bez poreza" json:"text_items_total_ht_title,omitempty"`
	TextItemsTaxTitle      string `default:"Porez" json:"text_items_tax_title,omitempty"`
	TextItemsDiscountTitle string `default:"Popust" json:"text_items_discount_title,omitempty"`
	TextItemsTotalTTCTitle string `default:"Ukupno" json:"text_items_total_ttc_title,omitempty"`

	TextTotalTotal      string `default:"Ukupno" json:"text_total_total,omitempty"`
	TextTotalDiscounted string `default:"Ukupno popust" json:"text_total_discounted,omitempty"`
	TextTotalTax        string `default:"Porez" json:"text_total_tax,omitempty"`
	TextTotalWithTax    string `default:"Ukupno sa porezom" json:"text_total_with_tax,omitempty"`

	BaseTextColor []int `default:"[35,35,35]" json:"base_text_color,omitempty"`
	GreyTextColor []int `default:"[82,82,82]" json:"grey_text_color,omitempty"`
	GreyBgColor   []int `default:"[232,232,232]" json:"grey_bg_color,omitempty"`
	DarkBgColor   []int `default:"[212,212,212]" json:"dark_bg_color,omitempty"`

	Font     string `default:"Helvetica"`
	BoldFont string `default:"Helvetica"`

	UnicodeTranslateFunc UnicodeTranslateFunc
}
