package merit

import "time"

type queryDate struct {
	time.Time
}

func (d *queryDate) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("20060102", string(b))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d queryDate) MarshalJSON() ([]byte, error) {
	return []byte(d.Time.Format("20060102")), nil
}
