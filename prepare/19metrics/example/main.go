package main

import (
	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"time"
)

/*
五种 Metrics 类型

Gauges ：最简单的度量指标，只有一个简单的返回值，或者叫瞬时状态
Counters：Counter 就是计数器，Counter 只是用 Gauge 封装了 AtomicLong
Meters：Meter度量一系列事件发生的速率(rate)，例如TPS。Meters会统计最近1分钟，5分钟，15分钟，还有全部时间的速率。
Histograms：Histogram统计数据的分布情况。比如最小值，最大值，中间值，还有中位数，75百分位, 90百分位, 95百分位, 98百分位, 99百分位, 和 99.9百分位的值(percentiles)。
Timer其实是 Histogram 和 Meter 的结合， histogram 某部分代码/调用的耗时， meter统计TPS。

*/

func gauges() {
	g := metrics.NewGauge()
	metrics.Register("bar", g)
	g.Update(1)

	var j int64
	j = 1
	for {
		time.Sleep(time.Second * 1)
		g.Update(j)
		j++
	}
}

// counters 与gauges 类似 只不过在操作上 gauges 是 update 而
// counter 是 inc 做加法 增加参数市值 dec 做减法
func counters() {
	c := metrics.NewCounter()
	//c := metrics.NewGaugeFloat64()
	metrics.Register("foo", c)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				c.Dec(19)
				time.Sleep(300e6)
			}
		}()
		go func() {
			for {
				c.Inc(47)
				time.Sleep(400e6)
			}
		}()
	}
	time.Sleep(50 * time.Second)
}

func meters() {
	m := metrics.NewMeter()
	metrics.Register("quux", m)
	var j int64
	j = 1
	for {
		time.Sleep(time.Second * 1)
		j++
		m.Mark(j)
	}
}

func histogram() {
	s := metrics.NewExpDecaySample(1028, 0.015)
	//s := metrics.NewUniformSample(1028)
	h := metrics.NewHistogram(s)
	metrics.Register("bang", h)
	go func() {
		for {
			h.Update(47)
			time.Sleep(400e6)
		}
	}()
	for {
		h.Update(19)
		time.Sleep(300e6)
	}
}

func timer() {
	t := metrics.NewTimer()
	metrics.Register("hooah", t)

	go func() {
		for {
			t.Time(func() { time.Sleep(400e6) })
		}
	}()

	for {
		t.Time(func() { time.Sleep(300e6) })
	}

}

func main() {
	// 结果输出
	output()

	//gauges()
	//counters()
	//meters()
	//histogram()
	timer()
}

func output() {
	// 输出到控制台
	go metrics.Log(metrics.DefaultRegistry, 1*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	//输出到influxdb
	/*go influxdb.InfluxDB(
	metrics.DefaultRegistry,
	time.Second*10,
	"http://127.0.0.1:8086",
	"mydb",
	"",
	"",
	"",
	true)*/

}
