package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"oms/pkg/db"
	"oms/pkg/domain"
	grpcPkg "oms/pkg/grpc"
	"oms/pkg/grpchandler"

	"github.com/joho/godotenv"
	strToDuration "github.com/xhit/go-str2duration/v2"
)

var (
	// Derived from ldflags -X
	buildRevision string
	buildVersion  string
	buildTime     string
	// general options
	versionFlag bool
	helpFlag    bool
	// timeout flag
	timeout int
	// grpc server port
	grpcServerPort string
	postgresURI    string
	// retryInterval for db connection
	retryInterval string
	// program controller
	done    = make(chan struct{})
	errGrpc = make(chan error)
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}
	flag.BoolVar(&versionFlag, "version", false, "show current version and exit")
	flag.BoolVar(&helpFlag, "help", false, "show usage and exit")
	flag.StringVar(&grpcServerPort, "grpcServerPort", ":5060", "grpc server port")
	flag.IntVar(&timeout, "timeout", 1, "timeout for db connection")
	postgresURI = os.Getenv("POSTGRESQL_DB_URL")
	if postgresURI == "" {
		log.Fatal("postgres db url not found")
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
	retryIntervalDuration, err := strToDuration.ParseDuration(retryInterval)
	if err != nil {
		log.Fatal("unable to parse string to duration : ", err)
	}
	postgreSQL := db.Connect2DB(postgresURI, retryIntervalDuration)
	defer postgreSQL.Close()

	d := domain.NewOrderCliet(postgreSQL, timeout)

	grpcHandler := grpchandler.NewGrpcHandler(d)
	grpcServer := grpcPkg.NewGrpcServer(grpcServerPort, grpcHandler)
	go func() {
		log.Printf("GRPC sever running on port: %v\n", grpcServerPort)
		errGrpc <- grpcServer.ListenAndServe()
	}()
	select {
	case err := <-errGrpc:
		log.Print("Grpc error", err)
	case <-done:
		log.Println("shutting down server ...")
	}
	time.AfterFunc(1*time.Second, func() {
		close(done)
		close(errGrpc)
	})
}
