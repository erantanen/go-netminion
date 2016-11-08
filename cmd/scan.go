package cmd

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// entrypoint is a host and port combination with test results
type entrypoint struct {
	Address string
	Port    int
	Open    bool
}

// scanCmd initiates a scan
var scanCmd = &cobra.Command{
	Use:   "scan [address...]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Expand all expression statements
		ports := explodePorts(strings.Split(portExpressions, ",")...)
		addresses := explodeAddresses(args...)

		// Loop through all possible entry points
		c := make(chan entrypoint)
		for _, address := range addresses {
			for _, port := range ports {
				go func(a string, p int) {
					c <- scan(a, p)
				}(address, port)
			}
		}

		// Loop until we timeout or receive all possible responses
		l := len(addresses) * len(ports)
		t := 0
		d := duration(timeout)
		start := time.Now()
		for i := 0; i < l; i++ {
			select {
			case r := <-c:
				t++
				status := "Closed"
				if r.Open {
					status = "Opened"
				}
				fmt.Printf("[ %s ] Host: %s, Port: %d\n", status, r.Address, r.Port)
			case <-time.After(d):
				// Timeout
			}
		}

		fmt.Printf("\nRecieved %d responses out of %d (%.2f%%) possibilities %s\n",
			t,
			l,
			(float64(t)/float64(l))*100,
			time.Now().Sub(start))
	},
}

// explodePorts takes an indeterminate amount of expressions
// and returns an array of included ports
func explodePorts(expressions ...string) []int {
	ports := []int{}

	for _, exp := range expressions {
		bounds := strings.SplitN(exp, "-", 2)

		if len(bounds) == 1 {
			if p, err := strconv.Atoi(exp); err == nil {
				ports = append(ports, p)
			}
		} else {
			lower, err := strconv.Atoi(bounds[0])
			if err != nil {
				continue
			}

			upper, err := strconv.Atoi(bounds[1])
			if err != nil {
				continue
			}

			for i := lower; i <= upper; i++ {
				ports = append(ports, i)
			}
		}
	}

	return ports
}

// explodeAddresses takes an indeterminate amount of expressions
// and returns an array of included addresses
func explodeAddresses(expressions ...string) []string {
	addresses := []string{}

	// TODO: Functionality not implemented yet!
	for _, exp := range expressions {
		addresses = append(addresses, exp)
	}

	return addresses
}

//duration  helps parse user input in to a meaningful value
func duration(timeout string) time.Duration {
	if ms, err := strconv.Atoi(timeout); err == nil {
		return time.Millisecond * time.Duration(ms)
	} else if d, err := time.ParseDuration(timeout); err == nil {
		return d
	}

	return time.Duration(time.Millisecond * 50)
}

// scan a single host and port and return the results
func scan(address string, port int) entrypoint {
	r := entrypoint{
		Address: address,
		Port:    port,
		Open:    false,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
	if err == nil {
		_ = conn.Close()
		r.Open = true
	}

	return r
}
