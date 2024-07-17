# anti-charts
Get the info for different crypto pairs from binance and perform analysis on it. You can also optionally make trades on Alpaca.

## Installation
1. Clone the project
```
$ git clone https://github.com/BrianMwangi21/anti-charts.git
```

2. Move into the directory, copy the .env.example to .env
```
$ cd anti-charts
$ cp .env.example .env
```

3. Fill the necessary fields in the .env file
```
BINANCE_API_KEY=
BINANCE_SECRET_KEY=
BINANCE_TESTNET_KEY=
BINANCE_TESTNET_SECRET_KEY=
ALPACA_API_KEY=
ALPACA_SECRET_KEY=
ALPACA_BASE_URL=
PERFORM_TRADES={True or False}
SPECIAL_CASES={True or False}
```

4. Build the project
```
$ make build
```

5. Run the project
```
$ make run
```
