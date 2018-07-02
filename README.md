# gomets
Simple utility to retrieve AWS instance metadata.

## Usage
```
$ wget http://files.bytester.net/gomets
$ chmod +x gomets
$ ./gomets
{
    "devpayProductCodes": null,
    "availabilityZone": "us-east-2b",
    "privateIp": "172.31.27.63",
    "version": "2017-09-30",
    "region": "us-east-2",
    "instanceId": "i-0cbb1d6622063e586",
    "billingProducts": null,
    "instanceType": "t2.medium",
    "accountId": "512345678967",
    "pendingTime": "2018-07-01T02:52:28Z",
    "imageId": "ami-922914f7",
    "kernelId": "",
    "ramdiskId": "",
    "architecture": "x86_64"
}
```

