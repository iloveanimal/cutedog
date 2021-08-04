# dogiscute


## example

```sh

package main

import (
	"time"

	dogErr "github.com/iloveanimal/cutedog/errors"
	dogLog "github.com/iloveanimal/cutedog/log"
)

func main() {

	logChan := make(chan dogLog.LogData)

	go dogLog.LogEngine(logChan)

	eh := dogErr.ErrorHandler{
		PushChan: logChan,
	}
	for {
		err := dogErr.GetGeneralError("someWhere", "test Error")
		eh.Handle(err)
		time.Sleep(time.Second)
	}
}
```


## to cloud example

```sh

package main

import (
	"time"

	dogErr "github.com/iloveanimal/cutedog/errors"
	dogLog "github.com/iloveanimal/cutedog/log"
)

func main() {

	logChan := make(chan dogLog.LogData)

	go dogLog.LogGCloudEngine(logChan)

	eh := dogErr.ErrorHandler{
		PushChan: logChan,
	}
	for {
		err := dogErr.GetGeneralError("someWhere", "test Error")
		eh.Handle(err)
		time.Sleep(time.Second)
	}
}
```