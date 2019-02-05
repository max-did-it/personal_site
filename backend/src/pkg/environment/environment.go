package environment

import "os"

// MailLogin is env for mail pkg
var MailLogin = os.Getenv("M_LOGIN")

// MailPass is env for mail pkg
var MailPass = os.Getenv("M_PASSWORD")

// MailTo is env for mail pkg
var MailTo = os.Getenv("M_TO")

// MailToName is env for mail pkg
var MailToName = os.Getenv("M_TO_NAME")

// MailFrom is env for mail pkg
var MailFrom = os.Getenv("M_FROM")

// MailFromName is env for mail pkg
var MailFromName = os.Getenv("M_FROM_NAME")
