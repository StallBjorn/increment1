package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type gauge float64
type counter int64

var httpAddr = flag.String("http", "127.0.0.1:8080", "standart url")

var mem runtime.MemStats
var pollInterval = 2 * time.Second
var reportInterval = 10 * time.Second

//type Metrics struct {
var (
	Alloc         gauge
	BuckHashSys   gauge
	Frees         gauge
	GCCPUFraction gauge
	GCSys         gauge
	HeapAlloc     gauge
	HeapIdle      gauge
	HeapInuse     gauge
	HeapObjects   gauge
	HeapReleased  gauge
	HeapSys       gauge
	LastGC        gauge
	Lookups       gauge
	MCacheInuse   gauge
	MCacheSys     gauge
	MSpanInuse    gauge
	MSpanSys      gauge
	Mallocs       gauge
	NextGC        gauge
	NumForcedGC   gauge
	NumGC         gauge
	OtherSys      gauge
	PauseTotalNs  gauge
	StackInuse    gauge
	StackSys      gauge
	Sys           gauge
	TotalAlloc    gauge
	PollCount     counter
	RandomValue   gauge
)

//}

func main() {
	flag.Parse()
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	i := 0
	for {
		for j := 0; j < 5; j++ {
			timer := time.NewTicker(pollInterval)
			_ = <-timer.C
			runtime.ReadMemStats(&mem)
			fmt.Println("renew")
		}

		//	datapack := Metrics{
		Alloc = gauge(mem.Alloc)
		BuckHashSys = gauge(mem.BuckHashSys)
		Frees = gauge(mem.Frees)
		GCCPUFraction = gauge(mem.GCCPUFraction)
		GCSys = gauge(mem.GCSys)
		HeapAlloc = gauge(mem.HeapAlloc)
		HeapIdle = gauge(mem.HeapIdle)
		HeapInuse = gauge(mem.HeapInuse)
		HeapObjects = gauge(mem.HeapObjects)
		HeapReleased = gauge(mem.HeapReleased)
		HeapSys = gauge(mem.HeapSys)
		LastGC = gauge(mem.LastGC)
		Lookups = gauge(mem.Lookups)
		MCacheInuse = gauge(mem.MCacheInuse)
		MCacheSys = gauge(mem.MCacheSys)
		MSpanInuse = gauge(mem.MSpanInuse)
		MSpanSys = gauge(mem.MSpanSys)
		Mallocs = gauge(mem.Mallocs)
		NextGC = gauge(mem.NextGC)
		NumForcedGC = gauge(mem.NumForcedGC)
		NumGC = gauge(mem.NumGC)
		OtherSys = gauge(mem.OtherSys)
		PauseTotalNs = gauge(mem.PauseTotalNs)
		StackInuse = gauge(mem.StackInuse)
		StackSys = gauge(mem.StackSys)
		Sys = gauge(mem.Sys)
		TotalAlloc = gauge(mem.TotalAlloc)
		//PollCount = PollCount

		i++
		PollCount = counter(i)

		_, err := http.Post(
			"http://"+*httpAddr+"/update/gauge/Allocs/"+fmt.Sprintf("%v", Alloc)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/BuckHashSys/"+fmt.Sprintf("%v", BuckHashSys)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/Frees/"+fmt.Sprintf("%v", Frees)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/GCCPUFraction/"+fmt.Sprintf("%v", GCCPUFraction)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/GCSys/"+fmt.Sprintf("%v", GCSys)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/HeapAlloc/"+fmt.Sprintf("%v", HeapAlloc)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/HeapIdle/"+fmt.Sprintf("%v", HeapIdle)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/HeapInuse/"+fmt.Sprintf("%v", HeapInuse)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/HeapObjects/"+fmt.Sprintf("%v", HeapObjects)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/HeapReleased/"+fmt.Sprintf("%v", HeapReleased)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/HeapSys/"+fmt.Sprintf("%v", HeapSys)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/LastGC/"+fmt.Sprintf("%v", LastGC)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/LastGC/"+fmt.Sprintf("%v", LastGC)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/Lookups/"+fmt.Sprintf("%v", Lookups)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/MCacheInuse/"+fmt.Sprintf("%v", MCacheInuse)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/MCacheSys/"+fmt.Sprintf("%v", MCacheSys)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/MSpanInuse/"+fmt.Sprintf("%v", MSpanInuse)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/MSpanSys/"+fmt.Sprintf("%v", MSpanSys)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/Mallocs/"+fmt.Sprintf("%v", Mallocs)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/NextGC/"+fmt.Sprintf("%v", NextGC)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/NumForcedGC/"+fmt.Sprintf("%v", NumForcedGC)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/NumGC/"+fmt.Sprintf("%v", NumGC)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/OtherSys/"+fmt.Sprintf("%v", OtherSys)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/PauseTotalNs/"+fmt.Sprintf("%v", PauseTotalNs)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/StackInuse/"+fmt.Sprintf("%v", StackInuse)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/StackSys/"+fmt.Sprintf("%v", StackSys)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/NextGC/"+fmt.Sprintf("%v", NextGC)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/Sys/"+fmt.Sprintf("%v", Sys)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/TotalAlloc/"+fmt.Sprintf("%v", TotalAlloc)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/PollCount/"+fmt.Sprintf("%v", PollCount)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(
			"http://"+*httpAddr+"/update/gauge/RandomValue/"+fmt.Sprintf("%v", RandomValue)+"",
			"text/plain",
			nil)
		if err != nil {
			log.Println(err)
		}
		//time.Sleep(reportInterval)
		/*fmt.Println(respalloc)
		fmt.Println("hiall", mem.Alloc)
		fmt.Println("http://" + *httpAddr + "/update/gauge/Allocs/" + strconv.Itoa(int(mem.Alloc)) + "")*/
	}
}
