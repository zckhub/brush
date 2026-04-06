package main

// 代码考核
// 题目 1
// 标题
// GPS数据流并发处理系统
// 题目描述
// 你正在开发一个GPS数据实时处理系统，系统需要处理来自多个终端的GPS数据流。每个GPS数据包包含以下信息：​
// 终端ID（string）​
// 时间戳（int64，Unix毫秒时间戳）​
// 纬度（float64）​
// 经度（float64）​
// 速度（float64，km/h）​
// 系统需要实现以下功能：​
// 并发数据接收：同时接收多个终端的数据​
// 数据聚合：按终端ID聚合最近1分钟的数据​
// 实时统计：计算每个终端的速度平均值​
// 异常检测：检测速度异常（超过120km/h）的数据点​
// 题目要求​
// 基础要求（必须完成）：​
// 实现GPSDataProcessor结构体，包含以下方法：​
// ProcessData(data GPSData)：处理单个GPS数据​
// GetTerminalStats(terminalID string) TerminalStats：获取终端统计信息​
// GetSpeedAlerts() []SpeedAlert：获取速度异常告警​
// 使用goroutine和channel实现并发处理​
// 保证线程安全​
// // ​
// type GPSData struct {​
//     TerminalID string​
//     Timestamp  int64​
//     Latitude   float64​
//     Longitude  float64​
//     Speed      float64​
// }​

// type TerminalStats struct {​
//     TerminalID    string​
//     DataCount     int​
//     AvgSpeed      float64​
//     LastUpdate    time.Time​
//     RecentData    []GPSData // 最近1分钟的数据​
// }​

// type SpeedAlert struct {​
//     TerminalID string​
//     Timestamp  int64​
//     Speed      float64
//     Reason     string
// }


import (
	"fmt"
	"sync"
	"time"
)

// GPSData 定义GPS数据包结构
type GPSData struct {
	TerminalID string
	Timestamp  int64
	Latitude   float64
	Longitude  float64
	Speed      float64
}

// TerminalStats 定义终端统计信息结构
type TerminalStats struct {
	TerminalID string
	DataCount  int
	AvgSpeed   float64
	LastUpdate time.Time
	RecentData []GPSData // 最近1分钟的数据
}

// SpeedAlert 定义速度告警结构
type SpeedAlert struct {
	TerminalID string
	Timestamp  int64
	Speed      float64
	Reason     string
}

// GPSDataProcessor GPS数据处理器
type GPSDataProcessor struct {
	dataChan    chan GPSData
	terminalMap map[string]*TerminalStats
	alerts      []SpeedAlert
	mu          sync.Mutex
}

// NewGPSDataProcessor 创建GPS数据处理器
func NewGPSDataProcessor() *GPSDataProcessor {
	p := &GPSDataProcessor{
		dataChan:    make(chan GPSData, 100), // 带缓冲防止阻塞
		terminalMap: make(map[string]*TerminalStats),
	}
	go p.run() // 启动后台处理goroutine
	return p
}

// ProcessData 处理单个GPS数据（并发接收入口）
func (p *GPSDataProcessor) ProcessData(data GPSData) {
	p.dataChan <- data
}

// run 后台循环处理数据
func (p *GPSDataProcessor) run() {
	for data := range p.dataChan {
		p.processSingleData(data)
	}
}

// processSingleData 处理单个数据的核心逻辑
func (p *GPSDataProcessor) processSingleData(data GPSData) {
	// 1. 检测速度异常
	if data.Speed > 120 {
		alert := SpeedAlert{
			TerminalID: data.TerminalID,
			Timestamp:  data.Timestamp,
			Speed:      data.Speed,
			Reason:     "Speed exceeds 120 km/h",
		}
		p.mu.Lock()
		p.alerts = append(p.alerts, alert)
		p.mu.Unlock()
	}

	// 2. 更新终端统计信息
	p.mu.Lock()
	defer p.mu.Unlock()

	// 获取或创建终端统计
	stats, exists := p.terminalMap[data.TerminalID]
	if !exists {
		stats = &TerminalStats{
			TerminalID: data.TerminalID,
			RecentData: make([]GPSData, 0),
		}
		p.terminalMap[data.TerminalID] = stats
	}

	// 追加当前数据
	stats.RecentData = append(stats.RecentData, data)

	// 清理1分钟前的过期数据
	cutoff := data.Timestamp - 60*1000
	filtered := make([]GPSData, 0, len(stats.RecentData))
	for _, d := range stats.RecentData {
		if d.Timestamp >= cutoff {
			filtered = append(filtered, d)
		}
	}
	stats.RecentData = filtered

	// 重新计算统计信息
	stats.DataCount = len(stats.RecentData)
	if stats.DataCount > 0 {
		totalSpeed := 0.0
		for _, d := range stats.RecentData {
			totalSpeed += d.Speed
		}
		stats.AvgSpeed = totalSpeed / float64(stats.DataCount)
	} else {
		stats.AvgSpeed = 0
	}
	stats.LastUpdate = time.Now()
}

// GetTerminalStats 获取指定终端的统计信息（返回深拷贝）
func (p *GPSDataProcessor) GetTerminalStats(terminalID string) TerminalStats {
	p.mu.Lock()
	defer p.mu.Unlock()

	stats, exists := p.terminalMap[terminalID]
	if !exists {
		return TerminalStats{TerminalID: terminalID}
	}

	// 深拷贝RecentData，防止外部修改内部状态
	recentCopy := make([]GPSData, len(stats.RecentData))
	copy(recentCopy, stats.RecentData)

	return TerminalStats{
		TerminalID: stats.TerminalID,
		DataCount:  stats.DataCount,
		AvgSpeed:   stats.AvgSpeed,
		LastUpdate: stats.LastUpdate,
		RecentData: recentCopy,
	}
}

// GetSpeedAlerts 获取所有速度告警（返回深拷贝）
func (p *GPSDataProcessor) GetSpeedAlerts() []SpeedAlert {
	p.mu.Lock()
	defer p.mu.Unlock()

	alertCopy := make([]SpeedAlert, len(p.alerts))
	copy(alertCopy, p.alerts)
	return alertCopy
}

func main() {
	processor := NewGPSDataProcessor()

	// 模拟3个终端并发发送数据
	terminalIDs := []string{"terminal-1", "terminal-2", "terminal-3"}
	var wg sync.WaitGroup

	for _, tid := range terminalIDs {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			baseTime := time.Now().UnixMilli()
			for i := 0; i < 10; i++ {
				speed := 60.0 + float64(i)*10
				// 让terminal-1产生一个超速数据
				if id == "terminal-1" && i == 9 {
					speed = 130.0
				}
				data := GPSData{
					TerminalID: id,
					Timestamp:  baseTime + int64(i)*1000, // 每秒1条数据
					Latitude:   31.2304 + float64(i)*0.001,
					Longitude:  121.4737 + float64(i)*0.001,
					Speed:      speed,
				}
				processor.ProcessData(data)
				time.Sleep(500 * time.Millisecond) // 模拟发送间隔
			}
		}(tid)
	}

	wg.Wait()
	time.Sleep(1 * time.Second) // 等待数据处理完成

	// 打印terminal-1的统计信息
	fmt.Println("=== Terminal-1 Stats ===")
	stats1 := processor.GetTerminalStats("terminal-1")
	fmt.Printf("Terminal ID: %s\n", stats1.TerminalID)
	fmt.Printf("Data Count: %d\n", stats1.DataCount)
	fmt.Printf("Avg Speed: %.2f km/h\n", stats1.AvgSpeed)
	fmt.Printf("Last Update: %s\n", stats1.LastUpdate.Format(time.RFC3339))
	fmt.Println("Recent Data:")
	for _, d := range stats1.RecentData {
		fmt.Printf("  Time: %d, Speed: %.2f\n", d.Timestamp, d.Speed)
	}

	// 打印所有速度告警
	fmt.Println("\n=== Speed Alerts ===")
	alerts := processor.GetSpeedAlerts()
	for _, a := range alerts {
		fmt.Printf("Terminal: %s, Time: %d, Speed: %.2f, Reason: %s\n",
			a.TerminalID, a.Timestamp, a.Speed, a.Reason)
	}
}
