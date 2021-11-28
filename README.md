# PortaDisk - Cheap Raspberry Pi Portable & Secure NAS Project

**Project Status:** Early work in progress

Yet another Raspberry Pi project, a concept for now (until I get all the needed parts). This repository mainly focuses on the software side.

## My personal feature wishlist

- Raspberry Pi Zero running Seafile on Docker.
- Disk and RAM usage reported using [Blinkt](https://thepihut.com/products/blinkt).
- Highest data security with LUKS encryption on the data drive.
- Weekly backups of both the data drive and root partition to a separate hard drive.
- A separate WiFi network for access outside your home, without proxy services.
- Data & backup drive unlocking through a web browser.
- All of this in a small package that can be carried in a backpack.

## Hardware

- Raspberry Pi Zero with a [USB hub HAT](https://www.waveshare.com/usb-hub-hat.htm).
- Another USB hub that has support for external power.
- Pi & the 2 hard drives connected to the second USB hub, which is connected to the HAT, which is connected to the Pi (hopefully this will work)
- Pi, USB hub, hard drives and a power supply crammed into a camera case

## Concept security considerations

- Modification of the `web-unlock` binary to intercept the password during unlocking might be possible
- HTTPS encryption should be a must when dealing with projects like this (caddy?)

## Installation guide

This guide assumes you have a fresh Raspbian Lite installation already set up with updates installed.

### Docker

```sh
curl -sSL https://get.docker.com | sh
sudo apt-get install -y python3 python3-pip libffi-dev libssl-dev python3-dev
sudo pip3 install docker-compose
sudo systemctl enable docker --now
```

### Seafile

1. Adjust the values in `software/seafile/docker-compose.yml` to your liking.
1. Run `docker-compose up`