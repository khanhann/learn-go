package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// The first problem that we have to solve is that we need a software that run on a timer. It need
// to connect to Xenia, read that database, identify all the data we haven't moved and pull it in.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Data is the structure of the data we are copying.
// For simplicity, just pretend that is is a string data.
type Data struct {
	Line string
}

// Xenia is a system we need to pull data from.
type Xenia struct {
	Host    string
	Timeout time.Duration
}

// Pull knows how to pull data out of Xenia.
// We could do func (*Xenia) Pull (*Data, error) that return the data and error. However, this
// would cost an allocation on every call we don't want that.
// Using the function below, we know data is a struct type and its size ahead of time. Therefore
// they could be on the stack.
func (*Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF

	case 5:
		return errors.New("Error reading data from Xenia")

	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

// Pillar is a system we need to store data into.
type Pillar struct {
	Host    string
	Timeout time.Duration
}

// Store knows how to store data into Pillar.
// We are using pointer semantics for consistency.
func (*Pillar) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

// System wraps Xenia and Pillar together into a single system.
// We have the API based on Xenia and Pillar. We want to build another API on top of this and use
// it as a foundation.
// One way is to have a type that have the behavior of being able to pull and store. We can do that
// through composition. System is based on the embedded value of Xenia and Pillar. And because of
// inner type promotion, System know how to pull and store.
type System struct {
	Xenia
	Pillar
}

// pull knows how to pull bulks of data from Xenia, leveraging the foundation that we built.
// We don't need to add method to System to do this. There is no state inside System that we want
// the system to maintain. Instead, we want the System to understand the behavior.
// Functions are a great way of writing API because functions can be more readable than any method
// can. We always want to start with an idea of writing API from the package level with function.
// When we write a function, all the input must be passed in. When we use a method, its signature
// doesn't indicate any level, what field or state that we are using on that value that we use to
// make the call.
func pull(x *Xenia, data []Data) (int, error) {
	// Range over the slice of data and share each element with the Xenial's Pull method.
	for i := range data {
		if err := x.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// store knows how to store bulks of data into Pillar.
// Similar to the function above.
// We might wonder if it is efficient. However, we are optimizing for correctness, not performance.
// When it is done, we will test it. If it is not fast enough, we will add more complexities to
// make it run faster.
func store(p *Pillar, data []Data) (int, error) {
	for i := range data {
		if err := p.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Copy knows how to pull and store data from the System.
// Now we can call the pull and store function, passing Xenia and Pillar through.
func Copy(sys *System, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(&sys.Xenia, data)
		if i > 0 {
			if _, err := store(&sys.Pillar, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

func main() {
	sys := System{
		Xenia: Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Pillar: Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
}
