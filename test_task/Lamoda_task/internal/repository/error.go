package repository

import "errors"

var DatabaseError = errors.New("Database error")

var NoArticular = errors.New("Article not found")

var NotEnoughQuantity = errors.New("Not enough product in storage")

var NoReservation = errors.New("No product reservation")

var InvalidStorage = errors.New("Invalid storage")
