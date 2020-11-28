package datastore

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"	

	"github.com/bornehill/golang-bootcamp-2020/domain/model"
	"github.com/bornehill/golang-bootcamp-2020/usecase/repository"
	ir "github.com/bornehill/golang-bootcamp-2020/interface/repository"
)

type DbRepository struct {
	Centres interface{ repository.CentreRepository }
}

func OpenDb() *DbRepository {
	centres := initialize()

	return &DbRepository{Centres: centres}
}

func initialize() repository.CentreRepository {
	filename := "./assets/centres.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	return ir.OpenCentreRepository(loadData(file))
}

func loadData(r io.Reader) *[]*model.Centre {
	reader := csv.NewReader(r)

	ret := make([]*model.Centre, 0, 0)

	for {
		row, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				log.Println("End of File")
			} else {
				log.Println(err)	
			}
			break
		}

		id, errId := strconv.Atoi(row[0])
		capacity, errCap := strconv.Atoi(row[6])
		openness, errOpen := strconv.Atoi(row[7])

		switch {
			case errId != nil:
				log.Println(errId)
			case errCap != nil:
				log.Println(errCap)
			case errOpen != nil:
				log.Println(errOpen)								
		}

		centre := &model.Centre{
			Id:			id,
			Capacity:	capacity,
			Openness:	openness,
			Name:		row[1],
			Address:	row[2],
			Email:		row[3],
			Phone:		row[4],
			Line:		row[5],
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, centre)
	}
	return &ret
}
