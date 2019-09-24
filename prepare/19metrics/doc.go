package _9metrics

/*
metrics 是什么？
当我们需要为某个系统某个服务做监控、做统计，就需要用到Metrics

五种 Metrics 类型

Gauges ：最简单的度量指标，只有一个简单的返回值，或者叫瞬时状态
Counters：Counter 就是计数器，Counter 只是用 Gauge 封装了 AtomicLong
Meters：Meter度量一系列事件发生的速率(rate)，例如TPS。Meters会统计最近1分钟，5分钟，15分钟，还有全部时间的速率。
Histograms：Histogram统计数据的分布情况。比如最小值，最大值，中间值，还有中位数，75百分位, 90百分位, 95百分位, 98百分位, 99百分位, 和 99.9百分位的值(percentiles)。
Timer其实是 Histogram 和 Meter 的结合， histogram 某部分代码/调用的耗时， meter统计TPS。

*/
