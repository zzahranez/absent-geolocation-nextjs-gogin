package helper

import (
	"fmt"
	"math"
)

func ConvertToRadians(officelat, officelong, userlat, userlong float64) (float64, float64, float64, float64) {
	officeLatRad := officelat * math.Pi / 180
	officeLongRad := officelong * math.Pi / 180
	userlatRad := userlat * math.Pi / 180
	userlongRad := userlong * math.Pi / 180

	return officeLatRad, officeLongRad, userlatRad, userlongRad
}

func CountDistanceWithHaversine(officelat, officelong, userlat, userlong float64) (distance float64) {
	// Earth Radius in meter
	var Radius = 6371000.0

	// Convert to radians
	latofficeRad, longOfficeRad, userLatRad, userLongRad := ConvertToRadians(officelat, officelong, userlat, userlong)

	// Delta PI And Delta Lamda
	delta_pi := latofficeRad - userLatRad
	delta_lamda := longOfficeRad - userLongRad

	// Delta Square
	delta_pi_square := math.Pow(math.Sin(delta_pi/2), 2)
	delta_lamda_square := math.Pow(math.Sin(delta_lamda/2), 2)

	// Count with Haversine Formula
	a := delta_pi_square + math.Cos(latofficeRad)*math.Cos(userLatRad)*delta_lamda_square

	atan := math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	c := 2 * atan

	distance = c * Radius
	return
}

func TestHaversineFormula() {
	// OFFICE KOORDINAT
	officeLatitiude := -1.6088414194250245
	officeLongtitude := 103.51784563414657

	// USERS KOORDINAT
	userLatitude := -1.6072776154640327
	userLongtitude := 103.52805565738703

	Jarak := CountDistanceWithHaversine(officeLatitiude, officeLongtitude, userLatitude, userLongtitude)
	fmt.Printf("Jarak adalah: %d meter\n", int(math.Round(Jarak)))
}
