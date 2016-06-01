# data-gov-sg-graphql-go
A GraphQL endpoint for https://data.gov.sg real-time APIs (Go/Golang)

Live demo available here: http://datagovsggraphqlgo-sogko.rhcloud.com/

![http://i.giphy.com/l0K49DoAtOPHpqNFK.gif](http://i.giphy.com/l0K49DoAtOPHpqNFK.gif)

## Implemented APIs
Original APIs reference available here: https://developers.data.gov.sg/datagovsg-apis/apis

__Environment__
- [x] https://api.data.gov.sg/v1/environment/2-hour-weather-forecast
- [x] https://api.data.gov.sg/v1/environment/24-hour-weather-forecast
- [x] https://api.data.gov.sg/v1/environment/4-day-weather-forecast
- [x] https://api.data.gov.sg/v1/environment/pm25
- [x] https://api.data.gov.sg/v1/environment/psi
- [x] https://api.data.gov.sg/v1/environment/uv-index

__Transport__
- [x] https://api.data.gov.sg/v1/transport/taxi-availability
- [x] https://api.data.gov.sg/v1/transport/traffic-images

## Motivation
- Something to demonstrate how `graphql-go` resolve fields concurrently.
- One approach to use GraphQL for existing REST(-ish?) APIs
- Just because.

## Notes
- `graphql-go` based on an [experimental branch](https://github.com/sogko/graphql/tree/sogko/experiment-parallel-resolve) that resolves fields concurrently. (The OpenShift deployment uses vendoring to support it)
- Written a quick HTTP client for `data.gov.sg` API that coalesces identical API requests into one single request. Yay go-routines and go-channels.
- Implemented GeoJSON GraphQL schema defined here https://github.com/sogko/graphql-schemas/tree/master/geojson
- The deployed GraphQL endpoint has a considerable latency due to the datacenter being in US-East region and the Data.gov.sg API being in Singapore.

# TODO
- [ ] Better documentation for GraphQL definitions
- [ ] Better TODO

# Contribute
- If you want to build/ is currently building / have already built something cool with this, let me know :)

# License 
Original API and retrieved dataset(s) are the property of [Singapore Government and its Statutory Boards.](https://developers.data.gov.sg/terms-use)

Everything else: MIT
