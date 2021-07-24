package tools

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	cpuUsageCgroup    *os.File
	memMaxUsageCgroup *os.File
	table             *os.File
	ctx               context.Context
	client            *mongo.Client
	collection        *mongo.Collection
)

var cpuUsage = make([]byte, 100)
var maxMemUsage = make([]byte, 100)

const measurementsFolder = "measurements"

func OpenFiles() {
	var err error
	cpuUsageCgroup, err = os.Open("/sys/fs/cgroup/cpuacct/cpuacct.usage") // For read access.
	if err != nil {
		log.Print(err)
		return
	}
	log.Println("/sys/fs/cgroup/cpuacct/cpuacct.usage opened successfuly")

	memMaxUsageCgroup, err = os.Open("/sys/fs/cgroup/memory/memory.max_usage_in_bytes") // For read access.
	if err != nil {
		log.Print(err)
		return
	}
	log.Println("/sys/fs/cgroup/memory/memory.max_usage_in_bytes opened successfuly")
}

func CloseFiles() {
	err := cpuUsageCgroup.Close()
	if err != nil {
		log.Print(err)
	}
	log.Println("/sys/fs/cgroup/cpuacct/cpuacct.usage closed successfuly")

	err = memMaxUsageCgroup.Close()
	if err != nil {
		log.Print(err)
	}
	log.Println("/sys/fs/cgroup/cpuacct/cpuacct.usage closed successfuly")

	tableName := table.Name()
	err = table.Close()
	if err != nil {
		log.Print(err)
	}
	log.Printf("%s closed successfuly\n", tableName)

	if err = client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func ReadCpuUsage() int64 {
	cpuUsageCgroup.Seek(0, 0)
	count, err := cpuUsageCgroup.Read(cpuUsage)
	if err != nil {
		log.Print(err)
		return 0
	}
	if count == 0 {
		log.Println("Can't read cpu.usage")
		return 0
	}

	value, err := strconv.ParseInt(strings.Split(string(cpuUsage), "\n")[0], 10, 64)
	if err != nil {
		log.Print(err)
		return 0
	}
	return value
}

func ReadMaxMemUsage() int64 {
	memMaxUsageCgroup.Seek(0, 0)
	count, err := memMaxUsageCgroup.Read(maxMemUsage)
	if err != nil {
		log.Print(err)
		return 0
	}
	if count == 0 {
		log.Println("Can't read memory.max_usage_in_bytes")
		return 0
	}

	value, err := strconv.ParseInt(strings.Split(string(maxMemUsage), "\n")[0], 10, 64)
	if err != nil {
		log.Print(err)
		return 0
	}
	return value
}

func LoginMongo() {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("mongo_uri")
	if len(mongoURI) == 0 {
		log.Fatalln("Mongo URI isn't provided")
	}

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("faas-measurements").Collection("functions")
}

func WriteMeasurements(startTime time.Time, funcName string, cpuUsage float64, memoryMaxUsage int64) {
	ctxx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctxx, bson.D{{"time", startTime.String()}, {"function_name", funcName}, {"cpu_used", cpuUsage}, {"max_memory", memoryMaxUsage}})
	if err != nil {
		log.Fatalln("Error while inserting data to collection")
	}
}
