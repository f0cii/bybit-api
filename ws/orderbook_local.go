package ws

import (
	"sort"
	"sync"
	"time"
)

type OrderBookLocal struct {
	ob map[string]*OrderBookL2
	m  sync.Mutex
}

func (o *OrderBookLocal) GetOrderBook() (ob OrderBook) {
	for _, v := range o.ob {
		switch v.Side {
		case "Buy":
			ob.Bids = append(ob.Bids, Item{
				Price:  v.Price,
				Amount: float64(v.Size),
			})
		case "Sell":
			ob.Asks = append(ob.Asks, Item{
				Price:  v.Price,
				Amount: float64(v.Size),
			})
		}
	}

	sort.Slice(ob.Bids, func(i, j int) bool {
		return ob.Bids[i].Price > ob.Bids[j].Price
	})

	sort.Slice(ob.Asks, func(i, j int) bool {
		return ob.Asks[i].Price < ob.Asks[j].Price
	})

	ob.Timestamp = time.Now()

	return
}

func NewOrderBookLocal() *OrderBookLocal {
	o := &OrderBookLocal{
		ob: make(map[string]*OrderBookL2),
	}
	return o
}

func (o *OrderBookLocal) LoadSnapshot(newOrderbook []*OrderBookL2) error {
	o.m.Lock()
	defer o.m.Unlock()

	o.ob = make(map[string]*OrderBookL2)

	for _, v := range newOrderbook {
		o.ob[v.Key()] = v
	}

	return nil
}

func (o *OrderBookLocal) Update(delta *OrderBookL2Delta) {
	o.m.Lock()
	defer o.m.Unlock()

	for _, elem := range delta.Delete {
		delete(o.ob, elem.Key())
	}

	for _, elem := range delta.Update {
		if v, ok := o.ob[elem.Key()]; ok {
			// price is same while id is same
			// v.Price = elem.Price
			v.Size = elem.Size
			v.Side = elem.Side
		}
	}

	for _, elem := range delta.Insert {
		o.ob[elem.Key()] = elem
	}
}
