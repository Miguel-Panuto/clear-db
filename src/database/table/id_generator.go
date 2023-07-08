package table

import "github.com/google/uuid"

var ids []string = []string{}

func newId() {
	u, err := uuid.NewRandom()
	if err != nil {
		return
	}
	ids = append(ids, u.String())
}

func getId() string {
	id := ids[0]

	ids = ids[1:]

	go newId()
	return id
}

func StartIds(howMuch int) {
	for i := 0; i < howMuch; i++ {
		u, err := uuid.NewRandom()
		if err != nil {
			i--
			continue
		}
		ids[i] = u.String()
	}
}
