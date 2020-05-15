package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slcsp/entities"
	"sort"
	"strconv"
)

const (
	NUM_PRICE_FIELDS   = 5
	NUM_ZIPCODE_FIELDS = 5
	NUM_ZIPCODE_TEST_FIELDS = 2
)

type priceKey struct {
	entities.RateArea
	entities.Level
}

func main() {
	priceMap, err := getPriceMap("testdata/plans.csv")
	if err != nil {
		log.Fatalf("Failed to get price map %s", err.Error())
	}

	zipCodeMap, err := getZipCodeMap("testdata/zips.csv")
	if err != nil {
		log.Fatalf("Failed to get zipCode map %s", err.Error())
	}

	zipCodesUnderTest, err := getZipCodesToTest("testdata/slcsp.csv")
	if err != nil {
		log.Fatalf("Failed to get zipCodes under test %s", err.Error())
	}

	OuterLoop:
	for _, zipCode := range zipCodesUnderTest {
		// Cannot determine price if zipCode has more than one rateArea
		if len(zipCodeMap[zipCode]) != 1 {
			fmt.Printf("%s,\n", zipCode)
			continue
		}
		for rateArea := range zipCodeMap[zipCode] {
			key := priceKey{
				rateArea,
				entities.Silver,
			}
			if len(priceMap[key]) <= 1 {
				fmt.Printf("%s,\n", zipCode)
				continue OuterLoop
			}

			var rates entities.Rates
			for rate := range priceMap[key] {
				rates = append(rates, rate)
			}

			sort.Sort(rates)
			fmt.Printf("%s,%.2f\n", zipCode, rates[1])
		}
	}
}

// getPriceMap reads in a CSV file with the assumed format:
// plan_id,state,metal_level,rate,rate_area (rateNumber)
// The header is assumed to be present
// returns a mapping between RateArea and unique Rates
func getPriceMap(planCsvFile string) (priceMap map[priceKey]map[entities.Rate]struct{}, err error)  {
	priceMap = make(map[priceKey]map[entities.Rate]struct{})

	f, err := os.Open(planCsvFile)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
	}()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	// Skip the header
	for _, record:= range records[1:] {
		if (len(record) != NUM_PRICE_FIELDS) {
			return nil, fmt.Errorf("Field count %d does not match expected %d",
				len(record), NUM_PRICE_FIELDS)
		}
		state := record[1]
		rateNumber, err := strconv.ParseUint(record[4], 10, 32)
		if err != nil {
			return nil, err
		}
		rate, err := strconv.ParseFloat(record[3], 32)
		if err != nil {
			return nil, err
		}
		level, err := entities.ParseLevel(record[2])
		if err != nil {
			return nil, err
		}

		key := priceKey{
			entities.RateArea{
				State: entities.StateCode(state),
				Number: entities.RateNumber(rateNumber),
			},
			level,
		}

		uniqueRates := priceMap[key]
		if uniqueRates == nil {
			uniqueRates = make(map[entities.Rate]struct{})
			priceMap[key] = uniqueRates
		}
		uniqueRates[entities.Rate(rate)] = struct{}{}
	}

	return //
}

// getZipCodeMap reads in a CSV file with the assumed format:
// zipcode,state,county_code,name,rate_area (rateNumber)
// The header is assumed to be present
// returns a mapping between ZipCode and unique RateAreas
func getZipCodeMap(zipCodeCsvFile string) (zipCodeMap map[entities.ZipCode]map[entities.RateArea]struct{}, err error) {
	zipCodeMap = make(map[entities.ZipCode]map[entities.RateArea]struct{})

	f, err := os.Open(zipCodeCsvFile)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
	}()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	// Skip the header
	for _, record := range records[1:] {
		if (len(record) != NUM_ZIPCODE_FIELDS) {
			return nil, fmt.Errorf("Field count %d does not match expected %d",
				len(record), NUM_ZIPCODE_FIELDS)
		}

		zipCode := entities.ZipCode(record[0])
		state := record[1]
		rateNumber, err := strconv.ParseUint(record[4], 10, 32)
		if err != nil {
			return nil, err
		}

		rateArea := entities.RateArea{
			State: entities.StateCode(state),
			Number: entities.RateNumber(rateNumber),
		}

		uniqueRateAreas := zipCodeMap[zipCode]
		if uniqueRateAreas == nil {
			uniqueRateAreas = make(map[entities.RateArea]struct{})
			zipCodeMap[zipCode] = uniqueRateAreas
		}
		uniqueRateAreas[rateArea] = struct{}{}
	}

	return //
}

// getZipCodeMap reads in a CSV file with the assumed format:
// zipcode,state,
// The header is assumed to be present
// returns zipCodes to test
func getZipCodesToTest(zipCodeTestCsvFile string) (zipCodesUnderTest []entities.ZipCode, err error) {
	f, err := os.Open(zipCodeTestCsvFile)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
	}()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	// Skip the header
	for _, record := range records[1:] {
		if (len(record) != NUM_ZIPCODE_TEST_FIELDS) {
			return nil, fmt.Errorf("Field count %d does not match expected %d",
				len(record), NUM_ZIPCODE_TEST_FIELDS)
		}
		zipCodesUnderTest = append(zipCodesUnderTest, entities.ZipCode(record[0]))
	}

	return //
}
