package data_processor

func Start(inc_data <-chan string, out_data chan uint64) {
	var addresses = make(map[string]bool)
	for {
		select {
		case data := <-inc_data:
			processAddresses(&addresses, data)
			out_data <- uint64(len(addresses))
		}
	}
}

func processAddresses(addresses *map[string]bool, data string) {

}
