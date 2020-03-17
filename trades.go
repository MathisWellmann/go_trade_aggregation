package go_trade_aggregation

import (
	"encoding/csv"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

// LoadTradesFromCSV will return all trades found in csv file with given filename
func LoadTradesFromCSV(filename string) ([]*Trade, error) {
	var out []*Trade

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := csv.NewReader(file)

	// ignore header as the format is known
	_, err = r.Read()
	if err != nil {
		return nil, err
	}

	for i := 0; i < i+1; i++ {
		row, err := r.Read()
		if err != nil {
			if err.Error() == "EOF" {
				// normal end of file
			} else {
				logrus.Error(err)
			}
			break
		}
		ts, _ := strconv.ParseInt(row[0], 10, 64)
		price, _ := strconv.ParseFloat(row[1], 64)
		size, _ := strconv.ParseFloat(row[2], 64)
		t := &Trade{
			Timestamp: ts,
			Price:     price,
			Size:      size,
		}
		out = append(out, t)
	}

	return out, nil
}
