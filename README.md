# newrelic-murmur

NewRelic instrumentation for Murmur.

## Synopsis

```shell
newrelic-murmur -h

Usage of newrelic-murmur:
  -host string
    	Murmur host (default "localhost")
  -interval int
    	Poll interval (seconds) (default 60)
  -license string
    	New Relic license key
  -port int
    	Murmur port (default 64738)
  -timeout int
    	Timeout (milliseconds) (default 1000)
  -verbose
    	Verbose
```

## Metrics

 - Maximum Bitrate
 - Maximum Users
 - Connected Users
 - Total Bandwidth(Estimated): Maximum Bitrate * Connected Users

## Environment variable

 - NEW_RELIC_LICENSE_KEY : New Relic license key. ```-license``` flag override it.

## Example

```shell
newrelic-murmur -license 0123456789abcdef0123456789abcdef01234567
```
