# <code>roadmap.sh</code> backend Projects

## Projects

### Project 1 - Number guessing Game

Number guessing Game.
The game is a simple command-line game where the user has to guess a number between 1 and 100. The game will give hints to the user if the guessed number is too high or too low. The game will also keep track of the number of guesses the user has made.
You can view how to play the game by reading the [Guessing game README.md](https://github.com/Rioba-Ian/roadmap-sh-backend/tree/main/go-number-guessing) file.

The project can be found on [roadmap.sh](https://roadmap.sh/projects/number-guessing-game)

##### Installation

You can install and run the Number Guessing Game using one of these methods:

**Method 1: Using Git and Go**

1. Clone the repository:

   ```bash
   git clone https://github.com/Rioba-Ian/roadmap-sh-backend.git
   cd roadmap-sh-backend/go-number-guessing
   ```

2. Build and run the game:
   ```bash
   go build -o guess-game
   ./guess-game
   ```

**Method 2: Using Go install**

If you have Go installed, you can directly install and run the game:

```bash
go install github.com/Rioba-Ian/go-number-guessing@latest
go-number-guessing
```

**Command Line Flags**

The game supports several command line flags:

- `--low`: Set the lowest number in the range (default: 1)
- `--high`: Set the highest number in the range (default: 100)
- `--time-limit`: Set the time limit in seconds (default: 15)

Example:

```bash
./guess-game --low=10 --high=50 --time-limit=30
```

### Project 2 - Weather API service

Weather API service.
The service is a simple REST API that provides weather information for a given location. The service will use the VisualCrossingWebServices API to fetch weather data. We use a cache to store the weather data for a given location. The key for the cache is the location name. The expiry time is set to 12 hours. The service looks as follows:

![](./weather-api-service/docs/weather-api-f8i1q.png)

#### Installation

You can install and run everything locally but first you need to install Docker and Go.

After installing Docker and Go, you can run the service using the following command:

You can then clone the repository and run the service using the following command:

```bash
git clone https://github.com/Rioba-Ian/roadmap-sh-backend.git
cd roadmap-sh-backend/weather-api-service
go mod download
go mod tidy
```

You can then run the service using the following command:

```bash
docker-compose up
```

The service will be available at http://localhost:8080, you can test it out on the browser or using curl

```bash
curl http://localhost:8080/weather?city=London
```

In the docker container you will see these logs as an example

```bash
weather-api-service-app-1    | The city is: Pretoria, Gauteng, South Africa
weather-api-service-app-1    | cache-miss
weather-api-service-app-1    | cache-hit
```

For a detailed documentation you can visit the projects readme.
