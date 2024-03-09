[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-orange.svg)](https://sonarcloud.io/summary/new_code?id=rodrigobsimon2_togodo)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white)](https://conventionalcommits.org)

# 🚀 Togodo - Go TODO App

Togodo is a sleek and efficient TODO app written in Go, designed to empower you in managing your tasks effortlessly. It offers essential features for creating and listing tasks, each with distinct priorities.

## 🌟 Features

- ✅ Create tasks with a description and priority.
- 📋 List tasks based on priority.
- ⌨️ Easy-to-use command-line interface.

## 🚦 Getting Started

Get up and running with Togodo in just a few steps:

### 🛠️ Prerequisites

- [Go](https://golang.org/) installed on your machine.
- [Docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/) for running the app with a MySQL database.

### 🚀 Installation

1. **Clone the repository:**

   `$ git clone https://github.com/your-username/togodo.git`

2. **Navigate to the project directory:**

    `$ cd togodo`

3. **Run the Docker containers:**

    `$ docker-compose up -d`

# 💻 Usage

Run the Togodo app with the following command:

`$ go run cmd/togodo-cli/main.go -create -description "Your task description" -priority 2`

Or you can use the docker image:

`$ docker pull rodrigobsimon/togodo:latest`

`$ docker run --rm --name togodo-cli --network togodo-cli-network -it rodrigobsimon/togodo-cli -create -description="Walk Cat" -priority=0 -list`

Perhaps you want to build locally:

`$ sh scripts/build.sh`

Use the -list flag to view the list of tasks:

`$ go run cmd/togodo-cli/main.go -list`

# 🎏 Flags
- `-create`: Create a new task.
- `-list`: List all tasks.
- `-description`: Specify the task description.
- `-priority`: Set the task priority (0 - Not important, No urgent; 1 - Not important, Urgent; 2 - Important, Not urgent; 3 - Important, Urgent).

# 🤝 Contributing

Feel free to contribute to Togodo by opening issues or submitting pull requests. Your feedback and suggestions are highly appreciated! 🌈

# 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.