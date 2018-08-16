package types

import (
	"time"
	"strconv"
	"math"
)

type Response struct {
	Result interface{} `json:"result"`
	Time ResponseTime `json:"time"`
}

type UnixMicroTime time.Time
type SecDuration time.Duration

func (t *UnixMicroTime) UnmarshalJSON(data []byte) error {
	ts, err := strconv.ParseFloat(string(data), 10)

	if err != nil {
		return err
	}

	i, f := math.Modf(ts)

	*t = UnixMicroTime(time.Unix(int64(i), int64(f * 10e6) * int64(time.Microsecond)))
	return nil
}

func (d *SecDuration) UnmarshalJSON(data []byte) error {
	df, err := strconv.ParseFloat(string(data), 10)

	if err != nil {
		return err
	}

	*d = SecDuration(int64(df * 10e9) * int64(time.Nanosecond))
	return nil
}

type ResponseTime struct {
	Start UnixMicroTime
	Finish UnixMicroTime
	Duration SecDuration
	Processing SecDuration
}

type ResponseError struct {
	Code string `json:"error"`
	Description string `json:"error_description"`
}

type IntResponse struct {
	Response
	Result int64 `json:"result,string"`
}

type StringResponse struct {
	Response
	Result string `json:"result"`
}