# Gourd

Analytics dashboard framework written in Golang and ChartJS

# Problem statement

People need to assemble meaningful visualizations of large, realtime datasets.
This problem is partially addressed by large analytics packages, but few of these provide real-time data displays, limiting the ability to proactively respond to new information. Gourd provides a fast, scalable, modern solution.

# Storyboard

1. Open page, view a list of streaming data source items (clickable)
2. Select the data sources to stream
3. Selected data charts immediately appear and start real-time streaming
4. Settings:
    - scale/zoom (x and y) (FE-only?)
    - display rate (throttle, FE-only?)
    - chart type: line, bar, area, scatter(?), Front End
    - data selections
5. Time-series data, but could be collapsed to show current value only

# Back end TODO

- [ ] Buffer for data channel (last 10 minutes or N data points?)
- [ ] All graph data from server onload (labels, axes, x & y)
- [ ] Multiple simultaneous graphs

- [ ] Example data source: Twitter stream (golang twitterstream client)
- [ ] performance testing:
    - [ ] # points on graph
    - [ ] update rate

- [x] heroku deployment

# Front end TODO

- [ ] Continuous websocket read/timed ui update on client side
  throttle)
- [ ] constant y axis scale
- [ ] Pause/resume
- [ ] Different graph types
- [ ] Gauges?
- [ ] fix animation

# Features

- [ ] pluggable, concurrent data loggers (define API or create admin panel)
- [ ] Analytics: min/max
- [ ] Analytics: mean

# Stack

- Revel framework
- AngularJS
- ChartJS
