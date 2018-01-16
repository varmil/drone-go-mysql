[![Build Status](http://beta.drone.io/api/badges/drone-demos/drone-go-mysql/status.svg)](http://beta.drone.io/drone-demos/drone-go-mysql)

Example project to demonstrate unit testing Go code that depends on a Mysql database. This project uses the [Drone](https://github.com/drone/drone) continuous integration server as the testing environment.

### varmil's note
* use `golang:1.8` instead of `golang:1.5`
* use `workspace` in .drone.yml
* use `root@tcp(database:3306)` instead of `root@tcp(127.0.0.1:3306)`
* use `panic()` in constructor of TodoManager
