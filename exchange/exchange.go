package exchange

type Exchange interface {
	LatestPrice(ticker string) float32
}
