package metrics

func FormatMetricName(metricName string) string {
	if len(metricName) == 0 {
		return metricName
	}

	n := len(metricName)
	validMetricsName := make([]byte, n)
	for i := 0; i < n; i++ {
		if c := metricName[i]; IsValidMetricName(c) {
			validMetricsName[i] = c
		} else {
			validMetricsName[i] = InvalidByteReplacement
		}
	}
	return string(validMetricsName)
}

func IsValidMetricName(in byte) bool {
	return IsLowerCase(in) || IsUpperCase(in) || IsDigit(in) || in == '_' || in == ':'
}

func IsLowerCase(in byte) bool {
	return in >= 'a' && in <= 'z'
}

func IsUpperCase(in byte) bool {
	return in >= 'A' && in <= 'Z'
}

func IsDigit(in byte) bool {
	return in >= '0' && in <= '9'
}
