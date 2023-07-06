package maib

// currencyEnum holds ISO4217 codes for currencies.
type currencyEnum uint16

const (
	// CurrencyMDL is ISO4217 code for Moldovan Lei.
	CurrencyMDL currencyEnum = 498

	// CurrencyEUR is ISO4217 code for Euro.
	CurrencyEUR currencyEnum = 978

	// CurrencyUSD is ISO4217 code for United States Dollar.
	CurrencyUSD currencyEnum = 840
)

// languageEnum holds 2 letter language identifiers.
type languageEnum string

const (
	// LanguageRomanian language identifier.
	LanguageRomanian languageEnum = "ro"

	// LanguageRussian language identifier.
	LanguageRussian languageEnum = "ru"

	// LanguageEnglish language identifier.
	LanguageEnglish languageEnum = "en"
)

// resultEnum holds possible values for RESULT field in response from MAIB EComm system.
type resultEnum string

const (
	// ResultOk - the transaction is successfully completed.
	ResultOk resultEnum = "OK"

	// ResultFailed - the transaction has failed.
	ResultFailed resultEnum = "FAILED"

	// ResultCreated - the transaction is just registered in the system.
	ResultCreated resultEnum = "CREATED"

	// ResultPending - the transaction is not completed yet.
	ResultPending resultEnum = "PENDING"

	// ResultDeclined - the transaction is declined by EComm.
	ResultDeclined resultEnum = "DECLINED"

	// ResultReversed - the transaction is reversed.
	ResultReversed resultEnum = "REVERSED"

	// ResultAutoReversed - the transaction is reversed by autoreversal.
	ResultAutoReversed resultEnum = "AUTOREVERSED"

	// ResultTimeout - the transaction was timed out.
	ResultTimeout resultEnum = "TIMEOUT"
)

// resultPSEnum holds possible values for RESULT_PS field in response from MAIB EComm system.
type resultPSEnum string

const (
	// ResultPSActive - the transaction was registered and payment is not completed yet.
	ResultPSActive resultPSEnum = "ACTIVE"

	// ResultPSFinished - payment was completed successfully.
	ResultPSFinished resultPSEnum = "FINISHED"

	// ResultPSCancelled - payment was cancelled.
	ResultPSCancelled resultPSEnum = "CANCELLED"

	// ResultPSReturned - payment was returned.
	ResultPSReturned resultPSEnum = "RETURNED"
)