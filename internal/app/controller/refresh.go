package controller

import (
	"context"
	"sort"

	"github.com/btmura/ponzi2/internal/app/model"
	"github.com/btmura/ponzi2/internal/stock"

	"golang.org/x/sync/errgroup"
)

func (c *Controller) stockUpdate(ctx context.Context, symbol string) controllerStockUpdate {
	g, gCtx := errgroup.WithContext(ctx)

	var hist *stock.History
	g.Go(func() error {
		h, err := c.stockDataFetcher.GetHistory(gCtx, &stock.GetHistoryRequest{Symbol: symbol})
		if err != nil {
			return err
		}
		hist = h
		return nil
	})

	if err := g.Wait(); err != nil {
		return controllerStockUpdate{
			symbol:    symbol,
			updateErr: err,
		}
	}

	return controllerStockUpdate{
		symbol: symbol,
		update: makeStockUpdate(symbol, hist.TradingSessions),
	}
}

func makeStockUpdate(symbol string, ts []*stock.TradingSession) *model.StockUpdate {
	ds := dailySessions(ts)
	ws := weeklySessions(ds)

	fillChangeValues(ds)
	fillChangeValues(ws)

	fillStochastics(ds)
	fillStochastics(ws)

	fillMovingAverages(ds)
	fillMovingAverages(ws)

	ds, ws = trimSessions(ds, ws)

	return &model.StockUpdate{
		Symbol:         symbol,
		DailySessions:  ds,
		WeeklySessions: ws,
	}
}

func dailySessions(ts []*stock.TradingSession) (ds []*model.TradingSession) {
	for _, s := range ts {
		ds = append(ds, &model.TradingSession{
			Date:   s.Date,
			Open:   s.Open,
			High:   s.High,
			Low:    s.Low,
			Close:  s.Close,
			Volume: s.Volume,
		})
	}
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Date.Before(ds[j].Date)
	})
	return ds
}

func weeklySessions(ds []*model.TradingSession) (ws []*model.TradingSession) {
	for _, s := range ds {
		diffWeek := ws == nil
		if !diffWeek {
			_, week := s.Date.ISOWeek()
			_, prevWeek := ws[len(ws)-1].Date.ISOWeek()
			diffWeek = week != prevWeek
		}

		if diffWeek {
			sc := *s
			ws = append(ws, &sc)
		} else {
			ls := ws[len(ws)-1]
			if ls.High < s.High {
				ls.High = s.High
			}
			if ls.Low > s.Low {
				ls.Low = s.Low
			}
			ls.Close = s.Close
			ls.Volume += s.Volume
		}
	}
	return ws
}

func fillChangeValues(ss []*model.TradingSession) {
	for i := range ss {
		if i > 0 {
			ss[i].Change = ss[i].Close - ss[i-1].Close
			ss[i].PercentChange = ss[i].Change / ss[i-1].Close
		}
	}
}

func fillStochastics(ss []*model.TradingSession) {
	const (
		k = 10
		d = 3
	)

	// Calculate fast %K for stochastics.
	fastK := make([]float32, len(ss))
	for i := range ss {
		if i+1 < k {
			continue
		}

		highestHigh, lowestLow := ss[i].High, ss[i].Low
		for j := 0; j < k; j++ {
			if highestHigh < ss[i-j].High {
				highestHigh = ss[i-j].High
			}
			if lowestLow > ss[i-j].Low {
				lowestLow = ss[i-j].Low
			}
		}
		fastK[i] = (ss[i].Close - lowestLow) / (highestHigh - lowestLow)
	}

	// Calculate fast %D (slow %K) for stochastics.
	for i := range ss {
		if i+1 < k+d {
			continue
		}
		ss[i].K = (fastK[i] + fastK[i-1] + fastK[i-2]) / 3
	}

	// Calculate slow %D for stochastics.
	for i := range ss {
		if i+1 < k+d+d {
			continue
		}
		ss[i].D = (ss[i].K + ss[i-1].K + ss[i-2].K) / 3
	}
}

func fillMovingAverages(ss []*model.TradingSession) {
	average := func(i, n int) (avg float32) {
		if i+1-n < 0 {
			return 0 // Not enough data
		}
		var sum float32
		for j := 0; j < n; j++ {
			sum += ss[i-j].Close
		}
		return sum / float32(n)
	}

	for i := range ss {
		ss[i].MovingAverage25 = average(i, 25)
		ss[i].MovingAverage50 = average(i, 50)
		ss[i].MovingAverage200 = average(i, 200)
	}
}

func trimSessions(ds, ws []*model.TradingSession) (trimDs, trimWs []*model.TradingSession) {
	const sixMonthWeeks = 4 /* weeks */ * 6 /* months */
	if len(ws) >= sixMonthWeeks {
		ws = ws[len(ws)-sixMonthWeeks:]
		for i := range ds {
			if ds[i].Date == ws[0].Date {
				ds = ds[i:]
				return ds, ws
			}
		}
	}
	return ds, ws
}
