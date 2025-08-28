package model
import ("gorm.io/gorm")
type Recommendation struct { gorm.Model; CloudProvider, ServiceID, ResourceType, Description string; PotentialSavings float64; IsActioned bool }
