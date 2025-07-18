package tasks

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/config"
	"github.com/RichardKnop/machinery/v2/tasks"

	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	"github.com/RichardKnop/machinery/v2/backends/result"
	amqpbroker "github.com/RichardKnop/machinery/v2/brokers/amqp"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
)

func StartServer() (*machinery.Server, error) {
	cnf := &config.Config{
		Broker:          "amqp://guest:guest@rabbitmq:5672/",
		DefaultQueue:    "machinery_tasks",
		ResultBackend:   "redis://redis:6379",
		ResultsExpireIn: 3600,
		AMQP: &config.AMQPConfig{
			Exchange:      "machinery_exchange",
			ExchangeType:  "direct",
			BindingKey:    "machinery_tasks",
			PrefetchCount: 3,
		},
	}

	// Create server instance
	broker := amqpbroker.New(cnf)
	backend := redisbackend.NewGR(cnf, []string{"redis:6379"}, 0)
	lock := eagerlock.New()
	server := machinery.NewServer(cnf, broker, backend, lock)

	// Register tasks
	tasksMap := map[string]interface{}{
		"add": Add,
	}

	return server, server.RegisterTasks(tasksMap)
}

func SendAddTask() (*result.AsyncResult, error) {
	server, err := StartServer()
	if err != nil {
		return nil, err
	}

	signature := &tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 5,
			},
		},
	}

	return server.SendTask(signature)
}
