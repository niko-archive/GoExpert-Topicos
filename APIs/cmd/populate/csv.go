package populate

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"gorm.io/gorm"
)

func LoadVehicleCSV(db *gorm.DB) {

	csvfile, err := os.Open("vehicle.csv")
	if err != nil {
		panic(err)
	}
	defer csvfile.Close()

	csvReader := csv.NewReader(bufio.NewReader(csvfile))
	csvReader.Read()

	allVehicles := []*entity.Vehicle{}
	start := time.Now()

	total := 0
	for {
		total++
		record, err := csvReader.Read()
		if err != nil {
			break
		}

		value, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Println(err)
		}

		brand := record[1]
		model := record[2]
		model_year := record[3]
		fuel := record[4]
		fipe_code := record[5]
		reference_month := record[6]
		vehicle_type := record[7]

		vehicle, err := entity.NewVehicle(
			value,
			brand,
			model,
			fuel,
			fipe_code,
			reference_month,
			vehicle_type,
			model_year,
		)

		if err != nil {
			log.Println(err)
			log.Println(record)

		}

		allVehicles = append(allVehicles, vehicle)
	}

	log.Println("Total:", total)
	log.Println("Total Entries:", len(allVehicles))

	tx := db.Begin()
	log.Println("Creating Vehicles...")
	for _, vehicle := range allVehicles {
		tx.Create(&vehicle)
	}

	log.Println("Done!")

	tx.Commit()
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)

}
