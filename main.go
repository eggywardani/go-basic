package main

import (
	_ "belajar-golang-dasar/database"
	"belajar-golang-dasar/helper"
	"container/list"
	"container/ring"
	"flag"
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Name string
	Age  int
}

type Sample struct {
	Name string `required:"true" max:"10"`
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {

	helper.SayHello("Eggy Andika Wardani")

	// result := database.GetDatabase()
	// fmt.Println(result)

	// package os

	fmt.Println("==========PACKAGE OS==========")
	args := os.Args
	fmt.Println("Argument :")
	fmt.Println(args)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", name)
	} else {
		fmt.Println("Error :", err)
	}

	fmt.Println("==========PACKAGE FLAG==========")
	host := flag.String("host", "localhost", "put your database host")
	username := flag.String("username", "root", "put your database username")
	password := flag.String("password", "root", "put your database username")

	flag.Parse()
	fmt.Println(*host, *username, *password)

	fmt.Println("==========PACKAGE STRING==========")

	fmt.Println(strings.Contains("Eggy Andika Wardani", "eggy"))
	fmt.Println(strings.Split("Eggy Andika Wardani", " "))
	fmt.Println(strings.ToLower("Eggy Andika Wardani"))
	fmt.Println(strings.ToUpper("Eggy Andika Wardani"))
	fmt.Println(strings.ToTitle("eggy andika wardani"))
	fmt.Println(strings.Trim("       Eggy andika wardani      ", " "))

	fmt.Println("==========PACKAGE STRCONV==========")

	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {

		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)

	if err == nil {
		fmt.Println(number)
	} else {

		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10)

	fmt.Println(value)

	fmt.Println(strconv.Itoa(1999))
	fmt.Println(strconv.Atoi("90999"))

	fmt.Println("==========PACKAGE MATH==========")

	fmt.Println(math.Round(1.3))
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Floor(1.9))
	fmt.Println(math.Ceil(1.1))

	fmt.Println(math.Max(11, 99))

	fmt.Println("==========PACKAGE CONTAINER/LIST==========")

	data := list.New()
	data.PushBack("Eggy")
	data.PushBack("Andika")
	data.PushBack("wardani")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println("==========PACKAGE CONTAINER/RING==========")
	data2 := ring.New(5)

	for i := 0; i < data2.Len(); i++ {
		data2.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data2 = data2.Next()
	}

	data2.Do(func(value interface{}) {
		fmt.Println(value)
	})

	fmt.Println("==========PACKAGE SORT==========")

	users := []User{
		{"Eggy", 20},
		{"Andika", 22},
		{"Wardani", 30},
		{"Rina", 21},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)

	fmt.Println("==========PACKAGE TIME==========")

	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2020, time.November, 23, 12, 53, 28, 1111, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-01"
	parse, _ := time.Parse(layout, "2020-10-12")
	fmt.Println(parse)

	fmt.Println("==========PACKAGE REFLECT==========")

	sample := Sample{""}
	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))

	fmt.Println(isValid(sample))

	fmt.Println("==========PACKAGE REGEXP==========")
	regex := regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eggo"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("ridho"))

}

func isValid(data interface{}) bool {

	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}
