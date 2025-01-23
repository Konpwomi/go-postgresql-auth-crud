package handlers

import (
    "strconv"
    "github.com/gofiber/fiber/v2"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var tasks = []Task{
	{Id: 1, Title: "Task 1", Description: "First task description", Status: "pending"},
	{Id: 2, Title: "Task 2", Description: "Second task description", Status: "completed"},
	{Id: 3, Title: "Task 3", Description: "Third task description", Status: "pending"},
}

func GetTasks(c *fiber.Ctx) error {
	return c.JSON(tasks)
}

func GetTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, task := range tasks {
		if task.Id == id {
			return c.JSON(task)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Task not found")
}

func CreateTask(c *fiber.Ctx) error {
	var newTask Task
	if err := c.BodyParser(&newTask); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid data")
	}
	newTask.Id = len(tasks) + 1
	tasks = append(tasks, newTask)
	return c.Status(fiber.StatusCreated).JSON(newTask)
}

func UpdateTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	var updatedTask Task
	if err := c.BodyParser(&updatedTask); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid data")
	}
	for index, task := range tasks {
		if task.Id == id {
			tasks[index].Title = updatedTask.Title
			tasks[index].Description = updatedTask.Description
			tasks[index].Status = updatedTask.Status
			return c.JSON(tasks[index])
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Task not found")
} 

func DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for index, task := range tasks {
		if task.Id == id {
            tasks = append(tasks[:index], tasks[index+1:]...)
            return c.SendString("Task successfully deleted")
        }
    }
    return c.Status(fiber.StatusNotFound).SendString("Task not found")
}