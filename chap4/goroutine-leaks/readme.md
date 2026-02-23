### Goroutine's paths of termination

1. it has completed its work
2. it cannot continue its work due to an unrecoverable error
3. it's told to stop working

### Rule

Parent goroutine is responsible for _creating_ and _stopping_ child goroutine
And _stopping_ achieved by sending signal to child via read-only channel (done channel)
