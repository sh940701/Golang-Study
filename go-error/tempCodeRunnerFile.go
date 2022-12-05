f !scanner.Scan() {
		return 0, 0, fmt.Errorf("Fail to scan")
	}