# go-limiter
Go limiter, is used to limit the number of calles on an endpoint, based on a the needed number of requests per second.

## Overview
The case of busy loop on a certin resource or endpoint, should be handled, in this package this is done by letting the user define the number of hits per seconds on an end point, and then the code starts monitoring the resorce, and decide if there is a need to sleep for amount of time.
### Usage

Create an object of the limiter using the following constructor,
it takes ```reqPerSec``` the limit bound requests per seconds :

```
NewLimitWindow(reqPerSec int) *LimitWindow 
```

Check in your code the amount of time the app needes to sleep using the  ```Check()``` function:

```
func (l *LimitWindow) Check() time.Duration
```
```Check()``` is a non-blocking function performs some calculation on the queue that is used to store the request, to make sure that the number of   ```Requests Per Seconds``` didn't exceed the limit, and returns the sleep time.

```CheckWithSleep()``` is a blocking function performs some calculation on the queue that is used to store the request, to make sure that the number of   ```Requests Per Seconds``` didn't exceed the limit, and sleeps the calulcated amount of  time.

## Authors

* **Muhammed Imad** - *Gabriel Systems* 
