

# API list


|idx|url |example|note|
|---|----|---------|--------|
|1  |https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest| curl -H "X-CMC_PRO_API_KEY: b5cd3d03-b51f-434a-a12a-0bdc68383c5e" -H "Accept: application/json" -G https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest | btc banner|
|2  |https://blockchain.info/charts/avg-block-size?timespan=1days&format=json|Curl -H “Accept: application/json” -G https://blockchain.info/charts/avg-block-size?timespan=1days&format=json||
|3  |https://blockchain.info/rawblock/$blockhash| curl https://blockchain.info/rawblock/000000000000000000194e816bc839bf0e9b0a8338fffed8787a01136b582b95||


## BtcBanner

1. https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest

|source data fields| our page fields  |  note |
|------------------|------------------|-------|
|price             |market price      |       |
|volume_24h        |volume            |       |
|market_cap        |market cap        |       |
|circulating_supply|circulating supply|       |
|max_supply        |max supply        |       |


2. https://blockchain.info/charts/avg-block-size?timespan=1days&format=json

|source data fields| our page fields  |  note |
|------------------|------------------|-------|
|blocksize per day |blocksize per day |       |

3. https://blockchain.info/rawblock/$blockhash


|source data fields| our page fields  |  note |
|------------------|------------------|-------|
|blocksize per day |height            |       |
|blocksize per day |mined time        |       |
|blocksize per day |miner             |       |
|blocksize per day |block hash        |       |
|blocksize per day |block size        |       |
