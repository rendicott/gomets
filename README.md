# gomets
Simple utility to retrieve AWS instance metadata and other useful info with a oneliner.

## Usage
```
$ wget http://files.bytester.net/gomets
$ chmod +x gomets
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
[ec2-user@ip-192-168-46-110 tmp]$ ./gomets -tags
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

