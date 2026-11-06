package prometheus

func histogramQuantile(q, metric, selector, rate string) string {
	return "histogram_quantile(" + q + ", sum by (le) (rate(" + metric + "_bucket" + selector + rate + ")))"
}

func histogramAverage(metric, selector, rate string) string {
	return "sum(rate(" + metric + "_sum" + selector + rate + ")) / sum(rate(" + metric + "_count" + selector + rate + "))"
}
