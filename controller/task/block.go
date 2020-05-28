package task

import (
	"api/dao"
	"api/public"
	"sync"
	"time"
)

type block struct {
	Number int
	Size   int
	hash   []string
	sync   sync.Mutex
}

var examplesBlock *block

func BlockInit() {
	blockLast := dao.Blok{}
	if err := (&blockLast).Last(); err != nil {
		blockLast.Number = 0
	}
	examplesBlock = &block{
		Number: blockLast.Number + 1,
		Size:   235,
	}
	examplesBlock.server()
}

func getBlock() *block {
	return examplesBlock
}

func (b *block) server() {
	t := time.NewTicker(5 * time.Second) //TODO 5秒钟出块
	for {
		<-t.C
		b.insert()
	}
}

func (b *block) insert() {
	var (
		number int
		size   int
		hash   []string
		block  dao.Blok
		tx     string
	)
	b.sync.Lock()
	number = b.Number
	size = b.Size
	hash = b.hash
	b.Number++
	b.Size = 235 //TODO 初始块235
	b.sync.Unlock()
	//TODO 预留挖矿空间
	tx = "hzcblock" + public.RandString(59, public.LOWER_CASE, public.NUMBER)
	block = dao.Blok{Number: number, Hash: tx, Size: size}
	(&block).Create()
	(&dao.Hash{BlockNumber: number, BlockHash: block.Hash}).Updates(hash)
}

func (b *block) add(hash string) {
	b.sync.Lock()
	b.hash = append(b.hash, hash)
	b.Size += 133 //TODO 单笔交易块增加133数据大小
	b.sync.Unlock()
	return
}
