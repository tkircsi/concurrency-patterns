package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello GoLand")

	// NumberGenerator
	/*	runeStream := generators.NewRuneGenerator(context.Background(), "Jónapot kívánok!")
		for r := range runeStream {
			fmt.Printf("%c", r)
		}*/

	// Fibonacci generator
	/*	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		fibStream := generators.NewFibonacciGenerator(ctx, 1000)
		for i := range fibStream {
			fmt.Printf("%d, ", i)
		}*/

	// Int generator
	// ctx, cancel := context.WithCancel(context.Background())
	/*	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		intStream := generators.NewIntGenerator(ctx, 10)
		defer cancel()

		for i := range intStream {
			fmt.Printf("%d ,", i)
			time.Sleep(time.Millisecond * 500)
		}*/

	// Or channel
	/*{
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go func() {
			fmt.Printf("%v\n", <-patterns.Or(
				ctx,
				cancel,
				time.After(time.Second*2),
				time.After(time.Second*3),
				time.After(time.Second*4),
			))
		}()

		fmt.Println("Some other time-consuming job...")
		time.Sleep(time.Second * 5)
		fmt.Println("Done")
	}*/

	// Check Http status (error handling)
	/*{
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		start := time.Now()
		statusStream := snippets.CheckHttpStatus(
			ctx,
			"https://www.google.hu",
			"https://www.fundamenta.hu",
			"https://badhost.com",
			"https://ebank.fundamenta.hu/informacio/tsz",
		)

		for stat := range statusStream {
			if stat.Error != nil {
				fmt.Printf("%s - %v\n", stat.Url, stat.Error)
				continue
			}
			fmt.Printf("%s\t%v\t%dms\n", stat.Url, stat.Response.Status, stat.Duration)
		}
		ms := time.Since(start) / time.Millisecond
		fmt.Printf("Ellapsed: %dms\n", ms)
		fmt.Println("Done")
	}*/

	// Pipelines
	/*{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		intStream := generators.NewIntSliceGenerator(ctx, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
		mul := pipelines.Multiply(ctx, intStream, 3)
		add := pipelines.Add(ctx, mul, 5)

		for n := range add {
			// check context timeout
			// time.Sleep(time.Second)
			fmt.Printf("%d, ", n)
		}

	}*/

	// Pipeline with generic stages
	/*{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		integers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		toInterfaceSlice := func(vv []int) []interface{} {
			iv := make([]interface{}, len(vv))
			for i := 0; i < len(iv); i++ {
				iv[i] = vv[i]
			}
			return iv
		}
		values := toInterfaceSlice(integers)
		rep := pipelines.Repeate(ctx, values...)
		take := pipelines.Take(ctx, rep, 25)
		results := pipelines.ToInt(ctx, take)

		for v := range results {
			fmt.Printf("%d, ", v)
		}
	}*/

	// Repeate a function
	/*{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		rand.Seed(time.Now().UnixNano())

		randFn := func() interface{} {
			return rand.Int()
		}

		randStream := pipelines.RepeateFn(ctx, randFn)
		results := pipelines.Take(ctx, randStream, 6)

		for v := range results {
			fmt.Println(v)
		}
	}*/

	// Simple pipeline (PrimeFinder takes a long time)
	/*{
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		rand.Seed(time.Now().UnixNano())
		randFn := func() interface{} {
			return rand.Intn(1000000000)
		}

		start := time.Now()
		randIntStream := stages.RepeateFn(ctx, randFn)
		fmt.Println("Primes:")
		for prime := range stages.Take(ctx, stages.PrimeFinder(ctx, randIntStream), 10) {
			fmt.Printf("\t%d\n", prime)
		}
		fmt.Printf("Search took: %v\n", time.Since(start))
	}*/

	// Fan-out Fan-in pipeline
	/*{
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		rand.Seed(time.Now().UnixNano())
		randFn := func() interface{} {
			return rand.Intn(1000000000)
		}

		start := time.Now()

		randIntStream := stages.RepeateFn(ctx, randFn)

		numFinders := runtime.NumCPU()
		fmt.Printf("Number of CPUs: %d\n", numFinders)
		fmt.Println("Fan-out Fan-in Primes:")
		finders := patterns.FanOut(ctx, numFinders, randIntStream, stages.PrimeFinder)
		primeStream := patterns.FanIn(ctx, finders...)

		for prime := range stages.Take(ctx, primeStream, 10) {
			fmt.Printf("\t%d\n", prime)
		}
		fmt.Printf("Search took: %v\n", time.Since(start))
	}*/

	// Without OrDone
	/*{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		intStream := generators.NewIntGenerator(ctx, 100)
		// intStream is closed after 5 seconds, but readable.
		// IntToInterface stage will run indefinitely because it can read closed channel
		// and does not verify channel read flags.
		ifStream := stages.IntToInterface(context.Background(), intStream)
		primeStream := stages.PrimeFinder(context.Background(), ifStream)

		for v := range primeStream {
			fmt.Printf("%d, ", v)
			time.Sleep(time.Second)
		}
	}*/

	// With OrDone
	/*{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		intStream := generators.NewIntGenerator(ctx, 100)
		ifStream := stages.IntToInterface(context.Background(), intStream)
		//orDoneStream := patterns.OrDone(context.Background(), ifStream)
		primeStream := stages.PrimeFinder(context.Background(), ifStream)

		for v := range primeStream {
			fmt.Printf("%d, ", v)
			time.Sleep(time.Second)
		}
	}*/

	// Tee channel
	/*{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		intStream := generators.NewIntGenerator(ctx, 100)
		ifStream := stages.IntToInterface(ctx, intStream)
		out1, out2 := patterns.Tee(ctx, ifStream)

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			for v := range out1 {
				fmt.Printf("out1: %d\n", v)
				//time.Sleep(time.Second)
			}
		}()
		go func() {
			defer wg.Done()
			for v := range out2 {
				fmt.Printf("out2: %d\n", v)
				//time.Sleep(time.Second)
			}
		}()
		wg.Wait()
	}*/
}
