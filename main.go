package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
)

var (
	Secret string
)

func init() {
	flag.StringVar(&Secret, "secret", "", "TOTP secret")
	flag.Parse()

	if flag.NArg() > 0 {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	if Secret == "" {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		Secret = string(b)
	}

	t, err := totp.GenerateCode(Secret, time.Now())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(t)
}
