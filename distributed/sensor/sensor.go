package main
import (
  "flag"
  "math/rand"
  "time"
  "strconv"
)

var name = flag.String("name", "sensor", "name of the sensor")
var freq = flag.Uint("freq", 5, "update frequency in cycles/sec")
var max = flag.Float64("max", 5., "Maximum value for generated reading")
var min = flag.Float64("min", 1., "minimum value for generated reading")
var stepSize = flag.Float64("step", 0.1, "maximum allowable change per measurment")

var r = rand.New(rand.NewSource(time.Now().UnixNano()))


func main(){
  flag.Parse()
  dur,_ := time.ParseDuration(strconv.Itoa(1000/int(*freq)) + "ms")
  signal := time.Tick(dur)
  for range signal{

  }
}
