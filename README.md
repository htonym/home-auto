# Home Automation Project

## Client

## Checklist

- [ ] Enable I2C on the raspberry pi
- [ ] Add the systemd service

### systemd Service

File name: /etc/systemd/system/dht-client.service

Use the IP address is sending from raspberry pi.

```sh
[Unit]
Description=Client to record temperature and humidity
After=network.target

[Service]
Type=simple
User=tony
ExecStart=/home/tony/home-auto-client/client  --room-id=4 --interval=300 --host="http://192.168.4.181:8080"
Restart=on-failure

[Install]
WantedBy=multi-user.target

```

Then run systemd commands:

```sh
sudo chmod 644 /etc/systemd/system/dht-client.service
sudo systemctl daemon-reload
sudo systemctl enable dht-client.service
sudo systemctl start dht-client.service
sudo systemctl status dht-client.service
```
