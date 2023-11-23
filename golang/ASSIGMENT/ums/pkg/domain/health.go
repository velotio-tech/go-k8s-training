package domain

import (
	"context"
	"log"
	"time"
)

// GetHealth checks if the the database connection is live or not
func (d *UserCliet) GetHealth() bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(d.Timeout)*time.Second)
	defer cancel()
	err := d.DB.PingContext(ctx)
	if err != nil {
		log.Println("failed to connect to database : ", err)
		return false
	}
	return true
}
