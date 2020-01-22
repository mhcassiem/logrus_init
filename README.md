# logrus_init
project to hold initialization code for logrus


## logger_init
Contains function Init_logger which returns a logrus logger.

Example use case:

```
package main

import (
	"github.com/mhcassiem/logrus_init/logger_init"
)

func main()  {
	log := logger_init.Init_logger()
	log.Infoln("Hello test")
}
```

Dependencies: 

```
github.com/x-cray/logrus-prefixed-formatter
github.com/sirupsen/logrus
```