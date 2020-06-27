<div align="center">
<h2>OpenPushOver</h2> 
<p>multiplatform Pushover alerts notifier client</p>
</div>

<div align="center">
<img src="https://pushover.net/images/icon-72.png" alt="Pushover logo">
<img src = "https://image.flaticon.com/icons/png/32/14/14980.png" style="position:relative; bottom:1rem; left:0.5rem">
<img width="80" height="80" src="https://cdn.worldvectorlogo.com/logos/golang-gopher.svg" alt="golang gopher">
</div>


---

- Requirements

    for Linux

    - GNU/Linux only
    - libnotify (GNU/Linux) for notifications

- Features

    - Supports proxys
    - Supports basic end to end encryption
    - Supports multiple pushover accounts

## Sample Config
- You need to create the cache directory

- CheckFrequencySeconds can not be less than 5 seconds and defaults to 5 seconds if set to anything less.

- Default configuration file is located in the same directory as the exec however can be overridden using the -config flag.

- DeviceUUID should be generated by setting the "Register" flag to true in the config and will be automatically unset afterwards. The "Force" flag should be set to true if you intend to replace an already existing device.

- Key should be what you intend to use to receive encrypted messages and should obviously be the same on both ends.

```json
{
    "Globals": {
        "CacheDir" : "./cache",
        "DeviceName": "Fusion",
        "CheckFrequencySeconds": 5
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
            "Force" : false,
            "Username": "email",
            "Password": "password",
            "Key": "testkey123456789",
            "Proxy": "Tor"
        }
    ]
}
```


> This is fork of `TheCreeper/OpenPushOver` with some modifications
