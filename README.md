# Bitcoin arbitrage

WORK IN PROGRESS - so it's not working yet.

Hobby project to learn Go and see if bitcoin arbitrage trading is any good.

Feel free to use, but at your own risk :-)

## What is arbitrage trading?
Arbitrage trading ([Wikipedia](https://en.wikipedia.org/wiki/Arbitrage)) is profiting from price differences between markets by simultaneously buying from one market and selling in an other market.

The concept is very simple: say you could buy 1 BTC at Kraken for 100$ and sell 1 BTC at Bitfinex for 110$ at the same time. You'd still have the same number of BTCs, however you're left with 10$ more after the trade (not considering trading fees).

There is a catch though: you couldn't first buy 1 BTC at Kraken, transfer the BTC to Bitfinex, and then sell it again. The transfer is not instant and by the time you transferred the BTC and sell at Bitfinex the arbitrage opportunity is most likely gone. So to make an almost instant trade like in the example, you'd need dollars in your Kraken account and BTCs in your Bitfinex account to make two transfers simultaneously.

### Example calculation
A more precise calculation for the example above:
