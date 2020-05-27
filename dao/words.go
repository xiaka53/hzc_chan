package dao

import "api/public"

type Words struct {
	Word string `json:"word" orm:"cloumn(word);primary_key" description:"word"`
}

func (w *Words) TableName() string {
	return "words"
}

func (w *Words) Rand(num int) (list []string) {
	public.ChanPool.Table(w.TableName()).Order("order by RAND()").Pluck("word", &list)
	return
}
