package geo

import (
	"context"
	"fmt"
	"microservice/geo/internal/models"
	"microservice/geo/internal/service"
	pb "microservice/geo/protos/gen/go"
)

type Geo interface {
	GeoSearch(input string) ([]*models.Address, error)
	GeoCode(lat, lng string) ([]*models.Address, error)
}

type ServerGeo struct {
	pb.UnimplementedGeoServiceServer
	geo service.GeoService
}

func (s *ServerGeo) SearchAddress(context context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {

	address, err := s.geo.GeoSearch(req.Input)
	if err != nil {
		return nil, fmt.Errorf("err get address:%v", err)
	}

	var grpcAddresses []*pb.Address
	for _, addr := range address {
		grpcAddress := &pb.Address{
			PostalCode:           addr.PostalCode.(string),
			Country:              addr.Country,
			CountryIsoCode:       addr.CountryISOCode,
			FederalDistrict:      addr.FederalDistrict.(string),
			RegionFiasId:         addr.RegionFIASID,
			RegionKladrId:        addr.RegionKLADRID,
			RegionIsoCode:        addr.RegionISOCode,
			RegionWithType:       addr.RegionWithType,
			RegionType:           addr.RegionType,
			RegionTypeFull:       addr.RegionTypeFull,
			Region:               addr.Region,
			AreaFiasId:           addr.AreaFIASID.(string),
			AreaKladrId:          addr.AreaKLADRID.(string),
			AreaWithType:         addr.AreaWithType.(string),
			AreaType:             addr.AreaType.(string),
			AreaTypeFull:         addr.AreaTypeFull.(string),
			Area:                 addr.Area.(string),
			CityFiasId:           addr.CityFIASID,
			CityKladrId:          addr.CityKLADRID,
			CityWithType:         addr.CityWithType,
			CityType:             addr.CityType,
			CityTypeFull:         addr.CityTypeFull,
			City:                 addr.City,
			CityArea:             addr.CityArea.(string),
			CityDistrictFiasId:   addr.CityDistrictFIASID.(string),
			CityDistrictKladrId:  addr.CityDistrictKLADRID.(string),
			CityDistrictWithType: addr.CityDistrictWithType.(string),
			CityDistrictType:     addr.CityDistrictType.(string),
			CityDistrictTypeFull: addr.CityDistrictTypeFull.(string),
			CityDistrict:         addr.CityDistrict.(string),
			StreetFiasId:         addr.StreetFIASID,
			StreetKladrId:        addr.StreetKLADRID,
			StreetWithType:       addr.StreetWithType,
			StreetType:           addr.StreetType,
			StreetTypeFull:       addr.StreetTypeFull,
			Street:               addr.Street,
			SteadFiasId:          addr.SteadFIASID.(string),
			SteadCadnum:          addr.SteadCadnum.(string),
			SteadType:            addr.SteadType.(string),
			SteadTypeFull:        addr.SteadTypeFull.(string),
			Stead:                addr.Stead.(string),
			HouseFiasId:          addr.HouseFIASID.(string),
			HouseKladrId:         addr.HouseKLADRID.(string),
			HouseCadnum:          addr.HouseCadnum.(string),
			HouseType:            addr.HouseType.(string),
			HouseTypeFull:        addr.HouseTypeFull.(string),
			House:                addr.House.(string),
			BlockType:            addr.BlockType.(string),
			BlockTypeFull:        addr.BlockTypeFull.(string),
			Block:                addr.Block.(string),
			Entrance:             addr.Entrance.(string),
			Floor:                addr.Floor.(string),
			FlatFiasId:           addr.FlatFIASID.(string),
			FlatCadnum:           addr.FlatCadnum.(string),
			FlatType:             addr.FlatType.(string),
			FlatTypeFull:         addr.FlatTypeFull.(string),
			Flat:                 addr.Flat.(string),
			FlatArea:             addr.FlatArea.(string),
			SquareMeterPrice:     addr.SquareMeterPrice.(string),
			FlatPrice:            addr.FlatPrice.(string),
			PostalBox:            addr.PostalBox.(string),
			FiasId:               addr.FIASID,
			FiasCadastreNumber:   addr.FIASCadastreNumber,
			FiasLevel:            addr.FIASLevel,
			FiasActualityState:   addr.FIASActualityState,
			KladrId:              addr.KLADRID,
			GeonameId:            addr.GeonameID,
			CapitalMarker:        addr.CapitalMarker,
			Okato:                addr.OKATO,
			Oktmo:                addr.OKTMO,
			TaxOffice:            addr.TaxOffice,
			TaxOfficeLegal:       addr.TaxOfficeLegal,
			Timezone:             addr.Timezone.(string),
			GeoLat:               addr.GeoLat,
			GeoLon:               addr.GeoLon,
			BeltwayHit:           addr.BeltwayHit.(string),
			BeltwayDistance:      addr.BeltwayDistance.(string),
			Metro:                addr.Metro.(string),
			Divisions:            addr.Divisions.(string),
			QcGeo:                addr.QCGeo,
			QcComplete:           addr.QCComplete.(string),
			QcHouse:              addr.QCHouse.(string),
			HistoryValues:        addr.HistoryValues,
			UnparsedParts:        addr.UnparsedParts.(string),
			Source:               addr.Source.(string),
			Qc:                   addr.QC.(string),
		}
		grpcAddresses = append(grpcAddresses, grpcAddress)
	}
	return &pb.SearchResponse{Addresses: grpcAddresses}, nil
}

func (s *ServerGeo) GeocodeAddress(context context.Context, req *pb.GeocodeRequest) (*pb.GeocodeResponse, error) {
	return nil, nil
}
