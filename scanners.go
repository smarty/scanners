// Package scanners provides scanners for text files that encode
// data as CSV, space-delimited fields, or fixed-width columns.
//
// All three scanners either emulate or wrap a bufio.Scanner,
// and incorporate the bufio.Scanner style of defining a scan-loop,
// looping, and then checking for errors after the scan-loop has
// completed:
//
//	scanner := SomeNewScanner()
//
//	for scanner.Scan() {
//	    scanner.GetSomeValues()
//	}
//
//	if err := scanner.Err(); err != nil {
//	    log.Fatal(err)
//	}
package scanners
