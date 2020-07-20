package data_processor

func Start(inc_data <-chan string, out_data chan<- string) {
	for {
		select {
		case data := <-inc_data:
			out_data <- data
		}
	}
}
