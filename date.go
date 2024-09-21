package merit

import "time"

type QueryDate struct {
	time.Time
}

func (d *QueryDate) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("20060102", string(b))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d QueryDate) MarshalJSON() ([]byte, error) {
	return []byte(d.Time.Format("20060102")), nil
}

func QueryDateNow() QueryDate {
	return QueryDate{time.Now()}
}
