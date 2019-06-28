This is a tool to get the temperature with a raspberry from an ds18b20 temperature sensor.

# Stuff to do on the raspberry
```
sudo echo dtoverlay=w1-gpio-pullup,gpiopin=4 >> /boot/config.txt
sudo modprobe w1_gpio && sudo modprobe w1_therm
sudo modprobe wire
sudo modprobe w1-gpio
sudo modprobe w1-therm
```
After completing those steps the temperature sensor should be working.

I added the following to the crontab on the raspberry:
```
@reboot sleep 30 && /home/pi/git/go-temp/go-temp
```

# Prometheus
My Prometheus config looks like this:
```
global:
  evaluation_interval: 30s
  scrape_interval: 15s
scrape_configs:
  - job_name: livingroom
    static_configs:
      - targets:
        - "RASPBERRY:8080"
```

# Grafana
After adding Prometheus as a Datasource in Grafana, one can create a graph like the following:
[GRAFANA](pics/grafana.png)
