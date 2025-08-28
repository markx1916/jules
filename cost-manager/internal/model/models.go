package model
import ("time"; "gorm.io/gorm")
type CostData struct { gorm.Model; CloudProvider, AccountID, ServiceName string; CostActual float64; BillingDate time.Time }
type Budget struct { gorm.Model; Name, ScopeType, ScopeValue, Period string; Amount float64 }
type Alert struct { gorm.Model; AlertType, Description, Severity string; IsResolved bool }
