package common

type Error struct {
}

func (e Error) CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
