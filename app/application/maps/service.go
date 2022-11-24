package maps

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/kelvins/geocoder"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
)

type maps struct {
	repo              gateways.MapRepo
	addressRepo       gateways.AddressRepo
	merchantRepo      gateways.MerchantRepo
	merchantStoreRepo gateways.MerchantStoreRepo
	repoType          string
	taskType          string
	GoogleAPIKey      string
	WG                *sync.WaitGroup
	DataChan          chan any
	DoneChan          chan bool
	ErrorChan         chan error
}

func NewMaps(wg *sync.WaitGroup) gateways.MapService {
	dataChan := make(chan any, 1024)
	doneChan := make(chan bool)
	errorChan := make(chan error)
	return &maps{
		taskType:     "",
		GoogleAPIKey: config.Google().APIKey,
		WG:           wg,
		DataChan:     dataChan,
		DoneChan:     doneChan,
		ErrorChan:    errorChan,
	}
}
func (m *maps) SetRepo(repo gateways.MapRepo) gateways.MapService {
	m.repo = repo
	return m
}
func (m *maps) SetAddressRepo(repo gateways.AddressRepo) gateways.MapService {
	m.addressRepo = repo
	return m
}
func (m *maps) SetMerchantRepo(repo gateways.MerchantRepo) gateways.MapService {
	m.merchantRepo = repo
	return m
}

func (m *maps) SetMerchantStoreRepo(repo gateways.MerchantStoreRepo) gateways.MapService {
	m.merchantStoreRepo = repo
	return m
}

func (m *maps) ExecuteTask(data any, taskType, repoType string) {
	m.WG.Add(1)
	m.DataChan <- data
	m.taskType = taskType
	m.repoType = repoType
}

func (m *maps) Listen() {
	for {
		select {
		case data := <-m.DataChan:
			go m.parseTask(data, m.ErrorChan)
		case err := <-m.ErrorChan:
			fmt.Println(err)
		case <-m.DoneChan:
			return
		}
	}
}

func (m *maps) Done() {
	m.DoneChan <- true
}

func (m *maps) CloseChannels() {
	close(m.DataChan)
	close(m.ErrorChan)
	close(m.DoneChan)
}

func (m *maps) parseTask(data any, errorChan chan error) {
	defer m.WG.Done()
	switch m.taskType {
	case "geocoding":
		if err := m.getCoordinate(data); err != nil {
			errorChan <- err
		}
	}
}

func (m *maps) getCoordinate(data any) error {

	r, id, err := m.formatAddress(data)
	if err != nil {
		return err
	}
	if r == nil {
		return nil
	}
	geocoder.ApiKey = m.GoogleAPIKey
	location, errr := geocoder.Geocoding(
		geocoder.Address{
			Street:   r.Street,
			Number:   r.Number,
			City:     r.City,
			District: r.District,
			State:    r.State,
			Country:  r.Country,
		},
	)
	if errr != nil {
		return err
	}

	if m.repoType == "address" {
		return m.addressRepo.SaveCoordinate(
			&services.Coordinate{
				Latitude:  location.Latitude,
				Longitude: location.Longitude,
			}, id,
		)
	}

	if m.repoType == "merchant" {
		return m.merchantRepo.SaveCoordinate(
			&services.Coordinate{
				Latitude:  location.Latitude,
				Longitude: location.Longitude,
			}, id,
		)
	}
	if m.repoType == "store" {
		return m.merchantStoreRepo.SaveCoordinate(
			&services.Coordinate{
				Latitude:  location.Latitude,
				Longitude: location.Longitude,
			}, id,
		)
	}

	return nil
}

func (m *maps) formatAddress(data any) (*services.GeocodingData, int, error) {
	if m.repoType == "address" {
		a := data.(*ent.Address)
		num, _ := strconv.Atoi(a.StreetNumber)
		return &services.GeocodingData{
			Street:   a.StreetName,
			Number:   num,
			City:     a.City,
			District: a.District,
			State:    a.Region,
			Country:  a.Country,
		}, a.ID, nil
	}
	if m.repoType == "merchant" {
		mr := data.(*ent.Merchant)
		num, _ := strconv.Atoi(mr.Edges.Store.Address.StreetNumber)
		return &services.GeocodingData{
			Street:   mr.Edges.Store.Address.StreetName,
			Number:   num,
			City:     mr.Edges.Store.Address.City,
			District: mr.Edges.Store.Address.District,
			State:    mr.Edges.Store.Address.Region,
			Country:  mr.Edges.Store.Address.Country,
		}, mr.Edges.Store.ID, nil
	}
	if m.repoType == "store" {
		ms := data.(*ent.MerchantStore)
		num, _ := strconv.Atoi(ms.Address.StreetNumber)
		return &services.GeocodingData{
			Street:   ms.Address.StreetName,
			Number:   num,
			City:     ms.Address.City,
			District: ms.Address.District,
			State:    ms.Address.Region,
			Country:  ms.Address.Country,
		}, ms.ID, nil
	}

	return nil, 0, nil
}
