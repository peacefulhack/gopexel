package utils

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"os"
	"strings"
)

func GetAllEnv(length int) ([]string, error) {
	var value []string
	err := godotenv.Load("./.env")
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	envVars := os.Environ()
	for i, val := range envVars {
		if i > length-1 {
			break
		}
		realVal := strings.Split(val, "=")
		if len(realVal) <= 1 {
			return nil, status.Errorf(codes.Internal, "Please add all value in .env there was some missing")
		}
		value = append(value, realVal[1])
	}
	return value, nil
}

func Randomizer(max, min int) int {
	randNum := rand.Intn(max-min+1) + min
	return randNum
}
