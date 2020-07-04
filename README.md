# IEX CLI

This is a command line interface for the [Investors Exchange API](https://iextrading.com/developer/docs/) written in Go. The IEX API returns data related to the IEX stock exchange.

## Installation
Download the binary [here](https://github.com/covertbert/iex-cli/tags) and make sure it is executable and in your `$PATH`.

## Usage examples
### Chart data
View chart data for a specified stock over a specified period of time. e.g. `iex-cli chart --range "1y" AAPL`.

![Chart Screenshot](docs/screenshots/ohlc-chart.png?raw=true "Chart Screenshot")

### Key Statistics
View key statistics for a stock with `iex-cli stats AAPL`.

![Chart Screenshot](docs/screenshots/stats.png?raw=true "Chart Screenshot")

### News
View general news with `iex-cli news` or stock-specific news with `iex-cli news AAPL`.

![Chart Screenshot](docs/screenshots/news.png?raw=true "Chart Screenshot")

To see a more detailed list of commands, run `iex-cli`.
