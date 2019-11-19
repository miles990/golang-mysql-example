package models

type Account struct {
	UID             string `json:"uid" db:"uid"`	
	// Source          string `json:"source" db:"source"`
	// Platform        string `json:"platform" db:"platform"`
	// Key             string `json:"key" db:"key"`
	// Secret          string `json:"secret" db:"secret"`
	Name			string `json:"name" db:"name"`	
	CreateTimestamp int64  `json:"createtimestamp" db:"createtimestamp"`
}