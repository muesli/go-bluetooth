# go-bluetooth

Golang bluetooth client based on bluez DBus interfaces

See here for reference https://git.kernel.org/cgit/bluetooth/bluez.git/tree/doc

## Status

The current API is unstable and may change in the future.

The features implemented are

- [x] Discovery
- [x] Adapter support
- [x] Device support (SensorTag example)
- [x] GATT Service and characteristics interface
- [x] Adapter on/off via `rfkill`
- [x] Handle systemd `bluetooth.service` unit
- [x] Expose `hciconfig` basic API
- [x] Expose bluetooth services via bluez GATT API
- [ ] HCI protocol communication
- [ ] Pairing support

## Examples

Check `examples/` folder for an overview of the API

## Setup

The library has been tested with

- golang `1.9` (minimum `v1.7`)
- bluez bluetooth `v5.48` (minimum supported `v5.43`)

### bluez upgrade

Bluez, the linux bluetooth implementation, has introduced GATT support from `v5.43`

Ensure you are using an up to date version with `bluetoothd -v`

See in `scripts/` how to upgrade bluez

### Development notes

-   Give access to `hciconfig` to any user (may have [security implications](https://www.insecure.ws/linux/getcap_setcap.html))

    ```
    sudo setcap 'cap_net_raw,cap_net_admin+eip' `which hciconfig`
    ```
- Create a dbus profile

    ```sh
    ln -s `pwd`/scripts/dbus-dev.conf /etc/dbus1/system.d/go-bluetooth.config
    ```
- Monitor activity

    `sudo dbus-monitor --system "type=error"`

- View `bluetoothd` debug messages

    `sudo bluetoothd -Edn P hostname`

- Enable LE advertisement (to use a single pc, you will need 2 bluetooth adapter)

  ```bash
    sudo btmgmt -i 0 power off
    sudo btmgmt -i 0 name "my go app"
    sudo btmgmt -i 0 le on    
    sudo btmgmt -i 0 connectable on
    sudo btmgmt -i 0 advertising on
    sudo btmgmt -i 0 power on
  ```

## TODO List / Help wanted

-   Add docs with examples
-   Add Device read / write and custom data converters
-   Unit tests coverage
-   Integrate hci communication from `github.com/[currentlabs|go-ble]/ble`

## References

- https://git.kernel.org/cgit/bluetooth/bluez.git/tree/doc
- https://www.bluetooth.com/specifications/gatt/services
- http://events.linuxfoundation.org/sites/events/files/slides/Bluetooth%20on%20Modern%20Linux_0.pdf
- https://github.com/nettlep/gobbledegook

## License

MIT License
