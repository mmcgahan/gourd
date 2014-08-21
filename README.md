# Gourd

Analytics dashboard framework written in golang and AngularJS

# Todo:
- Restore buffer for data stream
- Continuous read, timed write/ui-update on client side (JS setTimeout for throttle)
- fixed y axis scale
- Multiple simultaneous graphs
- All data from server (labels, axes, x & y)
- Pause/resume

- Twitter stream (twitterstream)
- Different graph types
- Gauges?
- fix animation
- performance testing:
    - # points on graph
    - update rate

- heroku deployment


# Features

- pluggable, concurrent data loggers
- custom data processors

- admin panel for visualizations

# Stack

- Nginx reverse proxy to static and go server
- Revel framework
- AngularJS
- ChartJS
