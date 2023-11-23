package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"ums/pkg/db"
	"ums/pkg/domain"
	"ums/pkg/handlers"
	"ums/pkg/helper"
	"ums/pkg/routes"
	"ums/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	strToDuration "github.com/xhit/go-str2duration/v2"
	pb "github.myproto.com"

	"google.golang.org/grpc"
)

var (
	// Derived from ldflags -X
	buildRevision string
	buildVersion  string
	buildTime     string
	// general options
	versionFlag bool
	helpFlag    bool
	// rest server port
	restServerPort string
	// timeout flag
	timeout int
	// order service url
	orderServiceURL string
	// retryInterval for db connection
	retryInterval string
	// program controller
	done        = make(chan struct{})
	errRest     = make(chan error)
	postgresURI string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}
	flag.BoolVar(&versionFlag, "version", false, "show current version and exit")
	flag.BoolVar(&helpFlag, "help", false, "show usage and exit")
	flag.StringVar(&restServerPort, "port", ":8060", "rest server port")
	flag.IntVar(&timeout, "timeout", 1, "timeout for db connection")
	orderServiceURL = os.Getenv("ORDER_SERVICE_URL")
	if orderServiceURL == "" {
		log.Fatal("token service url not found")
	}

	if retryInterval = os.Getenv("RETRY_INTERVAL"); len(retryInterval) == 0 {
		log.Println("retry interval for db connection not found !")
		return
	}
}

func setBuildVariables() {
	if buildRevision == "" {
		buildRevision = "dev"
	}
	if buildVersion == "" {
		buildVersion = "dev"
	}
	if buildTime == "" {
		buildTime = time.Now().UTC().Format(time.RFC3339)
	}
}

func parseFlags() {
	flag.Parse()
	if helpFlag {
		flag.Usage()
		os.Exit(0)
	}
	if versionFlag {
		log.Printf("%s %s %s\n", buildRevision, buildVersion, buildTime)
		os.Exit(0)
	}
}

func handleInterrupts() {
	log.Println("start handle interrupts")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	sig := <-interrupt
	log.Printf("caught sig: %v", sig)
	// close resource here
	done <- struct{}{}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	setBuildVariables()
	parseFlags()
	go handleInterrupts()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	restServer := gin.Default()

	retryIntervalDuration, err := strToDuration.ParseDuration(retryInterval)
	if err != nil {
		log.Fatal("unable to parse string to duration : ", err)
	}
	postgreSQL := db.Connect2DB(postgresURI, retryIntervalDuration)
	defer postgreSQL.Close()
	// Order Service connection
	orderServiceConnection, err := grpc.Dial(orderServiceURL, grpc.WithInsecure())
	if err != nil {
		log.Println("failed to start order service : ", err)
		os.Exit(1)
		return
	}
	orderClient := pb.NewOrderManagementClient(orderServiceConnection)
	orderService := service.NewOrderServiceClient(orderClient)
	s := service.NewService(orderService)
	d := domain.NewUserCliet(postgreSQL, timeout)
	hlpr := helper.NewUserServiceHelper()

	h := handlers.NewUserHandler(d, hlpr, s)

	r := routes.NewRoutes(h)
	routes.AttachRoutes(restServer, r)
	go func() {
		errRest <- restServer.Run(restServerPort)
	}()

	select {
	case err := <-errRest:
		log.Printf("ListenAndServe error: %v", err)
	case <-done:
		log.Println("shutting down server ...")
	}
	time.AfterFunc(1*time.Second, func() {
		close(done)
		close(errRest)
	})
}
