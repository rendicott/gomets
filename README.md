# gomets
Simple utility to retrieve AWS instance metadata and other useful info with a oneliner.

## Installation
Grab a release from the releases section and do a wget or something like:

```
wget https://github.com/rendicott/gomets/releases/download/v0.2/gomets
chmod +x gomets
```

## Usage
```
$ ./gomets
{
    "Metadata": {
        "devpayProductCodes": null,
        "availabilityZone": "us-east-1a",
        "privateIp": "`192.168.46.110",
        "version": "2017-09-30",
        "region": "us-east-1",
        "instanceId": "i-096dc6418ba3aabc2",
        "billingProducts": null,
        "instanceType": "t2.micro",
        "accountId": "123456789012",
        "pendingTime": "2018-11-01T02:59:16Z",
        "imageId": "ami-e5d3dd9a",
        "kernelId": "",
        "ramdiskId": "",
        "architecture": "x86_64"
    },
    "Tags": []
}
```

If you add the `-tags` flag it will attempt to do an EC2 DescribeTags on its own instance-id.
```
$ ./gomets -tags
{
    "Metadata": {
        "devpayProductCodes": null,
        "availabilityZone": "us-east-1a",
        "privateIp": "192.168.46.110",
        "version": "2017-09-30",
        "region": "us-east-1",
        "instanceId": "i-096dc6418ba3aabc2",
        "billingProducts": null,
        "instanceType": "t2.micro",
        "accountId": "123456789012",
        "pendingTime": "2018-11-01T02:59:16Z",
        "imageId": "ami-e5d3dd9a",
        "kernelId": "",
        "ramdiskId": "",
        "architecture": "x86_64"
    },
    "Tags": [
        {
            "Key": "Name",
            "Value": "ansible-tester"
        },
        {
            "Key": "env",
            "Value": "dev"
        },
        {
            "Key": "appId",
            "Value": "9999999"
        }
    ]
}
```

Then you can do things like grab the instance `Name` tag with a `jq` oneliner:

```
$ ./gomets -tags | jq '.Tags[] | select(.Key=="Name") | .Value'
"ansible-tester"
```