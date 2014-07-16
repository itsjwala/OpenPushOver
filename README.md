OpenPushOver Client
=====================

[![Build Status](https://drone.io/github.com/TheCreeper/OpenPushOver/status.png)](https://drone.io/github.com/TheCreeper/OpenPushOver/latest)

![ScreenShot](http://firebit.co.uk/imgsrc/06521.png)

- Requirements

    - GNU/Linux only
    - libnotify (GNU/Linux) for notifications

- Features

    - Supports proxys
    - Supports basic end to end encryption

##Sample Config

- CheckFrequencySeconds can not be less than 10 seconds and defaults to 10 seconds if set to anything less.

- Default configuration file is located in the same directory as the exec however can be overridden using the -config flag.

- DeviceUUID should be generated by setting the "Register" flag to true in the config and will be automatically unset afterwards.

- Key should be what you intend to use to receive encrypted messages and should obviously be the same on both ends.

```json
{
    "Globals": {
        "DeviceName": "Fusion",
        "CheckFrequencySeconds": 10
    },
    "Proxys": [
        {
            "Name": "Tor",
            "Type": "socks5",
            "Address": "127.0.0.1:9050",
            "Username": "",
            "Password": "",
            "Timeout": 1
        }
    ],
    "Accounts": [
        {
            "DeviceUUID": "",
            "Register": true,
            "Username": "email",
            "Password": "password",
            "Key": "testkey123456789",
            "Proxy": "Tor"
        }
    ]
}
```