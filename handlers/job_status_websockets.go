package handlers

import (
	"Phylogeny/tasks"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func JobStatusWebSocket(JobQueue *tasks.JobQueue) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		defer c.Close()

		jobID := c.Params("jobID")
		if jobID == "" {
			log.Println("Job ID is required")
			return
		}

		// Subscribe to job status updates
		ch := JobQueue.Broadcast.Subscribe(jobID)
		defer JobQueue.Broadcast.Unsubscribe(jobID, ch)

		for {
			select {
			case status, ok := <-ch:
				if !ok {
					return
				}

				// Send the status update to the WebSocket client
				if err := c.WriteMessage(websocket.TextMessage, []byte(status)); err != nil {
					log.Println("WebSocket write error:", err)
					return
				}
			}
		}
	})
}
