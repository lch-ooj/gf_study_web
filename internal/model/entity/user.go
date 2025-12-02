package entity

import (
    "github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
    Id        uint        `json:"id"        description:"User ID"`
    Passport  string      `json:"passport"  description:"User Passport"`
    Password  string      `json:"password"  description:"User Password"`
    Nickname  string      `json:"nickname"  description:"User Nickname"`
    CreatedAt *gtime.Time `json:"createdAt" description:"Created Time"`
    UpdatedAt *gtime.Time `json:"updatedAt" description:"Updated Time"`
}