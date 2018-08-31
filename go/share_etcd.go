// Service finding:
// Service Requestor
const (
  services = "my_services/"
  webCrawlerService = "web_crawler"
  machineLearningService = "machine_learning"
  persistenceService = "persistence"
)
func findServicesAndBind() error {
  items, err := etcd.Get(services, etcd.WithPrefix)
  if err != nil { return err }
  for _, item := range items {
    bindService(item.Key, item.Value)
  }
  return nil
}
func watchServices() {
  // watch all services, return event channel
  ch := etcd.Watch(services, etcd.WithPrefix)
  // wait for new event channel in a go-routine
  go func() {
    for evt := range ch {
      // get service when each event arrives
      service := evt.Key[len(services):]
      switch evt.Type {
      case etcd.PutEvent: // service added/updated
        bindService(service, evt.Value)
      case etcd.DeleteEvent: // service deleted
        unBindSerivce(service)
      }
    }
  }()
}
func bindService(service, data string) {
  switch service {
  case webCrawlerService:
    bindCrawlerService(data)
  case machineLearningService:
    bindMachineLearningService(data)
  case persistenceService:
    bindPersistenceService(data)
  }
}
func main() {
  // find existed services and bind
  findServicesAndBind()
  // watch services for future update/delete
  watchServices()
  // other code ...
}

// Service Provider -- Web Crawler Service
func main() {
  // register service to etcd with heart beat
  heartBeat := 5 * time.Second
  leaseID, cancel, err := etcd.Timeout(heartBeat)
  etcd.Put("my_services/web_crawler", 
    "172.10.2.34:5678", etcd.WithLease(leaseID))
  go func() {
    tick := time.Tick(heartBeat / 2)
    for _ := range tick {
      // extend the life inside the heart beat
      // if crash, etcd will remove the node
      etcd.KeepAliveOnce(leaseID)
    }
  }()
  // other code ...
}

// Publish
// register topic with last queue ID
maxID :=
etcd.Put("topic/talk", maxID)

// publish a message to the topic
locker := etcd.Lock("lock/topic/talk")
lastMsgID, err := etcd.Get("topic/talk")
lastID, err := strconv.Atoi(lastMsgID)
if err != nil {
  lastID = 0
}
lastID++
etcd.Put(fmt.Sprintf("topic/talk/%05d", lastID), "hello")
etcd.Put("topic/talk", strconv.Itoa(lastID))
locker.Unlock()

// Subscription
topics, err := etcd.Get("topic/", etcd.WithPrefix)
for _, topic := range topics {
  // consume logic ...
  etcd.Delete(topic.Key) // delete after consume
}

ch := etcd.Watch("topic/talk/", etcd.WithPrefix)
go func() {
  for evt := range ch {
      // get service when each event arrives
      msgID := evt.Key[len("topic/talk/"):]
      if evt.Type == etcd.PutEvent {
        consume(msgID, evt.Value) // consume message
        etcd.Delete(evt.Key) // delete after consume
      }
    }
}()


type Server struct {
  ID string
  WorkLoad int
}
func main() {
  go healthService()
  // other code ...
}
func healthService() {
  heartBeat := 5 * time.Second
  leaseID, cancel, err := etcd.Timeout(heartBeat)
  server := Server{
    ID: "web_crawler/8",
    WorkLoad: calculateWorkLoad(),
  }
  etcd.Put(server.ID, json.Marshal(server), 
    etcd.WithLease(leaseID))
  tick := time.Tick(heartBeat / 2)
  for _ := range tick {
    etcd.KeepAliveOnce(leaseID)
    // update workload periodically
    server.WorkLoad = calculateWorkLoad()
    etcd.Put(server.ID, json.Marshal(server), 
      etcd.WithLease(leaseID))
  }
}

// version 1: share memory with lock
func RequestMinWorkLoadServer() {
  server := FindMinWorkLoadServer()
  request(server)
}
minWorkLoadServer := Server{
  WorkLoad: MaxInt32,
}
var locker sync.Mutex
func FindMinWorkLoadServer() Server {
  locker.Lock()
  defer locker.Unlock()
  return minWorkLoadServer
}
func watchServers() {
  // watch all web crawler servers, return event channel
  ch := etcd.Watch("web_crawler/", etcd.WithPrefix)
  go func() {
    for evt := range ch {
      switch evt.Type {
      case etcd.PutEvent: // server added/updated
        var server Server
        err := json.Unmarshal(evt.Value, &server)
        // prepare min work load each
        locker.Lock()
        if server.WorkLoad < minWorkLoadServer.WorkLoad {
          minWorkLoadServer = server
        }
        locker.Unlock()
      case etcd.DeleteEvent: // server deleted
        unBindSerivce(evt.Key)
      }
    }
  }()
}

// version 2: communication with channel instead of sharing memory (without lock)


func RequestMinWorkLoadServer() {
  server := FindMinWorkLoadServer()
  request(server)
}
minWorkLoadServerChan := make(chan Server)
func FindMinWorkLoadServer() Server {
  return <- minWorkLoadServerChan
}
func minWorkLoadService(serverChan <-chan Server, 
             minWorkLoadServerChan chan<- Server) {
  minWorkLoadServer := Server{
    WorkLoad: MaxInt32,
  }
  for {
    select {
    case minWorkLoadServerChan <- minWorkLoadServer:
    case server := <- serverChan:
      if server.WorkLoad < minWorkLoadServer.WorkLoad {
        minWorkLoadServer = server
      } 
    }
  }
}
func watchServers() {
  chServer := make(chan Server)
  // prepare min work load each time when server updated
  go minWorkLoadService(chServer, minWorkLoadServerChan)
  // watch all web crawler servers, return event channel
  ch := etcd.Watch("web_crawler/", etcd.WithPrefix)
  go func() {
    for evt := range ch {
      switch evt.Type {
      case etcd.PutEvent: // server added/updated
        var server Server
        err := json.Unmarshal(evt.Value, &server)
        chServer <- server
      case etcd.DeleteEvent: // server deleted
        unBindSerivce(evt.Key)
      }
    }
  }()
}


etcd.Put("job/1/status", "Start")


ch := etcd.Watch("job/1/status")
for evt := range ch {
  switch evt.Type {
  case etcd.PutEvent:
    switch evt.Value {
    case "Start":
      // handle on start ...
    case "Done":
      // handle on done ...
    }
  }
}
etcd.Put("job/1/status", "Done")


locker := etcd.Lock("/lock/job/1")
// do something
locker.Unlock()

locker, err := etcd.LockWithTimeout(
  "/lock/job/1", context.Timeout(time.Second))
if err == nil {
  ...
} else {
  ...
}

// atomic if else
if no lock {
  put lock
} else {
  wait for unlock
}

// create etcd client
cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
// create etcd concurrency session
s1, err := concurrency.NewSession(cli)
// create election
e1 := concurrency.NewElection(s1, "/my-election/")
// election campaign, if 
err := e1.Campaign(context.Background(), "e1")

// config management
etcd.Put("/etc/max_timeout", []byte("10"))

// get config
maxTimeout, err := etcd.Get("/etc/max_timeout")

// watch config
ch := etcd.Watch("/etc/max_timeout")
for evt := range ch {
  maxTimeout := evt.Value
}