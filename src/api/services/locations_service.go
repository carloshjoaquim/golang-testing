package services

import (
	"github.com/carloshjoaquim/golang-testing/src/api/domain/locations"
	"github.com/carloshjoaquim/golang-testing/src/api/providers/locations_provider"
	"github.com/carloshjoaquim/golang-testing/src/api/utils/errors"
)

type locationsService struct{}
type locationsServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

var (
	LocationService locationsServiceInterface
)

func init() {
	LocationService = &locationsService{}
}

func (s *locationsService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return locations_provider.GetCountry(countryId)
}
