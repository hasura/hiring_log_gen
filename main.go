package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"

	leakybucket "github.com/ermanimer/design-patterns/leaky-bucket"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var (
	ops   = []string{"insert", "update", "delete", "select"}
	names = []string{"actor", "advertisement", "afternoon", "airport", "akeye", "ambulance", "animal", "answer", "apple", "army", "australia", "balloon", "banana", "battery", "beach", "beard", "bed", "belgium", "boy", "branch", "breakfast", "brother", "camera", "candle", "car", "caravan", "carpet", "cartoon", "china", "church", "crayon", "crowd", "daughter", "death", "denmark", "diamond", "dinner", "disease", "doctor", "dog", "dream", "dress", "e", "easter", "egg", "egganegg", "ening", "eye", "football", "forest", "fountain", "france", "furniture", "garage", "garden", "gas", "ghost", "girl", "glass", "gold", "grass", "greece", "guitar", "hair", "hamburger", "helicopter", "helmet", "holiday", "honey", "horse", "hospital", "house", "hydrogen", "ice", "insect", "insurance", "iron", "island", "jackal", "jelly", "jewellery", "jordan", "juice", "kangaroo", "king", "kitchen", "kite", "knife", "lamp", "lawyer", "leather", "library", "lighter", "lion", "lizard", "lock", "london", "lunch", "machine", "magazine", "magician", "megg", "nail", "nieye", "ntegg", "ocean", "oil", "orange", "oxygen", "oyster", "painting", "parrot", "pencil", "piano", "pillow", "pizza", "planet", "plastic", "portugal", "potato", "queen", "quill", "rain", "rainbow", "raincoat", "refrigerator", "restaurant", "river", "rocket", "room", "rose", "russia", "sandwich", "school", "scooter", "shampoo", "shoe", "soccer", "spoon", "stone", "sugar", "sweden", "teacher", "telephone", "television", "tent", "thailand", "tomato", "toothbrush", "traffic", "train", "truck", "whale", "zebra", "zoo"}
)

func randOp() string {
	op := ops[rand.Intn(len(ops))]
	name := names[rand.Intn(len(names))]
	return op + "_" + name
}

func randf(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randi(min, max int) int {
	return min + rand.Intn(max-min)
}

func hash(s string) int {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		panic(err)
	}
	return int(h.Sum32())
}

// Trace is the data about an operation that's being logged
//
// Example:
//
//{
//  "project_id": "ebe07caf-9595-45c2-a8fc-f51a4f6e6b10",
//  "timestamp": "2022-04-05T15:14:31.136485+02:00",
//  "operation": {
//    "name": "Random",
//    "runtime": 0.01,
//    "request_id": "9d90d2d8-1595-4f4d-9c0e-9606038de8f9",
//    "response_size": 123,
//    "request_size": 124,
//    "http_status": 200
//  }
//}
type Trace struct {
	ProjectID uuid.UUID `json:"project_id"`
	Timestamp time.Time `json:"timestamp"`
	Op        struct {
		Name         string    `json:"name"`
		Runtime      float64   `json:"runtime"`
		RequestID    uuid.UUID `json:"request_id"`
		ResponseSize int       `json:"response_size"`
		RequestSize  int       `json:"request_size"`
		HTTPStatus   int       `json:"http_status"`
	} `json:"operation"`
}

// generator generates log traces, based on its configuration:
//
// * nProjs: number of different projects generating traces. Common values between 1 and 100
// * rate: max log traces that will be generated per second. Common values between 10 and 1000
// * speed: indicates how fast time should pass. For instance:
// 		* if speed was 1, then each trace timestamp will be that of Time.now()
//		* if it was 10, each millisecond a new trace will have a timestamp 10ms newer than the millisecond before.
// 		* if it was 3600, every second a new trace will have a timestamp of one hour later.
type generator struct {
	nProjs int
	rate   int
	speed  int
}

// run generates the log traces using a different goroutine for each project
func (g generator) run(ctx context.Context) {
	start := time.Now()

	pipe := make(chan Trace)
	defer close(pipe)

	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	wg := sync.WaitGroup{}
	wg.Add(g.nProjs)
	for i := 0; i < g.nProjs; i++ {
		go g.genProjectLogs(ctx, start, uuid.New(), pipe, &wg)
	}

	lb := leakybucket.NewLeakyBucket(g.rate)
	defer lb.Stop()
	lb.Start()

	ticker := time.NewTicker(10 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			if lb.IsFull() {
				continue
			}
			select {
			case trace := <-pipe:
				b, err := json.Marshal(trace)
				if err != nil {
					log.Fatalf("Unexpected error logging trace: %stop", b)
				}
				log.Println(string(b))
			default:
			}
		case <-ctx.Done():
			wg.Wait()
			return
		}
	}
}

func (g generator) genProjectLogs(ctx context.Context, start time.Time, pid uuid.UUID, in chan Trace, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case in <- g.genProjectTrace(pid, start):
		case <-ctx.Done():
			return
		}
	}
}

func (g generator) genProjectTrace(pid uuid.UUID, start time.Time) Trace {
	now := time.Now()
	diff := now.Sub(start)
	escaledNow := now.Add(time.Duration(diff.Nanoseconds() * int64(g.speed)))
	opName := randOp()
	outlier := hash(opName)%100 == 0

	runtime := randf(0, 1)
	responseSize := randi(10, 1000)
	requestSize := randi(1, 100)
	status := 200

	if outlier {
		runtime *= 10
		responseSize *= 10
		requestSize *= 10
		if randi(0, 10) > 2 {
			status = randi(500, 503)
		}
	}

	return Trace{
		ProjectID: pid,
		Timestamp: escaledNow,
		Op: struct {
			Name         string    `json:"name"`
			Runtime      float64   `json:"runtime"`
			RequestID    uuid.UUID `json:"request_id"`
			ResponseSize int       `json:"response_size"`
			RequestSize  int       `json:"request_size"`
			HTTPStatus   int       `json:"http_status"`
		}{
			Name:         randOp(),
			Runtime:      runtime,
			RequestID:    uuid.New(),
			ResponseSize: requestSize,
			RequestSize:  responseSize,
			HTTPStatus:   status,
		},
	}
}

func main() {
	nProjs := flag.Uint("projects", 25, "number of projects. Valid between 1 and 1000.")
	rate := flag.Uint("rate", 1000, "traces to emit per second, accross all projects. Valid between 1 and 10000")
	speed := flag.Uint("speed", 360, "relative speed of time. Valid between 1 -realtime- and 86400 -each second counts for a day-.")
	flag.Parse()

	var errs []error
	if *nProjs < 1 || *nProjs > 1000 {
		errs = append(errs, fmt.Errorf("wrong number of projects: %d", *nProjs))
	}
	if *rate < 1 || *rate > 10000 {
		errs = append(errs, fmt.Errorf("wrong rate: %d", *rate))
	}
	if *speed < 1 || *speed > 86400 {
		errs = append(errs, fmt.Errorf("wrong speed: %d", *speed))
	}
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e.Error())
		}
		flag.Usage()
		os.Exit(1)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		select {
		case <-c:
			cancelFunc()
		}
	}()

	g := generator{
		nProjs: int(*nProjs),
		rate:   int(*rate),
		speed:  int(*speed),
	}
	g.run(ctx)
}
